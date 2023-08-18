package auth_service

//import "../simple-demo-main/model"
import "github.com/RaymondCode/simple-demo/model"

type Auth struct {
	Username string
	Password string
}

func (a *Auth) Check() (bool, error, int) {
	return model.CheckAuth(a.Username, a.Password)
}
