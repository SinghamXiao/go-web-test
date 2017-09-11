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

func (this *RedisClient) StartAndGC(config string) error {
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

	this.key = cf["key"]
	this.connInfo = cf["conn"]
	this.dbNum, _ = strconv.Atoi(cf["dbNum"])
	this.password = cf["password"]

	this.connectInit()

	c := this.p.Get()
	defer c.Close()

	return c.Err()
}

// connect to redis.
func (this *RedisClient) connectInit() {
	dialFunc := func() (c redis.Conn, err error) {
		c, err = redis.Dial("tcp", this.connInfo)
		if err != nil {
			return nil, err
		}

		if this.password != "" {
			if _, err := c.Do("AUTH", this.password); err != nil {
				c.Close()
				return nil, err
			}
		}

		_, selectErr := c.Do("SELECT", this.dbNum)
		if selectErr != nil {
			c.Close()
			return nil, selectErr
		}
		return
	}
	// initialize a new pool
	this.p = &redis.Pool{
		MaxIdle:     16,
		MaxActive:   1024,
		IdleTimeout: 180 * time.Second,
		Dial:        dialFunc,
	}
}

// actually do the redis cmds
func (this *RedisClient) do(commandName string, args ...interface{}) (reply interface{}, err error) {
	c := this.p.Get()
	defer c.Close()

	return c.Do(commandName, args...)
}

func (this *RedisClient) Del(key string) (int, error) {
	return redis.Int(this.do("DEL", key))
}

func (this *RedisClient) Exists(key string) (int, error) {
	return redis.Int(this.do("EXISTS", key))
}

func (this *RedisClient) ExpireAt(key string, t uint64) (int, error) {
	return redis.Int(this.do(" EXPIREAT", key))
}

func (this *RedisClient) Set(key string, val interface{}) error {
	_, err := this.do("SET", key, val)
	return err
}

func (this *RedisClient) Get(key string) (string, error) {
	return redis.String(this.do("GET", key))
}

func (this *RedisClient) HGet(key, field string) (string, error) {
	return redis.String(this.do("HGET", key))
}

func (this *RedisClient) HMGet(key, field string) (string, error) {
	return redis.String(this.do("HGET", key))
}

func (this *RedisClient) HGetAll(key string) ([]string, error) {
	return redis.Strings(this.do("HGETALL", key))
}

func (this *RedisClient) HSet(key, field, value string) (int, error) {
	return redis.Int(this.do("HSET", key, field, value))
}

func (this *RedisClient) HMSet(key, field, value string) (int, error) {
	return redis.Int(this.do("HMSET", key, field, value))
}

func NewRedisClient(config string) (*RedisClient, error) {
	this := new(RedisClient)
	err := this.StartAndGC(config)
	if err != nil {
		return nil, err
	}

	return this, nil
}
