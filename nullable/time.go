package nullable

import (
	"database/sql"
	"time"
)

// Time is an alias for mysql.NullTime data type
type Time struct {
	sql.NullTime
}

// Time Nullable conformance
func (n Time) IsNull() bool {
	return !n.Valid
}

// Time convenience initializer
func NewTime(t time.Time) Time {
	return Time{
		sql.NullTime{
			Time: t, Valid: true,
		},
	}
}

// Time convenience initializer for invalid (nil)
func NullTime() Time {
	return Time{
		sql.NullTime{
			Valid: false,
		},
	}
}

// MarshalJSON for Time
func (n Time) MarshalJSON() ([]byte, error) {
	return marshalJSON(n.Time.Format(time.RFC3339), n.Valid)
}

// UnmarshalJSON for Time
func (n *Time) UnmarshalJSON(b []byte) error {
	return unmarshalJSON(b, &n.Time, &n.Valid)
}
