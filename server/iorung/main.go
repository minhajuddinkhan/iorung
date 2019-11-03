package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/davecgh/go-spew/spew"
	socketio "github.com/googollee/go-socket.io"
	"github.com/minhajuddinkhan/iorung"
	iorpc "github.com/minhajuddinkhan/iorung/io.rpc"
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

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("empty jwt secret")
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
		JWTSecret: jwtSecret,
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

	l, err := iorpc.NewIOListener(&conf)
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
