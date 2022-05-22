package app

import (
	"Scoltest/internal/config"

	orderhttp "Scoltest/internal/order/api/http"
	pgorder "Scoltest/internal/order/db/pg"
	orderservice "Scoltest/internal/order/service"
	prodhttp "Scoltest/internal/product/api/http"
	pgprod "Scoltest/internal/product/db/pg"
	prod "Scoltest/internal/product/service"
	http2 "Scoltest/internal/user/api/http"
	pguser "Scoltest/internal/user/db/pg"
	user2 "Scoltest/internal/user/service"
	"Scoltest/pkg/client/postgres"
	"Scoltest/pkg/loger"
	"context"
	"errors"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net"
	"net/http"
	"time"
)

type Server struct {
	config *config.Config
	logger *loger.Logger
	//router *httprouter.Router
	ctx context.Context
}

func (s *Server) Start() {

	router := httprouter.New()
	conn, err := postgres.Connect(context.Background(), postgres.ConnectionData{
		Login:    s.config.Base.Login,
		Password: s.config.Base.Pass,
		Host:     s.config.Base.Host,
		Port:     s.config.Base.Port,
		BaseNane: s.config.Base.Basename,
	})

	if err != nil {
		s.logger.Fatalln("Error connect to base ", err.Error(), conn)
		return
	}
	s.ctx = context.Background()
	userStorage, err := pguser.New(s.ctx, s.logger, conn)
	if err != nil {
		s.logger.Fatalln("Error create user storage ", err.Error())
	}
	userService := user2.NewService(userStorage, s.logger)
	user := http2.NewHandler(s.logger, userService)
	user.Register(router)

	productStorage, err := pgprod.New(s.ctx, s.logger, conn)
	if err != nil {
		s.logger.Fatalln("Error create product storage ", err.Error())
	}
	productService := prod.NewService(productStorage, s.logger)
	product := prodhttp.NewHandler(s.logger, productService)
	product.Register(router)

	orderStorage, err := pgorder.New(s.ctx, s.logger, conn)
	if err != nil {
		s.logger.Fatalln("Error create order storage ", err.Error())
	}
	orderService := orderservice.NewService(orderStorage, s.logger)
	order := orderhttp.NewHandler(s.logger, orderService)
	order.Register(router)

	listener, err := net.Listen("tcp", s.config.Listen.Port)
	if err != nil {
		s.logger.Fatalln(fmt.Sprintf("Error create listener : %s", err.Error()))
	}

	server := http.Server{
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	s.logger.Infoln(fmt.Sprintf("Start listen server port : %s", s.config.Listen.Port))
	s.logger.Fatalln(server.Serve(listener))

}

func New(config *config.Config, logger *loger.Logger) (*Server, error) {

	if config == nil {
		return nil, errors.New("config is nil")
	}

	if logger == nil {
		return nil, errors.New("loger is nil")
	}
	return &Server{config: config, logger: logger}, nil
}
