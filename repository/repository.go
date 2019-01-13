package repository

type Repository struct{}

func (r Repository) GetData() string {
	return "Data from Repository"
}
