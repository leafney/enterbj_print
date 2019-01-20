package car

import (
	"enterbj_print/models"
	"errors"
	"github.com/Leafney/req"
	"github.com/tidwall/gjson"
	"log"
)

// 获取applyid
func GetEnterBjApply(licenseNo string,flowingNo string) (applyInfo models.ApplyModel,err error) {

	url:="https://enterbj.zhongchebaolian.com/enterbj_print/filetransfer/print/getEnterBjApply"

	header:=req.Header{
		"Accept": "application/json, text/javascript, */*; q=0.01",
		"Accept-Encoding":"gzip, deflate, br",
		"Accept-Language":"zh-CN,zh;q=0.9",
		"Content-Type":"application/x-www-form-urlencoded; charset=UTF-8",
		"Origin":"https://enterbj.zhongchebaolian.com",
		"Referer":"https://enterbj.zhongchebaolian.com/enterbj_print/jsp/print.jsp",
		"User-Agent":"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36",
		"X-Requested-With":"XMLHttpRequest",
	}
	param:=req.Param{
		"licenseNo":licenseNo,
		"flowingNo":flowingNo,
	}
	req.Debug=true
	r,err:=req.Post(url,header,param)
	if err!=nil{
		log.Fatal(err)
	}

	respStr,err:=r.ToString()
	if err!=nil{
		log.Fatal(err)
	}

	rStatus:= gjson.Get(respStr,"status").String()

	if rStatus=="1"{
	//	请求成功
		applyInfo=models.ApplyModel{
			ApplyId:gjson.Get(respStr,"applyid").String(),
			DriverName:gjson.Get(respStr,"drivername").String(),
			LicenseNo:gjson.Get(respStr,"licenseno").String(),
			EndTime:gjson.Get(respStr,"endtime").String(),
			StartTime:gjson.Get(respStr,"starttime").String(),
			CreateTime:gjson.Get(respStr,"createtime").String(),
			Address:gjson.Get(respStr,"address").String(),
			PaperId:gjson.Get(respStr,"paperid").String(),
			CarType:gjson.Get(respStr,"cartype").String(),
		}

		return applyInfo,nil
	}else{
	//	请求失败
		err=errors.New("未查询到进京证信息")
	}

	return applyInfo,err
}

// 获取进京证图片
func GetBjImgByApplyid(applyInfo models.ApplyModel) (imgPath string,err error) {
	url:="https://enterbj.zhongchebaolian.com/enterbj_print/filetransfer/print/generateEnterBjImgByApplyid"

	header:=req.Header{
		"Accept":"*/*",
		"Accept-Encoding":"gzip, deflate, br",
		"Accept-Language":"zh-CN,zh;q=0.9",
		"Content-Type":"application/x-www-form-urlencoded; charset=UTF-8",
		"Origin":"https://enterbj.zhongchebaolian.com",
		"Referer":"https://enterbj.zhongchebaolian.com/enterbj_print/jsp/print.jsp",
		"User-Agent":"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36",
		"X-Requested-With":"XMLHttpRequest",
	}

	param:=req.Param{
		"licenseno":applyInfo.LicenseNo,
		"drivername":applyInfo.DriverName,
		"starttime":applyInfo.StartTime,
		"endtime":applyInfo.EndTime,
		"applyid":applyInfo.ApplyId,
		"createtime":applyInfo.CreateTime,
		"address":applyInfo.Address,
		"paperid":applyInfo.PaperId,
		"cartype":applyInfo.CarType,
	}
	req.Debug=true

	r,err:= req.Post(url,header,param)
	if err!=nil{
		log.Fatal(err)
	}

	respStr,err:=r.ToString()
	if err!=nil{
		log.Fatal(err)
	}

	resCode:=gjson.Get(respStr,"rescode").String()
	if resCode=="200"{
		imgPath=gjson.Get(respStr,"imgPath").String()
		return imgPath,nil
	}else{
	err=errors.New("返回失败")
	}
	return imgPath,err
}

// 下载图片
func DownBjImg(imgPath string)  {
	r,_:=req.Get(imgPath)
	r.ToFile("")
}

