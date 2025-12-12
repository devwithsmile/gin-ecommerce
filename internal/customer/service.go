package customer

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Signup(ctx context.Context, req SignupRequest) (Customer, error)
	Login(ctx context.Context, req LoginRequest) (Customer, error)
	GetCustomer(email string) (Customer, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) Signup(ctx context.Context, req SignupRequest) (Customer, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 12)
	if err != nil {
		return Customer{}, errors.New("error in password business logic")
	}
	customer := Customer{
		Email:    req.Email,
		Name:     req.Name,
		Password: string(hashedPassword),
	}
	return s.repo.Create(ctx, customer)
}

func (s *service) Login(ctx context.Context, req LoginRequest) (Customer, error) {
	return Customer{}, nil
}

func (s *service) GetCustomer(email string) (Customer, error) {
	return s.repo.GetByEmail(email)
}
