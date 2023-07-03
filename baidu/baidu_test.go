package baidu

import (
	"encoding/base64"
	"fmt"
	"os"
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

func TestBaiduOSS_UploadBase64(t *testing.T) {
	filePath := "dAmZNXqEdXTvUvSlyrPTLtvCmjCkTp.png" // 替换为你的本地文件路径

	// 读取本地文件
	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// 将文件内容转换为 Base64 字符串
	base64String := base64.StdEncoding.EncodeToString(fileBytes)
	fmt.Println("Base64 encoded string:", base64String)
	baiduOss := BaiduOSS{}
	path, err := baiduOss.UploadBase64(&base64String, ".png")
	if err != nil {
		t.Error("upload base err --> ", err)
	}
	fmt.Println(path)
}
