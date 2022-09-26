package service

type IUserService interface {
	GetUsers() ([]UserDto, error)
	CreateUser(userDto UserDto) error
}

type UserDto struct {
	ID       int64  `json:"id,omitempty"`
	FullName string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Gender   string `json:"gender,omitempty"`
	Status   string `json:"status,omitempty"`
}

type UserResponse struct {
	ID   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type UserRequest struct {
	ID     int64  `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Email  string `json:"email,omitempty"`
	Gender string `json:"gender,omitempty"`
	Status string `json:"status,omitempty"`
}

type UserService struct {
	userClient IUserClient
}

func (service UserService) CreateUser(userDto UserDto) error {
	var userRequest UserRequest
	userRequest.ID = userDto.ID
	userRequest.Name = userDto.FullName
	userRequest.Email = userDto.Email
	userRequest.Gender = userDto.Gender
	userRequest.Status = userDto.Status

	err := service.userClient.CreateUser(userRequest)
	if err != nil {
		return err
	}
	return nil
}

func NewUserService(userClient IUserClient) *UserService {
	return &UserService{userClient: userClient}
}

func (service UserService) GetUsers() ([]UserDto, error) {
	response, err := service.userClient.GetUsers()

	if err != nil {
		return nil, err
	}

	var result []UserDto
	for _, userResponse := range response {
		userDto := UserDto{
			ID:       userResponse.ID,
			FullName: userResponse.Name,
		}
		result = append(result, userDto)
	}

	return result, nil
}
