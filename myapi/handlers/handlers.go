package handlers

import (
	"fmt"
	"io"
	"net/http"
)

// /hello のハンドラ
func HelloHandler(w http.ResponseWriter, req *http.Request) {
	// もし、req の中の Method フィールドが GET だったら
	if req.Method == http.MethodGet {
		// 通常通りにレスポンスを返す
		io.WriteString(w, "Hello, world!\n")
	} else {
		// Invalid methodというレスポンスを、405番のステータスコードと共に返す
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}

// /article のハンドラ
func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	// もし、req の中の Method フィールドが POST だったら
	if req.Method == http.MethodPost {
		// 通常通りにレスポンスを返す
		io.WriteString(w, "Posting Article...\n")
	} else {
		// Invalid methodというレスポンスを、405番のステータスコードと共に返す
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}

// /article/list のハンドラ
func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	// もし、req の中の Method フィールドが GET だったら
	if req.Method == http.MethodGet {
		// 通常通りにレスポンスを返す
		io.WriteString(w, "Article List\n")
	} else {
		// Invalid methodというレスポンスを、405番のステータスコードと共に返す
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}

// /article/1 のハンドラ
func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleID := 1
	resString := fmt.Sprintf("Article No.%d\n", articleID)
	// もし、req の中の Method フィールドが GET だったら
	if req.Method == http.MethodGet {
		// 通常通りにレスポンスを返す
		io.WriteString(w, resString)
	} else {
		// Invalid methodというレスポンスを、405番のステータスコードと共に返す
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}

// /article/nice のハンドラ
func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	// もし、req の中の Method フィールドが POST だったら
	if req.Method == http.MethodPost {
		// 通常通りにレスポンスを返す
		io.WriteString(w, "Posting Nice...\n")
	} else {
		// Invalid methodというレスポンスを、405番のステータスコードと共に返す
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}

// /comment のハンドラ
func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	// もし、req の中の Method フィールドが POST だったら
	if req.Method == http.MethodPost {
		// 通常通りにレスポンスを返す
		io.WriteString(w, "Posting Comment...\n")
	} else {
		// Invalid methodというレスポンスを、405番のステータスコードと共に返す
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}
