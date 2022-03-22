package oss

import (
	"github.com/CarrotVegeta/aliyunoss/upload"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type OssClient struct {
	Client *oss.Client
}
type Config struct {
	EndPoint        string
	AccessKeyId     string
	AccessKeySecret string
}

func NewClient(c *Config) (*OssClient, error) {
	// 创建OSSClient实例。
	// yourEndpoint填写Bucket对应的Endpoint，以华东1（杭州）为例，填写为https://oss-cn-hangzhou.aliyuncs.com。其它Region请按实际情况填写。
	// 阿里云账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM用户进行API访问或日常运维，请登录RAM控制台创建RAM用户。
	client, err := oss.New(c.EndPoint, c.AccessKeyId, c.AccessKeySecret)
	if err != nil {
		return nil, err
	}
	oc := &OssClient{Client: client}
	return oc, nil
}

func (oc *OssClient) Uploader(u upload.Upload) error {
	u.SetClient(oc.Client)
	err := u.Uploader()
	if err != nil {
		return err
	}
	return nil
}
func (oc *OssClient) CreateBucket(bucketName string) error {
	// 创建名为examplebucket的存储空间，并设置存储类型为低频访问oss.StorageIA、读写权限ACL为公共读oss.ACLPublicRead、数据容灾类型为同城冗余存储oss.RedundancyZRS。
	err := oc.Client.CreateBucket(bucketName, oss.StorageClass(oss.StorageIA), oss.ACL(oss.ACLPublicRead), oss.RedundancyType(oss.RedundancyZRS))
	if err != nil {
		return err
	}
	return nil
}
func (oc *OssClient) IsExist(bucketName string) (bool, error) {
	// 判断存储空间是否存在。
	isExist, err := oc.Client.IsBucketExist(bucketName)
	if err != nil {
		return false, err
	}
	return isExist, nil
}
