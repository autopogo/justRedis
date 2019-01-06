/* 
package justSessions presents a configuration structure you can use to open and close a database connection to redis. It provides functions for basic session management.
*/
package justSessions

import (

	redis "github.com/go-redis/redis"
  "errors"
  log "github.com/autopogo/justLogging"
)



// DBInst is a configuration structure for SQL. The auth will be zero'd once its opened.
type SessionsConfig struct {
	db *redis.Client
}


// The only unique error for this package so far
var (
 ErrStmtConflict = errors.New("justSQL: Tried to create two statements of the same name")
)

// Open opens the database connection, and makes the maps of precompiled statements
func (d *SessionsConfig) Open() error {
	// TODO set logger
	d.db = redis.NewClient(&redis.Options{
    Addr:     "localhost:6379",
    Password: "", // no password set
    DB:       0,  // use default DB
	})
	pong, err := d.db.Ping().Result()
	if err != nil {
		log.Errorf("Sessions, Open failed: redis ping: %v, err: %v", pong, err)
		panic("Panic'ed due to redis")
	}
	log.Enterf("Sessions, Open: Redis ping: %v", pong)

	return nil
}

// Close closes the database
func (d *SessionsConfig) Close() error {
	err := d.db.Close()
	if (err != nil) {
		log.Errorf("Sessions, Close: Error closing: %v", err)
		return err
	}
	return nil
}

