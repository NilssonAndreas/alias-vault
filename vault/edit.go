package vault

import (
	"encoding/json"
	"time"
)

func GetAliasByName(name string) (*Alias, error) {
	row := db.QueryRow("SELECT command, tags, created_at FROM aliases WHERE alias = ?", name)

	var command string
	var tagsJson string
	var createdAt time.Time

	err := row.Scan(&command, &tagsJson, &createdAt)
	if err != nil {
		return nil, err
	}

	var tags []string
	if err := json.Unmarshal([]byte(tagsJson), &tags); err != nil {
		return nil, err
	}

	return &Alias{
		Alias:     name,
		Command:   command,
		Tags:      tags,
		CreatedAt: createdAt,
	}, nil
}
