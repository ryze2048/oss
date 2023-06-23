package obs

import (
	"fmt"
	"github.com/ryze2048/oss/common"
	"time"
)

func (o *Obs) GetPath(info string) string {
	return fmt.Sprintf("%s/%s/%s%s", o.Path, time.Now().Format("2006-01-02"), common.GenerateRandomFilename(16), common.GetFileExtension(info))
}

func (o *Obs) GetResponseContentTypeSuffix(contentType string) string {
	return fmt.Sprintf("%s/%s/%s%s", o.Path, time.Now().Format("2006-01-02"), common.GenerateRandomFilename(16), common.GetFileExtensionToContentType(contentType))
}
