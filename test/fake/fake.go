package main

import (
	"database/sql"
	"fmt"
	"github.com/gitsang/golog"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
	"sync"
)

const (
	Username   = "root"
	Password   = ""
	CnDB       = "tidb-cn.l7i.top"
	UsDB       = "tidb-us.l7i.top"
	GlobalDb   = "tidb-global.l7i.top"
	DbHookPort = "9436"
	Network    = "tcp"

	CNStart = 100000000
	USStart = 200000000
)

var cfg = Config{
	Databases: []*Database{
		{
			Name: "sang",
			Init: false,
			Tables: []*Table{
				{Name: "tab01", Cnt: 100000, Init: true},
				{Name: "tab02", Cnt: 100000, Init: true},
				{Name: "tab03", Cnt: 1000, Init: true},
				{Name: "tab04", Cnt: 1000, Init: true},
				{Name: "tab05", Cnt: 1000, Init: true},
				{Name: "tab06", Cnt: 1000, Init: true},
				{Name: "tab07", Cnt: 1000, Init: true},
				{Name: "tab08", Cnt: 1000, Init: true},
				{Name: "tab09", Cnt: 1000, Init: true},
				{Name: "tab10", Cnt: 1000, Init: true},
				{Name: "tab11", Cnt: 1000, Init: true},
				{Name: "tab12", Cnt: 1000, Init: true},
				{Name: "tab13", Cnt: 1000, Init: true},
				{Name: "tab14", Cnt: 1000, Init: true},
				{Name: "tab15", Cnt: 1000, Init: true},
				{Name: "tab16", Cnt: 1000, Init: true},
				{Name: "tab17", Cnt: 1000, Init: true},
				{Name: "tab18", Cnt: 1000, Init: true},
				{Name: "tab19", Cnt: 1000, Init: true},
				{Name: "tab20", Cnt: 1000, Init: true},
			},
		},
	},
}

type Table struct {
	Name string
	Init bool
	Cnt  int
}

type Database struct {
	Name   string
	Init   bool
	Tables []*Table
}

type Config struct {
	Databases []*Database
}

func dbConn(username, password, network, dbhost, dbport, database string) (*sql.DB, error) {
	src := fmt.Sprintf("%s:%s@%s(%s:%s)/%s", username, password, network, dbhost, dbport, database)
	db, err := sql.Open("mysql", src)
	if err != nil {
		log.Error("dbConn failed", zap.Error(err), zap.Any("src", src))
		return nil, err
	}
	log.Debug("dbConn success", zap.Error(err), zap.Any("src", src))

	return db, nil
}

func dropDatabase(db *sql.DB, dbname string) error {
	createDbSql := fmt.Sprintf("drop database if exists %s", dbname)
	_, err := db.Exec(createDbSql)
	if err != nil {
		log.Error("error", zap.Error(err))
		return err
	}
	log.Info("success", zap.String("db", dbname))

	return nil
}

func createDatabase(db *sql.DB, dbname string) error {
	createDbSql := fmt.Sprintf("create database if not exists %s", dbname)
	_, err := db.Exec(createDbSql)
	if err != nil {
		log.Error("error", zap.Error(err))
		return err
	}
	log.Debug("success", zap.String("db", dbname))

	return nil
}

func dropTable(db *sql.DB, dbname, tabname string) error {
	dropSql := fmt.Sprintf("drop table if exists %s.%s", dbname, tabname)
	_, err := db.Exec(dropSql)
	if err != nil {
		log.Error("drop error", zap.Error(err), zap.String("db", dbname), zap.String("tab", tabname))
		return err
	}
	log.Debug("drop success", zap.String("db", dbname), zap.String("tab", tabname))

	return nil
}

func createTable(db *sql.DB, dbname, tabname string) error {
	createSql := fmt.Sprintf("create table if not exists %s.%s (id int primary key, c varchar(10))", dbname, tabname)
	_, err := db.Exec(createSql)
	if err != nil {
		log.Error("createTable error", zap.Error(err), zap.String("db", dbname), zap.String("tab", tabname))
		return err
	}
	log.Debug("createTable success", zap.String("db", dbname), zap.String("tab", tabname))

	return nil
}

func insert(db *sql.DB, dbname, tabname string, start, len int, content string) error {
	insertSql := fmt.Sprintf("insert into %s.%s (id, c) values(?, ?);", dbname, tabname)
	stmt, err := db.Prepare(insertSql)
	if err != nil {
		log.Error("insert prepare error", zap.Error(err))
		return err
	}

	for i := start; i < start+len; i++ {
		_, err = stmt.Exec(i, content)
		if err != nil {
			log.Error("insert exec error", zap.Error(err))
			return err
		}
		log.Debug("insert", zap.Any("i", i), zap.Any("content", content),
			zap.Any("db", dbname), zap.Any("tab", tabname))
	}

	return nil
}

