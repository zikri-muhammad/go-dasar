package repository

import (
	"context"
	"database/sql"
	"go/apiExample/entity"
)

type commentRepositoryImpl struct {
	DB *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	return &commentRepositoryImpl{
		DB: db,
	}
}
func (repo *commentRepositoryImpl) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	sqlExec := "INSERT INTO comments(email, comment) VALUES(?,?)"
	result, err := repo.DB.ExecContext(ctx, sqlExec, comment.Email, comment.Comment)
	if err != nil {
		return comment, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return comment, err
	}

	comment.Id = int32(id)
	return comment, nil
}

func (repo *commentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comment, error) {
	var comments []entity.Comment
	rows, err := repo.DB.QueryContext(ctx, "SELECT id, email, comment FROM comments")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var comment entity.Comment
		if err := rows.Scan(&comment.Id, &comment.Email, &comment.Comment); err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	return comments, nil
}
