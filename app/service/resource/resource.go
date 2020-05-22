package resource

import (
	"context"
	"sync"

	"github.com/20326/vega/app/model"
	"github.com/jinzhu/gorm"
	"github.com/phuslu/log"
)

// New returns a new ResourceService.
func New(db *gorm.DB) model.ResourceService {
	return &resourceService{
		db:    db,
		mutex: &sync.Mutex{},
	}
}

type resourceService struct {
	db    *gorm.DB
	mutex *sync.Mutex
}

// Find returns a resource from the datastore.
func (s *resourceService) Find(ctx context.Context, id uint64) (*model.Resource, error) {
	out := &model.Resource{}

	if err := s.db.First(out, id).Error; nil != err {
		return nil, err
	}
	return out, nil
}

// List returns a list of resources from the datastore.
func (s *resourceService) List(ctx context.Context) ([]*model.Resource, error) {
	var err error
	var out []*model.Resource

	if err = s.db.Model(&model.Resource{}).Order("`id` DESC").Find(&out).Error; nil != err {
		log.Error().Err(err).Msg("get resources failed")
	}

	return out, err
}

// Create persists a new resource to the datastore.
func (s *resourceService) Create(ctx context.Context, resource *model.Resource) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	tx := s.db.Begin()
	if err := tx.Create(resource).Error; nil != err {
		tx.Rollback()

		return err
	}
	tx.Commit()

	return nil
}

// Update persists an updated resource to the datastore. col map[string]interface{}
func (s *resourceService) Update(ctx context.Context, resource *model.Resource) error {
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

	if err = tx.Save(resource).Error; nil != err {
		return err
	}
	return nil
}

// Delete deletes a resource from the datastore.
func (s *resourceService) Delete(ctx context.Context, resource *model.Resource) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	out := &model.Resource{}

	tx := s.db.Begin()
	if err := s.db.Where("`id` = ?", resource.ID).First(out).Error; nil != err {
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

// Count returns a count of active resources.
func (s *resourceService) Count(ctx context.Context) (int, error) {
	var err error
	var out int

	if err = s.db.Model(&model.Resource{}).Order("`id` DESC").Count(&out).Error; nil != err {
		log.Error().Err(err).Msg("get resources failed")
	}

	return out, err
}
