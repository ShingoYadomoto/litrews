package model

import (
	"time"

	"github.com/ShingoYadomoto/litrews/src/db"
	"github.com/ShingoYadomoto/litrews/src/helper"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

type User struct {
	ID         int       `db:"id" json:"id" form:"id"`
	Name       string    `db:"name" json:"name" form:"name"`
	Password   string    `db:"password" json:"password" form:"password"`
	Email      string    `db:"email" json:"email" form:"email"`
	CreateTime time.Time `db:"create_time" json:"create_time"`
	UpdateTime time.Time `db:"update_time" json:"update_time"`
	DeleteFlg  bool      `db:"delete_flg" json:"delete_flg"`
	Topics     []*Topic  `db:"-"`
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
            user.*,
            p.id "user_profile.id",
            p.display_name "user_profile.display_name",
            p.create_time "user_profile.create_time",
            p.update_time "user_profile.update_time",
            IFNULL(a.id, 0) "user_avatar.id",
            IFNULL(a.file_name, "") "user_avatar.file_name",
            IFNULL(a.width, 0) "user_avatar.width",
            IFNULL(a.height, 0) "user_avatar.height",
            IFNULL(a.create_time, CAST('1111-11-11 11:11:11' AS DATE)) "user_avatar.create_time",
            IFNULL(a.update_time, CAST('1111-11-11 11:11:11' AS DATE)) "user_avatar.update_time"
          FROM
            user
          INNER JOIN
            user_profile AS p
          ON
            user.id = p.user_id
          LEFT JOIN
            user_avatar AS a
          ON
            user.id = a.user_id
          WHERE
            user.id = ?`
	err = m.db.Get(user, q, id)
	if err != nil {
		err = errors.Wrap(err, "couldn't get user")
		return
	}

	// ManyToMany, HasMany のリレーションはjoinとは別で取ってくるのが良さそう
	//roleModel := NewRoleModel(m.db)
	//user.Roles, err = roleModel.GetRolesByUserID(id)

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
