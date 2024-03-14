package types

type Client interface {
	Admin() Admin
	Users() Users
	Domains() Domains
	Lists() Lists
	Aliases() Aliases
}

type Credentials struct {
	Username string
	Password string
}
