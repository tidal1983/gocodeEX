package dao

import (
	"models/enity"
	"models/postgre"
)

func SelectSlice() []*enity.Userinfo {
	//查询數據
	db, err := postgre.OpenDB()
	postgre.CheckErr(err)
	rows, err := db.Query("SELECT * FROM userinfo order by uid")
	postgre.CheckErr(err)
	var datas []*enity.Userinfo
	for rows.Next() {
		// var userinfo enity.Userinfo
		userinfo := new(enity.Userinfo)
		var err = rows.Scan(&userinfo.UID, &userinfo.Username, &userinfo.Department, &userinfo.Password, &userinfo.Created)
		datas = append(datas, userinfo)
		postgre.CheckErr(err)
	}
	return datas
}
func CheckdataSlice(username string) []*enity.Userinfo {
	//查询數據
	db, err := postgre.OpenDB()
	postgre.CheckErr(err)
	rows, err := db.Query("SELECT * FROM userinfo where username =$1 ", username)
	postgre.CheckErr(err)
	var datas []*enity.Userinfo
	for rows.Next() {
		// var userinfo enity.Userinfo
		userinfo := new(enity.Userinfo)
		var err = rows.Scan(&userinfo.UID, &userinfo.Username, &userinfo.Department, &userinfo.Password, &userinfo.Created)
		datas = append(datas, userinfo)
		postgre.CheckErr(err)
	}
	return datas
}
func SelectSingle(uid string) enity.Userinfo {
	//查询數據
	db, err := postgre.OpenDB()
	postgre.CheckErr(err)

	var userinfo enity.Userinfo
	err = db.QueryRow("SELECT * FROM userinfo where uid = $1 ", uid).
		Scan(&userinfo.UID, &userinfo.Username, &userinfo.Department, &userinfo.Password, &userinfo.Created)
	postgre.CheckErr(err)

	return userinfo
}
