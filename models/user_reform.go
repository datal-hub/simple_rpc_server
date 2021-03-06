// Code generated by gopkg.in/reform.v1. DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type userTableType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *userTableType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("users").
func (v *userTableType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *userTableType) Columns() []string {
	return []string{"uuid", "login", "dttm"}
}

// NewStruct makes a new struct for that view or table.
func (v *userTableType) NewStruct() reform.Struct {
	return new(User)
}

// NewRecord makes a new record for that table.
func (v *userTableType) NewRecord() reform.Record {
	return new(User)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *userTableType) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// UserTable represents users view or table in SQL database.
var UserTable = &userTableType{
	s: parse.StructInfo{Type: "User", SQLSchema: "", SQLName: "users", Fields: []parse.FieldInfo{{Name: "Uuid", Type: "string", Column: "uuid"}, {Name: "Login", Type: "string", Column: "login"}, {Name: "Dttm", Type: "time.Time", Column: "dttm"}}, PKFieldIndex: 0},
	z: new(User).Values(),
}

// String returns a string representation of this struct or record.
func (s User) String() string {
	res := make([]string, 3)
	res[0] = "Uuid: " + reform.Inspect(s.Uuid, true)
	res[1] = "Login: " + reform.Inspect(s.Login, true)
	res[2] = "Dttm: " + reform.Inspect(s.Dttm, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *User) Values() []interface{} {
	return []interface{}{
		s.Uuid,
		s.Login,
		s.Dttm,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *User) Pointers() []interface{} {
	return []interface{}{
		&s.Uuid,
		&s.Login,
		&s.Dttm,
	}
}

// View returns View object for that struct.
func (s *User) View() reform.View {
	return UserTable
}

// Table returns Table object for that record.
func (s *User) Table() reform.Table {
	return UserTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *User) PKValue() interface{} {
	return s.Uuid
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *User) PKPointer() interface{} {
	return &s.Uuid
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *User) HasPK() bool {
	return s.Uuid != UserTable.z[UserTable.s.PKFieldIndex]
}

// SetPK sets record primary key.
func (s *User) SetPK(pk interface{}) {
	if i64, ok := pk.(int64); ok {
		s.Uuid = string(i64)
	} else {
		s.Uuid = pk.(string)
	}
}

// check interfaces
var (
	_ reform.View   = UserTable
	_ reform.Struct = (*User)(nil)
	_ reform.Table  = UserTable
	_ reform.Record = (*User)(nil)
	_ fmt.Stringer  = (*User)(nil)
)

func init() {
	parse.AssertUpToDate(&UserTable.s, new(User))
}
