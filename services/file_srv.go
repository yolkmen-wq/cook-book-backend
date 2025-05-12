package services

import (
	"cook-book-admin-backend/respositories"
	"mime/multipart"
)

type FileService interface {
	UploadFile(objectName string, file multipart.File) (string, error)
}

type fileService struct {
	fileRepo respositories.FileRepository
}

func NewFileService(fileRepo *respositories.FileRepository) FileService {
	return &fileService{fileRepo: *fileRepo}
}

func (fs *fileService) UploadFile(objectName string, file multipart.File) (string, error) {
	return fs.fileRepo.UploadFile(objectName, file)
}
