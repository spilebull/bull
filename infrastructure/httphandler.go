package infrastructure

/*
HTTPハンドラーは、以下を行います
1.HTTPリクエストを受け入れ、それを解析して検証する。(contextや、path、bodyからのデータ取得)
2.ステップ1で得たデータを使ってコントローラーを呼び出してBusinessLogicを実行する。
3.コントローラーが返す内容に応じて、適切なHTTPレスポンスを送信する。
*/

import (
	"net/http"
	"server/adapter/controller"
	"server/domain"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (server *Server) createUser(c *gin.Context) {
	request := domain.User{}
	if err := c.ShouldBind(&request); err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypeBind)
		return
	}
	userController := controller.NewUserController(server.Store)
	err := userController.CreateUser(&request)
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"code": http.StatusCreated, "message": http.StatusText(http.StatusCreated)})
}

func (server *Server) updateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		_ = c.Error(err)
		return
	}
	request := domain.User{}
	if err := c.ShouldBind(&request); err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypeBind)
		return
	}
	userController := controller.NewUserController(server.Store)
	err = userController.UpdateUser(&request, id)
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": http.StatusText(http.StatusOK)})
}

func (server *Server) getUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		_ = c.Error(err)
		return
	}
	userController := controller.NewUserController(server.Store)
	user, err := userController.GetUserByID(id)
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": http.StatusText(http.StatusOK), "data": user})
}

func (server *Server) getUserList(c *gin.Context) {
	userController := controller.NewUserController(server.Store)
	users, err := userController.GetUsers()
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": http.StatusText(http.StatusOK), "data": users})
}

func (server *Server) deleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		_ = c.Error(err)
		return
	}
	userController := controller.NewUserController(server.Store)
	err = userController.DeleteUserByID(id)
	if err != nil {
		_ = c.Error(err).SetType(gin.ErrorTypePublic)
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
