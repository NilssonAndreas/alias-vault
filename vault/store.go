package vault

import (
	"database/sql"
	"encoding/json"
	"os"
	"path/filepath"
	"time"

	_ "modernc.org/sqlite"
)

type Alias struct {
	Alias     string
	Command   string
	Tags      []string
	CreatedAt time.Time
}

var db *sql.DB

func getDbPath() string {
	homeDir, _ := os.UserHomeDir()
	return filepath.Join(homeDir, ".aliasvault.db")
}
func init() {
	var err error
	db, err = sql.Open("sqlite", getDbPath())
	if err != nil {
		panic(err)
	}
	createTable()
}

func createTable() {
	query := `CREATE TABLE IF NOT EXISTS aliases (
		alias TEXT PRIMARY KEY,
		command TEXT,
		tags TEXT,
		created_at TIMESTAMP
	)`
	_, err := db.Exec(query)
	if err != nil {
		panic(err)
	}
}

func SaveAlias(alias, command string, tags []string) error {
	tagsJson, err := json.Marshal(tags)
	if err != nil {
		return err
	}

	_, err = db.Exec(`
		INSERT INTO aliases (alias, command, tags, created_at)
		VALUES (?, ?, ?, ?)
		ON CONFLICT(alias) DO UPDATE SET
			command = excluded.command,
			tags = excluded.tags
	`, alias, command, string(tagsJson), time.Now())

	return err
}
