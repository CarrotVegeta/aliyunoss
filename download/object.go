package download

import (
	"errors"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type Object struct {
	client     *oss.Client
	BucketName string
	ObjectKey  string
	FilePath   string
}

func (o *Object) SetClient(client *oss.Client) {
	o.client = client
}
func (o *Object) Downloader() error {
	// 填写Bucket名称，例如examplebucket。
	bucket, err := o.client.Bucket(o.BucketName)
	if err != nil {
		return errors.New(fmt.Sprintf("get object to file err : %v", err))
	}
	// 下载文件到本地文件，并保存到指定的本地路径中。如果指定的本地文件存在会覆盖，不存在则新建。
	// 如果未指定本地路径，则下载后的文件默认保存到示例程序所属项目对应本地路径中。
	// 依次填写Object完整路径（例如exampledir/exampleobject.txt）和本地文件的完整路径(例如D:\\localpath\\examplefile.txt)。Object完整路径中不能包含Bucket名称。
	err = bucket.GetObjectToFile(o.ObjectKey, o.FilePath)
	if err != nil {
		return errors.New(fmt.Sprintf("get object to file err : %v", err))
	}
	return nil
}
