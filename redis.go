/*
package justRedis presents a configuration structure you can use to open and close a database connection to redis, and turn redis commands into go functions accessible by everyone.
*/
package justRedis

import (
	"errors"
	log "github.com/autopogo/justLogging"
	redis "github.com/go-redis/redis"
)

// DBInst is a configuration structure for SQL. The auth will be zero'd once its opened.
type RedisConfig struct {
	db      *redis.Client
	scripts map[string]*redis.Script
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
	if err != nil {
		log.Errorf("justRedis, Close: Error closing: %v", err)
		return err
	}
	return nil
}

func (d *RedisConfig) AddScript(script string) error {

	return nil
}

func (d *RedisConfig) AddScriptFromFile(script string, name string) error {
	return nil
}

// it's like, to build it, it needs to find the functions it needs, and build a script with it, outputting it
