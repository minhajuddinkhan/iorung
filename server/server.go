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
	"github.com/minhajuddinkhan/iorung/socketpool"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		allowHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization"

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
		w.Header().Set("Access-Control-Allow-Methods", "POST, PUT, PATCH, GET, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", allowHeaders)

		next.ServeHTTP(w, r)
	})
}

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
		return nil
	})

	go server.Serve()
	defer server.Close()

	// conf := config.New()
	gp := socketpool.NewGamePool()
	l, err := iorpc.NewIOListener(&conf, gp)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		err := http.Serve(l, nil)
		if err != nil {
			log.Fatal(err)
		}
	}()

	redis, err := auth.NewAuthRedis(&conf)
	if err != nil {
		log.Fatal(err)
	}
	server.OnEvent("/", "authenticate", game.JoinGame(redis, gp))

	http.Handle("/socket.io/", corsMiddleware(server))
	fmt.Println("LISTENING ON PORT", conf.Port)
	return http.ListenAndServe(fmt.Sprintf(":%s", conf.Port), nil)

}
