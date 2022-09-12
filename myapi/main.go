package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	// httpリクエストを受けて、それに対するhttpレスポンス内容をコネクションに書き込む
	// helloHandlerという名前で、ハンドラを定義
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, World!!\n")
	}
	// 定義したhelloHandlerを使うように登録する
	http.HandleFunc("/", helloHandler)
	log.Println("server start at port 8080")

	// サーバーの起動
	// プログラムが終了するような重大なエラーが返ってきた場合に（ListenAndServeの返り値がエラー）、該当エラーを出力し処理を終了する（log.Fatal）
	log.Fatal(http.ListenAndServe(":8080", nil))
}
