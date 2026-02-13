package service

import (
	"ProjectAndroidTest/dao"
	"ProjectAndroidTest/model"
	"ProjectAndroidTest/pkg"
	"errors"
	"log"
	"time"
)

func LogonHandler(logonDto model.LogonDTO) (model.LogonResponseDTO, pkg.Response) {
	if logonDto.Password == "" || logonDto.Account == "" {
		log.Println("logon:account or password is empty!")
		return model.LogonResponseDTO{}, pkg.Response{Status: "400", Info: "logon:account or password is empty!"}
	}
	err, userData := dao.LogonHandler(logonDto)
	if err != nil {
		if errors.Is(err, pkg.ErrAccNotFound) {
			return model.LogonResponseDTO{}, pkg.Response{Status: "400", Info: pkg.ErrAccNotFound.Error()}
		} else if errors.Is(err, pkg.ErrPassError) {
			return model.LogonResponseDTO{}, pkg.Response{Status: "400", Info: pkg.ErrPassError.Error()}
		}
		return model.LogonResponseDTO{}, pkg.InternalError(err)
	}

	//发Token
	accessToken, err := pkg.CreateAccessToken(userData.Account, userData.Role)
	if err != nil {
		return model.LogonResponseDTO{}, pkg.InternalError(err)
	}
	refreshToken, err := pkg.CreateRefreshToken(userData.Account, userData.Role)
	if err != nil {
		return model.LogonResponseDTO{}, pkg.InternalError(err)
	}
	return model.LogonResponseDTO{
		UserName:     userData.UserName,
		ProfileImage: userData.ProfileImage,
		OperationDTO: model.OperationDTO{
			ExpirationTime: time.Now().Add(+7 * 24 * time.Hour),
			UserRole:       userData.Role,
			AccessToken:    accessToken,
			RefreshToken:   refreshToken,
		},
	}, pkg.Response{Status: "200", Info: "success"}
}

func RegisterHandler(registerDto model.RegisterDTO) pkg.Response {
	if registerDto.UserName == "" || registerDto.Password == "" || registerDto.Account == "" {
		log.Println("register:account or password is empty!")
		return pkg.Response{Status: "400", Info: "account or password is empty!"}
	}
	err := dao.RegisterHandler(registerDto)
	if err != nil {
		if errors.Is(err, pkg.ErrAccExist) {
			return pkg.Response{Status: "400", Info: pkg.ErrAccExist.Error()}
		}
		return pkg.InternalError(err)
	}
	return pkg.Response{Status: "200", Info: "success"}
}
func DeregisterHandler(logonDto model.LogonDTO) pkg.Response {
	if logonDto.Password == "" || logonDto.Account == "" {
		log.Println("deregister:account or password is empty!")
		return pkg.Response{Status: "400", Info: "the account/password is empty!"}
	}
	err := dao.DeregisterHandler(logonDto)
	if err != nil {
		if errors.Is(err, pkg.ErrAccNotFound) {
			return pkg.Response{Status: "400", Info: pkg.ErrAccNotFound.Error()}
		} else if errors.Is(err, pkg.ErrPassError) {
			return pkg.Response{Status: "400", Info: pkg.ErrPassError.Error()}
		}
		return pkg.InternalError(err)
	}
	return pkg.Response{Status: "200", Info: "success"}
}
func ChangeProfileHandler(changeProfileDto model.ChangeProfileDTO) pkg.Response {
	if changeProfileDto.ProfileImage == "" || changeProfileDto.Name == "" || changeProfileDto.Account == "" {
		return pkg.Response{Status: "400", Info: "the image/name/account is empty!"}
	}
	err := dao.ChangeProfileHandler(changeProfileDto)
	if err != nil {
		if errors.Is(err, pkg.ErrAccNotFound) {
			return pkg.Response{Status: "400", Info: pkg.ErrAccNotFound.Error()}
		}
		return pkg.InternalError(err)
	}
	return pkg.Response{Status: "200", Info: "success"}
}
