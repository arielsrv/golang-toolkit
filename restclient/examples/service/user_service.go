package service

type IUserService interface {
	GetUsers() ([]UserDto, error)
	GetUser(userID int64) (*UserDto, error)
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
	ID     int64  `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Email  string `json:"email,omitempty"`
	Gender string `json:"gender,omitempty"`
	Status string `json:"status,omitempty"`
}

type UserRequest struct {
	Name   string `json:"name,omitempty"`
	Email  string `json:"email,omitempty"`
	Gender string `json:"gender,omitempty"`
	Status string `json:"status,omitempty"`
}

type UserService struct {
	userClient IUserClient
}

func (service UserService) CreateUser(userDto UserDto) (UserDto, error) {
	userRequest := &UserRequest{
		Name:   userDto.FullName,
		Email:  userDto.Email,
		Gender: userDto.Gender,
		Status: userDto.Status,
	}

	id, err := service.userClient.CreateUser(*userRequest)
	if err != nil {
		return userDto, err
	}

	userDto.ID = id

	return userDto, nil
}

func NewUserService(userClient IUserClient) *UserService {
	return &UserService{userClient: userClient}
}

func (service UserService) GetUser(userID int64) (*UserDto, error) {
	response, err := service.userClient.GetUser(userID)
	if err != nil {
		return nil, err
	}

	userDto := UserDto{
		ID:       response.ID,
		FullName: response.Name,
		Email:    response.Email,
		Gender:   response.Gender,
		Status:   response.Status,
	}

	return &userDto, nil
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
