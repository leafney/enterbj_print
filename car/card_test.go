package car

import "testing"

func TestA(t *testing.T) {

	applyInfo, err := GetEnterBjApply("", "") //flowingNo 不要忘记前面的A
	if err != nil {
		// 没有获取到查询序列号所对应的进京证信息，直接退出
		t.Log(err)
		return
	}
	t.Log("applyInfo: ", applyInfo)

	bjImg, err := GetBjImgByApplyid(applyInfo)
	if err != nil {
		t.Log(err)
	}

	t.Log("bjImg: ", bjImg)
	DownloadBjImg(bjImg, "../data/bbb.jpg")

	t.Log("ok")
}
