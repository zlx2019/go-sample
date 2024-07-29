// @Title database.go
// @Description 数据库
// @Author Zero - 2024/7/27 23:27:22

package database

import (
	"fmt"
	"go-sample/configs"
	"go-sample/internal/global"
	"go-sample/internal/setup/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

// Setup 初始化Gorm
func Setup() {
	dbConf := global.Conf.DB
	gdb, err := gorm.Open(mysql.New(withMysqlConfig(dbConf)), withGormConfig(dbConf))
	if err != nil {
		logs.Logger.FatalSf("failed to connect database: %s", err.Error())
	}
	// 配置连接池
	if db, err := gdb.DB(); err == nil {
		db.SetMaxOpenConns(dbConf.Pool.MaxOpenConn)
		db.SetMaxIdleConns(dbConf.Pool.MaxIdleConn)
		db.SetConnMaxLifetime(dbConf.Pool.MaxLifeTime)
	}
	// 自动创建表
	if dbConf.CreateTables {
		if err = autoCreateTables(gdb); err != nil {
			logs.Logger.FatalSf("failed to auto create tables: %s", err.Error())
		}
	}
	global.Dc = gdb
	logs.Logger.Info("connected database success.")
}

// 根据实体结构，自动创建表结构.
func autoCreateTables(db *gorm.DB) error {
	return db.AutoMigrate()
}

func withGormConfig(db configs.DataBase) *gorm.Config {
	var _logger logger.Interface
	if db.Debug {
		// 输出日志信息
		//time.ParseDuration(db.SlowSql)
		logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
			SlowThreshold:             db.SlowSql, // 慢sql阈值
			Colorful:                  false,      // 关闭彩色日志
			IgnoreRecordNotFoundError: true,       // 忽略查询不存在错误 `ErrRecordNotFound `
			ParameterizedQueries:      false,      // 关闭sql日志参数显示
			LogLevel:                  logger.Info,
		})
	} else {
		// 不输出SQL日志
		_logger = logger.Default.LogMode(logger.Silent)
	}
	return &gorm.Config{
		Logger:                                   _logger,
		DisableForeignKeyConstraintWhenMigrating: true,  // 关闭自动创建外键约束
		SkipDefaultTransaction:                   false, // 关闭增删改操作默认开启事务，提升插入性能
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
	}
}

// Mysql Config
func withMysqlConfig(db configs.DataBase) mysql.Config {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=10s",
		db.Username, db.Password, db.Host, db.Port, db.DbName)
	return mysql.Config{
		DSN: dsn,
	}
}