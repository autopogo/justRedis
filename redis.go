/* 
package justRedis presents a configuration structure you can use to open and close a database connection to redis, and turn redis commands into go functions accessible by everyone.
*/
package justRedis

import (

	redis "github.com/go-redis/redis"
  "errors"
  log "github.com/autopogo/justLogging"
)



// DBInst is a configuration structure for SQL. The auth will be zero'd once its opened.
type RedisConfig struct {
	db *redis.Client
}


// The only unique error for this package so far
var (
 ErrStmtConflict = errors.New("justRedis: Tried to create two statements of the same name")
)

// Open opens the database connection, and makes the maps of precompiled statements
func (d *RedisConfig) Open() error {
	// TODO set logger
	d.db = redis.NewClient(&redis.Options{
    Addr:     "localhost:6379",
    Password: "", // no password set
    DB:       0,  // use default DB
	})
	pong, err := d.db.Ping().Result()
	if err != nil {
		log.Errorf("justRedis, Open failed: redis ping: %v, err: %v", pong, err)
		panic("Panic'ed due to redis")
	}
	log.Enterf("justRedis, Open: Redis ping: %v", pong)

	return nil
}

// Close closes the database
func (d *RedisConfig) Close() error {
	err := d.db.Close()
	if (err != nil) {
		log.Errorf("justRedis, Close: Error closing: %v", err)
		return err
	}
	return nil
}

