package main

type AliyunProvider struct {
	typ ProviderType
}

func NewAliyunProvider(typ ProviderType) *AliyunProvider {
	return &AliyunProvider{typ: typ}
}

func (a AliyunProvider) Type() ProviderType {
	return a.typ
}

func (a AliyunProvider) RunInstance(r *RunInstanceRequest) (*RunInstanceResponse, error) {
	// 包装阿里云的sdk
	return nil, nil
}
