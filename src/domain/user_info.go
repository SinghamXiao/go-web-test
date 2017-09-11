package domain

import (
	"encoding/json"
	"github.com/tidwall/gjson"
)

type UserInfo struct {
	Username string `json:"username"`

	Password string `json:"password"`

	ConfirmPassword string `json:"confirmPassword"`
}

func (this *UserInfo) encode() []byte {
	bytes, err := json.Marshal(this)
	if err != nil {
		return make([]byte, 0)
	}
	return bytes
}

func Decode(data []byte) *UserInfo {
	userInfo := new(UserInfo)
	err := gjson.Unmarshal(data, userInfo)
	if err != nil {
		return nil
	}
	return userInfo
}
