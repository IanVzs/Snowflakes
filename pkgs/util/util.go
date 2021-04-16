package util

import "github.com/IanVzs/Snowflakes/pkgs/setting"

// Setup Initialize the util
func init() {
	jwtSecret = []byte(setting.AppSetting.JwtSecret)
}
