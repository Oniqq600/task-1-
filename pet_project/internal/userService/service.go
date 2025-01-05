package orm

type UserService struct {
	repo UsersRepository
}

func NewUserService(repo *userRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetAllUsers() ([]Users, error) {
	return s.repo.GetUsers()
}

func (s *UserService) CreateUser(user Users) (Users, error) {
	return s.repo.PostUser(user)
}

func (s *UserService) UpdateUserByID(id uint, user Users) (Users, error) {
	return s.repo.PatchUserByID(id, user)
}

func (s *UserService) DeleteUserByID(id uint) error {
	return s.repo.DeleteUserByID(id)
}
