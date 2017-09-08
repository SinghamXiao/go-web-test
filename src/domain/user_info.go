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

func (userInfo *UserInfo) encode() []byte {
	s, err := json.Marshal(userInfo)
	if err != nil {
		return make([]byte, 0)
	}
	return s
}

func Decode(data []byte) *UserInfo {
	userInfo := new(UserInfo)
	gjson.Unmarshal(data, userInfo)
	return userInfo
}
