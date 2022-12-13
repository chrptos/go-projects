package handlers

import (
	"fmt"
	"io"
	"net/http"
)

// hello のハンドラ（GET）
func HelloHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		io.WriteString(w, "Hello, world!\n")
	} else {
		// Invalid methodというレスポンスを、405番のステータスコードと共に返す
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}

// article のハンドラ（POST: ブログ記事の投稿）
func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		io.WriteString(w, "Posting Article...\n")
	} else {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}

// article/list のハンドラ（GET: ブログ一覧を取得）
func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		io.WriteString(w, "Article List\n")
	} else {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}

// article/1 のハンドラ（GET: 指定したIDのブログ記事を取得する）
func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleID := 1
	resString := fmt.Sprintf("Article No.%d\n", articleID)
	if req.Method == http.MethodGet {
		io.WriteString(w, resString)
	} else {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}

// article/nice のハンドラ（POST: 記事にいいねをつける）
func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		io.WriteString(w, "Posting Nice...\n")
	} else {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}

// comment のハンドラ（POST: 記事にコメントをする）
func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		io.WriteString(w, "Posting Comment...\n")
	} else {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}
