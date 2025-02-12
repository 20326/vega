package setting

import (
	"context"
	"sync"

	"github.com/20326/vega/app/model"
	"github.com/jinzhu/gorm"
)

// New returns a new SettingService.
func New(db *gorm.DB) model.SettingService {
	return &settingService{
		db:    db,
		mutex: &sync.Mutex{},
	}
}

type settingService struct {
	db    *gorm.DB
	mutex *sync.Mutex
}

// Find returns a setting from the datastore.
func (s *settingService) Find(ctx context.Context, id uint64) (*model.Setting, error) {
	out := &model.Setting{}

	if err := s.db.First(out, id).Error; nil != err {
		return nil, err
	}
	return out, nil
}

// Find returns a setting from the datastore.
func (s *settingService) FindName(ctx context.Context, name string) (*model.Setting, error) {
	out := &model.Setting{}

	if err := s.db.Where("`name` = ?", name).First(out).Error; nil != err {
		return nil, err
	}
	return out, nil
}

// Find returns a setting from the datastore.
func (s *settingService) FindLike(ctx context.Context, like string) ([]*model.Setting, error) {
	var out []*model.Setting

	if err := s.db.Where(" `name` LIKE ? ", like+"%").Find(&out).Error; nil != err {
		return nil, err
	}
	return out, nil
}

// List returns a list of settings from the datastore.
func (s *settingService) List(ctx context.Context) ([]*model.Setting, error) {
	var err error
	var out []*model.Setting

	if err = s.db.Model(&model.Setting{}).Order("`id` DESC").Find(&out).Error; nil != err {
	}

	return out, err
}

// Create persists a new setting to the datastore.
func (s *settingService) Create(ctx context.Context, setting *model.Setting) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	tx := s.db.Begin()
	if err := tx.Create(setting).Error; nil != err {
		tx.Rollback()

		return err
	}
	tx.Commit()

	return nil
}

// Update persists an updated setting to the datastore. col map[string]interface{}
func (s *settingService) Update(ctx context.Context, setting *model.Setting) error {
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

	if err = tx.Save(setting).Error; nil != err {
		return err
	}
	return nil
}

// Update persists an updated setting to the datastore. col map[string]interface{}
func (s *settingService) Updates(ctx context.Context, settings []*model.Setting) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	var err error
	tx := s.db.Begin()
	for _, setting := range settings {
		// add exists
		if _, err = s.FindName(ctx, setting.Name); nil != err {
			if err := tx.Create(setting).Error; nil != err {
				continue
			}
		} else {
			if err := tx.Model(&model.Setting{}).
				Where("`name` = ?", setting.Name).
				Select("value").
				Updates(map[string]interface{}{"value": setting.Value}).Error; nil != err {
				tx.Rollback()

				return err
			}
		}
	}
	tx.Commit()

	return nil
}

// Delete deletes a setting from the datastore.
func (s *settingService) Delete(ctx context.Context, id uint64) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	out := &model.Setting{}

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

// Count returns a count of active settings.
func (s *settingService) Count(ctx context.Context) (int, error) {
	var err error
	var out int

	if err = s.db.Model(&model.Setting{}).Order("`id` DESC").Count(&out).Error; nil != err {
	}

	return out, err
}
