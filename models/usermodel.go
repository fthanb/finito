package models

import (
	"database/sql"

	"github.com/fthanb/web-pplbo/config"
	"github.com/fthanb/web-pplbo/entities"
)

type UserModel struct {
	db *sql.DB
}

func NewUserModel() *UserModel {
	conn, err := config.DBConn()

	if err != nil {
		panic(err)
	}

	return &UserModel{
		db: conn,
	}
}

func (u UserModel) Where(user *entities.User, fieldName, fieldValue string) error {
	row, err := u.db.Query("select id, nama, nim, password from user where "+fieldName+" = ? limit 1", fieldValue)

	if err != nil {
		return err
	}

	defer row.Close()

	for row.Next() {
		row.Scan(&user.Id, &user.Nama, &user.Nim, &user.Password)
	}

	return nil
}

func (u UserModel) Create(user entities.User) (int64, error) {
	result, err := u.db.Exec("insert into user (nama, nim, password) values(?,?,?)", user.Nama, user.Nim, user.Password)

	if err != nil {
		return 0, err
	}
	lastInsertId, _ := result.LastInsertId()
	return lastInsertId, nil
}
