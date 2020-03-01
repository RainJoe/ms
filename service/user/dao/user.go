package dao

import (
	"context"

	pb "github.com/RainJoe/ms/service/user/proto/user"
	"github.com/garyburd/redigo/redis"
)

const (
	createUserSQL = "INSERT INTO users (u_pretty_id, u_name, u_password, u_create_time, u_update_time) VALUES ($1, $2, $3, $4, $5);"
	getUserSQL    = "SELECT u_pretty_id, u_name, u_create_time, u_update_time FROM users WHERE u_pretty_id = $1 AND u_delete_time = 0;"
	getAllUserSQL = "SELECT u_pretty_id, u_name, u_create_time, u_update_time FROM users WHERE u_delete_time = 0 LIMIT $1 OFFSET $2;"
	authUserSQL   = "SELECT u_pretty_id, u_name, u_create_time, u_update_time FROM users WHERE u_name = $1 AND u_password = $2 AND u_delete_time = 0;"
)

//CreateUser create user
func (d *Dao) CreateUser(ctx context.Context, user *pb.User) error {
	stmt, err := d.DB.PrepareContext(ctx, createUserSQL)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, user.PrettyId, user.Name, user.Password, user.CreateTime, user.UpdateTime)
	if err != nil {
		return err
	}
	return nil
}

//GetUser get user
func (d *Dao) GetUser(ctx context.Context, user *pb.User) (*pb.User, error) {
	err := d.DB.QueryRowContext(ctx, getUserSQL, user.PrettyId).Scan(&user.PrettyId, &user.Name, &user.CreateTime, &user.UpdateTime)
	if err != nil {
		return nil, err
	}
	return user, nil
}

//GetAllUsers get all user
func (d *Dao) GetAllUsers(ctx context.Context, req *pb.Request) ([]*pb.User, error) {
	rows, err := d.DB.QueryContext(ctx, getAllUserSQL, req.PageSize, req.PageNum-1)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	users := make([]*pb.User, 0)
	for rows.Next() {
		var user pb.User
		err := rows.Scan(&user.PrettyId, &user.Name, &user.CreateTime, &user.UpdateTime)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

//ValidateUser validate user
func (d *Dao) ValidateUser(ctx context.Context, name, password string) (*pb.User, error) {
	var user pb.User
	err := d.DB.QueryRowContext(ctx, authUserSQL, name, password).Scan(&user.PrettyId, &user.Name, &user.CreateTime, &user.UpdateTime)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

//SetUserToken2Redis set token
func (d *Dao) SetUserToken2Redis(key, value string) error {
	conn := d.redisPool.Get()
	defer conn.Close()
	_, err := conn.Do("SET", key, value)
	if err != nil {
		return err
	}
	return nil
}

//GetUserTokenFromRedis get token
func (d *Dao) GetUserTokenFromRedis(key string) (string, error) {
	conn := d.redisPool.Get()
	defer conn.Close()
	result, err := redis.String(conn.Do("GET", key))
	if err != nil {
		return "", err
	}
	return result, nil
}

//DelUserTokenFromRedis delete token
func (d *Dao) DelUserTokenFromRedis(key string) error {
	conn := d.redisPool.Get()
	defer conn.Close()
	_, err := conn.Do("DEL", key)
	if err != nil {
		return err
	}
	return nil
}
