// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// Userprivilege is the golang structure for table userprivilege.
type Userprivilege struct {
	UserId       int    `json:"userId"       orm:"user_id"       ` //
	EmailAddress string `json:"emailAddress" orm:"email_address" ` //
	Privilege    string `json:"privilege"    orm:"privilege"     ` //
}
