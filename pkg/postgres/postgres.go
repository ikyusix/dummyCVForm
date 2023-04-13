package postgres

import (
	"database/sql"
	"dummyCVForm/pkg/logger"
	"dummyCVForm/utils/config"
	_ "github.com/lib/pq"
)

var (
	db *dbPgsql
)

type dbPgsql struct {
	dbPq *sql.DB
}

func InitDBConnection() error {
	err := InitConnectionDB()
	if err != nil {
		return err
	}
	return nil
}

func InitConnectionDB() error {
	logger.Log.Println("start init postgres")
	db = new(dbPgsql)
	conn, errdb := sql.Open("postgres", config.MyConfig.Db)
	if errdb != nil {
		return errdb
	}
	if err := conn.Ping(); err != nil {
		return err
	}

	db.dbPq = conn

	return nil
}

func GetConnectionDB() (*sql.DB, error) {
	return db.GetConnection()
}

func (dpq *dbPgsql) GetConnection() (*sql.DB, error) {
	return dpq.dbPq, nil
}

func CloseDBConnection() {
	closeConnection()
}

func closeConnection() {
	logger.Log.Println("Closing DB connection...")
	db.dbPq.Close()
}
