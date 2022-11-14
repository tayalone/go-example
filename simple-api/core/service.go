package core

/*
Srv is Struct of instance of Service
*/
type Srv struct {
	r Repo
}

var s Srv

/*New return instace of Srv
 */
func New(r Repo) Service {
	s.r = r
	return &s
}

/*GetAll return all of todo from Repo*/
func (s *Srv) GetAll() []Todo {
	return s.r.Get()
}

/*GetByID return Todo or error*/
func (s *Srv) GetByID(id uint) (Todo, error) {
	td, err := s.r.GetByPk(id)
	if err != nil {
		return Todo{}, err
	}
	return td, nil
}

/*Make return Todo which created from repo*/
func (s *Srv) Make(title string, desc string) Todo {
	return s.r.Create(title, desc)
}

/*UpdateByID do update Todo Data */
func (s *Srv) UpdateByID(id uint, title *string, desc *string, done *bool) error {
	err := s.r.UpdateByPk(id, title, desc, done)
	return err
}

/*RemoveByID remove Todo By PK */
func (s *Srv) RemoveByID(id uint) error {
	return s.r.RemoveByPK(id)
}
