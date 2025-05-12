package article_repo

import (
	"cook-book-admin-backend/models"
	"fmt"
	"gorm.io/gorm"
)

type CarouselRepository struct {
	db *gorm.DB
}

func NewCarouselRepository(db *gorm.DB) *CarouselRepository {
	return &CarouselRepository{db: db}
}

// GetCarouses 获取轮播图列表
func (cr *CarouselRepository) GetCarouses(req *models.GetCarouselsRequest) ([]models.Carousel, int64, int, int, error) {
	var carousels []models.Carousel
	var total int64

	// 构建查询条件
	db := cr.db.Table("carousel")

	if req.CarouselName != "" {
		db = db.Where("carousel_name LIKE ?", "%"+req.CarouselName+"%")
	}

	// 计算总数
	if err := db.Count(&total).Error; err != nil {
		fmt.Println("获取轮播图总数失败", err)
		return nil, 0, 0, 0, err
	}

	// 获取分页轮播图列表
	if req.PageNum != 0 && req.PageSize != 0 {
		if err := db.Limit(req.PageSize).Offset((req.PageNum - 1) * req.PageSize).Find(&carousels).Error; err != nil {
			fmt.Println("获取轮播图列表失败", err)
			return nil, 0, 0, 0, err
		}
	} else {
		if err := db.Find(&carousels).Error; err != nil {
			fmt.Println("获取轮播图列表失败", err)
			return nil, 0, 0, 0, err
		}
	}
	for key, value := range carousels {
		// 获取轮播图位置名称
		var positionName string
		if err := cr.db.Table("dict_data").Select("dict_label").Where(" dict_type = 'sys_carousel_position' AND dict_value = ?", value.Position).Find(&positionName).Error; err != nil {
			fmt.Println("获取轮播图位置名称失败", err)
			return nil, 0, 0, 0, err
		}
		carousels[key].PositionName = positionName
	}

	return carousels, total, req.PageSize, req.PageNum, nil
}

// CreateCarousel 创建轮播图
func (cr *CarouselRepository) CreateCarousel(carousel *models.Carousel) error {
	if err := cr.db.Table("carousel").Omit("position_name").Create(carousel).Error; err != nil {
		fmt.Println("创建轮播图失败", err)
		return err
	}
	return nil
}

// UpdateCarousel 更新轮播图
func (cr *CarouselRepository) UpdateCarousel(carousel *models.Carousel) error {
	if err := cr.db.Updates(carousel).Error; err != nil {
		fmt.Println("更新轮播图失败", err)
		return err
	}
	return nil
}

// DeleteCarousel 删除轮播图
func (cr *CarouselRepository) DeleteCarousel(id int64) error {
	if err := cr.db.Table("carousel").Delete(&models.Carousel{}, id).Error; err != nil {
		fmt.Println("删除轮播图失败", err)
		return err
	}
	return nil
}

// GetCarouselItems 获取轮播图项列表
func (cr *CarouselRepository) GetCarouselItems(req *models.GetCarouselItemsRequest) ([]models.CarouselItem, int64, int, int, error) {
	var carouselItems []models.CarouselItem
	var total int64

	// 构建查询条件
	db := cr.db.Table("carousel_items")

	if req.Name != "" {
		db = db.Where("name LIKE ?", "%"+req.Name+"%")
	}

	// 计算总数
	if err := db.Count(&total).Error; err != nil {
		fmt.Println("获取轮播图项总数失败", err)
		return nil, 0, 0, 0, err
	}

	// 获取分页轮播图列表
	if req.PageNum != 0 && req.PageSize != 0 {
		if err := db.Limit(req.PageSize).Offset((req.PageNum - 1) * req.PageSize).Find(&carouselItems).Error; err != nil {
			fmt.Println("获取轮播图项列表失败", err)
			return nil, 0, 0, 0, err
		}
	} else {
		if err := db.Find(&carouselItems).Error; err != nil {
			fmt.Println("获取轮播图项列表失败", err)
			return nil, 0, 0, 0, err
		}
	}

	return carouselItems, total, req.PageSize, req.PageNum, nil
}

// CreateCarouselItem 创建轮播图项
func (cr *CarouselRepository) CreateCarouselItem(carouselItem *models.CarouselItem) error {
	if err := cr.db.Table("carousel_items").Create(carouselItem).Error; err != nil {
		fmt.Println("创建轮播图项失败", err)
		return err
	}
	return nil
}

// UpdateCarouselItem 更新轮播图项
func (cr *CarouselRepository) UpdateCarouselItem(carouselItem *models.CarouselItem) error {
	if err := cr.db.Updates(carouselItem).Error; err != nil {
		fmt.Println("更新轮播图项失败", err)
		return err
	}
	return nil
}

// DeleteCarouselItem 删除轮播图项
func (cr *CarouselRepository) DeleteCarouselItem(id int64) error {
	if err := cr.db.Table("carousel_items").Delete(&models.CarouselItem{}, id).Error; err != nil {
		fmt.Println("删除轮播图项失败", err)
		return err
	}
	return nil
}
