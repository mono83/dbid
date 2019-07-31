package dbid

// SchemaLocator returns schema name (database table name)
// for entity, that implements this interface.
type SchemaLocator interface {
	Schema() string
}
