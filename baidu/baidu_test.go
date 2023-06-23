package baidu

import (
	"fmt"
	"testing"
)

func TestBaiduOSS_UploadLocal(t *testing.T) {
	baiduOss := BaiduOSS{}
	path, err := baiduOss.UploadLocal("dAmZNXqEdXTvUvSlyrPTLtvCmjCkTp.png")
	if err != nil {
		t.Error("err --> ", err)
	}
	fmt.Println(path)
}

func TestBaiduOSS_UploadLink(t *testing.T) {
	baiduOss := BaiduOSS{}
	path, err := baiduOss.UploadLink("https://img.alicdn.com/imgextra/i3/O1CN01SF1OYM1YFKVYx2SJt_!!6000000003029-2-tps-224-224.png")
	if err != nil {
		t.Error("err --> ", err)
	}
	fmt.Println(path)
}

func TestBaiduOSS_GetTemporaryLink(t *testing.T) {
	baiduOss := BaiduOSS{}
	path, err := baiduOss.GetTemporaryLink("2023-06-20/33a502b91607c4207aef41a32867e9b3.png", 360)
	if err != nil {
		t.Error("err --> ", err)
	}
	fmt.Println(path)
}
