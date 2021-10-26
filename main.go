package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type (
	// Todo -
	Todo struct {
		ID   int    `json:"id,omitempty"`
		Msg  string `json:"msg,omitempty"`
		Done bool   `json:"done,omitempty"`
	}
	// Todos -
	Todos struct {
		Todos []Todo `json:"todos,omitempty"`
	}
	// ClientRequest -
	ClientRequest struct {
		ID   int    `json:"id,omitempty"`
		Type string `json:"type,omitempty"`
		Todo Todo   `json:"todos,omitempty"`
	}
	//ClientResposne -
	ClientResposne struct {
		Todos []Todo `json:"todos,omitempty"`
	}
)

func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error while upgrading the request and resopnse:", err.Error())
		return
	}

	for {

		request := &ClientRequest{}
		err := conn.ReadJSON(request)
		if err != nil {
			log.Println("Error while reading the request:", err.Error())
			return
		}
		log.Println("Msg from client :", request)

		response := &ClientResposne{Todos: []Todo{}}

		switch request.Type {
		case "ADD":
			response = addTodo(request, response)
		case "DELETE":
			response.Todos = removeTodo(response.Todos, request.ID)
		}

		if err := conn.WriteJSON(response); err != nil {
			log.Println("Erro while responding:", err.Error())
			return
		}

	}
}

func addTodo(request *ClientRequest, response *ClientResposne) *ClientResposne {
	request.Todo.ID++
	response.Todos = append(response.Todos, request.Todo)
	return response
}

func removeTodo(todos []Todo, id int) []Todo {
	var temp Todos
	for _, v := range todos {
		if id != v.ID {
			temp.Todos = append(temp.Todos, v)
		}
	}
	return temp.Todos
}

func main() {

}
