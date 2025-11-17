package config

import (
	"user-management/delivery/rest/handler"
	"user-management/delivery/rest/middleware"
	"user-management/delivery/rest/routes"
	"user-management/internal/event/subscriber"
	"user-management/internal/repository"
	"user-management/internal/service"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type Bootstrap struct {
	DB        *gorm.DB
	Cache     *memcache.Client
	Logger    zerolog.Logger
	Validator *validator.Validate
	Router    *chi.Mux
	Ch        *amqp.Channel
}

func Initialize(deps *Bootstrap) {
	// publisher

	// cache

	// repository
	userRepo := repository.NewUserRepository(deps.DB)

	// service
	userServ := service.NewUserService(deps.Logger, deps.Validator, userRepo)

	// handler
	userHand := handler.NewUserHandler(userServ)

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
}
