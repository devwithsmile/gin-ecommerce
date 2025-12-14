package customer

import (
	"devwithsmile/gin-ecommerce/internal/users"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Customer struct {
	users.BaseUser
	users.Creds
}

// NewCustomer is the constructor - this is standard Go practice
func NewCustomer(email, name, password string) (Customer, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return Customer{}, err
	}

	return Customer{
		BaseUser: users.BaseUser{
			ID:        uuid.NewString(),
			Email:     email,
			Name:      name,
			CreatedAt: time.Now(),
		},
		Creds: users.Creds{
			PasswordHash: string(hashedPassword),
		},
	}, nil
}
