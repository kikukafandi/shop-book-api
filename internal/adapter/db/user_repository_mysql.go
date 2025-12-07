package db

import (
	"context"

	"kikukafandi/book-shop-api/internal/domain"

	"gorm.io/gorm"
)

// UserModel is the database model for User.
type UserModel struct {
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"size:255;not null"`
	Email    string `gorm:"size:255;not null;unique"`
	Password string `gorm:"size:255;not null"`
	Role     string `gorm:"size:50;not null"`
}

// TableName returns the table name for UserModel.
func (UserModel) TableName() string {
	return "users"
}

// UserRepositoryMySQL implements domain.UserRepository using MySQL/GORM.
type UserRepositoryMySQL struct {
	db *gorm.DB
}

// NewUserRepositoryMySQL creates a new UserRepositoryMySQL.
func NewUserRepositoryMySQL(db *gorm.DB) *UserRepositoryMySQL {
	return &UserRepositoryMySQL{db: db}
}

// Save saves a user to database.
func (r *UserRepositoryMySQL) Save(ctx context.Context, user domain.User) (domain.User, error) {
	model := toUserModel(user)

	if err := r.db.WithContext(ctx).Create(&model).Error; err != nil {
		return domain.User{}, err
	}

	return toUserDomain(model), nil
}

// FindByID finds a user by ID.
func (r *UserRepositoryMySQL) FindByID(ctx context.Context, id uint) (domain.User, error) {
	var model UserModel

	if err := r.db.WithContext(ctx).First(&model, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return domain.User{}, domain.ErrUserNotFound
		}
		return domain.User{}, err
	}

	return toUserDomain(model), nil
}

// FindByEmail finds a user by email.
func (r *UserRepositoryMySQL) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	var model UserModel

	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&model).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return domain.User{}, domain.ErrUserNotFound
		}
		return domain.User{}, err
	}

	return toUserDomain(model), nil
}

// FindAll returns all users.
func (r *UserRepositoryMySQL) FindAll(ctx context.Context) ([]domain.User, error) {
	var models []UserModel

	if err := r.db.WithContext(ctx).Find(&models).Error; err != nil {
		return nil, err
	}

	users := make([]domain.User, len(models))
	for i, model := range models {
		users[i] = toUserDomain(model)
	}

	return users, nil
}

// Update updates a user in database.
func (r *UserRepositoryMySQL) Update(ctx context.Context, user domain.User) (domain.User, error) {
	model := toUserModel(user)

	if err := r.db.WithContext(ctx).Save(&model).Error; err != nil {
		return domain.User{}, err
	}

	return toUserDomain(model), nil
}

// Delete deletes a user from database.
func (r *UserRepositoryMySQL) Delete(ctx context.Context, id uint) error {
	if err := r.db.WithContext(ctx).Delete(&UserModel{}, id).Error; err != nil {
		return err
	}
	return nil
}

// ExistsByEmail checks if a user with given email exists.
func (r *UserRepositoryMySQL) ExistsByEmail(ctx context.Context, email string) (bool, error) {
	var count int64

	if err := r.db.WithContext(ctx).Model(&UserModel{}).Where("email = ?", email).Count(&count).Error; err != nil {
		return false, err
	}

	return count > 0, nil
}

// toUserModel converts domain.User to UserModel.
func toUserModel(user domain.User) UserModel {
	return UserModel{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		Role:     user.Role,
	}
}

// toUserDomain converts UserModel to domain.User.
func toUserDomain(model UserModel) domain.User {
	return domain.User{
		ID:       model.ID,
		Name:     model.Name,
		Email:    model.Email,
		Password: model.Password,
		Role:     model.Role,
	}
}
