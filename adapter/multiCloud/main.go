package main

import "fmt"

// 生产实践：多云管理平台

type ProviderType string

const (
	ProviderTypeAliyun  ProviderType = "aliyun"
	ProviderTypeTencent ProviderType = "tencent"
)

// 创建云主机请求参数
type RunInstanceRequest struct {
}

// 创建云主机返回结果
type RunInstanceResponse struct {
}

// 定义各个云厂商云主机统一操作接口
type Provider interface {
	Type() ProviderType
	// 创建云主机
	RunInstance(r *RunInstanceRequest) (*RunInstanceResponse, error)
}

func NewProvider(typ ProviderType) Provider {
	switch typ {
	case ProviderTypeAliyun:
		return NewAliyunProvider(typ)
	case ProviderTypeTencent:
		return NewTencentProvider(typ)
	default:
		panic("不支持的provider")
	}
	return nil
}

func main() {

	var p Provider
	// 阿里云
	p = NewProvider(ProviderTypeAliyun)
	resp, err := p.RunInstance(&RunInstanceRequest{})
	fmt.Println(resp, err)

	// 腾讯云
	p = NewProvider(ProviderTypeTencent)
	resp, err = p.RunInstance(&RunInstanceRequest{})
	fmt.Println(resp, err)
}
