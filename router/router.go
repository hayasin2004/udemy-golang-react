package router

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"os"
	"udemy-golang-react/controller"
)

func NewRouter(uc controller.IUserController, tc controller.ITaskController) *echo.Echo {
	//ユーザー関係のエンドポイント
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		//アクセスを許可するドメインの追加
		AllowOrigins: []string{"http://localhost:3000", os.Getenv("FE_URL")},
		//許可するヘッダー一覧の入力　→XCSRTokenFを含めることでHeader経由でCSRFトークンを受け取ることができる
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders, echo.HeaderXCSRFToken},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE"},
		AllowCredentials: true,
	}))

	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		CookiePath:     "/",
		CookieDomain:   os.Getenv("API_DOMAIN"),
		CookieHTTPOnly: true,
		CookieSameSite: http.SameSiteNoneMode,
		//↑　これ本番環境用

		//↓これはポストマンで実行する奴
		//CookieSameSite: http.SameSiteDefaultMode,
		//	CookieMaxAge : 60 →　これでCSRFトークンの有効期限を変えれる。　今回の場合は60秒
	}))

	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.LogIn)
	e.POST("/logout", uc.LogOut)
	e.GET("/csrf", uc.CsrfToken)
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
