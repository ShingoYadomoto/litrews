package model

import (
	"time"

	"github.com/ShingoYadomoto/litrews/src/db"
	"github.com/ShingoYadomoto/litrews/src/helper"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

type User struct {
	ID        int       `db:"id" json:"id" form:"id"`
	Name      string    `db:"name" json:"name" form:"name"`
	Password  string    `db:"password" json:"password" form:"password"`
	Email     string    `db:"email" json:"email" form:"email"`
	CreatedAt time.Time `db:"created_at" json:"create_at"`
	UpdatedAt time.Time `db:"updated_at" json:"update_at"`
	Topics    []*Topic  `db:"-"`
}

type UserModel struct {
	db db.AbstractDB
}

func NewUserModel(db db.AbstractDB) *UserModel {
	return &UserModel{
		db: db,
	}
}

func (m *UserModel) GetUserByID(id int) (user *User, err error) {
	user = new(User)

	q := `SELECT
            user.*
          FROM
            user
          WHERE
            user.id = ?`
	err = m.db.Get(user, q, id)
	if err != nil {
		err = errors.Wrap(err, "couldn't get user")
		return
	}

	// ManyToMany, HasMany のリレーションはjoinとは別で取ってくるのが良さそう
	topicModel := NewTopicModel(m.db)
	user.Topics, err = topicModel.GetTopicsByUserID(id)

	return
}

func (m *UserModel) GetUserAuthInfoByID(id int) (user *User, err error) {
	user = new(User)

	q := `SELECT
            *
          FROM
            user
          WHERE
            id = ?`
	err = m.db.Get(user, q, id)
	if err != nil {
		err = errors.Wrap(err, "couldn't get userAuthInfo")
		return
	}

	return
}

func (m *UserModel) CreateUser(user *User) (err error) {
	user.Password, err = helper.ToHash(user.Password)
	if err != nil {
		err = errors.Wrap(err, "password is invalid")
		return
	}

	q := `INSERT INTO
            user (name, email, password)
          VALUES
            (:name, :email, :password)`
	_, err = m.db.NamedExec(q, user)
	if err != nil {
		err = errors.Wrap(err, "couldn't create user")
		return
	}

	return
}

func (m *UserModel) UpdateUser(user *User) (err error) {
	q := `UPDATE
            user
          SET
            name = :name, email = :email
          WHERE
            id = :id`
	_, err = m.db.NamedExec(q, user)
	if err != nil {
		err = errors.Wrap(err, "couldn't update user")
		return
	}

	return
}

func (m *UserModel) GetUserByAuthInfo(user *User) (newUser *User, err error) {
	newUser = new(User)

	q := `SELECT
            *
          FROM
            user
          WHERE
            name = :name OR email = :email`
	rows, err := m.db.NamedQuery(q, user)
	for rows.Next() {
		err = rows.StructScan(newUser)
		if err != nil {
			err = errors.Wrap(err, "incorrect password or username or email")
			return
		}
	}

	return
}

func (m *UserModel) UpdateUserPassword(userID int, planePassword string) (err error) {
	hashedPassword, err := helper.ToHash(planePassword)
	if err != nil {
		err = errors.Wrap(err, "password is invalid")
		return
	}
	q := `UPDATE
            user
          SET
             password = ?
          WHERE
            id = ?`
	_, err = m.db.Exec(q, hashedPassword, userID)
	if err != nil {
		err = errors.Wrap(err, "couldn't update user's password")
		return
	}

	return
}

func (m *UserModel) CompareHashAndPlain(hashedPassword string, plainPassword string) (err error) {
	err = helper.CompareHashAndPlain(hashedPassword, plainPassword)
	if err != nil {
		err = errors.Wrap(err, "password is incorrect")
		return
	}

	return
}
