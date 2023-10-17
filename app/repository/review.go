package repository

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/sekke276/greendeco.git/app/models"
	"github.com/sekke276/greendeco.git/platform/database"
)

type ReviewRepository interface {
	Create(m *models.CreateReview) error
	FindById(id uuid.UUID) (*models.Review, error)
	FindReviewsByProductId(id *uuid.UUID, q *models.ReviewQuery) ([]*models.Review, error)
}

const (
	ReviewTable = "reviews"
)

type ReviewRepo struct {
	db *database.DB
}

var _ ReviewRepository = (*ReviewRepo)(nil)

func NewReviewRepo(db *database.DB) ReviewRepository {
	return &ReviewRepo{
		db: db,
	}
}

func (repo *ReviewRepo) Create(m *models.CreateReview) error {
	query := fmt.Sprintf(`INSERT INTO "%s" (user_id, product_id, content, star) VALUES ($1,$2,$3,$4)`, ReviewTable)
	if _, err := repo.db.Exec(query, m.UserId, m.ProductId, m.Content, m.Star); err != nil {
		return err
	}

	return nil
}

func (repo *ReviewRepo) FindById(id uuid.UUID) (*models.Review, error) {
	query := fmt.Sprintf(`SELECT * FROM "%s" WHERE id = $1`, ReviewTable)
	review := &models.Review{}
	if err := repo.db.Select(review, query, id); err != nil {
		return nil, err
	}
	return review, nil
}

func (repo *ReviewRepo) FindReviewsByProductId(id *uuid.UUID, q *models.ReviewQuery) ([]*models.Review, error) {
	result := []*models.Review{}
	limit := q.Limit
	limit += 1
	pageOffset := q.BaseQuery.Limit * (q.BaseQuery.OffSet - 1)
	firstQuery := fmt.Sprintf(`SELECT * FROM "%s" `, ReviewTable)
	query := repo.newReviewQueryBuilder(firstQuery).
		SetProduct(id).
		SetStar(q.Star).
		SetUser(q.UserId).
		SortBy(q.Sort, q.SortBy).
		Build()

	query = fmt.Sprintf(query+"LIMIT %d OFFSET %d", limit, pageOffset)
	if err := repo.db.Select(&result, query); err != nil {
		return nil, err
	}

	return result, nil
}
