package service

import (
	"ProjectAndroidTest/dao"
	"ProjectAndroidTest/model"
	"ProjectAndroidTest/pkg"
)

func HomePageHandler() ([10]model.HomeDTO, pkg.Response) {
	//调用dao
	data, err := dao.HomePageHandler()
	//说明数据库出错了
	if err != nil {
		return [10]model.HomeDTO{}, pkg.InternalError(err)
	}

	//给DTO赋值
	var resultDTO [10]model.HomeDTO
	for i, item := range data {
		resultDTO[i].Title = item.Title
		resultDTO[i].Image = item.Image
		resultDTO[i].Desc = item.Desc
	}

	return resultDTO, pkg.Response{Status: "200", Info: "acquire message success!"}
}
func HomeSearchHandler(request model.SearchDTO) ([]model.HomeDTO, pkg.Response) {
	//空值判断
	if request.Content == "" {
		return nil, pkg.Response{Status: "400", Info: "content is empty"}
	}

	//调用dao，上传搜索关键词
	data, err := dao.HomeSearchHandler(request)
	//说明数据库出错了
	if err != nil {
		return nil, pkg.InternalError(err)
	}

	//给DTO添加搜索结果
	var resultDTO []model.HomeDTO
	for _, item := range data {
		resultDTO = append(resultDTO, model.HomeDTO{Title: item.Title, Desc: item.Desc, Image: item.Image})
	}
	//说明没有搜到结果
	if len(resultDTO) == 0 {
		return nil, pkg.Response{Status: "200", Info: "no result found!"}
	}

	return resultDTO, pkg.Response{Status: "200", Info: "search success!"}
}
