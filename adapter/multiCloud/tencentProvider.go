package main

type TencentProvider struct {
	typ ProviderType
}

func NewTencentProvider(typ ProviderType) *TencentProvider {
	return &TencentProvider{typ: typ}
}

func (t *TencentProvider) Type() ProviderType {
	return t.typ
}

func (t *TencentProvider) RunInstance(r *RunInstanceRequest) (*RunInstanceResponse, error) {
	// 包装腾讯云的sdk
	return nil, nil
}
