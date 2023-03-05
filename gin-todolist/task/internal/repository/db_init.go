package repository

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func InitDB() {
	host := viper.GetString("mysql.host")
	port := viper.GetString("mysql.port")
	database := viper.GetString("mysql.database")
	username := viper.GetString("mysql.username")
	password := viper.GetString("mysql.password")
	charset := viper.GetString("mysql.charset")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username, password, host, port, database, charset)
	if err := Database(dsn); err != nil {
		panic(err)
	}
}

func Database(dsn string) error {
	var ormLogger logger.Interface
	if gin.Mode() == "debug" {
		ormLogger = logger.Default.LogMode(logger.Info)
	} else {
		ormLogger = logger.Default
	}
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,  // 禁用datetime的精度, mysql5.6之前不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除再新建的方式, 因为mysql5.7之前不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 change 重命名 column, 因为mysql8之前不支持重命名 column
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}), &gorm.Config{
		Logger: ormLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return err
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(20)                  // 设置连接池, 空闲
	sqlDB.SetMaxOpenConns(100)                 // 设置最大连接数
	sqlDB.SetConnMaxLifetime(time.Second * 30) // 连接最长保持时间
	DB = db
	migration() // 迁移: 将Go定义的数据结构映射到数据库的表, 就不用自己慢慢 create table 啦 ~
	return nil
}
