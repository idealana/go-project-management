package repositories

import (
    "github.com/idealana/go-project-management/config"
    "github.com/idealana/go-project-management/models"
)

type ListRepository interface {
    Create(list *models.List) error
}

type listRepository struct {
}

func NewListRepository() ListRepository {
    return &listRepository{}
}

func (r *listRepository) Create(list *models.List) error {
    return config.DB.Create(list).Error
}

func (r *listRepository) Update(list *models.List) error {
    return config.DB.
        Model(&models.List{}).
        Where("public_id = ?", list.PublicID).
        Updates(map[string]interface{}{
            "title": list.Title,
        }).
        Error
}

func (r *listRepository) Delete(id uint) error {
    return config.DB.Delete(&models.List{}, id).Error
}

func (r *listRepository) UpdatePosition(boardPublicID string, position []string) error {
    return config.DB.
        Model(&models.ListPosition{}).
        Where("board_internal_id = (SELECT internal_id FROM boards WHERE public_id = ?)", boardPublicID).
        Update("list_order", position).
        Error
}
