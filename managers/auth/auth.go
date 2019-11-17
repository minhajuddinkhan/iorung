package auth

func (mgr *manager) Authenticate(token string) (uint, uint, error) {
	return mgr.authRedis.Get(token)
}
