package models

import "github.com/globalsign/mgo/bson"

type UserInfo struct {
	Id       bson.ObjectId      `bson:"id"`
	UserId   string             `bson:"user_id" json:"userId"`     // 对用户名做MD5后的值
	UserName string             `bson:"user_name" json:"userName"` // 用户名（英文字母或数字，不可随意更改）
	NickName string             `bson:"nick_name" json:"nickName"` // 昵称 （可以更改）
	PassWord string             `bson:"pass_word" json:"passWord"` // 密码
	Email    string             `bson:"email" json:"email"`        // 邮箱
	CarList  []CarInfoListModel `bson:"car_list" json:"carList"`   // 车辆信息
}

type CarInfoListModel struct {
	LicenseNo string          `bson:"license_no" json:"licenseNo"` // 车牌号
	CardList  []CardListModel `bson:"card_list" json:"cardList"`   // 进京证申请记录
}

type CardListModel struct {
	FlowingNo    string      `bson:"flowing_no" json:"flowingNo"`        // 序列号
	FlowingImg   string      `bson:"flowing_img" json:"flowingImg"`      // 进京证序列号图片
	ImgPath      string      `bson:"img_path" json:"imgPath"`            // 进京证图片的官网链接
	ImgLocalPath string      `bson:"img_local_path" json:"imgLocalPath"` // 进京证图片的本地链接
	ImgBinary    bson.Binary `bson:"img_binary" json:"imgBinary"`        // 进京证图片的二进制存储
	Status       int         `bson:"status" json:"status"`               // 进京证状态 /* 1:flowNo存在了 2:申请信息获取到了 3：进京证图片获取到了 4：信息已完善*/
	ApplyInfo    ApplyModel  `bson:"apply_info" json:"applyInfo"`        // 进京证申请信息
	CreateTime   int64       `bson:"create_time" json:"createTime"`      // 添加时间
}

type ApplyModel struct {
	ApplyId    string `bson:"apply_id" json:"applyId"`       // 申请编号
	DriverName string `bson:"driver_name" json:"driverName"` // 驾驶人
	LicenseNo  string `bson:"license_no" json:"licenseNo"`   // 车牌号
	EndTime    string `bson:"end_time" json:"endTime"`       // 进京证结束日期
	StartTime  string `bson:"start_time" json:"startTime"`   // 进京证开始日期
	CreateTime string `bson:"create_time" json:"createTime"` // 进京证申请日期
	Address    string `bson:"address" json:"address"`        // 单位或个人住址
	PaperId    string `bson:"paper_id" json:"paperId"`       // 进京序列号
	CarType    string `bson:"car_type" json:"carType"`       // 车辆号牌类型
}
