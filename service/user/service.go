package user

import (
	"github.com/phanletrunghieu/notifier/domain"
)

// Service interface for project service
type Service interface {
	Create(p *domain.User) error
	Update(p *domain.User) (*domain.User, error)
	Find(p *domain.User) (*domain.User, error)
	FindAll() ([]domain.User, error)
	Delete(p *domain.User) error
}
