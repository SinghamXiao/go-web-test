package service

import (
	"github.com/garyburd/redigo/redis"
	"encoding/json"
	"strconv"
	"errors"
	"time"
)

var (
	// DefaultKey the collection name of redis for user.
	DefaultKey = ""
)

type RedisClient struct {
	p        *redis.Pool
	connInfo string
	dbNum    int
	key      string
	password string
}

func (rc *RedisClient) StartAndGC(config string) error {
	var cf map[string]string = make(map[string]string)
	json.Unmarshal([]byte(config), &cf)

	if _, ok := cf["key"]; !ok {
		cf["key"] = DefaultKey
	}
	if _, ok := cf["conn"]; !ok {
		return errors.New("config has no conn key")
	}
	if _, ok := cf["dbNum"]; !ok {
		cf["dbNum"] = "0"
	}
	if _, ok := cf["password"]; !ok {
		cf["password"] = ""
	}
	rc.key = cf["key"]
	rc.connInfo = cf["conn"]
	rc.dbNum, _ = strconv.Atoi(cf["dbNum"])
	rc.password = cf["password"]

	rc.connectInit()

	c := rc.p.Get()
	defer c.Close()

	return c.Err()
}

// connect to redis.
func (rc *RedisClient) connectInit() {
	dialFunc := func() (c redis.Conn, err error) {
		c, err = redis.Dial("tcp", rc.connInfo)
		if err != nil {
			return nil, err
		}

		if rc.password != "" {
			if _, err := c.Do("AUTH", rc.password); err != nil {
				c.Close()
				return nil, err
			}
		}

		_, selectErr := c.Do("SELECT", rc.dbNum)
		if selectErr != nil {
			c.Close()
			return nil, selectErr
		}
		return
	}
	// initialize a new pool
	rc.p = &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 180 * time.Second,
		Dial:        dialFunc,
	}
}

// actually do the redis cmds
func (rc *RedisClient) do(commandName string, args ...interface{}) (reply interface{}, err error) {
	c := rc.p.Get()
	defer c.Close()

	return c.Do(commandName, args...)
}

func (rc *RedisClient) Set(key string, val interface{}) error {
	_, err := rc.do("SET", key, val)
	return err
}

func (rc *RedisClient) Get(key string) (string, error) {
	return redis.String(rc.do("GET", key))
}

func NewRedisClient(config string) (*RedisClient, error) {
	rc := new(RedisClient)
	err := rc.StartAndGC(config)
	if err != nil {
		return nil, err
	}

	return rc, nil
}
