package car

import "testing"

func TestA(t *testing.T)  {

	applyInfo,err:= GetEnterBjApply("","")
	if err!=nil{
		t.Log(err)
	}

	bjImg,err:= GetBjImgByApplyid(applyInfo)
	if err!=nil{
		t.Log(err)
	}

	t.Log(bjImg)

}