package services

import (
	"log"
	"reflect"

	pool "github.com/X3ne/go-scylla-grpc/services/scylla-workers"

	"github.com/bwmarrin/snowflake"
	"github.com/scylladb/gocqlx/v2/qb"
	"github.com/scylladb/gocqlx/v2/table"
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

	hash, err := HashPassword(user.Password, &params{
		memory:      64 * 1024,
		iterations:  3,
		parallelism: 2,
		keyLength:   32,
		saltLength:  16,
	})
	if err != nil {
		return err
	}

	user.Password = string(hash)

	if err := Session.Query(personTable.Insert()).BindStruct(user).ExecRelease(); err != nil {
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
	q := Session.Query(stmt, names).BindMap(qb.M{"id": id})

	log.Println(q.String())

	pool.HandleQuery(q, reflect.TypeOf(User{}))

	if err := q.GetRelease(&user); err != nil {
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
	q := Session.Query(stmt, names).BindMap(qb.M{"username": username})

	log.Println(q.String())

	pool.HandleQuery(q, reflect.TypeOf(User{}))

	if err := q.GetRelease(&user); err != nil {
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
	q := Session.Query(stmt, names)

	pool.HandleQuery(q, reflect.TypeOf([]*User{}))

	if err := q.SelectRelease(&users); err != nil {
		return nil, err
	}

	return users, nil
}

func UpdateUser(user *User) error {
	if err := Session.Query(personTable.Update()).BindStruct(user).ExecRelease(); err != nil {
		return err
	}

	return nil
}

func DeleteUser(id string) error {
	if err := Session.Query(personTable.Delete()).BindStruct(User{Id: id}).ExecRelease(); err != nil {
		return err
	}

	return nil
}
