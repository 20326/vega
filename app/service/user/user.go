package user

import (
	"context"
	"sync"

	"github.com/20326/vega/app/model"
	"github.com/20326/vega/pkg/pagination"
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

	if err := s.db.Preload("Roles").First(out, id).Error; nil != err {
		return nil, err
	}
	return out, nil
}

// FindName returns a user from the datastore by username.
func (s *userService) FindName(ctx context.Context, username string) (*model.User, error) {
	out := &model.User{}

	if err := s.db.Preload("Roles").Where("`username` = ?", username).First(out).Error; nil != err {
		return nil, err
	}

	return out, nil
}

// FindToken returns a user from the datastore by token.
func (s *userService) FindToken(ctx context.Context, token string) (*model.User, error) {
	out := &model.User{}

	if err := s.db.Preload("Roles").Where("`token` = ?", token).First(out).Error; nil != err {
		return nil, err
	}

	return out, nil
}

// FindWhere returns a list of users by query params from the datastore.
func (s *userService) FindWhere(query model.PageQuery, roles []string) (out []*model.User, pagination pagination.Pagination) {
	offset := (query.PageNo - 1) * query.PageSize
	count := 0

	var err error

	tx := s.db.Model(&model.User{}).Preload("Roles")
	if 0 < len(roles) {
		subQuery := s.db.Model(&model.UserRole{}).Select("user_id").Where("role_id IN (?)", roles).QueryExpr()
		tx = tx.Where("id IN (?)", subQuery)
		// tx = tx.Preload("Roles").Joins("INNER JOIN vega_user_role on vega_user_role.user_id = vega_users.id AND vega_user_role.role_id in (?)", roles)
	} else {
		tx = tx.Preload("Roles")
	}
	if "" != query.Where {
		tx = tx.Where(query.Where, query.WhereArgs...)
	}
	if err = tx.Count(&count).Offset(offset).Limit(query.PageSize).
		Order("`id` DESC").Find(&out).Error; nil != err {
	}

	//for _, user := range out {
	//	user.FillRoleList()
	//}

	pagination = pagination.NewPagination(query.PageNo, query.PageSize, count)

	return
}

// List returns a list of users from the datastore.
func (s *userService) List(ctx context.Context) ([]*model.User, error) {
	var err error
	var out []*model.User

	if err = s.db.Model(&model.User{}).
		Preload("Roles").
		Order("`id` DESC").
		Find(&out).Error; nil != err {
	}

	//for _, user := range out {
	//	user.FillRoleList()
	//}

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

// Update columns persists an updated user to the datastore. col map[string]interface{}
func (s *userService) Updates(ctx context.Context, user *model.User, values interface{}) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	tx := s.db.Begin()
	if err := s.db.Model(user).Updates(values).Error; nil != err {
		tx.Rollback()

		return err
	}
	tx.Commit()
	return nil
}

// Delete deletes a user from the datastore.
func (s *userService) Delete(ctx context.Context, id uint64) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	out := &model.User{}

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

// RelatedClear will clear users.
func (s *userService) RelatedClear(ctx context.Context, user *model.User) {
	s.db.Model(&user).Association("Roles").Clear()
}

// Count returns a count of active users.
func (s *userService) Count(ctx context.Context) (int, error) {
	var err error
	var out int

	if err = s.db.Model(&model.User{}).Order("`id` DESC").Count(&out).Error; nil != err {
	}

	return out, err
}
