/*
mainパッケージ
*/
package main

import (
	db "github.com/Tomoya185-miyawaki/gin-todo/infrastructure"
	router "github.com/Tomoya185-miyawaki/gin-todo/router"
)

func main() {
	db := db.Init()
	defer db.Close()
	router.Bind()
}
