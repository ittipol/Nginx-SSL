package database

import (
	"context"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type sqlLogger struct {
	logger.Interface
}

func (sqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, _ := fc()

	fmt.Printf("%v \n===============================\n", sql)
}

func GetDbConnection(username string, password string, host string, port int, db string, dryRun bool) *gorm.DB {
	// root:1234@tcp(127.0.0.1:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, db)

	// dial := postgres.Open(dsn)
	dial := mysql.Open(dsn)

	conn, err := gorm.Open(dial, &gorm.Config{
		Logger: &sqlLogger{},
		DryRun: dryRun,
	})

	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Database Connected")

	return conn
}
