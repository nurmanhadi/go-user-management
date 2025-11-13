package subscriber

import (
	"user-management/internal/service"
	"user-management/pkg/dto"
	"user-management/pkg/global"

	"github.com/goccy/go-json"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rs/zerolog"
)

type UserSubscriber struct {
	ch          *amqp.Channel
	logger      zerolog.Logger
	userService *service.UserService
}

func NewUserSubscriber(ch *amqp.Channel, logger zerolog.Logger, userService *service.UserService) *UserSubscriber {
	return &UserSubscriber{
		ch:          ch,
		logger:      logger,
		userService: userService,
	}
}
func (e *UserSubscriber) UserRegistered() {
	queue, err := e.ch.QueueDeclare(global.EventUserRegistered, true, false, false, false, nil)
	if err != nil {
		panic(err)
	}
	msgs, err := e.ch.Consume(
		queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}
	go func() {
		for x := range msgs {
			payload := new(dto.EventUserPayload[dto.EventUserData])
			if err := json.Unmarshal(x.Body, payload); err != nil {
				e.logger.Error().Err(err).Msg("failed unmarshal to payload")
				return
			}
			err := e.userService.UserCreate(payload)
			if err != nil {
				return
			}
		}
	}()
}
func (e *UserSubscriber) UserAvatar() {
	queue, err := e.ch.QueueDeclare(global.EventUserAvatar, true, false, false, false, nil)
	if err != nil {
		panic(err)
	}
	msgs, err := e.ch.Consume(
		queue.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}
	go func() {
		for x := range msgs {
			payload := new(dto.EventUserPayload[dto.EventUserAvatar])
			if err := json.Unmarshal(x.Body, payload); err != nil {
				e.logger.Error().Err(err).Msg("failed unmarshal to payload")
				return
			}
			err := e.userService.UserUpdateAvatar(payload)
			if err != nil {
				return
			}
		}
	}()
}
