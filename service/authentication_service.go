package service

var (
	AuthService *AuthenticationServiceInterface = AuthenticationService{}
)

type AuthenticationService struct {

}

type AuthenticationServiceInterface interface {
	Login(userName string, password string)
	Signup(userName string, password string, firstName string , email string)
	Signout(userName string)
}

func (*AuthenticationService) Login(userName string, password string) {

}

func (*AuthenticationService) Signup(userName string, password string, firstName string , email string) {

}

func (*AuthenticationService) Signout(userName string) {

}

