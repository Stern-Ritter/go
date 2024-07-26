package model

type User struct {
	ID       uint64
	Name     string
	Email    string
	Password string
}

type CreateUserDto struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserDto struct {
	ID    uint64 `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func CreateUserDtoToUser(dto CreateUserDto) User {
	return User{
		Name:     dto.Name,
		Email:    dto.Email,
		Password: dto.Password,
	}
}

func UserToUserDto(user User) UserDto {
	return UserDto{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}
