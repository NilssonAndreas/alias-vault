package vault

import (
	"encoding/json"
	"strings"
)

func SearchAliases(query string) ([]Alias, error) {
	rows, err := db.Query("SELECT alias, command, tags, created_at FROM aliases")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []Alias
	for rows.Next() {
		var a Alias
		var tagsStr string
		if err := rows.Scan(&a.Alias, &a.Command, &tagsStr, &a.CreatedAt); err != nil {
			continue
		}
		_ = json.Unmarshal([]byte(tagsStr), &a.Tags)

		// Enkel sök i alias, command, taggar
		if strings.Contains(strings.ToLower(a.Alias), query) ||
			strings.Contains(strings.ToLower(a.Command), query) ||
			tagsContain(a.Tags, query) {
			results = append(results, a)
		}
	}

	return results, nil
}

func tagsContain(tags []string, query string) bool {
	for _, tag := range tags {
		if strings.Contains(strings.ToLower(tag), query) {
			return true
		}
	}
	return false
}
