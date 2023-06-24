package aliyun

import (
	"fmt"
	"github.com/ryze2048/oss/common"
	"path/filepath"
	"testing"
	"time"
)

func TestAliyunOSS_UploadLocal(t *testing.T) {
	dirPath := generateDirectory("")
	fmt.Println("目录路径:", dirPath)
	// generateDirectoryV2("test", "s.png")
}

func TestUrl(t *testing.T) {
	var oss = AliyunOSS{}
	key, err := oss.UploadLink("https://main.qcloudimg.com/raw/6960df3fe22c8fd12c415e86a1188d5f.png")
	if err != nil {
		t.Error("err --> ", err)
	}
	fmt.Println(key)
}

func TestAliyunOSS_DownloadByte(t *testing.T) {
	var oss = AliyunOSS{}
	var byteInfo = make([]byte, 0)
	var err error

	if byteInfo, err = oss.DownloadByte("test/2023-06-19/545fff172e80e525288a8e394847caf2.png"); err != nil {
		t.Error("err --> ", err)
	}
	fmt.Println(byteInfo)
}

func TestAliyunOSS_GetTemporaryLink(t *testing.T) {
	var oss = AliyunOSS{}
	var path string
	var err error
	if path, err = oss.GetTemporaryLink("test/2023-06-19/545fff172e80e525288a8e394847caf2.png", 5*60); err != nil {
		t.Error("err --> ", err)
	}
	fmt.Println(path)
}

func generateDirectory(x string) string {
	dirPath := filepath.Join(x, "y")
	return filepath.Join(dirPath, "z.png")
}

func generateDirectoryV2(BasePath string, info string) {
	var dirPath = filepath.Join(BasePath, time.Now().Format("2006-01-02"))
	dirPath = filepath.Join(dirPath, common.GenerateRandomFilename(16)+common.GetFileExtension(info))
	fmt.Println(dirPath)
}

func TestAliyunOSS_GetPath(t *testing.T) {
	a := AliyunOSS{}
	fmt.Println(a.GetPath("info.png"))
}
