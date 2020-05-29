package user

import (
	"context"
	"sync"

	"github.com/20326/vega/app/model"
	"github.com/jinzhu/gorm"
)

// New returns a new UserService.
func New(db *gorm.DB) model.UserService {
	return &userService{
		db:    db,
		mutex: &sync.Mutex{},
	}
}

type userService struct {
	db    *gorm.DB
	mutex *sync.Mutex
}

// Find returns a user from the datastore.
func (s *userService) Find(ctx context.Context, id uint64) (*model.User, error) {
	out := &model.User{}

	if err := s.db.First(out, id).Error; nil != err {
		return nil, err
	}
	return out, nil
}

// FindName returns a user from the datastore by username.
func (s *userService) FindName(ctx context.Context, username string) (*model.User, error) {
	out := &model.User{}

	if err := s.db.Where("`username` = ?", username).First(out).Error; nil != err {
		return nil, err
	}

	return out, nil
}

// FindToken returns a user from the datastore by token.
func (s *userService) FindToken(ctx context.Context, token string) (*model.User, error) {
	out := &model.User{}

	if err := s.db.Where("`token` = ?", token).First(out).Error; nil != err {
		return nil, err
	}

	return out, nil
}

// List returns a list of users from the datastore.
func (s *userService) List(ctx context.Context) ([]*model.User, error) {
	var err error
	var out []*model.User

	if err = s.db.Model(&model.User{}).Order("`id` DESC").Find(&out).Error; nil != err {
	}

	return out, err
}

// Create persists a new user to the datastore.
func (s *userService) Create(ctx context.Context, user *model.User) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	tx := s.db.Begin()
	if err := tx.Create(user).Error; nil != err {
		tx.Rollback()

		return err
	}
	tx.Commit()

	return nil
}

// Update persists an updated user to the datastore. col map[string]interface{}
func (s *userService) Update(ctx context.Context, user *model.User) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	var err error

	tx := s.db.Begin()
	defer func() {
		if err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()

	if err = tx.Save(user).Error; nil != err {
		return err
	}
	return nil
}

// Delete deletes a user from the datastore.
func (s *userService) Delete(ctx context.Context, user *model.User) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	out := &model.User{}

	tx := s.db.Begin()
	if err := s.db.Where("`id` = ?", user.ID).First(out).Error; nil != err {
		tx.Rollback()

		return err
	}
	if err := tx.Delete(out).Error; nil != err {
		tx.Rollback()

		return err
	}
	tx.Commit()

	return nil
}

// Count returns a count of active users.
func (s *userService) Count(ctx context.Context) (int, error) {
	var err error
	var out int

	if err = s.db.Model(&model.User{}).Order("`id` DESC").Count(&out).Error; nil != err {
	}

	return out, err
}
