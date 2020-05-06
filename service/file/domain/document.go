package domain

type Document struct {
	ID          string
	Name        string
	Path        string
	Slug        string
	Type        string
	Description string
	Info        string
	Cdn         bool
}
