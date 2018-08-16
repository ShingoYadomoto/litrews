package handler

import (
	"net/http"

	"github.com/ShingoYadomoto/litrews/src/context"
	"github.com/ShingoYadomoto/litrews/src/model"
	"github.com/ShingoYadomoto/litrews/src/request"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func Signin(c echo.Context) (err error) {
	req := new(request.Signin)
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

	requestUser := new(model.User)
	if err = c.Bind(requestUser); err != nil {
		log.Error(err)
		return c.Render(http.StatusOK, "error", err)
	}

	user, err := userModel.GetUserByAuthInfo(requestUser)
	if err != nil {
		log.Error(err)
		return c.Render(http.StatusOK, "error", err)
	}
	//email, nameが間違っている場合
	if user.ID == 0 {
		s.AddFlash("サインイン情報が間違っています。")
		s.Save(c.Request(), c.Response())
		return c.Redirect(http.StatusFound, "/")
	}

	err = userModel.CompareHashAndPlain(user.Password, requestUser.Password)
	//passwordが間違っている場合
	if err != nil {
		log.Error(err)
		s.AddFlash("サインイン情報が間違っています。")
		s.Save(c.Request(), c.Response())
		return c.Redirect(http.StatusFound, "/")
	}

	s.Values["user_id"] = user.ID
	s.AddFlash("サインインが完了しました。")
	s.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusFound, "/user")
}

func Signup(c echo.Context) (err error) {
	req := new(request.Signup)
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

	user := new(model.User)
	if err = c.Bind(user); err != nil {
		log.Error(err)
		return c.Render(http.StatusOK, "error", err)
	}

	err = userModel.CreateUser(user)
	if err != nil {
		log.Error(err)
		return c.Render(http.StatusOK, "error", err)
	}

	signedUpUser, err := userModel.GetUserByAuthInfo(user)
	if err != nil {
		log.Error(err)
		return c.Render(http.StatusOK, "error", err)
	}

	s.Values["user_id"] = signedUpUser.ID
	s.AddFlash("サインアップが完了しました。")
	s.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusFound, "/user")
}

func Signout(c echo.Context) (err error) {
	s := c.(*context.CustomContext).GetSession()

	s.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	}

	s.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusFound, "/")
}
