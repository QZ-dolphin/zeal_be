// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Userprivilege is the golang structure of table userprivilege for DAO operations like Where/Data.
type Userprivilege struct {
	g.Meta       `orm:"table:userprivilege, do:true"`
	UserId       interface{} //
	EmailAddress interface{} //
	Privilege    interface{} //
}
