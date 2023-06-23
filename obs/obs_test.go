package obs

import (
	"fmt"
	"testing"
)

func TestObs_UploadLocal(t *testing.T) {
	o := Obs{}
	var path string
	var err error
	path, err = o.UploadLocal(`dAmZNXqEdXTvUvSlyrPTLtvCmjCkTp.png`)
	if err != nil {
		t.Error("upload err --> ", err)
	}
	fmt.Println(path)
}

func TestObs_UploadLink(t *testing.T) {
	o := Obs{}
	var path string
	var err error
	path, err = o.UploadLink(`https://fanyiapp.cdn.bcebos.com/cms/image/7b9655bb7a32a6c1a472c3591428855d.jpg`)
	if err != nil {
		t.Error("upload err --> ", err)
	}
	fmt.Println(path)
}

func TestObs_GetTemporaryLink(t *testing.T) {
	o := Obs{}
	var path string
	var err error
	path, err = o.GetTemporaryLink(`test/2023-06-20/81210a4a9a007a351e93c8e2bca57af3.jpg`, 3600)
	if err != nil {
		t.Error("upload err --> ", err)
	}
	fmt.Println(path)
}
