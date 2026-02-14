package dao

import (
	"ProjectAndroidTest/model"
	"ProjectAndroidTest/pkg"
	"errors"
	"log"
	"time"

	"gorm.io/gorm"
)

func LogonHandler(logonDto model.LogonDTO) (error, model.UserModel) {
	//检验账号是否存在
	err0, errCode := SearchAccount(logonDto.Account)
	if errCode == -1 {
		return err0, model.UserModel{}
	} else if errCode == 0 {
		return pkg.ErrAccNotFound, model.UserModel{}
	}

	//测验密码
	err1, userData := VerifyPassword(logonDto.Account, logonDto.Password)
	if err1 != nil {
		return err1, model.UserModel{}
	}
	return nil, userData
}
func RegisterHandler(registerDto model.RegisterDTO) error {
	hashPassword, err := pkg.GeneratePassword(registerDto.Password, 10)
	if err != nil {
		log.Println("register:generate hash error!", err.Error())
		return err
	}

	//检验账号是否存在
	err1, errCode := SearchAccount(registerDto.Account)
	if errCode == -1 {
		return err1
	} else if errCode == 1 {
		return pkg.ErrAccExist
	}

	userData := model.UserModel{
		UserName:     registerDto.UserName,
		Account:      registerDto.Account,
		Password:     hashPassword,
		ProfileImage: model.DefaultImage}
	//检查软删除
	var deletedModel model.UserModel
	err2 := dataBase.Unscoped().Where("account = ?", registerDto.Account).First(&deletedModel).Error
	if err2 == nil {
		deletedModel.UserName = registerDto.UserName
		deletedModel.Account = registerDto.Account
		deletedModel.Password = hashPassword
		deletedModel.ProfileImage = model.DefaultImage
		deletedModel.Role = 0
		deletedModel.CreatedAt = time.Now()
		deletedModel.UpdatedAt = time.Time{}
		deletedModel.DeletedAt = gorm.DeletedAt{}
		return dataBase.Unscoped().Save(&deletedModel).Error
	}
	if !errors.Is(err2, gorm.ErrRecordNotFound) {
		log.Println("register:create account error:", err2)
		return err2
	}

	if result := dataBase.Create(&userData); result.Error != nil {
		log.Println("register:create account error:", result.Error.Error())
		return result.Error
	}
	return nil
}
func DeregisterHandler(logonDto model.LogonDTO) error {

	err0, errCode := SearchAccount(logonDto.Account)
	if errCode == -1 {
		return err0
	} else if errCode == 0 {
		return pkg.ErrAccNotFound
	}

	err1, _ := VerifyPassword(logonDto.Account, logonDto.Password)
	if err1 != nil {
		return err1
	}

	result := dataBase.Where("account = ?", logonDto.Account).Delete(&model.UserModel{})
	if result.Error != nil {
		log.Println("deregister:delete account error:", result.Error.Error())
		return result.Error
	}
	result1 := dataBase.Where("author_account = ?", logonDto.Account).Delete(&model.DynamicModel{})
	if result1.Error != nil {
		log.Println("deregister:delete account error:", result1.Error.Error())
		return result1.Error
	}
	return nil
}
func ChangeProfileHandler(changeProfileDto model.ChangeProfileDTO) error {
	err0, errCode := SearchAccount(changeProfileDto.Account)
	if errCode == -1 {
		return err0
	} else if errCode == 0 {
		return pkg.ErrAccNotFound
	}
	result := dataBase.Model(&model.UserModel{}).Where("account = ?", changeProfileDto.Account).Updates(&model.UserModel{
		UserName:     changeProfileDto.Name,
		ProfileImage: changeProfileDto.ProfileImage,
	})
	if result.Error != nil {
		log.Println("changeProfile:update account error:", result.Error.Error())
		return result.Error
	}
	return nil
}

func SearchAccount(Account string) (error, int8) {
	var searchResult int64
	if result := dataBase.Model(&model.UserModel{}).Where("account = ?", Account).Count(&searchResult); result.Error != nil {
		log.Println("dataBase error:", result.Error.Error())
		return result.Error, -1
	}
	if searchResult != 0 {
		log.Println("account is exist!")
		return nil, 1
	}
	log.Println("account is not exist!")
	return nil, 0
}
func VerifyPassword(Account string, Password string) (error, model.UserModel) {
	var userData model.UserModel
	if result := dataBase.Model(&model.UserModel{}).Where("account = ?", Account).First(&userData); result.Error != nil {
		log.Println("cannot found this account!" + result.Error.Error())
		return result.Error, model.UserModel{}
	}
	err1 := pkg.ComparePassword(Password, userData.Password)
	if err1 != nil {
		log.Println("password error:", err1.Error())
		return pkg.ErrPassError, model.UserModel{}
	}
	return nil, userData
}
