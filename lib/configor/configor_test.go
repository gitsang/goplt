package config

import (
	"github.com/jinzhu/configor"
	"os"
	"testing"
)

var cfg = struct {
	Master struct {
		Name  string
		Addrs []string
	}

	File struct {
		FileName  string
		FileAddrs []string
	}

	Env struct {
		EnvName  string
		EnvAddrs []string
	}
}{}

func TestConfigFromDefault(t *testing.T) {
	cfg := struct {
		Default struct {
			Name     string   `default:"default_name"`
			Addrs    []string `default:"[def1.cfg.com:80, def2.cfg.com:80, def3.cfg.com:80]"`
			Contacts []struct {
				Name  string
				Email string
			} `default:"[{name: sam, email: sam@cfg.com}, {name: tom, email: tom@cfg.com}]"`
		}
	}{}

	cfgr := configor.New(&configor.Config{
		Environment: "default",
		Verbose:     true,
	})

	err := cfgr.Load(&cfg, "")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}

func TestConfigFromEnv(t *testing.T) {
	_ = os.Setenv("PREFIX_ENV_NAME",
		"name_from_env")
	_ = os.Setenv("PREFIX_ENV_ADDRS",
		"[env1.cfg.com:80, env2.cfg.com:80, env3.cfg.com:80]")
	_ = os.Setenv("PREFIX_ENV_CONTACTS",
		"[{name: sam, email: sam@cfg.com}, {name: tom, email: tom@cfg.com}]")

	cfg := struct {
		Env struct{
			Name string `required:"true"`
			Addrs []string `required:"true"`
			Contacts []struct {
				Name  string `required:"true"`
				Email string
			} `required:"true"`
		}
	}{}

	cfgr := configor.New(&configor.Config{
		Environment: "env",
		ENVPrefix:   "prefix",
		Verbose:     true,
	})

	err := cfgr.Load(&cfg, "")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}

func TestConfigor(t *testing.T) {

	err := configor.New(&configor.Config{
		Environment:          "test",
		ENVPrefix:            "TEST_CONFIGOR",
		Debug:                true,
		Verbose:              true,
		Silent:               true,
		AutoReload:           true,
		AutoReloadInterval:   10,
		AutoReloadCallback:   nil,
		ErrorOnUnmatchedKeys: true,
	}).Load(&cfg, "configor_test.yml")

	if err != nil {
		t.Error(err)
	}
	t.Log(cfg)
}
