package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/chrptos/myapi/models"
	"github.com/gorilla/mux"
)

// /article のハンドラ
func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	// バイトスライス reqBodybufferを何らかの形で用意
	// req.HeaderのGetメソッドで得られたContent-Lengthはstring型なので数値に変換する
	length, err := strconv.Atoi(req.Header.Get("Content-Length"))
	if err != nil {
		http.Error(w, "cannot get content length\n", http.StatusBadRequest)
		return
	}
	// 得られたlength分のバイト配列を作成する
	reqBodybuffer := make([]byte, length)
	// var reqBodybuffer []byte

	// ReadCloser(Reader+Closer)
	// Readメソッド：ボディの中身を、引数として渡したバイトスライスに格納する
	// Closeメソッド：ボディの中身を読み終わった時にクローズする
	if _, err := req.Body.Read(reqBodybuffer); !errors.Is(err, io.EOF) {
		// errorsパッケージのIs関数で、err と io.EOFが等しいか判定する（EOFは読み切ったかどうか
		http.Error(w, "fail to get request body\n", http.StatusBadRequest)
		return
	}
	// クローズ処理
	defer req.Body.Close()

	// json→Article構造体へデコード処理
	var reqArticle models.Article
	if err := json.Unmarshal(reqBodybuffer, &reqArticle); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
		return
	}

	// Article構造体→jsonエンコード処理
	article := models.Article1
	jsonData, err := json.Marshal(article)
	if err != nil {
		http.Error(w, "fail to encode\n", http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

// /article/list のハンドラ
func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	queryMap := req.URL.Query()

	var page int
	if p, ok := queryMap["page"]; ok && len(p) > 0 {
		var err error
		page, err = strconv.Atoi(p[0])
		if err != nil {
			http.Error(w, "Invalid query parameter", http.StatusBadRequest)
			return
		}
	} else {
		page = 1
	}

	articleList := []models.Article{models.Article1, models.Article2}
	jsonData, err := json.Marshal(articleList)
	if err != nil {
		errMsg := fmt.Sprintf("fail to encode json (page %d)\n", page)
		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

// /article/{id} のハンドラ
func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}
	resString := fmt.Sprintf("Article No.%d\n", articleID)
	io.WriteString(w, resString)

	article := models.Article1
	jsonData, err := json.Marshal(article)
	if err != nil {
		http.Error(w, "fail to encode\n", http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

// /article/nice のハンドラ
func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	article := models.Article1
	jsonData, err := json.Marshal(article)
	if err != nil {
		http.Error(w, "fail to encode\n", http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}

// /comment のハンドラ
func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	comment := models.Comment1
	jsonData, err := json.Marshal(comment)
	if err != nil {
		http.Error(w, "fail to encode\n", http.StatusInternalServerError)
		return
	}

	w.Write(jsonData)
}
