package user

type UserService interface {
	FindAll() ([]User, error)
	FindById(ID int) (User, error)
	Create(user UserRequest) (User, error)
	Update(userReqUpdate UserRequestUpdate) (User, error)
	Delete(ID int) (User, error)
}

type service struct {
	repository UserRepository
}

func NewService(repository UserRepository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]User, error) {
	users, err := s.repository.FindAll()
	return users, err
}

func (s *service) FindById(ID int) (User, error) {
	user, err := s.repository.FindById(ID)
	return user, err
}

func (s *service) Create(userReq UserRequest) (User, error) {
	var user User
	user.Name = userReq.Name
	user.Email = userReq.Email
	user.Password = userReq.Password

	user, err := s.repository.Create(user)
	return user, err
}

func (s *service) Update(userReqUpdate UserRequestUpdate) (User, error) {
	result, err := s.repository.FindById(int(userReqUpdate.Id))

	result.Name = userReqUpdate.Name
	result.Email = userReqUpdate.Email
	result.Password = userReqUpdate.Password

	user, err := s.repository.Update(result)
	return user, err
}

func (s *service) Delete(id int) (User, error) {
	user, err := s.repository.FindById(id)

	user, err = s.repository.Delete(user)
	return user, err
}
