package repository

import (
	"context"
	"fmt"
	"go/apiExample/entity"
	"go/apiExample/go_database"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestInsertComment(t *testing.T) {
	commentRepository := NewCommentRepository(go_database.GetConnection())

	ctx := context.Background()
	comment := entity.Comment{
		Email:   "ocu@ocudevs.com",
		Comment: "Test Test Repository Pattern",
	}

	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
