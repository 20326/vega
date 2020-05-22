package action

import (
	"context"
	"sync"

	"github.com/20326/vega/app/model"
	"github.com/jinzhu/gorm"
	"github.com/phuslu/log"
)

// New returns a new ActionService.
func New(db *gorm.DB) model.ActionService {
	return &actionService{
		db:    db,
		mutex: &sync.Mutex{},
	}
}

type actionService struct {
	db    *gorm.DB
	mutex *sync.Mutex
}

// Find returns a action from the datastore.
func (s *actionService) Find(ctx context.Context, id uint64) (*model.Action, error) {
	out := &model.Action{}

	if err := s.db.First(out, id).Error; nil != err {
		return nil, err
	}
	return out, nil
}

// List returns a list of actions from the datastore.
func (s *actionService) List(ctx context.Context) ([]*model.Action, error) {
	var err error
	var out []*model.Action

	if err = s.db.Model(&model.Action{}).Order("`id` DESC").Find(&out).Error; nil != err {
		log.Error().Err(err).Msg("get actions failed")
	}

	return out, err
}

// Create persists a new action to the datastore.
func (s *actionService) Create(ctx context.Context, action *model.Action) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	tx := s.db.Begin()
	if err := tx.Create(action).Error; nil != err {
		tx.Rollback()

		return err
	}
	tx.Commit()

	return nil
}

// Update persists an updated action to the datastore. col map[string]interface{}
func (s *actionService) Update(ctx context.Context, action *model.Action) error {
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

	if err = tx.Save(action).Error; nil != err {
		return err
	}
	return nil
}

// Delete deletes a action from the datastore.
func (s *actionService) Delete(ctx context.Context, action *model.Action) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	out := &model.Action{}

	tx := s.db.Begin()
	if err := s.db.Where("`id` = ?", action.ID).First(out).Error; nil != err {
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

// Count returns a count of active actions.
func (s *actionService) Count(ctx context.Context) (int, error) {
	var err error
	var out int

	if err = s.db.Model(&model.Action{}).Order("`id` DESC").Count(&out).Error; nil != err {
		log.Error().Err(err).Msg("get actions failed")
	}

	return out, err
}
