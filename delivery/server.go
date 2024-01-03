package delivery

import (
	"database/sql"
	"ecommerce/config"
	"ecommerce/delivery/controller"
	"ecommerce/repository"
	"ecommerce/usecase"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Server struct {
	userUC usecase.UserUseCase
	engine *gin.Engine
	host   string
}

func (s *Server) initRoute() {
	rg := s.engine.Group(config.ApiGroup)
	controller.NewUserController(s.userUC, rg).Route()
}

func (s *Server) Run() {
	s.initRoute()
	if err := s.engine.Run(s.host); err != nil {
		panic(fmt.Errorf("sever not running on host %s, because error %v", s.host, err.Error()))
	}
}

func NewServer() *Server {
	cfg, _ := config.NewConfig()
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.Host, cfg.Port, cfg.User, cfg.Pasword, cfg.Name)
	db, err := sql.Open(cfg.Driver, dsn)
	if err != nil {
		panic("connection error")
	}
	//inject repo
	userRepo := repository.NewUserREpository(db)

	//inject usecase
	userUc := usecase.NewUserUseCase(userRepo)
	engine := gin.Default()
	host := fmt.Sprintf(":%s", cfg.ApiPort)
	return &Server{
		engine: engine,
		host:   host,
		userUC: userUc,
	}
}
