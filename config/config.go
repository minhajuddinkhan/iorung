package config

import (
	"log"
	"os"
	"strconv"
)

type Redis struct {
	RedisURL string
}

//Conf webrung conf
type Conf struct {
	Port        string
	IORungPort  int
	JWTSecret   string
	AuthRedis   Redis
	SocketRedis Redis
}

func New() Conf {

	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		log.Fatal("empty http port from PORT env")
	}

	authRedis := os.Getenv("AUTH_REDIS_URL")
	if authRedis == "" {
		log.Fatal("empty auth redis url")
	}

	socketRedis := os.Getenv("SOCKET_REDIS_URL")
	if socketRedis == "" {
		log.Fatal("empty socket redis url")
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("empty jwt secret")
	}

	rpcPort, err := strconv.Atoi(os.Getenv("RPC_PORT"))
	if err != nil {
		log.Fatal(err)
	}
	return Conf{
		Port:       httpPort,
		IORungPort: rpcPort,
		AuthRedis: Redis{
			RedisURL: authRedis,
		},
		SocketRedis: Redis{
			RedisURL: socketRedis,
		},
		JWTSecret: jwtSecret,
	}

}
