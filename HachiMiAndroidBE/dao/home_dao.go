package dao

import (
	"ProjectAndroidTest/model"
	"log"
)

// 随机抽取十个值返回前端
func HomePageHandler() ([10]model.HomeModel, error) {
	var data [10]model.HomeModel
	result := dataBase.Model(&model.HomeModel{}).Order("RAND()").Limit(10).Find(&data)
	if result.Error != nil {
		log.Println("主页获取表数目失败!" + result.Error.Error())
		return [10]model.HomeModel{}, result.Error
	}
	return data, nil
}
func HomeSearchHandler(req model.SearchDTO) ([]model.HomeModel, error) {
	var data []model.HomeModel
	//简单的搜索逻辑
	err0 := dataBase.Where("title LIKE ?", "%"+req.Content+"%").Find(&data)
	if err0.Error != nil {
		log.Println("主页搜索出错!" + err0.Error.Error())
		return nil, err0.Error
	}
	return data, nil
}
