package device

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

// Create implement Create for Device service
func (s *pgService) Create(p *domain.Device) error {
	return s.db.Create(p).Error
}

// Update implement Update for Device service
func (s *pgService) Update(p *domain.Device) (*domain.Device, error) {
	old := domain.Device{Model: domain.Model{ID: p.ID}}
	if err := s.db.Find(&old).Error; err != nil {
		return nil, err
	}

	if p.Name != "" {
		old.Name = p.Name
	}

	if p.Token != "" {
		old.Token = p.Token
	}

	return &old, s.db.Save(&old).Error
}

// Find implement Find for Device service
func (s *pgService) Find(p *domain.Device) (*domain.Device, error) {
	res := p
	if err := s.db.Find(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}

// FindAll implement FindAll for Device service
func (s *pgService) FindAll() ([]domain.Device, error) {
	res := []domain.Device{}
	return res, s.db.Find(&res).Error
}

// Delete implement Delete for Device service
func (s *pgService) Delete(p *domain.Device) error {
	old := domain.Device{Model: domain.Model{ID: p.ID}}
	if err := s.db.Find(&old).Error; err != nil {
		return err
	}
	return s.db.Delete(old).Error
}
