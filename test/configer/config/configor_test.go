package config

import (
	"github.com/jinzhu/configor"
	"os"
	"testing"
)


func TestConfigor(t *testing.T) {
	_ = os.Setenv("SYNC_MANAGER_APP_NAME", "aaaaa")

	cfg := struct {
		App struct{
			Name string `default:"xxxx"`
		}
	}{}

	cfgr := configor.New(&configor.Config{
		ENVPrefix:   "SYNC_MANAGER",
		Verbose:     true,
	})

	err := cfgr.Load(&cfg, "")
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
}
