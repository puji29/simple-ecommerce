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
	userUC        usecase.UserUseCase
	productUC     usecase.ProductUseCase
	orderTableUC  usecase.OrderTableUseCase
	imageUC       usecase.ImageUseCase
	categoryUC    usecase.CategoryUseCase
	orderDetailUc usecase.OrderDetailUsecase
	engine        *gin.Engine
	host          string
}

func (s *Server) initRoute() {
	rg := s.engine.Group(config.ApiGroup)
	controller.NewUserController(s.userUC, rg).Route()
	controller.NewProductController(s.productUC, rg).Route()
	controller.NewOrderTableController(s.orderTableUC, rg).Route()
	controller.NewImageController(s.imageUC, rg).Route()
	controller.NewCategoryController(s.categoryUC, rg).Route()
	controller.NewOrderDetailController(s.orderDetailUc, rg).Route()
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
	productRepo := repository.NewProductRepository(db)
	orderTableRepo := repository.NewOrderTableRepository(db)
	imageRepo := repository.NewImageRepository(db)
	cateRepo := repository.NewCategoryRepository(db)
	orderDetailRepo := repository.NewOrderDetailRepository(db)

	//inject usecase
	userUc := usecase.NewUserUseCase(userRepo)
	ProductUc := usecase.NewProductUseCase(productRepo)
	OrderTableUC := usecase.NewOrderTableUseCase(orderTableRepo)
	imageUc := usecase.NewImageUsecase(imageRepo)
	orderDetailUc := usecase.NewOrderDeatailUsecase(orderDetailRepo)
	cateUc := usecase.NewCategoryUseCase(cateRepo)
	engine := gin.Default()
	host := fmt.Sprintf(":%s", cfg.ApiPort)
	return &Server{
		engine:        engine,
		host:          host,
		userUC:        userUc,
		productUC:     ProductUc,
		orderTableUC:  OrderTableUC,
		imageUC:       imageUc,
		categoryUC:    cateUc,
		orderDetailUc: orderDetailUc,
	}
}
