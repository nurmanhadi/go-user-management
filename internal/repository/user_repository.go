package repository

import (
	"user-management/internal/entity"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}
func (r *UserRepository) Save(user *entity.User) error {
	return r.db.Save(user).Error
}
func (r *UserRepository) FindByUsername(username string) (*entity.User, error) {
	user := new(entity.User)
	err := r.db.Where("username = ?", username).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (r *UserRepository) FindById(id int64) (*entity.User, error) {
	user := new(entity.User)
	err := r.db.Where("id = ?", id).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (r *UserRepository) CountByEmail(email string) (int64, error) {
	var count int64
	err := r.db.Model(&entity.User{}).Where("email = ?", email).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
func (r *UserRepository) CountByPhone(phone string) (int64, error) {
	var count int64
	err := r.db.Model(&entity.User{}).Where("phone = ?", phone).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}
