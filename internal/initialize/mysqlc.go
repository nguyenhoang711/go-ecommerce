package initialize

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/devpenguin/go-ecommerce/global"
	"go.uber.org/zap"
)


func checkErrorPanicc(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(err)
	}
}

func InitMysqlc() {
	m := global.Config.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", m.Username, m.Password, m.Host, m.Port, m.Dbname)
	db, err := sql.Open("mysql", dsn)
	checkErrorPanicc(err, "InitMysql initialization error")
	global.Logger.Info("Initializing MySQL successfully")
	global.Mdb = db

	SetPool()
}


func SetPool() {
	m := global.Config.Mysql
	sqlDb := global.Mdb
	sqlDb.SetConnMaxIdleTime(time.Duration(m.MaxIdleConns))
	sqlDb.SetMaxOpenConns(m.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Duration(m.ConnMaxLifeTime))
}