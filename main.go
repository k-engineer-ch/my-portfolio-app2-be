package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {

	router := gin.Default()
	v1 := router.Group("/v1")
	v1.GET("/testGet", getting)
	// router.POST("/somePost", posting)
	// router.PUT("/somePut", putting)
	// router.DELETE("/someDelete", deleting)
	// router.PATCH("/somePatch", patching)
	// router.HEAD("/someHead", head)
	// router.OPTIONS("/someOptions", options)

	router.Run(":8080")
}

func getting(c *gin.Context) {
	c.JSONP(http.StatusOK, gin.H{
		"message": "ok",
		"data":    "test",
	})
	db()
}

func db() {
	db, err := sql.Open("postgres", "postgres://dev_user:system001@localhost:5432/dev1_db?sslmode=disable")
	// db, err := sql.Open("postgres", "postgres://dev_user:system001@hostname:port/dev1_db user=dev_user password=system001 dbname=dev1_db sslmode=disable")
	checkErr(err)

	//データの挿入
	stmt, err := db.Prepare("INSERT INTO userinfo(username,departname,created) VALUES($1,$2,$3) RETURNING uid")
	checkErr(err)

	res, err := stmt.Exec("astaxie", "研究開発部門", "2012-12-09")
	checkErr(err)
	fmt.Println(res)

	// //pgはこの関数をサポートしていません。MySQLのインクリメンタルなIDのようなものが無いためです。
	// id, err := res.LastInsertId()
	// checkErr(err)

	// fmt.Println(id)

	// //データの更新
	// stmt, err = db.Prepare("update userinfo set username=$1 where uid=$2")
	// checkErr(err)

	// res, err = stmt.Exec("astaxieupdate", 1)
	// checkErr(err)

	// affect, err := res.RowsAffected()
	// checkErr(err)

	// fmt.Println(affect)

	// //データの検索
	// rows, err := db.Query("SELECT * FROM userinfo")
	// checkErr(err)

	// for rows.Next() {
	// 	var uid int
	// 	var username string
	// 	var department string
	// 	var created string
	// 	err = rows.Scan(&uid, &username, &department, &created)
	// 	checkErr(err)
	// 	fmt.Println(uid)
	// 	fmt.Println(username)
	// 	fmt.Println(department)
	// 	fmt.Println(created)
	// }

	// //データの削除
	// stmt, err = db.Prepare("delete from userinfo where uid=$1")
	// checkErr(err)

	// res, err = stmt.Exec(1)
	// checkErr(err)

	// affect, err = res.RowsAffected()
	// checkErr(err)

	// fmt.Println(affect)

	db.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
