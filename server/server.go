package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	engineio "github.com/googollee/go-engine.io"
	"github.com/googollee/go-engine.io/transport"
	"github.com/googollee/go-engine.io/transport/polling"
	"github.com/googollee/go-engine.io/transport/websocket"
	socketio "github.com/googollee/go-socket.io"
	"github.com/minhajuddinkhan/iorung/cache/auth"
	"github.com/minhajuddinkhan/iorung/config"
	"github.com/minhajuddinkhan/iorung/controllers/game"
	iorpc "github.com/minhajuddinkhan/iorung/io.rpc"
)

//Start starts the server
func Start(conf config.Conf) error {

	transports := []transport.Transport{
		polling.Default,
		websocket.Default,
	}
	server, err := socketio.NewServer(&engineio.Options{
		Transports:   transports,
		PingInterval: 1 * time.Second,
	})

	if err != nil {
		log.Fatal(err)
	}

	server.OnConnect("/", func(s socketio.Conn) error {
		fmt.Println("socket connected", s.ID())
		return nil
	})

	go server.Serve()
	defer server.Close()

	// conf := config.New()
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

	server.OnConnect("/", func(s socketio.Conn) error {
		fmt.Println("socket connectec", s.ID())
		return nil
	})

	redis, err := auth.NewAuthRedis(&conf)
	if err != nil {
		log.Fatal(err)
	}
	server.OnEvent("/", "authenticate", game.JoinGame(redis))

	http.Handle("/socket.io/", server)
	fmt.Println("LISTENING ON PORT", conf.Port)
	return http.ListenAndServe(fmt.Sprintf(":%s", conf.Port), nil)

}
