package auth

func (mgr *manager) Authenticate(token string) (string, string, error) {
	return mgr.authRedis.Get(token)
}
