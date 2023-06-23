package qiniu

import (
	"fmt"
	"github.com/ryze2048/oss/common"
	"time"
)

func (q *Qiniu) GetPath(info string) string {
	return fmt.Sprintf("%s/%s/%s%s", q.BasePath, time.Now().Format("2006-01-02"), common.GenerateRandomFilename(16), common.GetFileExtension(info))
}
func (q *Qiniu) GetResponseContentTypeSuffix(contentType string) string {
	return fmt.Sprintf("%s/%s/%s%s", q.BasePath, time.Now().Format("2006-01-02"), common.GenerateRandomFilename(16), common.GetFileExtensionToContentType(contentType))
}
