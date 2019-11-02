package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/davecgh/go-spew/spew"
	"github.com/minhajuddinkhan/iorung"
)

func main() {

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
	conf := iorung.Conf{
		Port: httpPort,
		AuthRedis: iorung.Redis{
			RedisURL: authRedis,
		},
		SocketRedis: iorung.Redis{
			RedisURL: socketRedis,
		},
	}

	r := mux.NewRouter()
	http.Handle("/", r)

	spew.Dump("LISTENING ON PORT", conf.Port)
	http.ListenAndServe(fmt.Sprintf(":%s", conf.Port), nil)

}
