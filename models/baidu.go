package models

import "github.com/globalsign/mgo/bson"

type BaiduAccessToken struct {
	Id           bson.ObjectId `bson:"id"`
	UserId       string        `bson:"user_id" json:"userId"`
	ClientId     string        `bson:"client_id" json:"clientId"`
	ClientSecret string        `bson:"client_secret" json:"clientSecret"`
	AccessToken  string        `bson:"access_token" json:"accessToken"`
	ExpiresTime  int64         `bson:"expires_time" json:"expiresTime"` // 失效时间戳
}

/*
新增接口key和secret
*/
func BaiduAddTokenInfo(t BaiduAccessToken) (err error) {
	t.Id = bson.NewObjectId()
	err = Insert(DbName, BdToken, t)
	return
}

/*
更新token及过期时间
*/
func BaiduUpToken(uId string, token string, exTime int64) (err error) {
	/*
			Todo 这个逻辑写在外面：
		   根据用户id查询该用户下的token信息，
		   判断token是否即将过期，
		   如果即将过期，重新申请token
	*/
	err = Update(DbName, BdToken, bson.M{"user_id": uId}, bson.M{"$set": bson.M{"access_token": token, "expires_time": exTime}})
	return err
}

/*
更新API Key和Secret Key
*/
func BaiduUpKeyAndSecret(uId string, cId string, cSecret string) (err error) {
	err = Update(DbName, BdToken, bson.M{"user_id": uId}, bson.M{"$set": bson.M{"client_id": cId, "client_secret": cSecret}})
	return
}

/*
获取用户绑定的Token信息
*/
func BaiduGetTokenInfo(uId string) (t BaiduAccessToken, err error) {
	err = FindOne(DbName, BdToken, bson.M{"user_id": uId}, nil, &t)
	if err == nil && t.UserId != "" {
		return t, nil
	} else {
		return t, err
	}
}
