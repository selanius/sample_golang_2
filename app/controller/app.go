package controller

import (
	"html/template"
	"log"
	"net/http"

	"../models"
	"../db"
)

// Index is method to render Top page.
func Index(rw http.ResponseWriter, request *http.Request) {

	log.Println("call Index")

	// htmlを表示
	err := parseTemplate().ExecuteTemplate(rw, "index.html", "")
	if err != nil {
		outputErrorLog("HTML 描画 エラー", err)
		log.Fatalln("エラーのため強制終了")
	}
}

// Home is method to render Name page.
func Home(rw http.ResponseWriter, request *http.Request){

	tmpl := parseTemplate()

	// db connection
	dbm := db.ConnDB()
	user := new(models.User)

	row := dbm.QueryRow("select id, name from users where id = ?", 1)
	if err := row.Scan(&user.ID, &user.Name); err != nil {
		// 変換に失敗した場合、INDEX画面表示
		toIndex(tmpl, rw, request, "ユーザ データ変換 失敗", err)
		return
	}

	// HOME画面表示
	err := tmpl.ExecuteTemplate(rw, "home.html", struct {
		User     *models.User
	}{
		User:     user,
	})
	if err != nil {
		outputErrorLog("HTML 描画 エラー", err)
	}
}

// parse HTML
func parseTemplate() *template.Template {
	tmpl, err := template.ParseGlob("../app/view/*.html")
	if err != nil {
		outputErrorLog("HTML パース 失敗", err)
	}
	return tmpl
}

// output error log
func outputErrorLog(message string, err error) {
	log.Println(message)
	log.Println(err)
}

// render Top page
func toIndex(tmpl *template.Template, rw http.ResponseWriter, request *http.Request, message string, err error) {
	outputErrorLog(message, err)
	http.Redirect(rw, request, "/index", http.StatusFound)
}
