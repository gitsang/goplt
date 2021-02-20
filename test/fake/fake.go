package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pingcap/log"
	"go.uber.org/zap"
	"sync"
)

const (
	Username   = "root"
	Password   = ""
	CnDB       = "10.120.24.130"
	UsDB       = "10.120.26.60"
	EuDb       = "10.120.25.163"
	DbHookPort = "9436"
	Network    = "tcp"

	CNStart = 100000000
	USStart = 200000000
	EuStart = 300000000
)

var cfg = Config{
	Databases: []*Database{
		{
			Name: "sang",
			Init: true,
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

	log.Info("start...")
	cnDb, err := dbConn(Username, Password, Network, CnDB, DbHookPort, "")
	if err != nil {
		log.Error("init connection failed", zap.Any("region", "cn"), zap.Error(err))
		return
	} else {
		log.Info("init connection success", zap.Any("region", "cn"))
	}
	usDb, err := dbConn(Username, Password, Network, UsDB, DbHookPort, "")
	if err != nil {
		log.Error("init connection failed", zap.Any("region", "us"), zap.Error(err))
		return
	} else {
		log.Info("init connection success", zap.Any("region", "us"))
	}
	euDb, err := dbConn(Username, Password, Network, EuDb, DbHookPort, "")
	if err != nil {
		log.Error("init connection failed", zap.Any("region", "eu"), zap.Error(err))
		return
	} else {
		log.Info("init connection success", zap.Any("region", "eu"))
	}

	if true {
		for _, db := range cfg.Databases {
			dbname := db.Name
			log.Info("start database...", zap.Any("db", dbname))

			// init db
			if db.Init {
				err = initDb(cnDb, dbname)
				if err != nil {
					log.Error("init connection failed", zap.Any("region", "cn"), zap.Any("db", dbname), zap.Error(err))
					return
				} else {
					log.Info("init connection failed", zap.Any("region", "cn"), zap.Any("db", dbname))
				}
				err = initDb(usDb, dbname)
				if err != nil {
					log.Error("init connection failed", zap.Any("region", "us"), zap.Any("db", dbname), zap.Error(err))
					return
				} else {
					log.Info("init connection failed", zap.Any("region", "us"), zap.Any("db", dbname))
				}
				err = initDb(euDb, dbname)
				if err != nil {
					log.Error("init connection failed", zap.Any("region", "eu"), zap.Any("db", dbname), zap.Error(err))
					return
				} else {
					log.Info("init connection failed", zap.Any("region", "eu"), zap.Any("db", dbname))
				}
			}

			// process table
			wg1 := sync.WaitGroup{}
			for _, tab := range db.Tables {
				tabname := tab.Name
				cnt := tab.Cnt
				init := tab.Init
				log.Info("start table ...", zap.Any("table", dbname + "/" + tabname))

				wg1.Add(1)
				go func(tab *Table) {
					defer wg1.Done()

					if init { // init table
						err = initTable(cnDb, dbname, tabname)
						if err != nil {
							log.Error("init table failed", zap.Any("region", "cn"), zap.Any("table", dbname + "/" + tabname), zap.Error(err))
							return
						} else {
							log.Info("init table success", zap.Any("region", "cn"), zap.Any("table", dbname + "/" + tabname))
						}
						err = initTable(usDb, dbname, tabname)
						if err != nil {
							log.Error("init table failed", zap.Any("region", "us"), zap.Any("table", dbname + "/" + tabname), zap.Error(err))
							return
						} else {
							log.Info("init table success", zap.Any("region", "us"), zap.Any("table", dbname + "/" + tabname))
						}
						err = initTable(euDb, dbname, tabname)
						if err != nil {
							log.Error("init table failed", zap.Any("region", "eu"), zap.Any("table", dbname + "/" + tabname), zap.Error(err))
							return
						} else {
							log.Info("init table success", zap.Any("region", "eu"), zap.Any("table", dbname + "/" + tabname))
						}
					}

					log.Info("start table insert...", zap.Any("table", dbname + "/" + tabname), zap.Any("cnt", cnt))
					{ // insert data
						wg2 := sync.WaitGroup{}
						wg2.Add(1)
						go func() {
							defer wg2.Done()
							err = insert(cnDb, dbname, tabname, CNStart, cnt, "cn")
							if err != nil {
								log.Error("insert failed", zap.Any("region", "cn"), zap.Any("table", dbname + "/" + tabname), zap.Error(err))
								return
							}
							log.Info("insert success", zap.Any("region", "cn"), zap.Any("table", dbname + "/" + tabname))
						}()
						wg2.Add(1)
						go func() {
							defer wg2.Done()
							err = insert(usDb, dbname, tabname, USStart, cnt, "us")
							if err != nil {
								log.Error("insert failed", zap.Any("region", "us"), zap.Any("table", dbname + "/" + tabname), zap.Error(err))
								return
							}
							log.Info("insert success", zap.Any("region", "us"), zap.Any("table", dbname + "/" + tabname))
						}()
						wg2.Add(1)
						go func() {
							defer wg2.Done()
							err = insert(euDb, dbname, tabname, EuStart, cnt, "eu")
							if err != nil {
								log.Error("insert failed", zap.Any("region", "eu"), zap.Any("table", dbname + "/" + tabname), zap.Error(err))
								return
							}
							log.Info("insert success", zap.Any("region", "eu"), zap.Any("table", dbname + "/" + tabname))
						}()
						wg2.Wait()
					}
				}(tab)
			}
			wg1.Wait()
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
				err = Delete(euDb, dbname, tabname, EuStart, EuStart + 100)
				if err != nil {
					return
				}
			}
		}
	}
}
