//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/mysql"
)

var Staff = newStaffTable("sakila", "staff", "")

type staffTable struct {
	mysql.Table

	// Columns
	StaffID    mysql.ColumnInteger
	FirstName  mysql.ColumnString
	LastName   mysql.ColumnString
	AddressID  mysql.ColumnInteger
	Picture    mysql.ColumnString
	Email      mysql.ColumnString
	StoreID    mysql.ColumnInteger
	Active     mysql.ColumnBool
	Username   mysql.ColumnString
	Password   mysql.ColumnString
	LastUpdate mysql.ColumnTimestamp

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type StaffTable struct {
	staffTable

	NEW staffTable
}

// AS creates new StaffTable with assigned alias
func (a StaffTable) AS(alias string) *StaffTable {
	return newStaffTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new StaffTable with assigned schema name
func (a StaffTable) FromSchema(schemaName string) *StaffTable {
	return newStaffTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new StaffTable with assigned table prefix
func (a StaffTable) WithPrefix(prefix string) *StaffTable {
	return newStaffTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new StaffTable with assigned table suffix
func (a StaffTable) WithSuffix(suffix string) *StaffTable {
	return newStaffTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newStaffTable(schemaName, tableName, alias string) *StaffTable {
	return &StaffTable{
		staffTable: newStaffTableImpl(schemaName, tableName, alias),
		NEW:        newStaffTableImpl("", "new", ""),
	}
}

func newStaffTableImpl(schemaName, tableName, alias string) staffTable {
	var (
		StaffIDColumn    = mysql.IntegerColumn("staff_id")
		FirstNameColumn  = mysql.StringColumn("first_name")
		LastNameColumn   = mysql.StringColumn("last_name")
		AddressIDColumn  = mysql.IntegerColumn("address_id")
		PictureColumn    = mysql.StringColumn("picture")
		EmailColumn      = mysql.StringColumn("email")
		StoreIDColumn    = mysql.IntegerColumn("store_id")
		ActiveColumn     = mysql.BoolColumn("active")
		UsernameColumn   = mysql.StringColumn("username")
		PasswordColumn   = mysql.StringColumn("password")
		LastUpdateColumn = mysql.TimestampColumn("last_update")
		allColumns       = mysql.ColumnList{StaffIDColumn, FirstNameColumn, LastNameColumn, AddressIDColumn, PictureColumn, EmailColumn, StoreIDColumn, ActiveColumn, UsernameColumn, PasswordColumn, LastUpdateColumn}
		mutableColumns   = mysql.ColumnList{FirstNameColumn, LastNameColumn, AddressIDColumn, PictureColumn, EmailColumn, StoreIDColumn, ActiveColumn, UsernameColumn, PasswordColumn, LastUpdateColumn}
	)

	return staffTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		StaffID:    StaffIDColumn,
		FirstName:  FirstNameColumn,
		LastName:   LastNameColumn,
		AddressID:  AddressIDColumn,
		Picture:    PictureColumn,
		Email:      EmailColumn,
		StoreID:    StoreIDColumn,
		Active:     ActiveColumn,
		Username:   UsernameColumn,
		Password:   PasswordColumn,
		LastUpdate: LastUpdateColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
