package main

import (
	"encoding/json"
	"github.com/gitsang/golog"
	"go.uber.org/zap"
	"reflect"
	"strings"
	"testing"
)

type ST2 struct {
	Id int `json:"id"`
	Port int `json:"port"`
}

type ST1 struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

type ST struct {
	Map map[ST1]ST2 `json:"map"`
	//Map map[string]ST2 `json:"map"`
}

type NST1 struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Id int `json:"id"`
	Port int `json:"port"`
}

type NST struct {
	List []NST1 `json:"map"`
}

func (st *ST) UnmarshalJSON(data []byte) error {
	rv := reflect.ValueOf(st)
	kind := rv.Kind()
	log.Info("", zap.Any("kind", kind.String()))

	return nil
}

func TestJsonMap(t *testing.T) {
	jstr0 := "{ \"map\": { { \"name\": \"js\", \"age\": 18 }: { \"id\": 10, \"port\": 1234 }, { \"name\": \"js2\", \"age\": 8 }: { \"id\": 12, \"port\": 1222 } } }"
	log.InfoS("origin", jstr0)

	jstr1 := strings.TrimSpace(jstr0)
	log.InfoS("trimspace", jstr1)

	jstr2 := strings.ReplaceAll(jstr1, "}:", "},")
	log.InfoS("replace", jstr2)

	jstr3 := strings.ReplaceAll(jstr2, "{ {", "[ {")
	log.InfoS("replace", jstr3)

	jstr4 := strings.ReplaceAll(jstr3, "} }", "} ]")
	log.InfoS("replace", jstr4)

	var nst NST
	e := json.Unmarshal([]byte(jstr4), &nst)
	if e != nil {
		log.Error("", zap.Error(e))
	}
	log.Info("", zap.Any("nst", nst))

	var st ST
	st.Map = make(map[ST1]ST2)
	var tst1 ST1
	var tst2 ST2
	for _, v := range nst.List {
		if v.Name != "" {
			tst1 = ST1{
				Name: v.Name,
				Age:  v.Age,
			}
		}
		if v.Name == "" {
			tst2 = ST2{
				Id:   v.Id,
				Port: v.Port,
			}

			st.Map[tst1] = tst2
		}
	}
	log.InfoS("st", st)
}

func TestStringSp(t *testing.T) {

}
