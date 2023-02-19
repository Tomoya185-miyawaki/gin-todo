/*
ヘルパー用のパッケージ
*/
package helpers

import (
	"github.com/Tomoya185-miyawaki/gin-todo/errors"
	"github.com/gin-gonic/gin"
)

// レスポンスをjsonで返す
func Response(c *gin.Context, result interface{}, err *errors.AppError) {
	if err.HasErrors() {
		c.JSON(400, err)
	} else {
		c.JSON(200, result)
	}
}
