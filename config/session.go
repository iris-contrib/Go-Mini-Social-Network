package config

import (
	"os"

	"github.com/gorilla/securecookie"
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
)

var manager *sessions.Sessions

func init() {
	// These should be initialized when godotenv inited,
	// the original repository has a bug here and the session secret
	// never loads because the "config" packages inits first.

	var (
		hashKey  = []byte(os.Getenv("SESSION_HASH_SECRET"))
		blockKey = []byte(os.Getenv("SESSION_BLOCK_SECRET"))
	)
	var secureCookie = securecookie.New(hashKey, blockKey)

	manager = sessions.New(sessions.Config{
		Cookie: "session_id",
		Encode: secureCookie.Encode,
		Decode: secureCookie.Decode,
	})

	// println("hashKey = " + string(hashKey))
	// println("blockKey = " + string(blockKey))
}

// GetSession to return the session
func GetSession(ctx iris.Context) *sessions.Session {
	return manager.Start(ctx)
}

// AllSessions function to return all the sessions
func AllSessions(ctx iris.Context) (string, string) {
	session := GetSession(ctx)
	id := session.GetString("id")
	username := session.GetString("username")
	return id, username
}
