package entities

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

// JSONStringSlice armazena []string como JSON no banco
type JSONStringSlice []string

// Scan implementa sql.Scanner - Converte JSON do banco para []string
func (j *JSONStringSlice) Scan(value interface{}) error {
	if value == nil {
		*j = nil
		return nil
	}

	switch v := value.(type) {
	case []byte:
		return json.Unmarshal(v, j)
	case string:
		return json.Unmarshal([]byte(v), j)
	default:
		return fmt.Errorf("tipo não suportado: %T", value)
	}
}

// Value implementa driver.Valuer - Converte []string para JSON
func (j JSONStringSlice) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}
	return json.Marshal(j)
}

// JSONIntSlice armazena []int como JSON no banco
type JSONIntSlice []int

func (j *JSONIntSlice) Scan(value interface{}) error {
	if value == nil {
		*j = nil
		return nil
	}

	switch v := value.(type) {
	case []byte:
		return json.Unmarshal(v, j)
	case string:
		return json.Unmarshal([]byte(v), j)
	default:
		return fmt.Errorf("tipo não suportado: %T", value)
	}
}

func (j JSONIntSlice) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}
	return json.Marshal(j)
}
