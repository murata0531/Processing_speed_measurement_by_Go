
package main

import (
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "github.com/jinzhu/gorm"
	"strconv"
	"fmt"
    "time"
)

type User struct {
	gorm.Model
    UserName string `json:"username"`
}

// DBの初期化
func dbInit() {
    db := gormConnect()
    // コネクション解放解放
    defer db.Close()
    db.AutoMigrate(&User{})
}

// データインサート処理
func dbInsert() {
    db := gormConnect()

    defer db.Close()
    // Insert処理
    for i := 0; i < 100; i++ {
		var uname = "user" + strconv.Itoa(i)
		db.Create(&User{UserName: uname})
	}
}

func main() {
	
	dbInit()
    
	start := time.Now()

	dbInsert()

	end := time.Now()
	fmt.Printf("end:%.17f[sec]\n",(end.Sub(start)).Seconds())
    
}

func DBMigrate(db *gorm.DB) *gorm.DB {
    db.AutoMigrate(&User{})
    return db
}


func gormConnect() *gorm.DB {
    DBMS := "mysql"
    USER := "test"
    PASS := "test"
    PROTOCOL := "tcp(mysql:3306)"
    DBNAME := "test"
    CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"
    db, err := gorm.Open(DBMS, CONNECT)
    if err != nil {
        panic(err.Error())
    }
    return db
}
