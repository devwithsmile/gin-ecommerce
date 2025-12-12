package customer

import (
	"context"
	"errors"
	"sync"
)

type Repository interface {
	Create(ctx context.Context, c Customer) (Customer, error)
	GetByEmail(email string) (Customer, error)
}

type memRepository struct {
	mu      sync.RWMutex
	records map[string]Customer // key: email
}

func NewMemRepository() Repository {
	return &memRepository{
		records: make(map[string]Customer),
	}
}

func (r *memRepository) Create(ctx context.Context, customer Customer) (Customer, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.records[customer.Email]; exists {
		return Customer{}, errors.New("customer already exists")
	}

	r.records[customer.Email] = customer
	return customer, nil
}

func (r *memRepository) GetByEmail(email string) (Customer, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	c, ok := r.records[email]
	if !ok {
		return Customer{}, errors.New("customer not found")
	}

	return c, nil
}
