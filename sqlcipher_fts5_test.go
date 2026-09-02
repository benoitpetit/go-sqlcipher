//go:build fts5 || sqlite_fts5
// +build fts5 sqlite_fts5

package sqlite3_test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSQLCipherFTS5(t *testing.T) {
	_, err := db.Exec("CREATE VIRTUAL TABLE IF NOT EXISTS encrypted_search USING fts5(content)")
	require.NoError(t, err)
}
