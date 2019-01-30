package models

import (
	"enterbj_print/utils"
	"fmt"
	"github.com/15125505/zlog/log"
	"github.com/globalsign/mgo/bson"
)

/*
用户相关
*/

/*
判断用户名是否存在
*/
func UserCheckUserNameIsHave(name string) (isHave bool) {
	var u UserInfo
	isHave = false

	err := FindOne(DbName, BjCard, bson.M{"user_name": name}, nil, &u)
	if err == nil && u.UserId != "" {
		isHave = true
	} else {
		log.Error(err)
	}
	return isHave
}

/*
判断用户邮箱是否存在
*/
func UserCheckEmailIsHave(email string) (isHave bool) {
	var u UserInfo
	isHave = false

	err := FindOne(DbName, BjCard, bson.M{"email": email}, nil, &u)
	if err == nil && u.UserId != "" {
		isHave = true
	} else {
		log.Error(err)
	}
	return isHave
}

/*
新增用户信息
*/
func UserAddUserInfo(u UserInfo) (err error) {
	u.Id = bson.NewObjectId()
	u.UserId = utils.GetMd5Hash(u.UserName)
	u.PassWord = utils.GetMd5Hash(fmt.Sprintf("enterbj%senterbj", u.PassWord))
	u.CarList = []CarInfoListModel{}
	err = Insert(DbName, BjCard, u)
	return
}

/*
根据用户id查询用户信息
*/
func UserGetUserInfo(uId string) (u UserInfo, err error) {
	err = FindOne(DbName, BjCard, bson.M{"user_id": uId}, nil, &u)
	if err == nil && u.UserId != "" {
		//u.CarInfo = []CarInfoModel{} // 不展示车辆相关的信息
		return u, nil
	} else {
		return u, err
	}
}

/*
新增用户的车辆牌照信息
*/
func UserAddLicenseNo(uId string, licenseNo string) (err error) {
	car := CarInfoListModel{
		LicenseNo: licenseNo,
		CardList:  []CardListModel{},
	}
	err = Update(DbName, BjCard, bson.M{"user_id": uId}, bson.M{"$push": bson.M{"car_list": car}})
	return err
}
