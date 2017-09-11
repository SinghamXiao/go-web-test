package domain

import (
	"encoding/json"
	"github.com/tidwall/gjson"
)

type RegisterInfo struct {
	Username string `json:"username"`

	Password string `json:"password"`

	ConfirmPassword string `json:"confirmPassword"`
}

func (this *RegisterInfo) encode() []byte {
	bytes, err := json.Marshal(this)
	if err != nil {
		return make([]byte, 0)
	}
	return bytes
}

func DecodeRegisterInfo(data []byte) *RegisterInfo {
	userInfo := new(RegisterInfo)
	err := gjson.Unmarshal(data, userInfo)
	if err != nil {
		return nil
	}
	return userInfo
}
