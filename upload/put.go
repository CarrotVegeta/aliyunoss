package upload

import (
	"errors"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

const (
	PutObjectFromFile = iota + 1
)

type Put struct {
	client     *oss.Client
	BucketName string `json:"bucket_name"`
	ObjectKey  string `json:"object_key"`
	FilePath   string `json:"file_path"`
	PutType    int    `json:"put_type"`
}

func (p *Put) SetClient(client *oss.Client) {
	p.client = client
}
func (p *Put) handleError(err string) error {
	return errors.New(fmt.Sprintf("pubobject err : %v", err))
}
func (p *Put) Uploader() error {
	switch p.PutType {
	case PutObjectFromFile:
		if err := p.putFile(); err != nil {
			return err
		}
	}
	return nil
}
func (p *Put) putFile() error {
	// 填写存储空间名称，例如examplebucket。
	bucket, err := p.client.Bucket(p.BucketName)
	if err != nil {
		return errors.New(fmt.Sprintf("putfile error : %v", err))
	}
	// 依次填写Object的完整路径（objectKey:例如exampledir/exampleobject.txt）和本地文件的完整路径（filePath:例如D:\\localpath\\examplefile.txt）。
	err = bucket.PutObjectFromFile(p.ObjectKey, p.FilePath)
	if err != nil {
		return errors.New(fmt.Sprintf("putfile error : %v", err))
	}
	return nil
}
