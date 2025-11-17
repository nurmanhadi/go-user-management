package config

import (
	grpcHand "user-management/delivery/grpc/handler"
	restHandler "user-management/delivery/rest/handler"
	"user-management/delivery/rest/middleware"
	"user-management/delivery/rest/routes"
	"user-management/internal/event/subscriber"
	"user-management/internal/protobuf/pb"
	"user-management/internal/repository"
	"user-management/internal/service"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/gorm"
)

type Bootstrap struct {
	DB        *gorm.DB
	Cache     *memcache.Client
	Logger    zerolog.Logger
	Validator *validator.Validate
	Router    *chi.Mux
	Ch        *amqp.Channel
	S         *grpc.Server
}

func Initialize(deps *Bootstrap) {
	// publisher

	// cache

	// repository
	userRepo := repository.NewUserRepository(deps.DB)

	// service
	userServ := service.NewUserService(deps.Logger, deps.Validator, userRepo)

	// handler
	userHand := restHandler.NewUserHandler(userServ)
	userGrpcHand := grpcHand.NewUserHandler(userServ)

	// middleware
	deps.Router.Use(middleware.ErrorHandler)

	// routes
	r := routes.Router{
		Router:      deps.Router,
		UserHandler: userHand,
	}
	r.New()

	// subcriber
	userSub := subscriber.NewUserSubscriber(deps.Ch, deps.Logger, userServ)
	userSub.UserRegistered()
	userSub.UserAvatar()

	// grpc register
	reflection.Register(deps.S)
	pb.RegisterUserServiceServer(deps.S, userGrpcHand)
}
