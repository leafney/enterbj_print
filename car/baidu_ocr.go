package car

import (
	"enterbj_print/models"
	"errors"
	"github.com/Leafney/req"
	"github.com/tidwall/gjson"
	"log"
)

/*
百度图像识别
*/

/*
获取请求的AccessToken
*/
func BaiduGetAccessToken(clientId string, clientSecret string) (t models.TokenInfo, err error) {
	url := "https://aip.baidubce.com/oauth/2.0/token"

	header := req.Header{
		"Content-Type": "application/json; charset=UTF-8",
	}
	param := req.Param{
		"grant_type":    "client_credentials",
		"client_id":     clientId,
		"client_secret": clientSecret,
	}
	req.Debug = true
	r, err := req.Post(url, header, param)
	if err != nil {
		log.Fatal(err)
	}

	respStr, err := r.ToString()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(respStr)

	rAccessToken := gjson.Get(respStr, "access_token").String()
	if rAccessToken != "" {
		//	请求正常
		rExpires := gjson.Get(respStr, "expires_in").Int()
		t.AccessToken = rAccessToken
		t.ExpiresIn = rExpires
		return t, nil
	} else {
		//	请求异常
		err = errors.New(gjson.Get(respStr, "error_description").String())
	}
	return
}

/*

 */
func BaiduOcrText(imgPath string, token string) {

}
