package dbid

import (
	"bytes"
	"errors"
	"fmt"
)

// DBX is database abstraction interface, compatible with jmoiron/sqlx
// Any other ORM can be supported but adapter is needed.
type DBX interface {
	// Get method reads single entry from database and writes
	// result into target. Must return error if no entry
	// found.
	Get(target interface{}, sql string, args ...interface{}) error

	// Select method reads multiple entries from database
	// and writes them into target slice.
	Select(target interface{}, sql string, args ...interface{}) error
}

var (
	sqlSelect = []byte("SELECT * FROM `")
	sqlWhere  = []byte("` WHERE `id` IN (")
	sqlComma  = []byte(",")
	sqlMark   = []byte("?")
	sqlClose  = []byte(")")
)

// XFind fetches one or more entities by ID from data source,
// backed in database abstraction
func XFind(dbx DBX, target interface{}, ids ...int) (err error) {
	if len(ids) == 0 {
		return errors.New("empty ids provided")
	}

	table, single, err := extract(target)
	if err != nil {
		return err
	}

	// Checking arguments count consistency
	if single && len(ids) != 1 {
		return fmt.Errorf("expected one ID for single entity find, but got %d", len(ids))
	}

	// Building query
	sb := bytes.NewBuffer(nil)
	sb.Write(sqlSelect)
	sb.Write([]byte(table))
	sb.Write(sqlWhere)
	for i := range ids {
		if i > 0 {
			sb.Write(sqlComma)
		}
		sb.Write(sqlMark)
	}
	sb.Write(sqlClose)

	// Invoking
	if single {
		err = dbx.Get(target, sb.String(), ids[0])
	} else {
		// Re-packing IDs into interface slice
		args := make([]interface{}, len(ids))
		for k, v := range ids {
			args[k] = v
		}

		err = dbx.Select(target, sb.String(), args...)
	}

	return
}
