package service

import (
	"net/http"
	"strconv"
	"strings"
	"time"
	"user-management/internal/entity"
	"user-management/internal/repository"
	"user-management/pkg/dto"
	"user-management/pkg/response"

	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog"
	"gorm.io/gorm"
)

type UserService struct {
	logger         zerolog.Logger
	validator      *validator.Validate
	userRepository *repository.UserRepository
}

func NewUserService(logger zerolog.Logger, validator *validator.Validate, userRepository *repository.UserRepository) *UserService {
	return &UserService{
		logger:         logger,
		validator:      validator,
		userRepository: userRepository,
	}
}

// message queue
func (s *UserService) UserCreate(payload *dto.EventUserPayload[dto.EventUserData]) error {
	if err := s.validator.Struct(payload); err != nil {
		s.logger.Warn().Err(err).Msg("failed to validate request")
		return err
	}
	user := &entity.User{
		AuthID:    payload.Data.UserId,
		Username:  payload.Data.Username,
		CreatedAt: payload.Data.Registered_at,
	}
	if err := s.userRepository.Save(user); err != nil {
		s.logger.Error().Err(err).Msg("failed save to database")
		return err
	}
	s.logger.Info().Str("event", payload.Event).Msg("user create success")
	return nil
}
func (s *UserService) UserUpdateAvatar(payload *dto.EventUserPayload[dto.EventUserAvatar]) error {
	if err := s.validator.Struct(payload); err != nil {
		s.logger.Warn().Err(err).Msg("failed to validate request")
		return err
	}
	user, err := s.userRepository.FindById(payload.Data.UserId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			s.logger.Warn().Err(err).Msg("user not found")
			return response.Except(http.StatusNotFound, "user not found")
		}
		s.logger.Error().Err(err).Msg("failed find by id to database")
		return err
	}
	user.AvatarURL = &payload.Data.AvatarUrl
	if err := s.userRepository.Save(user); err != nil {
		s.logger.Error().Err(err).Msg("failed save to database")
		return err
	}
	s.logger.Info().Str("user_id", strconv.Itoa(int(payload.Data.UserId))).Msg("user update avatar success")
	return nil
}

// api
func (s *UserService) UserChangeProfile(id string, request *dto.UserUpdateRequest) error {
	if err := s.validator.Struct(request); err != nil {
		s.logger.Warn().Err(err).Msg("failed to validate request")
		return err
	}
	newId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		s.logger.Error().Err(err).Msg("failed parse string to int64")
		return err
	}
	user, err := s.userRepository.FindById(newId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			s.logger.Warn().Err(err).Msg("user not found")
			return response.Except(http.StatusNotFound, "user not found")
		}
		s.logger.Error().Err(err).Msg("failed find by id to database")
		return err
	}
	if request.FirstName != nil {
		user.FirstName = request.FirstName
	}
	if request.LastName != nil {
		user.LastName = request.LastName
	}
	if request.Email != nil {
		newEmail := strings.ToLower(*request.Email)
		totalEmail, err := s.userRepository.CountByEmail(newEmail)
		if err != nil {
			s.logger.Error().Err(err).Msg("failed count by email to database")
			return err
		}
		if totalEmail > 0 {
			s.logger.Warn().Msg("email already exists")
			return response.Except(http.StatusConflict, "email already exists")
		}
		user.Email = &newEmail
	}
	if request.Phone != nil {
		user.Phone = request.Phone
	}
	if request.BirthDate != nil {
		t, err := time.Parse("2006-01-02", *request.BirthDate)
		if err != nil {
			s.logger.Error().Err(err).Msg("failed parse time format 2006-01-02")
			return err
		}
		user.BirthDate = &t
	}
	if request.Gender != nil {
		user.Gender = request.Gender
	}
	if request.Bio != nil {
		user.Bio = request.Bio
	}
	if request.Description != nil {
		user.Description = request.Description
	}
	if err := s.userRepository.Save(user); err != nil {
		s.logger.Error().Err(err).Msg("failed save to database")
		return err
	}
	s.logger.Info().Str("user_id", id).Msg("user change profile success")
	return nil
}
func (s *UserService) UserGetByUsername(username string) (*dto.UserResponse, error) {
	user, err := s.userRepository.FindByUsername(username)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			s.logger.Warn().Err(err).Msg("user not found")
			return nil, response.Except(http.StatusNotFound, "user not found")
		}
		s.logger.Error().Err(err).Msg("failed find by username to database")
		return nil, err
	}
	resp := &dto.UserResponse{
		ID:       user.ID,
		AuthID:   user.AuthID,
		Username: user.Username,
		Name: &dto.Username{
			FirstName: user.FirstName,
			LastName:  user.LastName,
		},
		Contact: &dto.UserContact{
			Email: user.Email,
			Phone: user.Phone,
		},
		About: &dto.UserAbout{
			Bio:         user.Bio,
			Description: user.Description,
			BirthDate:   user.BirthDate,
			Gender:      user.Gender,
		},
		Verification: &dto.UserVerification{
			EmailVerifiedAt: user.EmailVerifiedAt,
			PhoneVerifiedAt: user.PhoneVerifiedAt,
		},
		AvatarURL: user.AvatarURL,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
	s.logger.Info().Str("username", username).Msg("user get by username success")
	return resp, nil
}
