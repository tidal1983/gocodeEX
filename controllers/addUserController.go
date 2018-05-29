package controllers

import (
	"fmt"
	"html/template"
	"models/dao"
	"models/enity"
	"models/postgre"
	"net/http"
	"time"
)

func AddUserinfoController(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		t, _ := template.ParseFiles("views/adduserinfo.html")
		t.Execute(w, "")
	} else if r.Method == "POST" {
		r.ParseForm()
		var userinfo enity.Userinfo
		userinfo.Username = r.Form.Get("username")
		userinfo.Password = r.Form.Get("password")
		userinfo.Department = r.Form.Get("department")
		userinfo.Created = time.Now().Format(time.RFC3339)
		postgre.Insert(userinfo)
		t, _ := template.ParseFiles("views/show.html")
		t.Execute(w, "新增成功")
	}
}
func UpdateUserinfoController(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		r.ParseForm()
		uid := r.Form.Get("uid")
		fmt.Println(uid)
		t, _ := template.ParseFiles("views/updateuserinfo.html")
		t.Execute(w, dao.SelectSingle(uid))
	} else if r.Method == "POST" {
		r.ParseForm()
		uid := r.Form.Get("uid")
		userinfo := dao.SelectSingle(uid)
		userinfo.Password = r.Form.Get("password")
		userinfo.Department = r.Form.Get("department")
		postgre.Update(userinfo)
		t, _ := template.ParseFiles("views/show.html")
		t.Execute(w, "修改成功！")
	}
}
func SelectUserinfoController(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		t, _ := template.ParseFiles("views/userinfoGrid.html")
		total := ""
		for _, useinfo := range dao.SelectSlice() {

			var inputString = "<li>%v</li><li>%v</li><li>%v</li><input type='button' onclick='deletedata(%v)' value='刪除'/><br>"

			total += fmt.Sprintf(inputString, useinfo.Username, useinfo.Password, useinfo.Department, useinfo.UID)
		}
		t.Execute(w, template.HTML(total))
	}
}
func DeleteUserinfoController(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		r.ParseForm()
		uid := r.Form.Get("uid")
		var userinfo enity.Userinfo
		userinfo.UID = uid
		postgre.Delete(userinfo)
		w.Write([]byte("刪除成功"))
	}
}
