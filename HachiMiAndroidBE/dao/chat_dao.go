package dao

import (
	"ProjectAndroidTest/model"
	"errors"
	"log"
)

func ChatPageHandler() ([]model.DynamicModel, []model.UserModel, error) {
	var data []model.DynamicModel
	//随机获取10个值
	if err0 := dataBase.Model(&model.DynamicModel{}).Order("RAND()").Limit(10).Find(&data).Error; err0 != nil {
		log.Println("动态获取表数目失败!" + err0.Error())
		return nil, nil, err0
	}
	if len(data) < 10 {
		return nil, nil, errors.New("数据库现存数据少于十个")
	}

	//绑定十个值的用户数据
	var userData []model.UserModel
	err1 := BindUserData(&data, &userData)
	if err1 != nil {
		return nil, nil, err1
	}

	return data, userData, nil
}

func ChatSearchHandler(req model.SearchDTO) ([]model.DynamicModel, []model.UserModel, error) {
	var data []model.DynamicModel
	//搜索所有值
	err0 := dataBase.Model(&model.DynamicModel{}).Where("title LIKE ?", "%"+req.Content+"%").Find(&data)
	if err0.Error != nil {
		log.Println("动态搜索出错!" + err0.Error.Error())
		return nil, nil, err0.Error
	}
	if len(data) == 0 {
		return nil, nil, nil
	}
	//绑定所有值的用户数据
	var userData []model.UserModel
	err1 := BindUserData(&data, &userData)
	if err1 != nil {
		return nil, nil, err1
	}
	return data, userData, nil
}
func ChatUploadHandler(receiver model.DynamicUploadDTO) error {
	//上传数据到数据库
	test := model.DynamicModel{
		AuthorAccount: receiver.AuthorAccount,
		Title:         receiver.Title,
		Desc:          receiver.Desc,
		CoverImage:    receiver.CoverImage,
	}
	err1 := dataBase.Create(&test)
	if err1.Error != nil {
		log.Println("无法提交上传数据至主页！")
		return err1.Error
	}
	return nil
}
func BindUserData(data *[]model.DynamicModel, userData *[]model.UserModel) error {
	//获得绑定的用户数据
	var accountSlice = make([]string, 0, len(*data))
	for _, item := range *data {
		accountSlice = append(accountSlice, item.AuthorAccount)
	}
	var uniqueUsers []model.UserModel
	if err1 := dataBase.Model(&model.UserModel{}).Where("account IN ?", accountSlice).Find(&uniqueUsers).Error; err1 != nil {
		log.Println("动态绑定表数目失败!" + err1.Error())
		return err1
	}

	userMap := make(map[string]model.UserModel, len(uniqueUsers))
	for _, item := range uniqueUsers {
		userMap[item.Account] = item
	}
	*userData = make([]model.UserModel, len(*data))
	for i := range *data {
		account := (*data)[i].AuthorAccount
		user, ok := userMap[account]
		if ok {
			(*userData)[i].ProfileImage = user.ProfileImage
			(*userData)[i].UserName = user.UserName
			(*userData)[i].Account = user.Account
		}
	}
	return nil
}
