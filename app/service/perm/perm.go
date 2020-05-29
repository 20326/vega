package perm

import (
	"context"
	"sync"

	"github.com/20326/vega/app/model"
	"github.com/jinzhu/gorm"
)

// New returns a new PermService.
func New(db *gorm.DB) model.PermService {
	return &permService{
		db:    db,
		mutex: &sync.Mutex{},
	}
}

type permService struct {
	db    *gorm.DB
	mutex *sync.Mutex
}

// Find returns a perm from the datastore.
func (s *permService) Find(ctx context.Context, id uint64) (*model.Perm, error) {
	out := &model.Perm{}

	if err := s.db.First(out, id).Error; nil != err {
		return nil, err
	}
	return out, nil
}

// List returns a list of perms from the datastore.
func (s *permService) List(ctx context.Context) ([]*model.Perm, error) {
	var err error
	var out []*model.Perm

	if err = s.db.Model(&model.Perm{}).Order("`id` DESC").Find(&out).Error; nil != err {
	}

	return out, err
}

// Create persists a new perm to the datastore.
func (s *permService) Create(ctx context.Context, perm *model.Perm) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	tx := s.db.Begin()
	if err := tx.Create(perm).Error; nil != err {
		tx.Rollback()

		return err
	}
	tx.Commit()

	return nil
}

// Update persists an updated perm to the datastore. col map[string]interface{}
func (s *permService) Update(ctx context.Context, perm *model.Perm) error {
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

	if err = tx.Save(perm).Error; nil != err {
		return err
	}
	return nil
}

// Delete deletes a perm from the datastore.
func (s *permService) Delete(ctx context.Context, perm *model.Perm) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	out := &model.Perm{}

	tx := s.db.Begin()
	if err := s.db.Where("`id` = ?", perm.ID).First(out).Error; nil != err {
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

// Count returns a count of active perms.
func (s *permService) Count(ctx context.Context) (int, error) {
	var err error
	var out int

	if err = s.db.Model(&model.Perm{}).Order("`id` DESC").Count(&out).Error; nil != err {
	}

	return out, err
}