func Delete(db *sql.DB, dbname, tabname string, start, end int) error {
	deleteSql := fmt.Sprintf("delete from %s.%s where id >= %d and id <= %d", dbname, tabname, start, end)
	stmt, err := db.Prepare(deleteSql)
	if err != nil {
		log.Error("delete prepare error", zap.Error(err))
		return err
	}

	_, err = stmt.Exec()
	if err != nil {
		log.Error("delete exec error", zap.Error(err))
		return err
	}
	log.Debug("delete", zap.Any("db", dbname), zap.Any("tab", tabname))

	return nil
}

func initDb(db *sql.DB, dbname string) error {
	err := dropDatabase(db, dbname)
	if err != nil {
		return err
	}
	err = createDatabase(db, dbname)
	if err != nil {
		return err
	}
	return nil
}

func initTable(db *sql.DB, dbname, tabname string) error {
	err := dropTable(db, dbname, tabname)
	if err != nil {
		return err
	}
	err = createTable(db, dbname, tabname)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	var err error

	log.Info("start")
	cnDb, err := dbConn(Username, Password, Network, CnDB, DbHookPort, "")
	if err != nil {
		log.Error("init cn conn failed", err.Error())
		return
	} else {
		log.Info("init cn conn success")
	}
	usDb, err := dbConn(Username, Password, Network, UsDB, DbHookPort, "")
	if err != nil {
		log.Error("init us conn failed", err.Error())
		return
	} else {
		log.Info("init us conn success")
	}
	ggDb, err := dbConn(Username, Password, Network, GlobalDb, DbHookPort, "")
	if err != nil {
		log.Error("init gg conn failed", err.Error())
		return
	} else {
		log.Info("init gg conn success")
	}

	if true {
		for _, db := range cfg.Databases {
			dbname := db.Name
			log.Info("start db", dbname)

			// init db
			if db.Init {
				err = initDb(cnDb, dbname)
				if err != nil {
					log.Error("init cn db failed")
					return
				} else {
					log.Info("init cn db success")
				}
				err = initDb(usDb, dbname)
				if err != nil {
					log.Error("init us db failed")
					return
				} else {
					log.Info("init us db success")
				}
				err = initDb(ggDb, dbname)
				if err != nil {
					log.Error("init gg db failed")
					return
				} else {
					log.Info("init gg db success")
				}
			}

			// process table
			wg1 := sync.WaitGroup{}
			for _, tab := range db.Tables {
				tabname := tab.Name
				cnt := tab.Cnt
				init := tab.Init
				log.Info("start tab", dbname, tabname)

				wg1.Add(1)
				go func(tab *Table) {
					defer wg1.Done()

					log.Info("start tab init", dbname, tabname, init)
					if init { // init table
						err = initTable(cnDb, dbname, tabname)
						if err != nil {
							log.Error("init cn tab failed", dbname, tabname, err.Error())
							return
						} else {
							log.Info("init cn tab success", dbname, tabname)
						}
						err = initTable(usDb, dbname, tabname)
						if err != nil {
							log.Error("init us tab failed", dbname, tabname, err.Error())
							return
						} else {
							log.Info("init us tab success", dbname, tabname)
						}
						err = initTable(ggDb, dbname, tabname)
						if err != nil {
							log.Error("init gg tab failed", dbname, tabname, err.Error())
							return
						} else {
							log.Info("init gg tab success", dbname, tabname)
						}
					}

					log.Info("start tab insert", dbname, tabname, cnt)
					{ // insert data
						wg2 := sync.WaitGroup{}
						wg2.Add(1)
						go func() {
							defer wg2.Done()

							log.Info("start insert cn", dbname, tabname)
							err = insert(cnDb, dbname, tabname, CNStart, cnt, "cn")
							if err != nil {
								log.Error("insert cn error", err.Error())
								return
							}
						}()
						go func() {
							wg2.Add(1)
							defer wg2.Done()

							log.Info("start insert us", dbname, tabname)
							err = insert(usDb, dbname, tabname, USStart, cnt, "us")
							if err != nil {
								log.Error("insert us error", err.Error())
								return
							}
						}()
						wg2.Wait()
					}
				}(tab)

				log.Info("insert data success", dbname, tabname)
			}
			wg1.Wait()

			log.Info("finished db", dbname)
		}
	}

	if false {
		for _, db := range cfg.Databases {
			dbname := db.Name
			for _, tab := range db.Tables {
				tabname := tab.Name
				err = Delete(cnDb, dbname, tabname, CNStart, CNStart + 100)
				if err != nil {
					return
				}
				err = Delete(usDb, dbname, tabname, USStart, USStart + 100)
				if err != nil {
					return
				}
			}
		}
	}
}
