package obs

import (
	"encoding/base64"
	"fmt"
	"os"
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

func TestObs_UploadBase64(t *testing.T) {
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

	o := Obs{}

	path, err := o.UploadBase64(&base64String, ".png")
	if err != nil {
		t.Error("upload base err --> ", err)
	}
	fmt.Println(path)
}
