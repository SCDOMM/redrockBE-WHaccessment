package dao

import (
	"ProjectAndroidTest/model"
	"ProjectAndroidTest/pkg"
	"fmt"
	"log"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	dataBase *gorm.DB
)

func InitDataBase() {
	name := "root"
	password := "Aa3318752853"
	host := "127.0.0.1"
	port := "3306"
	dataBaseName := "test"
	dsn := name + ":" + password + "@tcp(" + host + ":" + port + ")/" + dataBaseName + "?charset=utf8mb4&parseTime=True&loc=Local"
	var err0 error
	dataBase, err0 = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err0 != nil {
		log.Println("数据库连接失败！")
		panic(err0)
	}
	err1 := dataBase.AutoMigrate(&model.HomeModel{})
	if err1 != nil {
		log.Println("数据库建表失败！")
		panic(err1)
	}
	err2 := dataBase.AutoMigrate(&model.DynamicModel{})
	if err2 != nil {
		log.Println("数据库建表失败！")
		panic(err1)
	}
	err3 := dataBase.AutoMigrate(&model.UserModel{})
	if err3 != nil {
		log.Println("数据库建表失败！")
		panic(err1)
	}
}
func HomeAddTestData() {
	for i := 0; i < 20; i++ {
		test := model.HomeModel{
			Title: "测试" + strconv.Itoa(i),
			Desc:  "这是一张图片",
			Image: "/9j/4AAQSkZJRgABAQAAAQABAAD/2wBDAAgGBgcGBQgHBwcJCQgKDBQNDAsLDBkSEw8UHRofHh0aHBwgJC4nICIsIxwcKDcpLDAxNDQ0Hyc5PTgyPC4zNDL/2wBDAQkJCQwLDBgNDRgyIRwhMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjL/wAARCAACAAQDASIAAhEBAxEB/8QAFQABAQAAAAAAAAAAAAAAAAAAAAf/xAAUEAEAAAAAAAAAAAAAAAAAAAAA/9oADAMBAAIRAxEAPwCdABmX/9k=",
		}
		err := dataBase.Create(&test)
		if err.Error != nil {
			log.Println("测试数据添加失败!")
			return
		}
	}
	fmt.Println("测试数据添加完成！")
}
func ChatAddTestData() {
	for i := 0; i < 20; i++ {
		test := model.DynamicModel{
			AuthorAccount: "114514" + strconv.Itoa(i),
			Title:         "测试" + strconv.Itoa(i),
			Desc:          "这是一张图片",
			CoverImage:    "/9j/4AAQSkZJRgABAQAAAQABAAD/2wBDAAgGBgcGBQgHBwcJCQgKDBQNDAsLDBkSEw8UHRofHh0aHBwgJC4nICIsIxwcKDcpLDAxNDQ0Hyc5PTgyPC4zNDL/2wBDAQkJCQwLDBgNDRgyIRwhMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjL/wAARCAACAAQDASIAAhEBAxEB/8QAFQABAQAAAAAAAAAAAAAAAAAAAAf/xAAUEAEAAAAAAAAAAAAAAAAAAAAA/9oADAMBAAIRAxEAPwCdABmX/9k=",
		}
		err := dataBase.Create(&test)
		if err.Error != nil {
			log.Println("测试数据添加失败!" + err.Error.Error())
			return
		}
	}
	fmt.Println("测试数据添加完成！")
}
func ReversoAddTestData() {
	hashPass, err := pkg.GeneratePassword("123456", 10)
	if err != nil {
		log.Println("测试数据添加失败!" + err.Error())
		return
	}
	for i := 0; i < 20; i++ {
		test := model.UserModel{
			UserName:     "测试" + strconv.Itoa(i),
			Account:      "114514" + strconv.Itoa(i),
			Password:     hashPass,
			ProfileImage: "/9j/4AAQSkZJRgABAQAAAQABAAD/2wBDAAgGBgcGBQgHBwcJCQgKDBQNDAsLDBkSEw8UHRofHh0aHBwgJC4nICIsIxwcKDcpLDAxNDQ0Hyc5PTgyPC4zNDL/2wBDAQkJCQwLDBgNDRgyIRwhMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjIyMjL/wAARCAACAAQDASIAAhEBAxEB/8QAFQABAQAAAAAAAAAAAAAAAAAAAAf/xAAUEAEAAAAAAAAAAAAAAAAAAAAA/9oADAMBAAIRAxEAPwCdABmX/9k=",
		}
		err := dataBase.Create(&test)
		if err.Error != nil {
			log.Println("测试数据添加失败!" + err.Error.Error())
			return
		}
	}
	fmt.Println("测试数据添加完成！")
}
