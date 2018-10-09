package device

import (
	"github.com/phanletrunghieu/notifier/domain"
)

// Service interface for project service
type Service interface {
	Create(p *domain.Device) error
	Update(p *domain.Device) (*domain.Device, error)
	Find(p *domain.Device) (*domain.Device, error)
	FindAll() ([]domain.Device, error)
	Delete(p *domain.Device) error
}
