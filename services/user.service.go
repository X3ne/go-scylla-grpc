package services

import (
	"github.com/bwmarrin/snowflake"
	"github.com/scylladb/gocqlx/qb"
	"github.com/scylladb/gocqlx/table"
)

var usersMetadata = table.Metadata{
	Name:    "users",
	Columns: []string{"id", "username", "password"},
	PartKey: []string{"id"},
	SortKey: []string{"username"},
}

type User struct {
	Id    string `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
}

var personTable = table.New(usersMetadata)

func CreateUser(user *User) error {
	node, err := snowflake.NewNode(1)
	if err != nil {
		return err
	}

	user.Id = node.Generate().String()

	if err := session.Query(personTable.Insert()).BindStruct(user).ExecRelease(); err != nil {
		return err
	}

	return nil
}

func GetUserById(id string) (*User, error) {
	stmt, names := qb.Select("users").
		Where(qb.Eq("id")).
		AllowFiltering().
		Limit(1).
		ToCql()

	var user User
	if err := session.Query(stmt, names).BindMap(qb.M{"id": id}).GetRelease(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUserByUsername(username string) (*User, error) {
	stmt, names := qb.Select("users").
		Where(qb.Eq("username")).
		AllowFiltering().
		Limit(1).
		ToCql()

	var user User
	if err := session.Query(stmt, names).BindMap(qb.M{"username": username}).GetRelease(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUsers() ([]*User, error) {
	stmt, names := qb.Select("users").
		AllowFiltering().
		Columns("id", "username").
		ToCql()

	var users []*User
	if err := session.Query(stmt, names).SelectRelease(&users); err != nil {
		return nil, err
	}

	return users, nil
}

func UpdateUser(user *User) error {
	if err := session.Query(personTable.Update()).BindStruct(user).ExecRelease(); err != nil {
		return err
	}

	return nil
}

func DeleteUser(id string) error {
	if err := session.Query(personTable.Delete()).BindStruct(User{Id: id}).ExecRelease(); err != nil {
		return err
	}

	return nil
}
