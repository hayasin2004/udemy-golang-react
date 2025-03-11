package controller

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"udemy-golang-react/model"
	useCase "udemy-golang-react/usecase"
)

type ITaskController interface {
	GetAllTasks(c echo.Context) error
	GetTaskById(c echo.Context) error
	CreateTask(c echo.Context) error
	UpdateTask(c echo.Context) error
	DeleteTask(c echo.Context) error
}

type taskController struct {
	tu useCase.ITaskUseCase
}

func NewTaskController(tu useCase.ITaskUseCase) ITaskController {
	return &taskController{tu}
}

// ユーザー情報のみを取得して、タスクを全件返している
func (tc *taskController) GetAllTasks(c echo.Context) error {

	//ユーザー認証
	user := c.Get("user").(*jwt.Token)
	//ClaimsはJWT内でユーザーやトークンに関する情報を格納する場所
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	//float64　→　userIdはany型なので型アサーション
	taskRes, err := tc.tu.GetAllTasks(uint(userId.(float64)))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, taskRes)
}

// タスクID , ユーザーID（jwt)を取得して、タスクの詳細をクライアントに返すもの
func (tc *taskController) GetTaskById(c echo.Context) error {
	//ユーザー認証
	user := c.Get("user").(*jwt.Token)

	//user.Claims.(jwt.mapClaims)　→　Reactでいうとprops , interfaceを同時に行っている
	//map[string]interface{}　→ jwt.MapClaimsの型指定。 [string]でキー部分をstring , 値の部分はanyで返される
	//つまり値の部分はまた個人的に指定する必要がある。　今回の場合だとfloat64の部分
	//map[string]　→　Go言語におけるキーを持つマップ型（連想配列とも呼ぶらしい）を指す。
	//Reactのmapとは違うらしい。 go言語でいうmapはキーと値のペアを格納するデータ構造。　→　辞書　、連想配列という物。
	//もしGo言語でループ取り出しをしたいのであれば、for文を使う。
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	//HTTPリクエストのURLパラメータのtaskIdを取得して文字列型でidに代入する　→ /tasks/123だった場合は"123"となる
	id := c.Param("taskId")

	//strconv.Atoi　→　Goの標準ライブラリ。 文字列型を整数型に変換するために使用される。 Atoi → ASCLL to Inteager(文字列から整数に変換)
	taskId, _ := strconv.Atoi(id)
	taskRes, err := tc.tu.GetTaskById(uint(userId.(float64)), uint(taskId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, taskRes)

}

func (tc *taskController) CreateTask(c echo.Context) error {
	//	ユーザー認証
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	//空のタスクオブジェクトの定義
	task := model.Task{}

	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	task.UserId = uint(userId.(float64))
	taskRes, err := tc.tu.CreateTask(task)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, taskRes)
}

func (tc *taskController) UpdateTask(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("taskId")
	taskId, _ := strconv.Atoi(id)
	task := model.Task{}

	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	taskRes, err := tc.tu.UpdateTask(task, uint(userId.(float64)), uint(taskId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, taskRes)
}

func (tc *taskController) DeleteTask(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("taskId")
	taskId, _ := strconv.Atoi(id)

	err := tc.tu.DeleteTask(uint(userId.(float64)), uint(taskId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
