package main

import (
	"database/sql"
	"fmt"

	"github.com/pingcap/log"
	"go.uber.org/zap"

	_ "github.com/go-sql-driver/mysql"
)

//数据库连接信息
const (
	RegionShanghai = "shanghai"
	RegionShenzhen = "shenzhen"
	RegionGlobal   = "global"
	Username       = "root"
	Password       = ""
	ShanghaiHost   = "10.200.112.223"
	ShenzhenHost   = "10.200.112.107"
	GlobalHost     = "10.200.112.212"
	DbHookPort     = 9436
	DataBase       = "csh"
	Table          = "tb"
	Network        = "tcp"
)

var (
	id           int
	content      string
	_id          int64
	_region      string
	_ts          int64
	_create_time int64
	_update_time int64
)

func DbHost(region string) string {
	switch region {
	case RegionShanghai:
		return ShanghaiHost
	case RegionShenzhen:
		return ShenzhenHost
	case RegionGlobal:
		return GlobalHost
	default:
		return ""
	}
}

func dbConn(region string) *sql.DB {
	src := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", Username, Password, Network, DbHost(region), DbHookPort, DataBase)
	db, err := sql.Open("mysql", src)
	if err != nil {
		panic(err)
	}
	return db
}

func QueryAll(region string) {
	db := dbConn(region)
	rows, _ := db.Query("select * from tb;")
	defer rows.Close()
	for rows.Next() {
		_ = rows.Scan(&_id, &_region, &_ts, &_create_time, &_update_time, &id, &content)
		log.Info("query from db",
			zap.Any("_id", _id), zap.Any("_region", _region), zap.Any("_ts", _ts),
			zap.Any("_create_time", _create_time), zap.Any("_update_time", _update_time),
			zap.Any("id:", id), zap.Any("content:", content),
		)
	}
}

func Query(region string, queryId int) {
	db := dbConn(region)
	rows, _ := db.Query("select * from tb where id = ?;", queryId)
	defer rows.Close()
	for rows.Next() {
		_ = rows.Scan(&_id, &_region, &_ts, &_create_time, &_update_time, &id, &content)
		log.Info("query from db",
			zap.Any("_id", _id), zap.Any("_region", _region), zap.Any("_ts", _ts),
			zap.Any("_create_time", _create_time), zap.Any("_update_time", _update_time),
			zap.Any("id:", id), zap.Any("content:", content),
		)
	}
}

func Insert(region string, insertId int, insertContent string) {
	db := dbConn(region)
	stmt, _ := db.Prepare("insert into tb (id, content) values(?, ?);")
	_, _ = stmt.Exec(insertId, insertContent)
	log.Info("insert finished")
}

func Update(region string, updateId int, updateContent string) {
	db := dbConn(region)
	stmt, _ := db.Prepare("update tb set content = ? where id = ?;")
	_, _ = stmt.Exec(updateContent, updateId)
	log.Info("update finished")
}

func Delete(region string, deleteId int) {
	db := dbConn(region)
	stmt, _ := db.Prepare("delete from tb where id = ?;")
	_, _ = stmt.Exec(deleteId)
	log.Info("delete finished")
}

func main() {
	Query(RegionShanghai, 10)
	Insert(RegionShanghai, 11, "a")
	QueryAll(RegionShanghai)
	Update(RegionShanghai, 11, "t")
	QueryAll(RegionShanghai)
	Delete(RegionShanghai, 11)
	QueryAll(RegionShanghai)
}
