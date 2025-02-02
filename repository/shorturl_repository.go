package repository

import (
	"database/sql"
	"fmt"
	"url-shortening-service/database"
	"url-shortening-service/utils"
)

type ShortUrl struct {
	Id          int    `json:"id"`
	Url         string `json:"url"`
	ShortCode   string `json:"shortCode"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
	AccessCount int    `json:"-"`
}

type ShorturlRepository struct {
	db *sql.DB
}

func NewShorturlRepository() *ShorturlRepository {
	return &ShorturlRepository{
		db: database.GetDB(),
	}
}

func (r *ShorturlRepository) GetShortUrlByShortCode(shortCode string) (ShortUrl, error) {
	var su ShortUrl

	row := r.db.QueryRow("CALL GetShorturl(?)", shortCode)
	if err := row.Scan(&su.Id, &su.Url, &su.ShortCode, &su.CreatedAt, &su.UpdatedAt, &su.AccessCount); err != nil {
		if err == sql.ErrNoRows {
			return su, fmt.Errorf("GetShortUrlByShortCode %v: no such short code ", shortCode)
		}
		return su, fmt.Errorf("GetShortUrlByShortCode %v: %v", shortCode, err)
	}

	return su, nil
}

func (r *ShorturlRepository) InsertShortUrl(url string) (ShortUrl, error) {
	var su ShortUrl
	var shortCode string = utils.GenerateRandomString(6)
	row := r.db.QueryRow("CALL InsertShorturl(?, ?)", url, shortCode)
	if err := row.Scan(&su.Id, &su.Url, &su.ShortCode, &su.CreatedAt, &su.UpdatedAt, &su.AccessCount); err != nil {
		if err == sql.ErrNoRows {
			return su, fmt.Errorf("InsertShortUrl %v: no such short code ", shortCode)
		}
		return su, fmt.Errorf("InsertShortUrl %v: %v", shortCode, err)
	}
	return su, nil
}

func (r *ShorturlRepository) UpdateShortUrl(url string, shortCode string) (ShortUrl, error) {
	var su ShortUrl
	row := r.db.QueryRow("CALL UpdateShorturl(?, ?)", url, shortCode)
	if err := row.Scan(&su.Id, &su.Url, &su.ShortCode, &su.CreatedAt, &su.UpdatedAt, &su.AccessCount); err != nil {
		if err == sql.ErrNoRows {
			return su, fmt.Errorf("UpdateShortUrl %v: no such short code ", shortCode)
		}
		return su, fmt.Errorf("UpdateShortUrl %v: %v", shortCode, err)
	}
	return su, nil
}

func (r *ShorturlRepository) DeleteShortUrl(shortCode string) (int64, error) {
	result, err := r.db.Exec("DELETE FROM ShortUrls WHERE shortCode = ?", shortCode)
	if err != nil {
		return 0, fmt.Errorf("DeleteShortUrl: %v", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("DeleteShortUrl: %v", err)
	}
	return rowsAffected, nil
}
