package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Expense struct {
	ID        int    `json:"id"`
	Date      string `json:"date"`
	Amount    string `json:"amount"`
	Category  string `json:"category"`
	Memo      string `json:"memo"`
	CreatedAt string `json:"created_at"`
}

func main() {

	router := gin.Default()

	// CORSミドルウェアを使用して、CORSポリシーを設定
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}
		c.Next()
	})

	v1 := router.Group("/v1")
	// v1.GET("/testGet", getting)
	v1.GET("/getAllExpenses", getAllExpenses)
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

func getAllExpenses(c *gin.Context) {
	// 接続
	db, err := sql.Open("postgres", "postgres://dev_user:system001@localhost:5432/dev1_db?sslmode=disable")
	checkErr(err)

	//データの検索
	rows, err := db.Query("SELECT * FROM expenses")
	checkErr(err)

	var expenses []Expense

	for rows.Next() {
		var expense Expense
		err := rows.Scan(&expense.ID, &expense.Date, &expense.Amount, &expense.Category, &expense.Memo, &expense.CreatedAt)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Internal Server Error",
			})
			return
		}
		expenses = append(expenses, expense)
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"data":    expenses,
	})

	// 接続切る
	db.Close()
}

func db() {
	db, err := sql.Open("postgres", "postgres://dev_user:system001@localhost:5432/dev1_db?sslmode=disable")
	// db, err := sql.Open("postgres", "postgres://dev_user:system001@hostname:port/dev1_db user=dev_user password=system001 dbname=dev1_db sslmode=disable")
	checkErr(err)

	//データの挿入
	// stmt, err := db.Prepare("INSERT INTO userinfo(username,departname,created) VALUES($1,$2,$3) RETURNING uid")
	stmt, err := db.Prepare("INSERT INTO expenses(date,amount,category,memo) VALUES($1,$2,$3,$4) RETURNING id")
	checkErr(err)

	// res, err := stmt.Exec("astaxie", "研究開発部門", "2012-12-09")
	res, err := stmt.Exec("2012-12-09", 1000, "food", "テスト登録")
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

	//データの検索
	// rows, err := db.Query("SELECT * FROM userinfo")
	// rows, err := db.Query("SELECT * FROM expenses")
	// クエリの実行
	var id int = 2
	var date string
	var amount float64
	var category string
	var memo string
	err = db.QueryRow("SELECT id, date, amount, category, memo FROM expenses WHERE id = $1", id).Scan(&id, &date, &amount, &category, &memo)
	checkErr(err)

	// 結果の表示
	fmt.Printf("ID: %d\n", id)
	fmt.Printf("Date: %s\n", date)
	fmt.Printf("Amount: %f\n", amount)
	fmt.Printf("Category: %s\n", category)
	fmt.Printf("Memo: %s\n", memo)

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
	// for rows.Next() {
	// 	var id int
	// 	var date string
	// 	var amount string
	// 	var category string
	// 	var memo string
	// 	var created_at string
	// 	err = rows.Scan(&id, &date, &amount, &category, &memo, &created_at)
	// 	checkErr(err)
	// 	fmt.Println(id)
	// 	fmt.Println(date)
	// 	fmt.Println(amount)
	// 	fmt.Println(category)
	// 	fmt.Println(memo)
	// 	fmt.Println(created_at)
	// }

	//データの削除
	stmt, err = db.Prepare("delete from expenses where id=$1")
	checkErr(err)

	res, err = stmt.Exec(2)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	db.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
