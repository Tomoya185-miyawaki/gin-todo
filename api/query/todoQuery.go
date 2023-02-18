/*
インターフェース用パッケージ
*/
package query

import "github.com/Tomoya185-miyawaki/gin-todo/domain/model"

// Todo用のインターフェース
type TodoQuery interface {
	// 全件取得
	FindAll() *model.Todos
}
