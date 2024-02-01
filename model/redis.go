package model

import (
	"github.com/gin-contrib/sessions/redis"
)

var UserSession redis.Store

func init() {
	// UserSession, _ = redis.NewStore(10, "tcp", "localhost:6379", "", []byte("1qaz@WSX"))
	UserSession, _ = redis.NewStore(10, "tcp", "192.168.1.107:6379", "", []byte("1qaz@WSX"))
}
