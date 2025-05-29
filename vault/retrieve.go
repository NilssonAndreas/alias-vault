package vault

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
)

func GetAlias(alias string) (*Alias, error) {
	row := db.QueryRow("SELECT alias, command, tags, created_at FROM aliases WHERE alias = ?", alias)
	var a Alias
	var tagsStr string
	if err := row.Scan(&a.Alias, &a.Command, &tagsStr, &a.CreatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("alias '%s' hittades inte", alias)
		}
		return nil, err
	}
	_ = json.Unmarshal([]byte(tagsStr), &a.Tags)
	return &a, nil
}
