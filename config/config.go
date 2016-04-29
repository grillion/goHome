package config

var (
	mongoHost string
	mongoDb string
	webAppRoot string
)

const SessionSeconds = 3600

func init(){
	mongoDb = "goHome"
	mongoHost = "192.168.10.10"
	webAppRoot = "/Users/mgrill/Workspaces/go/src/github.com/grillion/goHome/public"
}

func GetMongoHost() string {
	return mongoHost
}

func GetMongoDBName() string {
	return mongoDb
}

func GetWebAppRoot() string {
	return webAppRoot
}