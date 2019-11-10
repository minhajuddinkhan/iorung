package server

import (
	"fmt"
	"log"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
	"github.com/minhajuddinkhan/iorung/config"
	iorpc "github.com/minhajuddinkhan/iorung/io.rpc"
)

//Start starts the server
func Start(conf config.Conf) error {

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

	http.Handle("/socket.io/", server)
	fmt.Println("LISTENING ON PORT", conf.Port)
	return http.ListenAndServe(fmt.Sprintf(":%s", conf.Port), nil)

}
