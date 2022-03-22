package upload

import (
	"errors"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
)

type SplitUpload struct {
	client        *oss.Client
	BucketName    string `json:"bucket_name"`
	LocalFileName string `json:"local_file_name"`
	ObjectName    string `json:"object_name"`
}

func (su *SplitUpload) SetClient(client *oss.Client) {
	su.client = client
}
func (su *SplitUpload) handleError(err string) error {
	return errors.New(fmt.Sprintf("split upload err : %v", err))
}
func (su *SplitUpload) Uploader() error {
	// 获取存储空间。
	bucket, err := su.client.Bucket(su.BucketName)
	if err != nil {
		return su.handleError(err.Error())
	}
	// 将本地文件分片，且分片数量指定为3。
	chunks, err := oss.SplitFileByPartNum(su.LocalFileName, 3)
	fd, err := os.Open(su.LocalFileName)
	defer fd.Close()

	// 指定过期时间。
	//expires := time.Date(2049, time.January, 10, 23, 0, 0, 0, time.UTC)
	// 如果需要在初始化分片时设置请求头，请参考以下示例代码。
	options := []oss.Option{
		oss.MetadataDirective(oss.MetaReplace),
		//oss.Expires(expires),
		// 指定该Object被下载时的网页缓存行为。
		// oss.CacheControl("no-cache"),
		// 指定该Object被下载时的名称。
		// oss.ContentDisposition("attachment;filename=FileName.txt"),
		// 指定该Object的内容编码格式。
		// oss.ContentEncoding("gzip"),
		// 指定对返回的Key进行编码，目前支持URL编码。
		// oss.EncodingType("url"),
		// 指定Object的存储类型。
		// oss.ObjectStorageClass(oss.StorageStandard),
	}

	// 步骤1：初始化一个分片上传事件，并指定存储类型为标准存储。
	imur, err := bucket.InitiateMultipartUpload(su.ObjectName, options...)
	// 步骤2：上传分片。
	var parts []oss.UploadPart
	for _, chunk := range chunks {
		fd.Seek(chunk.Offset, os.SEEK_SET)
		// 调用UploadPart方法上传每个分片。
		part, err := bucket.UploadPart(imur, fd, chunk.Size, chunk.Number)
		if err != nil {
			return su.handleError(err.Error())
		}
		parts = append(parts, part)
	}

	// 指定Object的读写权限为公共读，默认为继承Bucket的读写权限。
	objectAcl := oss.ObjectACL(oss.ACLPublicRead)

	// 步骤3：完成分片上传，指定文件读写权限为公共读。
	_, err = bucket.CompleteMultipartUpload(imur, parts, objectAcl)
	if err != nil {
		return su.handleError(err.Error())
	}
	return nil
}
