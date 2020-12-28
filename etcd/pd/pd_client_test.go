package skip_pump_history_list

import (
	"context"
	pd "github.com/pingcap/pd/client"
	"go.uber.org/zap"
	"testing"
)

var TestPdAddrs = []string{"tidb-cn.l7i.top:2379"}

func TestPdConn(t *testing.T) {
	cli, err := pd.NewClient(TestPdAddrs, pd.SecurityOption{})
	if err != nil {
		t.Error("PdConn failed", zap.Any("PdHostPort", TestPdAddrs), zap.Error(err))
		t.FailNow()
	}
	t.Log("PdConn success", zap.Any("PdHostPort", TestPdAddrs))

	ctx := context.Background()
	clusterId := cli.GetClusterID(ctx)
	phyTs, logicTs, err := cli.GetTS(ctx)
	if err != nil {
		t.Error("get ts failed", zap.Error(err))
		t.FailNow()
	}
	t.Log("success", zap.Any("clusterId", clusterId),
		zap.Any("phyTs", phyTs), zap.Any("logicTs", logicTs))
}
