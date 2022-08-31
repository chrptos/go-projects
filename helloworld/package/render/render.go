package render

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func RenderTemplateOld(w http.ResponseWriter, tmpl string) {
	// 文字列でレンダリングしたいテンプレートファイルを受け取る
	// baseテンプレートも含まなければいけない
	// TODO: ディスクから読み取るのをやめる
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}
}

var tc = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	_, inMap := tc[t]
	if !inMap {
		// テンプレート作成処理
		log.Println("creating template and cache")
		err := createTemplateCache(t)
		if err != nil {
			log.Println(err)
		}
	} else {
		// キャッシュ作成済み
		log.Println("cached")
	}

	tmpl = tc[t]

	// ディスクリプタへ出力（Execute）
	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}

}

func createTemplateCache(t string) error {
	// 配列で処理対象を保持する
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.tmpl",
	}

	// テンプレートを新規作成
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}

	// テンプレートキャッシュへ追加する
	tc[t] = tmpl

	return nil
}
