/*
インターフェース用パッケージ
*/
package query

import (
	"github.com/Tomoya185-miyawaki/gin-todo/domain/model"
	"github.com/Tomoya185-miyawaki/gin-todo/errors"
)

// Todo用のインターフェース
type TodoQuery interface {
	// 全件取得
	FindAll() *model.Todos
	// IDをキーにして取得する
	FindById(id int) (*model.Todo, *errors.AppError)
	// 作成する
	Create(title string) (*string, *errors.AppError)
}
