package router

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"os"
	"udemy-golang-react/controller"
)

func NewRouter(uc controller.IUserController, tc controller.ITaskController) *echo.Echo {
	//ユーザー関係のエンドポイント
	e := echo.New()
	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.LogIn)
	e.POST("/logout", uc.LogOut)
	t := e.Group("/tasks")
	//token管理のミドルウェア
	t.Use(echojwt.WithConfig(echojwt.Config{
		//SigningKey　→ jwtを生成したとき同じコードが格納されている
		SigningKey: []byte(os.Getenv("SECRET")),
		//TokenLookup →　クライアントから送られてくるjwtをどこに保存されるのかを指定
		//今回はcookieにtokenという名前で保存されている
		TokenLookup: "cookie:token",
	}))
	//タスク関係のエンドポイントの設定
	t.GET("", tc.GetAllTasks)
	t.GET("/:taskId", tc.GetTaskById)
	t.POST("", tc.CreateTask)
	t.PUT("/:taskId", tc.UpdateTask)
	t.DELETE("/:taskId", tc.DeleteTask)
	return e
}
