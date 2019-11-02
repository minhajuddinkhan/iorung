package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/davecgh/go-spew/spew"
	"github.com/googollee/go-socket.io"
	"github.com/minhajuddinkhan/iorung"
	"github.com/minhajuddinkhan/iorung/rpcalls"
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

	rpcPort, err := strconv.Atoi(os.Getenv("RPC_PORT"))
	if err != nil {
		log.Fatal(err)
	}
	conf := iorung.Conf{
		Port:       httpPort,
		IORungPort: rpcPort,
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

	l, err := rpcalls.NewAuthListener(&conf)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		err := http.Serve(l, nil)
		if err != nil {
			log.Fatal(err)
		}
	}()

	http.Handle("/socket.io/", server)
	spew.Dump("LISTENING ON PORT", conf.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", conf.Port), nil))

}
