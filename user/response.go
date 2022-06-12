package user

type UserResponse struct {
	Id    int32  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func ConvertToResponse(user User) UserResponse {
	var userResponse UserResponse

	userResponse.Id = user.Id
	userResponse.Name = user.Name
	userResponse.Email = user.Email

	return userResponse
}
