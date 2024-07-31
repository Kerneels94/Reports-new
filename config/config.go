package config

var (
	userToken = ""
)

func SetUserToken(newUserToken string) {
	userToken = newUserToken
}
