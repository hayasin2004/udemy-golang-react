package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"time"
	"udemy-golang-react/model"
	useCase "udemy-golang-react/usecase"
)

//echo
//→リクエストパラメータ、ヘッダ、クエリ文字列、フォームデータにアクセスするもの
//→レスポンスの送信　（JsonやHTMLのレスポンスの送信が簡単に行える）
//→三位ドルウェアのサポート　→Contextを通じて、認証やロギングなどの共通処理を統一的に実行
//　→ログインの確認　イメージ　「リクエストが来る　→　共通のチェックをと通る　→　その後の処理」
//→リクエスト全体で共有したいデータを簡単に設定、処理できる
//　→　認証成功後にリクエストされたユーザーIDを保存しておくことで、後続の処理でユーザーIDを無駄な計算やデータ処理を防ぐことができる

type IUserController interface {
	SignUp(c echo.Context) error
	LogIn(c echo.Context) error
	LogOut(c echo.Context) error
	CsrfToken(c echo.Context) error
}

type userController struct {
	uu useCase.IUserUseCase
}

func NewUserController(uu useCase.IUserUseCase) IUserController {
	return &userController{uu}
}

func (uc *userController) SignUp(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	userRes, err := uc.uu.SignUp(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, userRes)
}

func (uc *userController) LogIn(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	tokenString, err := uc.uu.Login(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = tokenString
	cookie.Expires = time.Now().Add(24 * time.Hour)
	//cookie.Path →　これで指定されたパスの実に共有するもの。
	//今回の場合だと , "/"だからドメイン全体で使用。　"/index"だった場合は"/index/test1","/index/test2"など"/index"のパスのみ共有される。
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.Secure = true
	cookie.HttpOnly = true
	//クロスサイトリクエストを許可するモノ　→　他のサイトからこのサイトを見られて時にもcookieが送信される。
	cookie.SameSite = http.SameSiteNoneMode
	//cookieで作成したcookieをレスポンスに追加しクライアントに送信。
	c.SetCookie(cookie)
	//リクエストが正常に処理された時に204'no Content'を返す
	return c.NoContent(http.StatusOK)
}

// ログアウト
func (uc *userController) LogOut(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = ""
	cookie.Expires = time.Now()
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie)
	return c.NoContent(http.StatusOK)
}

func (uc *userController) CsrfToken(c echo.Context) error {
	token := c.Get("csrf").(string)
	return c.JSON(http.StatusOK, echo.Map{"csrf_token": token})
}
