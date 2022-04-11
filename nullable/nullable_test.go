package nullable

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestNullable(t *testing.T) {
	jsonStr := `{"time": "2022-01-04T12:30:05Z"}`
	//jsonStr := `{"time": null}`
	type TimeStruct struct {
		Time Nullable[time.Time] `json:"time"`
	}

	var s TimeStruct
	require.NoError(t, json.Unmarshal([]byte(jsonStr), &s))
	//require.False(t, s.Time.Valid)
	require.Equal(t, "", fmt.Sprintf("%v", s.Time.Val))
}
