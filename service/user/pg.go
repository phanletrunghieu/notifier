package user

import (
	"github.com/jinzhu/gorm"
	"github.com/phanletrunghieu/notifier/domain"
)

type pgService struct {
	db *gorm.DB
}

// NewPGService .
func NewPGService(db *gorm.DB) Service {
	return &pgService{
		db: db,
	}
}

// Create implement Create for User service
func (s *pgService) Create(p *domain.User) error {
	return s.db.Create(p).Error
}

// Update implement Update for User service
func (s *pgService) Update(p *domain.User) (*domain.User, error) {
	old := domain.User{Model: domain.Model{ID: p.ID}}
	if err := s.db.Find(&old).Error; err != nil {
		return nil, err
	}

	old.Name = p.Name

	return &old, s.db.Save(&old).Error
}

// Find implement Find for User service
func (s *pgService) Find(p *domain.User) (*domain.User, error) {
	res := p
	if err := s.db.Find(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}

// FindAll implement FindAll for User service
func (s *pgService) FindAll() ([]domain.User, error) {
	res := []domain.User{}
	return res, s.db.Find(&res).Error
}

// Delete implement Delete for User service
func (s *pgService) Delete(p *domain.User) error {
	old := domain.User{Model: domain.Model{ID: p.ID}}
	if err := s.db.Find(&old).Error; err != nil {
		return err
	}
	return s.db.Delete(old).Error
}
