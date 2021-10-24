
package main

import (
    "github.com/gin-gonic/gin"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "github.com/jinzhu/gorm"
)

func main() {
    DBMigrate(DBConnect())
    r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "成功",
        })
    })
    r.Run(":8080")
}

type User struct {
    gorm.Model
    NickName string `json:"nickName"`
}

func DBMigrate(db *gorm.DB) *gorm.DB {
    db.AutoMigrate(&User{})
    return db
}


func DBConnect() *gorm.DB {
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
