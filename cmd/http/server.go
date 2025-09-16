package main

import (
	"net/http"

	httpHandler "github.com/davidsugianto/cloud-cost-tracker/internal/handler/http"
	"github.com/davidsugianto/cloud-cost-tracker/internal/handler/http/middleware"
	"github.com/davidsugianto/cloud-cost-tracker/internal/pkg/config"
	user "github.com/davidsugianto/cloud-cost-tracker/internal/usecase/user"
	"github.com/davidsugianto/go-pkgs/grace"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	*http.Server
	handler *httpHandler.Handler
	config  *config.Config
}

type Dependencies struct {
	UserUseCase user.Usecase
	Config      *config.Config
}

func New(deps Dependencies) *Server {
	return &Server{
		Server: &http.Server{},
		handler: httpHandler.New(httpHandler.Dependencies{
			UserUseCase: deps.UserUseCase,
		}),
		config: deps.Config,
	}
}

func (s *Server) v1Endpoint(r *gin.Engine) {
	g := r.Group("/v1")
	g.Use(gin.Recovery(), middleware.RequestID(), middleware.Logger())

	// health check
	g.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// authentication
	auth := g.Group("/auth")
	auth.POST("/login", s.handler.Login)
	auth.POST("/register", s.handler.SignUp)

	// cost
	// project
	// budget
	// sync

	// announcement.GET("", s.handler.GetAnnouncements)
	// resourceGroup.GET("/ownerless", s.handler.GetResourceGroupOwnerless)
	// resourceGroup.GET("/:env/:datacenter/:name/:project_id", s.handler.GetResourceGroupGeneralInformation)
}

func (s *Server) Run(port string) error {
	r := gin.Default()

	corsConfig := cors.Config{
		AllowOrigins:     s.config.CORS.AllowedOrigins,
		AllowMethods:     s.config.CORS.AllowedMethods,
		AllowHeaders:     s.config.CORS.AllowedHeaders,
		AllowCredentials: s.config.CORS.AllowCredentials,
	}
	r.Use(cors.New(corsConfig))

	s.v1Endpoint(r)

	s.Addr = port
	s.Handler = r

	return grace.ServeHTTP(s.Addr, s.Handler)
}
