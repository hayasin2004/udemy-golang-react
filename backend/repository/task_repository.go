package repository

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"udemy-golang-react/model"
)

type ITaskRepository interface {

	//複数のタスクを取得　→　配列で管理　、　第二引数でuserIdを渡すことで誰が投稿したかの情報を取得する（ログインしているユーザー）
	GetAllTasks(tasks *[]model.Task, userId uint) error

	//タスクを単体取得　→　配列では管理しない。　、第三引数でtaskIdを渡しそれに一致するものを取得してくる
	GetTaskById(task *model.Task, userId uint, taskId uint) error

	//タスクを新しく生成するだけの処理だから何もいらない
	CreateTask(task *model.Task) error

	//特定のタスクを更新するプログラム　→　ログインしているユーザー、タスクIdを渡す必要がある
	UpdateTask(task *model.Task, userId uint, taskId uint) error

	//特定のタスクを削除　→　今回はユーザー認証なしで削除？？
	DeleteTask(userId uint, taskId uint) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) ITaskRepository {
	return &taskRepository{db}
}

func (tr *taskRepository) GetAllTasks(tasks *[]model.Task, userId uint) error {
	//GetAllTaskを委員長として例えると
	//Joinsで生徒名簿をtaskRepositoryに結び付けてる。　イメージでいうとタスクを誰がやったのかを探してる
	if err := tr.db.Joins("User").Where("user_id=?", userId).Order("created_at").Find(tasks).Error; err != nil {
		return err
	}
	return nil
}

func (tr *taskRepository) GetTaskById(task *model.Task, userId uint, taskId uint) error {
	if err := tr.db.Joins("User").Where("user_id=?", userId).First(task, taskId).Error; err != nil {
		return err
	}
	return nil
}

func (tr *taskRepository) CreateTask(task *model.Task) error {
	if err := tr.db.Create(task).Error; err != nil {
		return err
	}
	return nil
}

// 更新するときのイメージ　→　特定のユーザーが所有する特定のタスクのタイトルを更新する処理
func (tr *taskRepository) UpdateTask(task *model.Task, userId uint, taskId uint) error {

	//clause　→　キー
	//.Claues(clause.Returning{}) →　更新操作の後に、更新されたレコードを返却するオプション
	//一部のDBではこの設定によって更新された内容を取得できるが必須では無い。
	//Update("title", task.Title)
	//　→　"title"の意味はデータべ―スのカラムの名前　、役割として更新操作の対象。
	//　→task.Titleは意味はGoの構造体でTask構造体の中に定義されたフィールドTitleの値を指す。→役割として更新数る新しいデータを表す
	result := tr.db.Model(task).Clauses(clause.Returning{}).Where("id=? AND user_id=?", taskId, userId).Update("title", task.Title)
	if result.Error != nil {
		return result.Error
	}
	//更新操作の結果、影響を受けた行数（RowsAffected)が0以下であればえら0としてタスクが存在しませんと返す。
	if result.RowsAffected < 1 {
		return fmt.Errorf("タスクが存在しません")
	}
	return nil
}

// Deleteはresultから文章が始まって、Createはif文から始まってる理由
// 着目点は結果オブジェクトが複数のプロパティを使用するかどうか。　→　Deleteは何かエラーが起きたのか、タスクを正しく消せたのかの複数
// Createは作成に関するエラーを見ればいい。　つまり一つの結果オブジェクト。
// 今回のTodoリストでいうと何か特定の物をいじるというときはresultのように一旦インスタンス化して持っておくという考えが多い。
// Createみたいに特定の物を探す操作が不要ならば、成功もしくは失敗のみの判定だけどいい。

func (tr *taskRepository) DeleteTask(userId uint, taskId uint) error {
	result := tr.db.Where("id=? AND user_id=?", taskId, userId).Delete(&model.Task{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("タスクが見つかりませんでした")
	}
	return nil
}
