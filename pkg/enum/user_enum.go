package enum

type UserRole string

const (
	Admin  UserRole = "admin"
	Member UserRole = "member"
	Guest  UserRole = "guest"
)