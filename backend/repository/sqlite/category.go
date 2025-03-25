package sqlite

import (
	"fmt"

	"backend/domain/models"
)

// GetCategories retrieves all categories from the database
func (db *DB) GetCategories() ([]models.Category, error) {
	query := `
		SELECT id, name, description, created_at
		FROM categories
		ORDER BY name ASC
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error querying categories: %v", err)
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var category models.Category
		err := rows.Scan(
			&category.ID,
			&category.Name,
			&category.Description,
			&category.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning category: %v", err)
		}
		categories = append(categories, category)
	}

	return categories, nil
}

// GetTags retrieves all tags from the database
func (db *DB) GetTags() ([]models.Tag, error) {
	query := `
		SELECT id, name, created_at
		FROM tags
		ORDER BY name ASC
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error querying tags: %v", err)
	}
	defer rows.Close()

	var tags []models.Tag
	for rows.Next() {
		var tag models.Tag
		err := rows.Scan(
			&tag.ID,
			&tag.Name,
			&tag.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning tag: %v", err)
		}
		tags = append(tags, tag)
	}

	return tags, nil
}

// CreateCategory creates a new category in the database
func (db *DB) CreateCategory(category *models.Category) error {
	query := `
		INSERT INTO categories (id, name, description, created_at)
		VALUES (?, ?, ?, ?)
	`

	_, err := db.Exec(query,
		category.ID,
		category.Name,
		category.Description,
		category.CreatedAt,
	)

	if err != nil {
		return fmt.Errorf("error creating category: %v", err)
	}

	return nil
}

// CreateTag creates a new tag in the database
func (db *DB) CreateTag(tag *models.Tag) error {
	query := `
		INSERT INTO tags (id, name, created_at)
		VALUES (?, ?, ?)
	`

	_, err := db.Exec(query,
		tag.ID,
		tag.Name,
		tag.CreatedAt,
	)

	if err != nil {
		return fmt.Errorf("error creating tag: %v", err)
	}

	return nil
}

// DeleteCategory deletes a category from the database
func (db *DB) DeleteCategory(id string) error {
	query := "DELETE FROM categories WHERE id = ?"
	_, err := db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error deleting category: %v", err)
	}
	return nil
}

// DeleteTag deletes a tag from the database
func (db *DB) DeleteTag(id string) error {
	query := "DELETE FROM tags WHERE id = ?"
	_, err := db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error deleting tag: %v", err)
	}
	return nil
}
