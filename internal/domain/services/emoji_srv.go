package services

import (
	"cook-book-backend/internal/domain/repositories"
	"cook-book-backend/internal/interfaces/dto"
)

type EmojiService interface {
	GetAllEmoji() ([]dto.EmojiResponse, int64, error)
}

type emojiSrv struct {
	emojiRepo repositories.EmojiRepository
}

func NewEmojiSrv(emojiRepo repositories.EmojiRepository) EmojiService {
	return &emojiSrv{
		emojiRepo: emojiRepo,
	}
}

func (e *emojiSrv) GetAllEmoji() ([]dto.EmojiResponse, int64, error) {
	return e.emojiRepo.GetAllEmoji()
}
