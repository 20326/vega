package config

type Resource struct {
	ID       string
	ActionID uint64
	Method   string
	Path     string
}

type Action struct {
	ID           string
	Action       string
	Describe     string
	Path         string
	Resources    []Resource
	DefaultCheck bool
	Status       int
	Deleted      int
}

type Permission struct {
	ID           string
	Name         string
	Describe     string
	Icon         string
	Path         string
	Actions      []Action
	DefaultCheck bool
	Status       int
	Deleted      int
}

type Role struct {
	ID           string
	Name         string
	Describe     string
}

type InitData struct {
	Roles       []Role
	Permissions []Permission
}
