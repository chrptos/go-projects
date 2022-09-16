package main

import (
	"log"
	"net/http"

	"github.com/chrptos/myapi/handlers"
	"github.com/gorilla/mux"
)

func main() {
	// Go の HTTP サーバーがデフォルトで持っているルータに、ルーティングルールを登録する
	// http.HandleFunc("/hello", handlers.HelloHandler)
	// http.HandleFunc("/article", handlers.PostArticleHandler)
	// http.HandleFunc("/article/list", handlers.ArticleListHandler)
	// http.HandleFunc("/article/1", handlers.ArticleDetailHandler)
	// http.HandleFunc("/article/nice", handlers.PostNiceHandler)
	// http.HandleFunc("/comment", handlers.PostCommentHandler)

	// Httpパッケージのルーターではなく、Gorillaパッケージのルーターを使用する
	r := mux.NewRouter()
	r.HandleFunc("/hello", handlers.HelloHandler)
	r.HandleFunc("/article", handlers.PostArticleHandler)
	r.HandleFunc("/article/list", handlers.ArticleListHandler)
	r.HandleFunc("/article/1", handlers.ArticleDetailHandler)
	r.HandleFunc("/article/nice", handlers.PostNiceHandler)
	r.HandleFunc("/comment", handlers.PostCommentHandler)

	log.Println("server start at port 8080")
	// ListenAndServeはサーバー内で使用されるルーターを指定する
	// Httpパッケージから、Gorillaパッケージのルーターに変えたことでnilからrへ変更
	log.Fatal(http.ListenAndServe(":8080", r))
}
