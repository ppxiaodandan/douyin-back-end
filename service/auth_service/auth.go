package auth_service

//import "../simple-demo-main/model"
import "github.com/ppxiaodandan/douyin-back-end/model"

type Auth struct {
	Username string
	Password string
}

func (a *Auth) Check() (bool, error, int) {
	return model.CheckAuth(a.Username, a.Password)
}
