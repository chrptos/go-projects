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

	// Httpパッケージのルーターではなく、Gorillaパッケージのルーターを使用する。
	// 指定されたメソッド以外を受け取った時は、ステータスコード405を返す。
	r := mux.NewRouter()
	r.HandleFunc("/hello", handlers.HelloHandler).Methods(http.MethodGet)
	r.HandleFunc("/article", handlers.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", handlers.ArticleListHandler).Methods(http.MethodGet)
	// r.HandleFunc("/article/1", handlers.ArticleDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", handlers.ArticleDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", handlers.PostNiceHandler).Methods(http.MethodPost)
	r.HandleFunc("/comment", handlers.PostCommentHandler).Methods(http.MethodPost)

	log.Println("server start at port 8080")
	// ListenAndServeはサーバー内で使用されるルーターを指定する
	// Httpパッケージから、Gorillaパッケージのルーターに変えたことでnilからrへ変更する
	log.Fatal(http.ListenAndServe(":8080", r))
}
