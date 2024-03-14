package types

type Users interface {
	GetUser()
	ListAllUsers()
	ListUsersByDomain(domain string)
	CreateUser()
}
