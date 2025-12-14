package users

type UserType int

const (
	CUSTOMER UserType = iota
	ADMIN
	OWNER
)

func (ut UserType) String() string {
	switch ut {
	case CUSTOMER:
		return "CUSTOMER"
	case ADMIN:
		return "ADMIN"
	case OWNER:
		return "OWNER"
	default:
		return "UNKNOWN"
	}
}
