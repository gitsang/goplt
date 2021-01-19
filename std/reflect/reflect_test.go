package reflect

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

type Config struct {
	ConfigName string `json:"config-name" necessary:"true"`
	ConfigSize int    `json:"config-size" necessary:"true"`
	ConfigBool bool   `json:"config-bool"`
}

func TestReflect(t *testing.T) {
	srcConf := Config{
		ConfigName: "name1",
		ConfigSize: 1,
		ConfigBool: false,
	}
	dstConf := Config{
		ConfigName: "name2",
		ConfigSize: 2,
		ConfigBool: true,
	}
	_ = os.Setenv("CONFIG_NAME", dstConf.ConfigName)
	_ = os.Setenv("config_size", strconv.Itoa(dstConf.ConfigSize))
	_ = os.Setenv("COnfIG_BOoL", strconv.FormatBool(dstConf.ConfigBool))

	elem := reflect.ValueOf(&srcConf).Elem()
	for i := 0; i < elem.NumField(); i++ {
		tag := elem.Type().Field(i).Tag.Get("json")
		key := strings.ReplaceAll(tag, "-", "_")
		filed := elem.Field(i)

		if filed.Type().Kind() == reflect.String {
			if v := os.Getenv(key); v != "" {
				filed.Set(reflect.ValueOf(v))
			}
		} else if filed.Type().Kind() == reflect.Int {
			if v := os.Getenv(key); v != "" {
				vInt, _ := strconv.Atoi(v)
				filed.Set(reflect.ValueOf(vInt))
			}
		} else if filed.Type().Kind() == reflect.Bool {
			if v := os.Getenv(key); v != "" {
				vBool, _ := strconv.ParseBool(v)
				filed.Set(reflect.ValueOf(vBool))
			}
		}

		necessary, _ := strconv.ParseBool(elem.Type().Field(i).Tag.Get("necessary"))
		if necessary {
			if filed.Type().Kind() == reflect.String {
				if filed.String() == "" {
					t.Error(tag, "is necessary but nil!", filed.String())
				}
			} else if filed.Type().Kind() == reflect.Int {
				if filed.Int() == 0 {
					t.Error(tag, "is necessary but nil!", filed.Int())
				}
			}
		} else {
			t.Log(tag, "is unnecessary")
		}
	}
	t.Log(srcConf)

	if srcConf != dstConf {
		t.FailNow()
	}
}

func TestFromFile(t *testing.T) {
	srcConf := Config{}

	body, err := ioutil.ReadFile("reflect_test.json")
	if err != nil {
		t.Error(err)
	}
	err = json.Unmarshal(body, &srcConf)
	if err != nil {
		t.Error(err)
	}

	t.Log(srcConf)
}
