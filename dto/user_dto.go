package dto

import "anxinyou/model"

type UserDto struct {
	Name      string `json:"name"`
	Telephone string `json:"telephone"`
	Portrait  string `json:"Portrait"`
}

func ToUserDto(user model.User) UserDto {
	return UserDto{
		Name:      user.Name,
		Telephone: user.Telephone,
		Portrait:  user.Portrait,
	}
}
