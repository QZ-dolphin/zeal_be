// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Userdata is the golang structure of table userdata for DAO operations like Where/Data.
type Userdata struct {
	g.Meta        `orm:"table:userdata, do:true"`
	UserId        interface{} //
	UserName      interface{} //
	Password      interface{} //
	EmailAddress  interface{} //
	IpAddress     interface{} //
	AvatarUrl     interface{} //
	CreatedTime   *gtime.Time //
	LastLoginTime *gtime.Time //
	DayName       interface{} //
	DayValue      interface{} //
}
