package users

type UserSv interface {
	
}

type UserServ struct {
	repo UserStore
}

func NewUserService(repo UserStore) *UserServ {
	return &UserServ{repo: repo}
}