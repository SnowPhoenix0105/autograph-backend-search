package server

import (
	"autograph-backend-search/server/common"
	"autograph-backend-search/server/handler"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Config struct {
	Host      string
	Port      int
	DebugMode bool
}

type Server struct {
	engine *gin.Engine
	config *Config
}

func New(config *Config) *Server {
	eng := gin.Default()

	eng.Use(common.LogRequest)
	eng.Use(common.SetUserInfo(config.DebugMode))
	eng.Use(cors.Default())

	eng.GET("/test/coffee", coffeeHandler)
	eng.GET("/search", handler.Search)
	eng.GET("/version", handler.GetVersion)
	eng.GET("/set-version", handler.SetVersion)
	eng.GET("/download", handler.Download)

	// 需要登录的路由
	adminGroup := eng.Group("admin")
	{
		adminGroup.Use(common.RejectNotLogin(config.DebugMode))

	}

	return &Server{
		engine: eng,
		config: config,
	}
}

func (s *Server) RunServer() error {
	return s.engine.Run(fmt.Sprintf("%s:%d", s.config.Host, s.config.Port))
}
