package domain

type ACL struct {
	ID        string
	Handler   string
	IsPublic  bool
	Title     string
	Permitted []Role
}
