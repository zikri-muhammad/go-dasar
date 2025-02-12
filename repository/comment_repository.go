package repository

import (
	"context"
	"go/apiExample/entity"
)

type CommentRepository interface {
	Insert(ctc context.Context, comment entity.Comment) (entity.Comment, error)
	// FindById(ctc context.Context, id int32) (entity.Comment, error)
	FindAll(ctc context.Context) ([]entity.Comment, error)
}
