package service

func NewUserService(svc UserService, mdw []UserMiddleware) UserService {
	for _, m := range mdw {
		svc = m(svc)
	}
	return svc
}
