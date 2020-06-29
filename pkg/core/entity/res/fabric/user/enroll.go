package user

import "github.com/chcp/bsn-sdk-go/pkg/core/entity/base"

type EnrollResData struct {
	base.BaseResModel
	Body *EnrollResDataBody `json:"body"`
}

type EnrollResDataBody struct {
	Cert string `json:"cert"`
}

func (f *EnrollResData) GetEncryptionValue() string {
	if f.Body == nil {
		return f.GetBaseEncryptionValue()
	}

	return f.GetBaseEncryptionValue() + f.Body.Cert
}