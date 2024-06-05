package repository

import (
	"backend/internal/database"
	"backend/internal/database/model"
	"backend/internal/web/structs"
)

func GetUsers() ([]structs.GetUsersResponse, error) {
	var data []structs.GetUsersResponse
	if err := database.GetDB().Model(&model.User{}).Scan(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}
