DBID
====

`dbid` is simple yet efficient tool, made to provide easy and efficient way 
to obtain entries from MySQL by ID - just use `dbid.XFind`.

## TL&DR;

##### Step one 

Create entity class, that implements `dbid.SchemaLocator`:

```go
type User struct {
	ID   int 
	Name string
}

// Schema is dbid.SchemaLocator interface implementation
// It returns database table name
func (User) Schema() string {
	return "users"
}
```

##### Step two 

Use it:

```go
package main

import (
    "fmt"
    _ "github.com/go-sql-driver/mysql" // Dont forget to load driver
    "github.com/jmoiron/sqlx"          // sqlx is compatible with dbid.DBX
    "github.com/mono83/dbid"
)

func main() {
    db, err := sqlx.Open("mysql", "...") // Provide DSN here
    if err != nil {
        panic(err)
    }

    // Reading user with ID 3
    var user User
    err = dbid.XFind(db, &user, 3)
    fmt.Println(err, user)

    // Reading multiple users with IDs 1,2,3,4,5
    var users []User
    err = dbid.XFind(db, &users, 1, 2, 3, 4, 5)
    fmt.Println(err, users)
}
``` 
