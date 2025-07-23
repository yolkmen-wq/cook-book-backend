package repositories

import (
	"cook-book-backend/internal/interfaces/dto"

	"gorm.io/gorm"
)

type EmojiRepository struct {
	db *gorm.DB
}

func NewEmojiRepository(db *gorm.DB) EmojiRepository {
	return EmojiRepository{db: db}
}

// GetAllEmoji 获取表情包
func (r EmojiRepository) GetAllEmoji() ([]dto.EmojiResponse, int64, error) {
	var list []dto.EmojiResponse
	var total int64

	err := r.db.Table("emojis").
		Count(&total).
		Order("created_at asc").
		Find(&list).Error

	if err != nil {
		return nil, 0, err
	}
	return list, total, nil
}
