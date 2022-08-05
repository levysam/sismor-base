package users

import (
	"github.com/go-jet/jet/v2/mysql"
)

var Users = newUsersTable("sismor", "users", "")

type UsersTable struct {
	mysql.Table

	//Columns
	ID       mysql.ColumnInteger
	Name     mysql.ColumnString
	Password mysql.ColumnString
	Email    mysql.ColumnString

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

// AS creates new UsersTable with assigned alias
func (a UsersTable) AS(alias string) UsersTable {
	return newUsersTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new UsersTable with assigned schema name
func (a UsersTable) FromSchema(schemaName string) UsersTable {
	return newUsersTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new UsersTable with assigned table prefix
func (a UsersTable) WithPrefix(prefix string) UsersTable {
	return newUsersTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new UsersTable with assigned table suffix
func (a UsersTable) WithSuffix(suffix string) UsersTable {
	return newUsersTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newUsersTable(schemaName, tableName, alias string) UsersTable {
	var (
		IDColumn       = mysql.IntegerColumn("id")
		NameColumn     = mysql.StringColumn("name")
		PasswordColumn = mysql.StringColumn("password")
		EmailColumn    = mysql.StringColumn("email")
		allColumns     = mysql.ColumnList{IDColumn, NameColumn, PasswordColumn, EmailColumn}
		mutableColumns = mysql.ColumnList{IDColumn, NameColumn, PasswordColumn, EmailColumn}
	)

	return UsersTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:       IDColumn,
		Name:     NameColumn,
		Password: PasswordColumn,
		Email:    EmailColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
