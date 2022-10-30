package main

import (
	"database/sql"
	"fmt"

	"github.com/chrptos/mysql/models"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)

	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	const sqlStr = `
		select title, contents, username, nice
		from articles;
	`

	rows, err := db.Query(sqlStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	articleArray := make([]models.Article, 0)
	for rows.Next() {
		var article models.Article
		// 引数に「データ読み出し結果を格納したい変数のポインタ」を指定することで、rowsの中に格納されている取得レコード内容を読み出すことができる。
		err := rows.Scan(&article.Title, &article.Contents, &article.UserName, &article.NiceNum)

		if err != nil {
			fmt.Println(err)
		} else {
			articleArray = append(articleArray, article)
		}
	}

	fmt.Printf("%+v\n", articleArray)
}
