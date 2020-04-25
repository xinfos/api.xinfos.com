package repository

import "api.xinfos.com/internal/model"

type repository interface {
	FindByID(id uint64) (*model.User, error)
	FindAll(map[string]interface{})
	Create()
	Update()
	Delete()
	TxCreate()
	TxUpdate()
	TxDelete()
}
