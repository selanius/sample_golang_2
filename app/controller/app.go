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
	users := []models.User{}

	rows, err := dbm.Query("select id, name from users")
	for rows.Next() {
		user := models.User{}
		if err = rows.Scan(&user.ID, &user.Name); err != nil {
			// 変換に失敗した場合、INDEX画面表示
			toIndex(tmpl, rw, request, "ユーザ データ変換 失敗", err)
			return
		}
		users = append(users, user)
	}

	// HOME画面表示
	err = tmpl.ExecuteTemplate(rw, "home.html", struct {
		Users []models.User
	}{
		Users: users,
	})
	if err != nil {
		outputErrorLog("HTML 描画 エラー", err)
	}
}

// Detail is method to render Detail page.
func Detail(rw http.ResponseWriter, request *http.Request){
	tmpl := parseTemplate()

	// ID取得
	id := request.URL.Query().Get("id")
	log.Println("id : ", id)

	// db connection
	dbm := db.ConnDB()
	user := new(models.User)
	row := dbm.QueryRow("select id, name from users where id = ?", id)
	if err := row.Scan(&user.ID, &user.Name); err != nil {
		// 変換に失敗した場合、INDEX画面表示
		toIndex(tmpl, rw, request, "ユーザ データ変換 失敗", err)
		return
	}

	// Detail画面表示
	err := tmpl.ExecuteTemplate(rw, "detail.html", struct {
		User     *models.User
	}{
		User:     user,
	})
	if err != nil {
		outputErrorLog("HTML 描画 エラー", err)
	}
}

// ユーザー情報更新
func Edit(rw http.ResponseWriter, request *http.Request){
	tmpl := parseTemplate()

	// 変更する名前とIDを取得
	err := request.ParseForm()
	if err != nil {
		// パースに失敗したらエラー
		toIndex(tmpl, rw, request, "フォーム パース エラー", err)
		return
	}
	name := request.Form.Get("name")
	log.Println("変更する名前：", name)

	id := request.Form.Get("id")
	log.Println("変更する対象ID：", id)

	// db connection
	dbm := db.ConnDB()
	_, err = dbm.Exec("update users set name = ? where id = ?", name, id)
	if err != nil {
		log.Println("ユーザー 更新 失敗")
	} else {
		log.Println("ユーザー 更新 成功")
		log.Println("更新したユーザーID：", id)
	}

	// HOME画面表示
	http.Redirect(rw, request, "/home", http.StatusFound)
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
