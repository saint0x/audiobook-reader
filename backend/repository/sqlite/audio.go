package sqlite

import (
	"fmt"
	"time"

	"backend/domain/models"
)

// SaveAudioSegment saves a new audio segment to the database
func (db *DB) SaveAudioSegment(segment *models.AudioSegment) error {
	query := `
		INSERT INTO audio_segments (
			id, book_id, content, audio_url,
			status, created_at, updated_at
		) VALUES (?, ?, ?, ?, ?, ?, ?)
	`

	_, err := db.Exec(query,
		segment.ID,
		segment.BookID,
		segment.Content,
		segment.AudioURL,
		segment.Status,
		segment.CreatedAt,
		segment.UpdatedAt,
	)

	if err != nil {
		return fmt.Errorf("error saving audio segment: %v", err)
	}

	return nil
}

// UpdateAudioSegment updates an existing audio segment in the database
func (db *DB) UpdateAudioSegment(segment *models.AudioSegment) error {
	query := `
		UPDATE audio_segments 
		SET content = ?, audio_url = ?, status = ?, updated_at = ?
		WHERE id = ?
	`

	segment.UpdatedAt = time.Now()
	_, err := db.Exec(query,
		segment.Content,
		segment.AudioURL,
		segment.Status,
		segment.UpdatedAt,
		segment.ID,
	)

	if err != nil {
		return fmt.Errorf("error updating audio segment: %v", err)
	}

	return nil
}

// GetAudioSegments retrieves all audio segments for a book
func (db *DB) GetAudioSegments(bookID string) ([]models.AudioSegment, error) {
	query := `
		SELECT id, book_id, content, audio_url, status, created_at, updated_at
		FROM audio_segments
		WHERE book_id = ?
		ORDER BY created_at ASC
	`

	rows, err := db.Query(query, bookID)
	if err != nil {
		return nil, fmt.Errorf("error querying audio segments: %v", err)
	}
	defer rows.Close()

	var segments []models.AudioSegment
	for rows.Next() {
		var segment models.AudioSegment
		err := rows.Scan(
			&segment.ID,
			&segment.BookID,
			&segment.Content,
			&segment.AudioURL,
			&segment.Status,
			&segment.CreatedAt,
			&segment.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning audio segment: %v", err)
		}
		segments = append(segments, segment)
	}

	return segments, nil
}

// GetAudioSegmentByID retrieves an audio segment by its ID
func (db *DB) GetAudioSegmentByID(id string) (*models.AudioSegment, error) {
	query := `
		SELECT id, book_id, content, audio_url, status, created_at, updated_at
		FROM audio_segments
		WHERE id = ?
	`

	segment := &models.AudioSegment{}
	err := db.QueryRow(query, id).Scan(
		&segment.ID,
		&segment.BookID,
		&segment.Content,
		&segment.AudioURL,
		&segment.Status,
		&segment.CreatedAt,
		&segment.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("error getting audio segment: %v", err)
	}

	return segment, nil
}

// DeleteAudioSegment deletes an audio segment from the database
func (db *DB) DeleteAudioSegment(id string) error {
	query := "DELETE FROM audio_segments WHERE id = ?"
	_, err := db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error deleting audio segment: %v", err)
	}
	return nil
}
