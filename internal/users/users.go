package users

import "time"

type BaseUser struct {
	ID    string
	Name  string
	Email string
	Phone string
	Address
	CreatedAt time.Time
}
type Address struct {
	Line1      string
	Line2      string
	City       string
	State      string
	PostalCode string
	Country    string
}

func (u *BaseUser) GetID() string {
	return u.ID
}
