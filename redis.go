/*
package justRedis presents a configuration structure you can use to open and close a database connection to redis, and turn redis commands into go functions accessible by everyone. It's helper, doesn't fully wrap
*/
package justRedis

import (
	"errors"
	log "github.com/autopogo/justLogging"
	redis "github.com/go-redis/redis"
	"io/ioutil"
)

// DBInst is a configuration structure for SQL. The auth will be zero'd once its opened.
type RedisConfig struct {
	Db      *redis.Client
	Scripts map[string]*redis.Script
}

// The only unique error for this package so far
var (
	ErrStmtConflict = errors.New("justRedis: Tried to create two scripts of the same name")
)

// Open opens the database connection, and makes the maps of precompiled statements
func (d *RedisConfig) Open(address string, password string, db int) error {
	// TODO set logger
	d.Db = redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password, // no password set
		DB:       db,       // use default DB
	})
	pong, err := d.Db.Ping().Result()
	if err != nil {
		log.Errorf("justRedis, Open failed: redis ping: %v, err: %v", pong, err)
		panic("Panic'ed due to redis")
	}
	log.Enterf("justRedis, Open: Redis ping: %v", pong)

	return nil
}

// Close closes the database
func (d *RedisConfig) Close() error {
	err := d.Db.Close()
	if err != nil {
		log.Errorf("justRedis, Close: Error closing: %v", err)
		return err
	}
	return nil
}

func (d *RedisConfig) AddScript(scriptText string, name string) (*redis.Script, error) {
	if d.Scripts == nil {
		d.Scripts = make(map[string]*redis.Script)
	}
	if script, ok := d.Scripts[name]; ok {
		log.Enterf("justRedis, .AddScript(): Tried to add script that already exists.")
		return script, ErrStmtConflict
	}
	d.Scripts[name] = redis.NewScript(scriptText)
	return d.Scripts[name], nil
}
func (d *RedisConfig) AddScriptFromFile(file string, name string) (*redis.Script, error) {
	dat, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	return d.AddScript(string(dat), name)
}

// it's like, to build it, it needs to find the functions it needs, and build a script with it, outputting it
