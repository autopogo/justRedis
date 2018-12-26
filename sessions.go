/* 
package justSessions presents a configuration structure you can use to open and close a database connection to redis. It provides functions for basic session management.
*/
package justSessions

import (


  "fmt"
  "errors"
  log "github.com/autopogo/justLogging"
)



// DBInst is a configuration structure for SQL. The auth will be zero'd once its opened.
type SessionsConfig struct {
}


// The only unique error for this package so far
var (
 ErrStmtConflict = errors.New("justSQL: Tried to create two statements of the same name")
)

// Open opens the database connection, and makes the maps of precompiled statements
func (d *SessionsConfig) Open() error {
	return nil
}

// Close closes the database
func (d *SessionsConfig) Close() error {
	return nil
}

