package db

import (
	//"blogger/dao/db"
	"blogger/model"
	"log"
	"testing"
)

func init() {
	//	parseTime=true 将mysql中时间类型，自动解析为 go 结构体中的时间类型
	//	不加报错
	dns := "RHW:RHW943359178@tcp(81.69.255.188:3306)/blogger?parseTime=True"
	err := Init(dns)
	if err != nil {
		panic(err)
	}
}

func TestInsertUserInfo(t *testing.T) {
	//	构建用户对象
	//	初始化结构体
	userInfo := &model.User{}
	userInfo.UserId = "abcdefg"
	userInfo.Username = "RHW"
	userInfo.Password = "123456"
	userInfo.Email = "943359178@qq.com"
	info, err := InsertUserInfo(userInfo)
	if err != nil {
		log.Fatal("get data from db failed, err: ", err)
		return
	}
	log.Println("insertId: ", info)
}
