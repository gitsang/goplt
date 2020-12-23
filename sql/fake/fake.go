package main

import (
	"database/sql"
	"fmt"
	"github.com/gitsang/golog"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap"
)

const (
	Username   = "root"
	Password   = ""
	CnDB       = "tidb-cn.l7i.top"
	UsDB       = "tidb-us.l7i.top"
	GlobalDb   = "tidb-global.l7i.top"
	DbHookPort = "9436"
	Network    = "tcp"
)

func dbConn(username, password, network, dbhost, dbport, database string) (*sql.DB, error) {
	src := fmt.Sprintf("%s:%s@%s(%s:%s)/%s", username, password, network, dbhost, dbport, database)
	db, err := sql.Open("mysql", src)
	if err != nil {
		log.Error("dbConn failed", zap.Error(err), zap.Any("src", src))
		return nil, err
	}
	log.Info("dbConn success", zap.Error(err), zap.Any("src", src))

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
	log.Info("success", zap.String("db", dbname))

	return nil
}

func dropTable(db *sql.DB, dbname, tabname string) error {
	dropSql := fmt.Sprintf("drop table if exists %s.%s", dbname, tabname)
	_, err := db.Exec(dropSql)
	if err != nil {
		log.Error("drop error", zap.Error(err), zap.String("db", dbname), zap.String("tab", tabname))
		return err
	}
	log.Info("drop success", zap.String("db", dbname), zap.String("tab", tabname))

	return nil
}

func createTable(db *sql.DB, dbname, tabname string) error {
	createSql := fmt.Sprintf("create table if not exists %s.%s (id int primary key, c varchar(10))", dbname, tabname)
	_, err := db.Exec(createSql)
	if err != nil {
		log.Error("createTable error", zap.Error(err), zap.String("db", dbname), zap.String("tab", tabname))
		return err
	}
	log.Info("createTable success", zap.String("db", dbname), zap.String("tab", tabname))

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
		log.Info("insert", zap.Any("i", i), zap.Any("content", content), zap.Any("db", dbname), zap.Any("tab", tabname))
	}

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

	cnDb, err := dbConn(Username, Password, Network, CnDB, DbHookPort, "")
	if err != nil {
		return
	}
	usDb, err := dbConn(Username, Password, Network, UsDB, DbHookPort, "")
	if err != nil {
		return
	}
	ggDb, err := dbConn(Username, Password, Network, GlobalDb, DbHookPort, "")
	if err != nil {
		return
	}

	{ // db: sang
		dbname := "sang"
		err = initDb(cnDb, dbname)
		if err != nil {
			return
		}
		err = initDb(usDb, dbname)
		if err != nil {
			return
		}
		err = initDb(ggDb, dbname)
		if err != nil {
			return
		}
		{ // tab1
			tabname := "tab1"
			{ // init table
				err = initTable(cnDb, dbname, tabname)
				if err != nil {
					return
				}
				err = initTable(usDb, dbname, tabname)
				if err != nil {
					return
				}
				err = initTable(ggDb, dbname, tabname)
				if err != nil {
					return
				}
			}
			{ // insert data
				err = insert(cnDb, dbname, tabname, 10000, 1000, "cn")
				if err != nil {
					return
				}
				err = insert(usDb, dbname, tabname, 20000, 1000, "us")
				if err != nil {
					return
				}
				err = insert(ggDb, dbname, tabname, 30000, 1000, "gg")
				if err != nil {
					return
				}
			}
		}
		{ // tab2
			tabname := "tab2"
			{ // init table
				err = initTable(cnDb, dbname, tabname)
				if err != nil {
					return
				}
				err = initTable(usDb, dbname, tabname)
				if err != nil {
					return
				}
				err = initTable(ggDb, dbname, tabname)
				if err != nil {
					return
				}
			}
			{ // insert data
				err = insert(cnDb, dbname, tabname, 10000, 1000, "cn")
				if err != nil {
					return
				}
				err = insert(usDb, dbname, tabname, 20000, 1000, "us")
				if err != nil {
					return
				}
				err = insert(ggDb, dbname, tabname, 30000, 1000, "gg")
				if err != nil {
					return
				}
			}
		}
	}

	{ // db: sang
		dbname := "sang1"
		err = initDb(cnDb, dbname)
		if err != nil {
			return
		}
		err = initDb(usDb, dbname)
		if err != nil {
			return
		}
		err = initDb(ggDb, dbname)
		if err != nil {
			return
		}
		{ // tab1
			tabname := "tab1"
			{ // init table
				err = initTable(cnDb, dbname, tabname)
				if err != nil {
					return
				}
				err = initTable(usDb, dbname, tabname)
				if err != nil {
					return
				}
				err = initTable(ggDb, dbname, tabname)
				if err != nil {
					return
				}
			}
			{ // insert data
				err = insert(cnDb, dbname, tabname, 10000, 1000, "cn")
				if err != nil {
					return
				}
				err = insert(usDb, dbname, tabname, 20000, 1000, "us")
				if err != nil {
					return
				}
				err = insert(ggDb, dbname, tabname, 30000, 1000, "gg")
				if err != nil {
					return
				}
			}
		}
		{ // tab2
			tabname := "tab2"
			{ // init table
				err = initTable(cnDb, dbname, tabname)
				if err != nil {
					return
				}
				err = initTable(usDb, dbname, tabname)
				if err != nil {
					return
				}
				err = initTable(ggDb, dbname, tabname)
				if err != nil {
					return
				}
			}
			{ // insert data
				err = insert(cnDb, dbname, tabname, 10000, 1000, "cn")
				if err != nil {
					return
				}
				err = insert(usDb, dbname, tabname, 20000, 1000, "us")
				if err != nil {
					return
				}
				err = insert(ggDb, dbname, tabname, 30000, 1000, "gg")
				if err != nil {
					return
				}
			}
		}
	}
}
