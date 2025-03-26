package system_mgmt_srv

import (
	"cook-book-admin-backend/models"
	"cook-book-admin-backend/respositories/system_mgmt_repo"
)

type DictMgmtService interface {
	GetDictList(req *models.GetDictListRequest) ([]models.Dict, int64, int, int, error)
	CreateDict(dict *models.Dict) error
	UpdateDict(dict *models.Dict) error
	DeleteDict(id int64) error
	GetDictData(req *models.GetDictDataListRequest) ([]models.DictData, int64, int, int, error)
	AddDictData(dictData *models.DictData) error
	UpdateDictData(dictData *models.DictData) error
	DeleteDictData(id int64) error
}

type dictMgmtService struct {
	dictRepo system_mgmt_repo.DictMgmtRepository
}

func NewDictMgmtService(dictRepo *system_mgmt_repo.DictMgmtRepository) DictMgmtService {
	return &dictMgmtService{*dictRepo}
}

// GetDictList gets the list of dictionaries
func (s *dictMgmtService) GetDictList(req *models.GetDictListRequest) ([]models.Dict, int64, int, int, error) {
	return s.dictRepo.GetDictList(req)
}

// CreateDict creates a new dictionary
func (s *dictMgmtService) CreateDict(dict *models.Dict) error {
	return s.dictRepo.CreateDict(dict)
}

// UpdateDict updates a dictionary
func (s *dictMgmtService) UpdateDict(dict *models.Dict) error {
	return s.dictRepo.UpdateDict(dict)
}

// DeleteDict deletes a dictionary
func (s *dictMgmtService) DeleteDict(id int64) error {
	return s.dictRepo.DeleteDict(id)
}

// GetDictData gets the list of dictionary data
func (s *dictMgmtService) GetDictData(req *models.GetDictDataListRequest) ([]models.DictData, int64, int, int, error) {
	return s.dictRepo.GetDictData(req)
}

// AddDictData adds a new dictionary data
func (s *dictMgmtService) AddDictData(dictData *models.DictData) error {
	return s.dictRepo.AddDictData(dictData)
}

// UpdateDictData updates a dictionary data
func (s *dictMgmtService) UpdateDictData(dictData *models.DictData) error {
	return s.dictRepo.UpdateDictData(dictData)
}

// DeleteDictData deletes a dictionary data
func (s *dictMgmtService) DeleteDictData(id int64) error {
	return s.dictRepo.DeleteDictData(id)
}
