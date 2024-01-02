package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"sync/atomic"
	"time"
)

func main() {

	dsn := "root:password@tcp(localhost:3306)/yl?" +
		"interpolateParams=true&charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:         logger.Default.LogMode(logger.Silent),
		TranslateError: true,
	})
	if err != nil {
		panic(" gorm.Open err: " + err.Error())
	}

	cfdb, err := db.DB()
	if err != nil {
		panic(" db.DB err: " + err.Error())
	}
	//
	//cfdb.SetConnMaxLifetime(-1)
	//cfdb.SetMaxIdleConns(100)
	//cfdb.SetMaxOpenConns(2000)
	//cfdb.SetConnMaxIdleTime(time.Minute)

	var a atomic.Int32

	for i := 1; i <= 300; i++ {
		go func() {
			tx := db.Begin()

			err = tx.Table("d_user").Where("id=5").Updates(map[string]interface{}{"name": fmt.Sprintf("yc-%d", i)}).Error
			if err != nil {
				fmt.Println(" == ", i, " tx.Updates()  err: ", err.Error())
				a.Add(1)
				return
			}
			<-time.After(300 * time.Second)

			tx.Commit()
		}()
	}

	<-time.After(3 * time.Second)
	fmt.Println(a.Load())
	fmt.Println(fmt.Sprintf("%+v", cfdb.Stats()))
	<-time.After(300 * time.Second)

}
