package services

import (
	"database/sql"
	"errors"
	"time"

	"sharing-vision-api/models"
)

type PostService struct {
	db *sql.DB
}

func NewPostService(db *sql.DB) *PostService {
	return &PostService{db: db}
}

func (s *PostService) CreatePost(title, content, category, status string) error {
	query := `
		INSERT INTO posts (title, content, category, created_date, updated_date, status)
		VALUES (?, ?, ?, ?, ?, ?)
	`
	now := time.Now()
	_, err := s.db.Exec(query, title, content, category, now, now, status)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostService) GetAllPosts(limit, offset int) ([]models.Post, error) {
	query := `
		SELECT id, title, content, category, created_date, updated_date, status
		FROM posts
		LIMIT ? OFFSET ?
	`
	rows, err := s.db.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	posts := []models.Post{}
	for rows.Next() {
		var p models.Post
		if err := rows.Scan(
			&p.ID,
			&p.Title,
			&p.Content,
			&p.Category,
			&p.CreatedDate,
			&p.UpdatedDate,
			&p.Status,
		); err != nil {
			return nil, err
		}
		posts = append(posts, p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

func (s *PostService) GetPostByID(id int) (*models.Post, error) {
	query := `
		SELECT id, title, content, category, created_date, updated_date, status
		FROM posts
		WHERE id = ?
	`
	row := s.db.QueryRow(query, id)

	var p models.Post
	err := row.Scan(
		&p.ID,
		&p.Title,
		&p.Content,
		&p.Category,
		&p.CreatedDate,
		&p.UpdatedDate,
		&p.Status,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, errors.New("post not found")
	}
	if err != nil {
		return nil, err
	}

	return &p, nil
}

func (s *PostService) UpdatePost(id int, title, content, category, status string) error {
	query := `
		UPDATE posts
		SET title = ?, content = ?, category = ?, status = ?, updated_date = ?
		WHERE id = ?
	`
	result, err := s.db.Exec(query, title, content, category, status, time.Now(), id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("post not found")
	}

	return nil
}

func (s *PostService) DeletePost(id int) error {
	query := `
		UPDATE posts
		SET status = 'thrash', updated_date = ?
		WHERE id = ?
	`

	result, err := s.db.Exec(query, time.Now(), id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("post not found")
	}

	return nil
}
