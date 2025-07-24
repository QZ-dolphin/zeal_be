// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Userdata is the golang structure for table userdata.
type Userdata struct {
	UserId        int         `json:"userId"        orm:"user_id"         ` //
	UserName      string      `json:"userName"      orm:"user_name"       ` //
	Password      string      `json:"password"      orm:"password"        ` //
	EmailAddress  string      `json:"emailAddress"  orm:"email_address"   ` //
	IpAddress     string      `json:"ipAddress"     orm:"ip_address"      ` //
	AvatarUrl     string      `json:"avatarUrl"     orm:"avatar_url"      ` //
	CreatedTime   *gtime.Time `json:"createdTime"   orm:"created_time"    ` //
	LastLoginTime *gtime.Time `json:"lastLoginTime" orm:"last_login_time" ` //
	DayName       string      `json:"dayName"       orm:"day_name"        ` //
	DayValue      string      `json:"dayValue"      orm:"day_value"       ` //
}
