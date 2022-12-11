package database

import (
	"Vadek/configuration"
	vLog "Vadek/loger"
	"database/sql"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
	"path/filepath"
	"time"
)

var Db *gorm.DB

func creatVadekDb(conf *configuration.Config) {
	var creatDb *sql.DB
	dataScriptFile := filepath.Join(conf.Vadek.WorkDir, "test.sql")
	dsf, err := os.ReadFile(dataScriptFile)
	creatDb, _ = sql.Open("sqlite3", conf.SQLite3.File)
	creatDb.SetMaxIdleConns(256)
	creatDb.Ping()
	_, err = creatDb.Exec(string(dsf))
	if err != nil {
		fmt.Println(err)
	}
}
func InitSqlite(conf *configuration.Config, gormLogger logger.Interface) (*gorm.DB, error) {
	sqliteConfig := conf.SQLite3
	if sqliteConfig == nil {
		//TODO error interface
		vLog.Warn("没有SQLITE3 文件")
		return nil, nil
	}
	vLog.Info("Try to open SQLite3 db", zap.String("path", sqliteConfig.File))
	db, err := gorm.Open(sqlite.Open(sqliteConfig.File), &gorm.Config{
		Logger:                   gormLogger,
		PrepareStmt:              true,
		SkipDefaultTransaction:   true,
		DisableNestedTransaction: true,
	})
	return db, err

}
func NewGormDB(conf *configuration.Config, gormLogger logger.Interface) *gorm.DB {
	var err error
	if conf.SQLite3 != nil && conf.SQLite3.Enable {
		Db, err = InitSqlite(conf, gormLogger)
		if err != nil {
			vLog.Fatal("open SQLite3 error", zap.Error(err))
		}
	}
	if Db == nil {
		vLog.Fatal("no available database")
	}
	vLog.Info("connect database success")
	sqlDb, err := Db.DB()
	if err != nil {
		vLog.Fatal("get database connection error")
	}
	sqlDb.SetMaxIdleConns(200)
	sqlDb.SetMaxOpenConns(300)
	sqlDb.SetConnMaxLifetime(time.Hour)
	//TODO
	SetDefault(Db)
	//dbMigrate()
	return Db
}

//func dbMigrate() {
//	db := Db.Session(&gorm.Session{
//		Logger: Db.Logger.LogMode(logger.Warn),
//	})
//	//err :=db.AutoMigrate()
//}
