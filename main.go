package main

import (
	"controllers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", controllers.Login)
	http.HandleFunc("/login", controllers.Login) //設置訪問的路由
	http.HandleFunc("/adduserinfo", controllers.AddUserinfoController)
	http.HandleFunc("/updateUserinfo", controllers.UpdateUserinfoController)
	http.HandleFunc("/selectUserinfo", controllers.SelectUserinfoController)
	http.HandleFunc("/deleteUser", controllers.DeleteUserinfoController)
	http.HandleFunc("/checkUser", controllers.CheckUserController)

	err := http.ListenAndServe(":9090", nil) //設置監聽的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	} else {

	}

}
