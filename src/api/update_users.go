
package main

import (
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "github.com/jinzhu/gorm"
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

// データ取得処理
func dbUpdate() {
    db := gormConnect()

    defer db.Close()

	users := []User{}
    db.Find(&users)

	// 取得した100行を一行ずつループして更新
	for _, user := range users {
		db.Model(&User{}).Where("id = ?", user.ID).Update("user_name", "update")
		fmt.Println(user.UserName)
	}
}

func main() {
	
	dbInit()
    
	start := time.Now()

	dbUpdate()

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
