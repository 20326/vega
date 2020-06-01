package role

import (
	"context"
	"sync"

	"github.com/20326/vega/app/model"
	"github.com/jinzhu/gorm"
)

// New returns a new RoleService.
func New(db *gorm.DB) model.RoleService {
	return &roleService{
		db:    db,
		mutex: &sync.Mutex{},
	}
}

type roleService struct {
	db    *gorm.DB
	mutex *sync.Mutex
}

// Find returns a role from the datastore.
func (s *roleService) Find(ctx context.Context, id uint64) (*model.Role, error) {
	out := &model.Role{}

	if err := s.db.First(out, id).Error; nil != err {
		return nil, err
	}
	return out, nil
}

// FindName returns a role from the datastore.
func (s *roleService) FindName(ctx context.Context, name string) (*model.Role, error) {
	out := &model.Role{}

	if err := s.db.Where("`name` = ?", name).First(out).Error; nil != err {
		return nil, err
	}
	return out, nil
}

// List returns a list of roles from the datastore.
func (s *roleService) List(ctx context.Context) ([]*model.Role, error) {
	var err error
	var out []*model.Role

	if err = s.db.Model(&model.Role{}).Order("`id` DESC").Find(&out).Error; nil != err {
	}

	return out, err
}

// Create persists a new role to the datastore.
func (s *roleService) Create(ctx context.Context, role *model.Role) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	tx := s.db.Begin()
	if err := tx.Create(role).Error; nil != err {
		tx.Rollback()

		return err
	}
	tx.Commit()

	return nil
}

// Update persists an updated role to the datastore. col map[string]interface{}
func (s *roleService) Update(ctx context.Context, role *model.Role) error {
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

	if err = tx.Save(role).Error; nil != err {
		return err
	}
	return nil
}

// Delete deletes a role from the datastore.
func (s *roleService) Delete(ctx context.Context, id uint64) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	out := &model.Role{}

	tx := s.db.Begin()
	if err := s.db.Where("`id` = ?", id).First(out).Error; nil != err {
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

// Count returns a count of active roles.
func (s *roleService) Count(ctx context.Context) (int, error) {
	var err error
	var out int

	if err = s.db.Model(&model.Role{}).Order("`id` DESC").Count(&out).Error; nil != err {
	}

	return out, err
}
