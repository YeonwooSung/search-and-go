package db

import (
	"strings"
	"time"

	"gorm.io/gorm"
)

// SearchIndex represents a search index entry in the database.
type SearchIndex struct {
	ID        string `gorm:"type:uuid;default:uuid_generate_v4()"`
	Value     string
	Urls      []CrawledUrl   `gorm:"many2many:token_urls;"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

/**
 * TableName returns the name of the table associated with the SearchIndex model.
 *
 * @return string
 */
func (s *SearchIndex) TableName() string {
	return "search_index"
}

/**
 * Save saves the search index to the database.
 * It takes a map of index values and their corresponding IDs, and a list of crawled URLs.
 * It creates or updates the search index entry for each value, and associates the relevant crawled URLs with the index entry.
 *
 * @param index map[string][]string
 * @param crawledUrls []CrawledUrl
 * @return error
 */
func (s *SearchIndex) Save(index map[string][]string, crawledUrls []CrawledUrl) error {
	for value, ids := range index {
		newIndex := &SearchIndex{
			Value: value,
		}
		if err := DBConn.Where(SearchIndex{Value: value}).FirstOrCreate(newIndex).Error; err != nil {
			return err
		}

		var urlsToAppend []CrawledUrl
		for _, id := range ids {
			for _, url := range crawledUrls {
				if url.ID == id {
					urlsToAppend = append(urlsToAppend, url)
					break
				}
			}
		}

		if err := DBConn.Model(&newIndex).Association("Urls").Append(&urlsToAppend); err != nil {
			return err
		}
	}
	return nil
}

/**
 * FullTextSearch performs a full-text search on the search index.
 * It takes a search value and returns a list of crawled URLs that match the search value.
 * The search value is split into terms, and each term is used to search for matching search index entries.
 * The associated crawled URLs from the matching search index entries are returned.
 *
 * @param value string
 * @return []CrawledUrl
 * @return error
 */
func (s *SearchIndex) FullTextSearch(value string) ([]CrawledUrl, error) {
	terms := strings.Fields(value)
	var urls []CrawledUrl

	for _, term := range terms {
		var searchIndexes []SearchIndex
		if err := DBConn.Preload("Urls").Where("value LIKE ?", "%"+term+"%").Find(&searchIndexes).Error; err != nil {
			return nil, err
		}

		for _, searchIndex := range searchIndexes {
			urls = append(urls, searchIndex.Urls...)
		}
	}
	return urls, nil
}
