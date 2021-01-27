package CfgAgent

import (
	"context"
	"encoding/json"
	"errors"
	"gitcode.yealink.com/server/server_framework/go-utils/ylog"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"reflect"
	"strconv"
	"sync"
	"time"
)

type ConfigCallback func(v string)

var (
	cfgagentAddr      string
	cfgagentApp       string
	cfgagentNamespace string
	globalConfLock    *sync.Mutex
	ConfigCallbacks   map[string]ConfigCallback
)

func init() {
	ConfigCallbacks = make(map[string]ConfigCallback)
}

func Init(mtx *sync.Mutex, addr, app, namespace string) {
	globalConfLock = mtx
	cfgagentAddr = addr
	cfgagentApp = app
	cfgagentNamespace = namespace
}

func RegisterCallback(key string, f ConfigCallback) {
	ConfigCallbacks[key] = f
}

func RunCallback(key, value string) {
	if f, exist := ConfigCallbacks[key]; exist {
		f(value)
	}
}

func load() error {
	conn, err := grpc.Dial(cfgagentAddr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		ylog.Error("dial failed", zap.Error(err))
		return err
	}
	defer func() {
		_ = conn.Close()
	}()
	cfgagent := NewCfgAgentClient(conn)

	resp, err := cfgagent.GetCfg(context.Background(), &GetCfgReq{
		App:       cfgagentApp,
		Namespace: cfgagentNamespace,
	})
	if err != nil {
		ylog.Error("get cfg failed", zap.Error(err))
		return err
	}

	for _, kv := range resp.Kv {
		ylog.Info("load kv", zap.Any("kv", kv))
		RunCallback(kv.Key, kv.Value)
	}

	return nil
}

func watch() {
	var conn *grpc.ClientConn
	var err error

	for {
		time.Sleep(time.Second)
		if conn != nil {
			_ = conn.Close()
			conn = nil
		}

		conn, err = grpc.Dial(cfgagentAddr, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			ylog.Error("dial failed", zap.Error(err))
			continue
		}
		cfgagent := NewCfgAgentClient(conn)

		stream, err := cfgagent.WatchCfg(context.Background(), &GetCfgReq{
			App:       cfgagentApp,
			Namespace: cfgagentNamespace,
		})
		if err != nil {
			ylog.Error("grpc get watch client failed", zap.Error(err))
			continue
		}

		// receiving loop
		for {
			resp, err := stream.Recv()
			if err != nil {
				ylog.Error("recv failed", zap.Error(err))
				break
			}

			for _, kv := range resp.Kv {
				ylog.Info("recv kv changed", zap.Any("kv", kv))
				RunCallback(kv.Key, kv.Value)
			}
		}
	}
}

func processTags(config interface{}) error {
	configValue := reflect.Indirect(reflect.ValueOf(config))
	if configValue.Kind() != reflect.Struct {
		return errors.New("invalid config, should be struct")
	}

	configType := configValue.Type()
	for i := 0; i < configType.NumField(); i++ {
		var (
			typeField  = configType.Field(i)
			valueField = configValue.Field(i)
			key = typeField.Tag.Get("cfgserver")
		)
		ylog.Debug("try load from cfgserver", zap.Any("key", key), zap.Any("kind", reflect.Indirect(valueField).Kind()))

		if !valueField.CanAddr() || !valueField.CanInterface() {
			continue
		}

		// register SetValueFunc
		if key != "" {
			kind := reflect.Indirect(valueField).Kind()
			switch kind {
			case reflect.String:
				RegisterCallback(key, func(v string) {
					globalConfLock.Lock()
					defer globalConfLock.Unlock()
					valueField.Set(reflect.ValueOf(v))
				})
			case reflect.Int:
				RegisterCallback(key, func(v string) {
					globalConfLock.Lock()
					defer globalConfLock.Unlock()
					vInt, _ := strconv.Atoi(v)
					valueField.Set(reflect.ValueOf(vInt))
				})
			case reflect.Bool:
				RegisterCallback(key, func(v string) {
					globalConfLock.Lock()
					defer globalConfLock.Unlock()
					vBool, _ := strconv.ParseBool(v)
					valueField.Set(reflect.ValueOf(vBool))
				})
			case reflect.Slice:
				fallthrough
			case reflect.Struct:
				ylog.Debug("read struct from cfgserver", zap.Any("key", key))
				RegisterCallback(key, func(v string) {
					globalConfLock.Lock()
					defer globalConfLock.Unlock()
					err := json.Unmarshal([]byte(v), valueField.Addr().Interface())
					if err != nil {
						ylog.Error("jsonUnmarshal failed", zap.Error(err))
					}
				})
			}
		}

		for valueField.Kind() == reflect.Ptr {
			valueField = valueField.Elem()
		}

		if valueField.Kind() == reflect.Struct {
			if err := processTags(valueField.Addr().Interface()); err != nil {
				return err
			}
		}

		if valueField.Kind() == reflect.Slice {
			for i := 0; i < valueField.Len(); i++ {
				if reflect.Indirect(valueField.Index(i)).Kind() == reflect.Struct {
					if err := processTags(valueField.Addr().Interface()); err != nil {
						return nil
					}
				}
			}
		}
	}

	return nil
}

func Load(config interface{}) error {
	err := processTags(config)
	if err != nil {
		ylog.Error("process tags failed", zap.Error(err))
		return err
	}

	err = load()
	if err != nil {
		ylog.Error("load cfg failed", zap.Error(err))
		return err
	}
	go watch()

	return nil
}
