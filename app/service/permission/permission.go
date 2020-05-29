package permission

import (
	"context"
	"sync"

	"github.com/20326/vega/app/model"
	"github.com/jinzhu/gorm"
)

// New returns a new PermissionService.
func New(db *gorm.DB) model.PermissionService {
	return &permissionService{
		db:    db,
		mutex: &sync.Mutex{},
	}
}

type permissionService struct {
	db    *gorm.DB
	mutex *sync.Mutex
}

// Find returns a permission from the datastore.
func (s *permissionService) Find(ctx context.Context, id uint64) (*model.Permission, error) {
	out := &model.Permission{}

	if err := s.db.First(out, id).Error; nil != err {
		return nil, err
	}
	return out, nil
}

// List returns a list of permissions from the datastore.
func (s *permissionService) List(ctx context.Context) ([]*model.Permission, error) {
	var err error
	var out []*model.Permission

	if err = s.db.Model(&model.Permission{}).Order("`id` DESC").Find(&out).Error; nil != err {
	}

	return out, err
}

// Create persists a new permission to the datastore.
func (s *permissionService) Create(ctx context.Context, permission *model.Permission) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	tx := s.db.Begin()
	if err := tx.Create(permission).Error; nil != err {
		tx.Rollback()

		return err
	}
	tx.Commit()

	return nil
}

// Update persists an updated permission to the datastore. col map[string]interface{}
func (s *permissionService) Update(ctx context.Context, permission *model.Permission) error {
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

	if err = tx.Save(permission).Error; nil != err {
		return err
	}
	return nil
}

// Delete deletes a permission from the datastore.
func (s *permissionService) Delete(ctx context.Context, id uint64) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	out := &model.Permission{}

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

// Count returns a count of active permissions.
func (s *permissionService) Count(ctx context.Context) (int, error) {
	var err error
	var out int

	if err = s.db.Model(&model.Permission{}).Order("`id` DESC").Count(&out).Error; nil != err {
	}

	return out, err
}
