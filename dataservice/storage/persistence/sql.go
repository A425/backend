package persistence

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var client *sqlx.DB

// DB ...
func DB() *sqlx.DB {
	return client
}

// Init ...
func Init(dbconf DBConfiguration) {
	var err error
	dbURL := fmt.Sprintf("%s:%d", dbconf.Host, dbconf.Port)
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=true",
		dbconf.User, dbconf.Password, dbURL, dbconf.DBName,
		dbconf.Encoding)

	client, err = sqlx.Connect("mysql", dsn)
	client.SetMaxIdleConns(dbconf.PoolSize)
	client.SetMaxOpenConns(dbconf.PoolSize)
	if err != nil {
		panic(fmt.Sprintf("connect mysql failed, err:%v, dsn:%s", err, dsn))
	}
}

// Release ...
func Release() error {
	return client.Close()
}
