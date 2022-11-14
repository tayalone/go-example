package core

/*Service is Method of Service*/
type Service interface {
	GetAll() []Todo
	GetByID(uint) (Todo, error)
	Make(string, string) Todo
	UpdateByID(uint, *string, *string, *bool) error
	RemoveByID(uint) error
}

/*Repo is Method of Repo*/
type Repo interface {
	Get() []Todo
	GetByPk(uint) (Todo, error)
	Create(string, string) Todo
	UpdateByPk(uint, *string, *string, *bool) error
	RemoveByPK(uint) error
}
