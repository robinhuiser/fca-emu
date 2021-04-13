package util

import (
	"fmt"

	"github.com/robinhuiser/fca-emu/ent"
)

func GetDatabaseClient(v string) (*ent.Client, error) {

	switch vendor := v; vendor {
	case "postgres":
		c := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
			GetEnvString("DB_HOST", "localhost"),
			GetEnvInt("DB_PORT", 5432),
			GetEnvString("DB_USER_NAME", ""),
			GetEnvString("DB_DATABASE_NAME", ""),
			GetEnvString("DB_USER_PASSWORD", ""))
		return ent.Open("postgres", c)

	case "mysql":
		c := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
			GetEnvString("DB_USER_NAME", ""),
			GetEnvString("DB_USER_PASSWORD", ""),
			GetEnvString("DB_HOST", "localhost"),
			GetEnvInt("DB_PORT", 3306),
			GetEnvString("DB_DATABASE_NAME", ""))
		return ent.Open("mysql", c)

	case "sqlite3":
		c := "file:ent?mode=memory&cache=shared&_fk=1"
		return ent.Open("sqlite3", c)
	}

	// If none of the supported databases above were configured, return error
	return nil, fmt.Errorf("non-supported database vendor %s", v)
}
