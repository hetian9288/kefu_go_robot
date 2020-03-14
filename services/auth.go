package services

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"kefu_go_robot/conf"
	"kefu_server/utils"
)

// AuthToken token
var AuthToken string

// AuthTokenRepository struct
type AuthTokenRepository struct{}

// GetAuthTokenRepositoryInstance get instance
func GetAuthTokenRepositoryInstance() *AuthTokenRepository {
	instance := new(AuthTokenRepository)
	return instance
}

// FetchToken auth
func (r *AuthTokenRepository) FetchToken() {
	config := new(conf.Cionfigs).GetConfigs()
	api := "/v1/auth/token/"
	path := config.GatewayHost + api
	AuthToken = ""
	// MD5
	m5 := md5.New()
	m5.Write([]byte(config.MiAppID + config.MiAppKey + config.MiAppSecret))
	secret := hex.EncodeToString(m5.Sum(nil))
	var request = map[string]string{}
	request["app_secret"] = secret
	response := utils.HTTPRequest(path, "POST", request, "")
	if response.Code != 200 {
		fmt.Println(response.Message)
		return
	}
	AuthToken = response.Data.(string)
}
