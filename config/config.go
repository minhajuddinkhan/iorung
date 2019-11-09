package config

import (
	"log"
	"net/url"
	"os"
	"strconv"
)

type Redis struct {
	RedisURL string
}

//DB DB config
type DB struct {
	Connection string
	DBName     string
}

//Conf webrung conf
type Conf struct {
	Port        string
	IORungPort  int
	JWTSecret   string
	AuthRedis   Redis
	SocketRedis Redis
	DB          DB
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

	dbURI := os.Getenv("MONGODB_URI")
	if dbURI == "" {
		log.Fatal("empty db uri")
	}
	parsedURI, err := url.Parse(dbURI)
	if err != nil {
		log.Fatal(err)
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
		DB: DB{
			Connection: dbURI,
			DBName:     parsedURI.Path[1:],
		},
	}

}
