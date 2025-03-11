package main

import (
	"udemy-golang-react/controller"
	"udemy-golang-react/db"
	"udemy-golang-react/repository"
	"udemy-golang-react/router"
	useCase "udemy-golang-react/usecase"
)

func main() {
	//コパイロット空の説明　→　イメージがつかないならこれで覚えて。
	//これは各専門家のリレーのバトン回しっていうイメージ。　→　大規模なウェブ開発でもロジックは違うけど、走順は同じ。・・・多分このイメージ大事

	//これはスタートライン。（データベース接続）
	db := db.NewDB()

	//第一走者（データベースの専門家）　→　データベース接続を受け取り、ユーザー情報の管理、取得する専門家。
	//データベースとの直接的なやり取り
	userRepository := repository.NewUserRepository(db)
	taskRepository := repository.NewTaskRepository(db)

	//第二走者（ビジネスロジックの専門家）　→　ユーザーリポジトリ―を利用して、ユーザー関連のビジネスロジックを実行するプランナー
	//データを処理し実行
	userUseCase := useCase.NewUserUseCase(userRepository)
	taskUseCase := useCase.NewTaskUseCase(taskRepository)

	//第三走者（司会者専門）ユーザーインターフェイスを管理し、ユーザーのリクエストを受け取り、ユースケースにリクエストを渡す者
	userController := controller.NewUserController(userUseCase)
	taskController := controller.NewTaskController(taskUseCase)

	//アンカー　→ユーザーコントローラーを使ってリクエストをルーティングするランナー
	e := router.NewRouter(userController, taskController)
	//ゴール地点
	e.Logger.Fatal(e.Start(":8080"))
}
