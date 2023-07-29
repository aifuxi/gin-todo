package cmd

import (
	_ "github.com/aifuxi/gin-todo/internal/model"
	"github.com/aifuxi/gin-todo/internal/router"
)

func RunServer() {
	r := router.New()

	r.Run("localhost:9000")
}
