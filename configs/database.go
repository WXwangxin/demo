package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"main/model"
	"time"
)

var Database *gorm.DB

func init() {
	var datetimePrecision = 2
	var _ error
	Database, _ = gorm.Open(mysql.New(mysql.Config{
		DSN:                       "root:12345678@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local", // data source name, refer https://github.com/go-sql-driver/mysql#dsn-data-source-name
		DefaultStringSize:         255,                                                                               // add default size for string fields, by default, will use db type `longtext` for fields without size, not a primary key, no index defined and don't have default values
		DisableDatetimePrecision:  true,                                                                              // disable datetime precision support, which not supported before MySQL 5.6
		DefaultDatetimePrecision:  &datetimePrecision,                                                                // default datetime precision
		DontSupportRenameIndex:    true,                                                                              // drop & create index when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,                                                                              // use change when rename column, rename rename not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false,                                                                             // smart configure based on used version
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "t_",  // users->t_users table name prefix, table for `User` would be `t_users`
			SingularTable: false, // users->user use singular table name, table for `User` would be `user` with this option enabled
			NoLowerCase:   false, // users->Users skip the snake_casing of names
			//NameReplacer:  strings.NewReplacer("Name", "username"), // name->username use name replacer to change struct/field name before convert it to db name
		},
	})

	/*	dsn := "root:12345678@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
		Database, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})*/
	// 获取通用数据库对象 sql.DB ，然后使用其提供的功能
	sqlDB, _ := Database.DB()
	//  用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(2)
	//  设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(20)
	//  设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	Database.AutoMigrate(&model.User{})
}
