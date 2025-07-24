// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserdataDao is the data access object for table userdata.
type UserdataDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns UserdataColumns // columns contains all the column names of Table for convenient usage.
}

// UserdataColumns defines and stores column names for table userdata.
type UserdataColumns struct {
	UserId        string //
	UserName      string //
	Password      string //
	EmailAddress  string //
	IpAddress     string //
	AvatarUrl     string //
	CreatedTime   string //
	LastLoginTime string //
	DayName       string //
	DayValue      string //
}

// userdataColumns holds the columns for table userdata.
var userdataColumns = UserdataColumns{
	UserId:        "user_id",
	UserName:      "user_name",
	Password:      "password",
	EmailAddress:  "email_address",
	IpAddress:     "ip_address",
	AvatarUrl:     "avatar_url",
	CreatedTime:   "created_time",
	LastLoginTime: "last_login_time",
	DayName:       "day_name",
	DayValue:      "day_value",
}

// NewUserdataDao creates and returns a new DAO object for table data access.
func NewUserdataDao() *UserdataDao {
	return &UserdataDao{
		group:   "default",
		table:   "userdata",
		columns: userdataColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UserdataDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UserdataDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *UserdataDao) Columns() UserdataColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UserdataDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UserdataDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UserdataDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
