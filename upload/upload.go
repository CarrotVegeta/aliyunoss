package upload

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type Upload interface {
	Uploader() error
	SetClient(client *oss.Client)
}
