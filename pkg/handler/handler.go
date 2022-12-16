package handler

import (
	"net/http"
	"time"

	"github.com/Marityr/gopitman/pkg/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://foo.com"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	// router.Use(logger.SetLogger())

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api/v1") //h.userIdentity
	{
		essay := api.Group("/essay")
		{
			essay.POST("/", h.createEssay)
		}

		// {
		// 	ws.GET("/", h.wsEndpoint)
		// }
	}

	router.GET("/ws", func(c *gin.Context) {
		wshandler(c.Writer, c.Request)
	})

	router.GET("/page", func(c *gin.Context) {
		w := c.Writer
		r := c.Request
		// c.HTML(200, "/static/index.htm", nil)
		p := "./static/index.htm"
		http.ServeFile(w, r, p)
	})

	return router
}
