package service

import (
	"backend/internal/web/repository"
	"backend/internal/web/structs"
)

func GetUsers() ([]structs.GetUsersResponse, error) {
	return repository.GetUsers()
}
