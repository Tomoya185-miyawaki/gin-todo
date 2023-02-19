/*
ヘルパー用のパッケージ
*/
package helpers

import (
	"net/http"

	"github.com/Tomoya185-miyawaki/gin-todo/errors"
	"github.com/gin-gonic/gin"
)

// レスポンスをjsonで返す
func Response(c *gin.Context, result interface{}, err *errors.AppError) {
	if err.HasErrors() {
		c.JSON(http.StatusBadRequest, err)
	} else {
		c.JSON(http.StatusOK, result)
	}
}
