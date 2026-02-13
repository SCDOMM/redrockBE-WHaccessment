package service

import (
	"ProjectAndroidTest/dao"
	"ProjectAndroidTest/model"
	"ProjectAndroidTest/pkg"
	"log"
)

func DeleteHomeHandler(homeDto model.HomeDTO) pkg.Response {
	if homeDto.Title == "" || homeDto.Desc == "" || homeDto.Image == "" {
		return pkg.Response{
			Status: "400",
			Info:   "the title/desc/iamge is empty!",
		}
	}
	condition := map[string]interface{}{
		"title": homeDto.Title,
		"desc":  homeDto.Desc,
	}
	log.Println(condition["title"], condition["desc"])
	err0 := dao.DeleteHomeHandler(condition)
	if err0 != nil {
		return pkg.InternalError(err0)
	}
	return pkg.Response{
		Status: "200",
		Info:   "success!",
	}
}

func HomeUploadHandler(receiver model.HomeDTO) pkg.Response {
	//空值判断
	if receiver.Desc == "" || receiver.Image == "" || receiver.Title == "" {
		return pkg.Response{Status: "400", Info: "content is empty"}
	}

	//调用dao，上传数据
	err := dao.HomeUploadHandler(receiver)
	//说明数据库出错
	if err != nil {
		return pkg.InternalError(err)
	}

	return pkg.Response{Status: "200", Info: "upload success!"}
}
func DeleteChatHandler(dynamicDto model.DynamicDTO) pkg.Response {
	if dynamicDto.AuthorName == "" || dynamicDto.ProfileImage == "" || dynamicDto.CoverImage == "" || dynamicDto.Title == "" || dynamicDto.Desc == "" {
		return pkg.Response{
			Status: "400",
			Info:   "the content is empty!",
		}
	}
	condition := map[string]interface{}{
		"title":          dynamicDto.Title,
		"desc":           dynamicDto.Desc,
		"author_account": dynamicDto.Account,
		"cover_image":    dynamicDto.CoverImage,
	}
	err0 := dao.DeleteChatHandler(condition)
	if err0 != nil {
		return pkg.InternalError(err0)
	}
	return pkg.Response{Status: "200", Info: "success!"}
}
