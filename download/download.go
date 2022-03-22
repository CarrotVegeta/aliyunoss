package download

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type Download interface {
	Downloader() error
	SetClient(client *oss.Client)
}
