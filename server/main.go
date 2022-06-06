package main

import (
	"log"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
	"github.com/shailesz/cli-chat-golang-server/src/models"
	"github.com/shailesz/cli-chat-golang-server/src/services"
)

func main() {
	services.InitConnection()

	server := socketio.NewServer(nil)

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		log.Println("connected:", s.ID())
		s.Join("chatroom")
		return nil
	})

	server.OnEvent("/", "auth", func(s socketio.Conn, user models.User) {
		status := user.Login(services.Conn)

		if status == 200 {
			res := models.AuthMessage{Status: 200, Data: user}
			s.Emit("auth", res)
		} else {
			res := models.AuthMessage{Status: 404, Data: user}
			s.Emit("auth", res)
		}
	})

	server.OnEvent("/", "signup", func(s socketio.Conn, user models.User) {
		status := user.Signup(services.Conn)

		if status == 200 {
			res := models.AuthMessage{Status: 200, Data: user}
			s.Emit("auth", res)
		} else {
			res := models.AuthMessage{Status: 404, Data: user}
			s.Emit("auth", res)
		}
	})

	server.OnEvent("/", "chat", func(s socketio.Conn, msg models.ChatMessage) {
		server.BroadcastToRoom("/", "chatroom", "message", msg)
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		log.Println("meet error:", e)
	})

	go func() {
		if err := server.Serve(); err != nil {
			log.Fatalf("socketio listen error: %s\n", err)
		}
	}()
	defer server.Close()

	http.Handle("/socket.io/", server)

	log.Println("Serving at localhost:8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
