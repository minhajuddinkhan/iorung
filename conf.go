package iorung

type Redis struct {
	RedisURL string
}

//Conf webrung conf
type Conf struct {
	Port        string
	IORungPort  int
	AuthRedis   Redis
	SocketRedis Redis
}
