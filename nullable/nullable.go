package nullable

import (
	"bytes"
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

type Nullable[T any] struct {
	Val   T
	Valid bool
}

func (n Nullable[T]) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return []byte(`null`), nil
	}

	return json.Marshal(n.Value)
}

func (n *Nullable[T]) UnmarshalJSON(b []byte) error {
	if bytes.Equal([]byte(`null`), b) {
		n.Valid = false
		return nil
	}

	err := json.Unmarshal(b, &n.Val)
	n.Valid = err == nil
	return err
}

func (n *Nullable[T]) Scan(value any) error {
	n.Valid = false

	if value == nil {
		return nil
	}

	v, ok := value.(T)
	if !ok {
		return fmt.Errorf("cannot scan value")
	}

	n.Val = v
	n.Valid = true
	return nil
}

func (n Nullable[T]) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Value, nil
}

type NullTime struct {
	sql.NullTime
}

// MarshalJSON for NullTime
func (n NullTime) MarshalJSON() ([]byte, error) {
	if !n.Valid {
		return []byte(`null`), nil
	}
	val := fmt.Sprintf("\"%s\"", n.Time.Format(time.RFC3339))
	return []byte(val), nil
}

// UnmarshalJSON for NullTime
func (n *NullTime) UnmarshalJSON(b []byte) error {
	if bytes.Equal([]byte(`null`), b) {
		n.Valid = false
		return nil
	}

	err := json.Unmarshal(b, &n.Time)
	n.Valid = err == nil
	return err
}
