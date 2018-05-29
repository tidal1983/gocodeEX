package controllers

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"models/dao"
	"net/http"
	"strconv"
	"time"
)

func CheckUserController(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		username := r.Form.Get("username")
		if len(dao.CheckdataSlice(username)) > 0 {
			w.Write([]byte("false"))
		} else {
			w.Write([]byte("success"))
		}
	}
}
func Login(w http.ResponseWriter, r *http.Request) {
	//獲取請求方法
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		t, _ := template.ParseFiles("views/login.html")
		t.Execute(w, token)
	} else if r.Method == "POST" {

		r.ParseForm()
		token := r.Form.Get("token")
		if token != "" {
			//驗證token的合法性
			usernameF := r.Form.Get("username")
			passowrd := r.Form.Get("password")
			for i, userinfo := range dao.CheckdataSlice(usernameF) {
				if i == 0 && passowrd == userinfo.Password {
					t, _ := template.ParseFiles("views/view.html")
					t.Execute(w, usernameF)
					break
				} else {
					t, _ := template.ParseFiles("views/show.html")
					t.Execute(w, "登入失敗，密碼錯誤")
				}
			}
		} else {
			//不存在token報錯
			t, _ := template.ParseFiles("views/login.html")
			t.Execute(w, "資料有誤")
		}

	}
}
