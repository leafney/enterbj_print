package models

import (
	"enterbj_print/utils"
	"github.com/globalsign/mgo/bson"
)

/*
进京证证件操作
*/

/*
用户为指定车辆新增申请序列号信息
@uId 用户id
@licenseNo 车牌号
@flowNo 申请序列号
@flowImg 申请序列号图片链接
*/
func CardAddNewFlowingInfo(uId string, licenseNo string, flowNo string, flowImg string) (err error) {
	card := CardListModel{
		FlowingNo:    flowNo,
		ApplyInfo:    ApplyModel{},
		FlowingImg:   flowImg,
		ImgPath:      "",
		ImgLocalPath: "",
		ImgBinary:    bson.Binary{},
		Status:       2,
		CreateTime:   utils.CreateUnix(),
	}
	err = Update(DbName, BjCard, bson.M{"user_id": uId, "car_list.license_no": licenseNo}, bson.M{"$push": bson.M{"car_list.$.card_list": card}})
	return
}

/*
用户通过序列号查询到的申请信息
*/
func CardAddNewApplyInfo(uId string, flowNo string, info ApplyModel) (err error) {
	//Todo 这里更新如果不能正常操作，需要考虑更改数据结构了
	err = Update(DbName, BjCard, bson.M{"user_id": uId, "car_list.card_list.flowing_no": flowNo}, bson.M{"$set": bson.M{"car_list.$.card_list.apply_info": info}})
	return
}
