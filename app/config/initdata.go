package config

type Resource struct {
	ID       string
	ActionID uint64
	Method   string
	Path     string
}

type Action struct {
	ID           string
	Name         string
	Describe     string
	Resources    []Resource
	DefaultCheck bool
	Status       int
	Deleted      int
}

type Permission struct {
	ID           string
	Name         string
	Label        string
	Describe     string
	Icon         string
	Path         string
	Actions      []Action
	DefaultCheck bool
	Status       int
	Deleted      int
}


type User struct {
	ID       string
	Username string
	Nickname string
	Password string
}

type Role struct {
	ID       string
	Name     string
	Label    string
	Describe string
	Users	 []User
}

type InitData struct {
	Roles       []Role
	Permissions []Permission
}
