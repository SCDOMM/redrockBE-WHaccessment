package dao

import (
	"ProjectAndroidTest/model"
	"fmt"
	"log"
)

func DeleteHomeHandler(condition map[string]interface{}) error {
	var homeModel model.HomeModel
	err0 := dataBase.Where(condition).Delete(&homeModel).Error
	if err0 != nil {
		log.Println("无法删除对应栏目!", err0.Error())
		return err0
	}
	return nil
}

func HomeUploadHandler(receiver model.HomeDTO) error {
	//上传数据到数据库
	test := model.HomeModel{
		Title: receiver.Title,
		Desc:  receiver.Desc,
		Image: receiver.Image,
	}
	err1 := dataBase.Create(&test)
	if err1.Error != nil {
		log.Println("无法提交上传数据至主页!" + err1.Error.Error())
		return err1.Error
	}
	return nil
}
func DeleteChatHandler(condition map[string]interface{}) error {
	var dynamicModel model.DynamicModel
	result := dataBase.Where(condition).Delete(&dynamicModel)
	err0 := result.Error
	if err0 != nil {
		log.Println("无法删除对应栏目!", err0.Error())
		return err0
	}
	if result.RowsAffected == 0 {
		log.Printf("条件：%+v", condition)
		return fmt.Errorf("未找到匹配条件的动态数据")
	}
	return nil
}
