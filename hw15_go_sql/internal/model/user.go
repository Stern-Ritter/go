package model

import "github.com/Stern-Ritter/go/hw15_go_sql/internal/utils"

type User struct {
	ID       int64
	Name     string
	Email    string
	Password string
}

type CreateUserDto struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateUserDto struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserDto struct {
	ID    int64  `json:"id"`
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

func UpdateUserDtoToUser(dto UpdateUserDto, user User) User {
	return User{
		Name:     utils.Coalesce(dto.Name, user.Name),
		Email:    utils.Coalesce(dto.Email, user.Email),
		Password: utils.Coalesce(dto.Password, user.Password),
	}
}

func UserToUserDto(user User) UserDto {
	return UserDto{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}
