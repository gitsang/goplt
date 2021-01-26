package config

import (
	CfgAgent "configor/config/cfgagent"
	"gitcode.yealink.com/server/server_framework/go-utils/ylog"
	"github.com/jinzhu/configor"
	"go.uber.org/zap"
	"sync"
)

type Config struct {
	Cfgagent struct {
		Addr      string `require:"true"`
		App       string `require:"true"`
		Namespace string `require:"true"`
	}
	Default struct {
		Name     string   `default:"default_name"`
		Addrs    []string `default:"[def1.cfg.com:80, def2.cfg.com:80, def3.cfg.com:80]"`
		Contacts []struct {
			Name  string
			Email string
		} `default:"[{name: sam, email: sam@cfg.com}, {name: tom, email: tom@cfg.com}]"`
	}
	Cfg struct {
		IdOne int `cfgserver:"id_one"`
		IdTwo int `cfgserver:"id_two"`
	}
}

var globalConf *Config
var globalConfLock *sync.Mutex

func init() {
	globalConf = new(Config)
}

func LoadConfig(file string) error {
	err := configor.Load(globalConf, file)
	if err != nil {
		return err
	}

	addr := globalConf.Cfgagent.Addr
	app := globalConf.Cfgagent.App
	namespace := globalConf.Cfgagent.Namespace
	CfgAgent.Init(globalConfLock, addr, app, namespace)

	err = CfgAgent.Load(globalConf)
	if err != nil {
		return err
	}

	return nil
}

func PrintConfig() {
	globalConfLock.Lock()
	defer globalConfLock.Unlock()
	ylog.Info("print config", zap.Any("config", globalConf))
}
