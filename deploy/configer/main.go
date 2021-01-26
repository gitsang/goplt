package main

import (
	"configor/config"
	"github.com/pingcap/log"
	"go.uber.org/zap"
)

func main() {
	// 1. read config from file and environment
	err := config.LoadConfig("")
	if err != nil {
		log.Error("load config from file failed", zap.Error(err))
	}

	// 3. read config from cfgserver
	// pretend to read

	// 1. start server
	//go func() {
	//	testCfg.cond.L.Lock()
	//	defer testCfg.cond.L.Unlock()

	//	for {
	//		fmt.Println("server 1", testCfg.Info.Id)
	//		testCfg.cond.Wait()
	//	}
	//}()

	//go func() {
	//	testCfg.cond.L.Lock()
	//	defer testCfg.cond.L.Unlock()

	//	for {
	//		fmt.Println("server 2", testCfg.Info.Id)
	//		testCfg.cond.Wait()
	//	}
	//}()

	// 1. listening config changed
	//for {
	//	select {
	//	case <-time.Tick(10 * time.Second):
	//	}

	//	// pretend to receive config change
	//	newConf := &TestConfig{
	//		Name: "jsr",
	//		Info: struct {
	//			Id  int
	//			Age int
	//		}{Id:  time.Now().Second(), Age: 25},
	//	}
	//	testCfg.UpdateConfig(newConf)
	//}
}
