package db

import (
	"context"
	"fmt"
	"github.com/gernhan/xml/caches"
	"github.com/patrickmn/go-cache"
	"log"
)

type TextModule struct {
	TextModuleName string
	TextTemplate   string
}

func FindAllByClientAndLanguageAndProductLineBrand(client, language, brand int64) ([]TextModule, error) {
	c := caches.GetCaches().StmCache
	// Check if the data is already cached.
	cacheKey := fmt.Sprintf("%s_%d_%d_%d", "findAllByClientAndLanguageAndProductLineBrand", client, language, brand)
	if data, found := c.Get(cacheKey); found {
		if textModules, ok := data.([]TextModule); ok {
			return textModules, nil
		}
	}
	query := `
		SELECT stm.text_module_name, stmt.text_template
		FROM s_text_modules_templates stmt
		JOIN s_text_modules stm ON stmt.text_module = stm.id
		WHERE stm.text_module_type = 9 AND stm.client = $1 AND stmt.language = $2 AND stm.product_line_brand = $3`

	// Execute the query and store the result in the texts variable.
	var texts []TextModule
	db := GetPool()
	rows, err := db.Query(context.Background(), query, client, language, brand)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var text TextModule
		err := rows.Scan(&text.TextModuleName, &text.TextTemplate)
		if err != nil {
			return nil, err
		}
		texts = append(texts, text)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	// Cache the data for future use.
	c.Set(cacheKey, texts, cache.DefaultExpiration)
	log.Println("Cache miss!")

	return texts, nil
}
