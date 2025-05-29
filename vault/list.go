package vault

import (
	"encoding/json"
)

func GetAllAliases() ([]Alias, error) {
	rows, err := db.Query("SELECT alias, command, tags, created_at FROM aliases ORDER BY created_at DESC")
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
		results = append(results, a)
	}

	return results, nil
}
