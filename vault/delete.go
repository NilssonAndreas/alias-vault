package vault

func DeleteAlias(alias string) error {
	_, err := db.Exec("DELETE FROM aliases WHERE alias = ?", alias)
	return err
}
