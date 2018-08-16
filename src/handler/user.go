package handler

import (
	"net/http"

	"github.com/ShingoYadomoto/litrews/src/context"
	cdb "github.com/ShingoYadomoto/litrews/src/db"
	"github.com/ShingoYadomoto/litrews/src/helper"
	"github.com/ShingoYadomoto/litrews/src/model"
	"github.com/ShingoYadomoto/litrews/src/request"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func ShowUser(c echo.Context) (err error) {
	db := c.(*context.CustomContext).GetDB()
	s := c.(*context.CustomContext).GetSession()

	userModel := model.NewUserModel(db)
	topicModel := model.NewTopicModel(db)

	userID := s.Values["user_id"].(int)

	user, err := userModel.GetUserByID(userID)
	if err != nil {
		log.Error(err)
		if user.ID == 0 {
			return c.Redirect(http.StatusFound, "/user/create")
		}
		return c.Render(http.StatusOK, "error", err)
	}

	topics, err := topicModel.GetAllTopics()
	if err != nil {
		log.Error(err)
		return c.Render(http.StatusOK, "error", err)
	}

	isCheckedTopicMap := make(map[int]bool, len(topics))
	for _, topic := range topics {
		for _, userTopic := range user.Topics {
			if userTopic.ID == topic.ID {
				isCheckedTopicMap[topic.ID] = true
			}
		}
	}

	return c.Render(http.StatusOK, "userShow", map[string]interface{}{
		"user":              user,
		"topics":            topics,
		"isCheckedTopicMap": isCheckedTopicMap,
	})
}

func UpdateUser(c echo.Context) (err error) {
	req := new(request.UserUpdate)
	err = c.Bind(req)
	if err != nil {
		log.Error(err)
		return c.Render(http.StatusOK, "error", err)
	}

	err = c.Validate(req)
	if err != nil {
		log.Error(err)
		c.(*context.CustomContext).SaveValidationErrors(err)
		return c.Redirect(http.StatusFound, c.Request().Referer())
	}

	db := c.(*context.CustomContext).GetDB()
	s := c.(*context.CustomContext).GetSession()

	userModel := model.NewUserModel(db)
	topicModel := model.NewTopicModel(db)

	userID := s.Values["user_id"].(int)

	oldUser, err := userModel.GetUserByID(userID)
	if err != nil {
		log.Error(err)
		return c.Render(http.StatusOK, "error", err)
	}

	newUser := new(model.User)
	newUser.ID = userID
	if err = c.Bind(newUser); err != nil {
		log.Error(err)
		return c.Render(http.StatusOK, "error", err)
	}

	// ↓ structがネストしていると、c.Bind()がうまくいかないので暫定処置
	topicIDs, err := helper.AtoiSlice(c.Request().Form["topic_ids[]"])
	if err != nil {
		log.Error(err)
		return c.Render(http.StatusOK, "error", err)
	}
	newUser.Topics, err = topicModel.GetTopicsByTopicIDs(topicIDs)
	if err != nil {
		log.Error(err)
		return c.Render(http.StatusOK, "error", err)
	}
	// ↑

	err = db.Transaction(func(sc cdb.SchemaContext) (err error) {
		topicModel := model.NewTopicModel(sc.DB())

		// userテーブルの更新
		err = userModel.UpdateUser(newUser)
		if err != nil {
			return
		}

		// users_topicsテーブルの更新
		if len(oldUser.Topics) == 0 {
			err = topicModel.CreateUserTopicsByUser(newUser)
			if err != nil {
				return
			}
		} else {
			err = topicModel.DeleteUserTopicByUserID(userID)
			if err != nil {
				return
			}
			err = topicModel.CreateUserTopicsByUser(newUser)
			if err != nil {
				return
			}
		}

		return
	})
	if err != nil {
		log.Error(err)
		return c.Render(http.StatusOK, "error", err)
	}

	s.AddFlash("ユーザー情報の更新が完了しました。")
	s.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusFound, "/user")
}
