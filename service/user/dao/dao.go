package dao

import (
	"database/sql"

	"github.com/RainJoe/ms/library/pg"
	rs "github.com/RainJoe/ms/library/redis"
	"github.com/RainJoe/ms/service/user/config"
	"github.com/garyburd/redigo/redis"
)

//Dao database access object
type Dao struct {
	DB        *sql.DB
	redisPool *redis.Pool
}

//New dao
func New(c *config.Config) (d *Dao) {
	return &Dao{
		DB:        pg.NewPostgres(&c.Pg),
		redisPool: rs.NewRedisPool(&c.Redis.Config),
	}
}

//Close dao
func (d *Dao) Close() {
	if d.DB != nil {
		d.DB.Close()
	}
	if d.redisPool != nil {
		d.redisPool.Close()
	}
}
