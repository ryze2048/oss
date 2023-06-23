package baidu

import (
	"fmt"
	"github.com/ryze2048/oss/common"
	"time"
)

func (b *BaiduOSS) GetPath(info string) string {
	return fmt.Sprintf("%s/%s/%s%s", b.BasePath, time.Now().Format("2006-01-02"), common.GenerateRandomFilename(16), common.GetFileExtension(info))
}

func (b *BaiduOSS) GetResponseContentTypeSuffix(contentType string) string {
	return fmt.Sprintf("%s/%s/%s%s", b.BasePath, time.Now().Format("2006-01-02"), common.GenerateRandomFilename(16), common.GetFileExtensionToContentType(contentType))
}
