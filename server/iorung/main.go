package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/googollee/go-socket.io"
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

	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}

	server.OnConnect("/", func(s socketio.Conn) error {
		fmt.Println("socket connected", s.ID())
		return nil
	})

	go server.Serve()
	defer server.Close()
	http.Handle("/socket.io/", server)
	spew.Dump("LISTENING ON PORT", conf.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", conf.Port), nil))

}
