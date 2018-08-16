package model

import (
	"time"

	"github.com/ShingoYadomoto/litrews/src/db"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type Topic struct {
	ID         int       `db:"id" json:"id" form:"id"`
	Name       string    `db:"name" json:"name" form:"name"`
	NameJa     string    `db:"name_ja" json:"name_ja" form:"name_ja"`
	CreateTime time.Time `db:"create_time" json:"create_time"`
	UpdateTime time.Time `db:"update_time" json:"update_time"`
}

type TopicModel struct {
	db db.AbstractDB
}

func NewRoleModel(db db.AbstractDB) *TopicModel {
	return &TopicModel{
		db: db,
	}
}

func (m *TopicModel) GetAllTopics() (topics []*Topic, err error) {
	topics = []*Topic{}

	q := `SELECT * FROM topic`
	err = m.db.Select(&topics, q)
	if err != nil {
		err = errors.Wrap(err, "couldn't get all topics")
		return
	}

	return
}

func (m *TopicModel) GetTopicByID(id int) (topic *Topic, err error) {
	topic = new(Topic)

	q := `SELECT * FROM topic WHERE id = ?`
	err = m.db.Get(topic, q, id)
	if err != nil {
		err = errors.Wrap(err, "couldn't get topic")
		return
	}

	return
}

func (m *TopicModel) GetTopicsByUserID(userID int) (topics []*Topic, err error) {
	topics = []*Topic{}

	q := `SELECT
            topic.*
          FROM
            topic
          INNER JOIN
            users_topics
          ON
            topic.id = users_topics.topic_id
          WHERE
            users_topics.user_id = ?`
	err = m.db.Select(&topics, q, userID)
	if err != nil {
		err = errors.Wrap(err, "couldn't get user_topics")
		return
	}

	return
}

func (m *TopicModel) GetTopicsByTopicIDs(topicIDs []int) (topics []*Topic, err error) {
	if len(topicIDs) == 0 {
		return
	}
	topics = []*Topic{}

	q, vs, err := sqlx.In(`SELECT * FROM topic WHERE id IN (?)`, topicIDs)
	if err != nil {
		err = errors.Wrap(err, "topicIDs []int is invalid")
		return
	}

	err = m.db.Select(&topics, q, vs...)
	if err != nil {
		err = errors.Wrap(err, "couldn't get user_topics")
		return
	}

	return
}

func (m *TopicModel) CreateUserTopicsByUser(user *User) (err error) {
	q := `INSERT INTO
            users_topics (user_id, topic_id)
          VALUES
            (?, ?)`
	// bulk insertできるメソッドが用意されてないので、暫定的にinsertぶん回し
	for _, topic := range user.Topics {
		_, err = m.db.Exec(q, user.ID, topic.ID)
		if err != nil {
			err = errors.Wrap(err, "couldn't get user_topics")
			return
		}
	}

	return
}

func (m *TopicModel) DeleteUserTopicByUserID(userID int) (err error) {
	q := `DELETE FROM
            users_topics
          WHERE
            user_id = ?`
	_, err = m.db.Exec(q, userID)
	if err != nil {
		err = errors.Wrap(err, "couldn't delete user's topics")
		return
	}

	return
}
