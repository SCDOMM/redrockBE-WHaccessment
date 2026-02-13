package service

import (
	"ProjectAndroidTest/dao"
	"ProjectAndroidTest/model"
	"ProjectAndroidTest/pkg"
)

func ChatPageHandler() ([]model.DynamicDTO, pkg.Response) {
	//调用dao
	dynamicData, userData, err := dao.ChatPageHandler()
	//说明数据库出错
	if err != nil {
		return nil, pkg.InternalError(err)
	}
	//绑定数据到DTO
	var dynamicDTO []model.DynamicDTO
	for i := 0; i < len(dynamicData); i++ {
		dynamicDTO = append(dynamicDTO, model.DynamicDTO{
			Title:        dynamicData[i].Title,
			CoverImage:   dynamicData[i].CoverImage,
			Desc:         dynamicData[i].Desc,
			AuthorName:   userData[i].UserName,
			ProfileImage: userData[i].ProfileImage,
			Time:         dynamicData[i].CreatedAt,
			Account:      userData[i].Account,
		})
	}
	return dynamicDTO, pkg.Response{Status: "200", Info: "acquire message success!"}
}
func ChatSearchHandler(request model.SearchDTO) ([]model.DynamicDTO, pkg.Response) {
	//调用dao
	dynamicData, userData, err := dao.ChatSearchHandler(request)
	//说明数据库出错
	if err != nil {
		return nil, pkg.InternalError(err)
	}
	//没搜到
	if len(dynamicData) == 0 {
		return nil, pkg.Response{Status: "200", Info: "no result found!"}
	}
	//绑定DTO
	var dynamicDTO []model.DynamicDTO
	for i := 0; i < len(dynamicData); i++ {
		dynamicDTO = append(dynamicDTO, model.DynamicDTO{
			Title:        dynamicData[i].Title,
			CoverImage:   dynamicData[i].CoverImage,
			Desc:         dynamicData[i].Desc,
			AuthorName:   userData[i].UserName,
			ProfileImage: userData[i].ProfileImage,
			Time:         dynamicData[i].CreatedAt,
			Account:      userData[i].Account,
		})
	}

	return dynamicDTO, pkg.Response{Status: "200", Info: "search success!"}
}
func ChatUploadHandler(receiver model.DynamicUploadDTO) pkg.Response {
	err := dao.ChatUploadHandler(receiver)
	if err != nil {
		return pkg.InternalError(err)
	}
	return pkg.Response{Status: "200", Info: "upload success!"}
}
