package customer

import (
	"context"
	"devwithsmile/gin-ecommerce/internal/auth"
	"errors"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Signup(ctx context.Context, req SignupRequest) (Customer, error)
	Login(ctx context.Context, req LoginRequest) (auth.TokenPair, error)
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
		ID:       uuid.NewString(),
		Email:    req.Email,
		Name:     req.Name,
		Password: string(hashedPassword),
	}
	return s.repo.Create(ctx, customer)
}

func (s *service) Login(ctx context.Context, req LoginRequest) (auth.TokenPair, error) {
	//find if customer with email exists
	customer, err := s.repo.GetByEmail(req.Email)
	if err != nil {
		return auth.TokenPair{}, errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(req.Password))
	// see if password is correct
	if err != nil {
		return auth.TokenPair{}, errors.New("invalid credentials")
	}
	// generate and return token
	tokenpair, err := auth.GenerateTokens(customer.ID)
	if err != nil {
		return auth.TokenPair{}, err
	}
	return *tokenpair, nil

}

func (s *service) GetCustomer(email string) (Customer, error) {
	return s.repo.GetByEmail(email)
}
