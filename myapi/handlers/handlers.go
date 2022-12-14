package handlers

import (
	"fmt"
	"io"
	"net/http"
)

// hello のハンドラ（GET）
func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}

// article のハンドラ（POST: ブログ記事の投稿）
func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Article...\n")
}

// article/list のハンドラ（GET: ブログ一覧を取得）
func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Article List\n")
}

// article/1 のハンドラ（GET: 指定したIDのブログ記事を取得する）
func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleID := 1
	resString := fmt.Sprintf("Article No.%d\n", articleID)
	io.WriteString(w, resString)
}

// article/nice のハンドラ（POST: 記事にいいねをつける）
func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Nice...\n")
}

// comment のハンドラ（POST: 記事にコメントをする）
func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Posting Comment...\n")
}
