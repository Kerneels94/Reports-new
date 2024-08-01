package config

var (
	userToken = ""
)

func SetUserToken(newUserToken string) {
	userToken = newUserToken
}

func GetUserToken() string {
	return userToken
}
