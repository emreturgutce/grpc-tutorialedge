package chat

import (
	"log"

	"golang.org/x/net/context"
)

var (
	todos []*Todo
)

type Server struct{}

func (s *Server) AddTodo(ctx context.Context, in *Todo) (*Message, error) {
	log.Printf("Receive message body from client: %s", in.String())
	todos = append(todos, &Todo{Id: in.Id, Title: in.Title, Description: in.Description})
	return &Message{}, nil
}

func (s *Server) GetTodos(ctx context.Context, in *Message) (*Todos, error) {
	log.Printf("Receive message body from client: %s %s", in.String(), ctx)
	return &Todos{Todos: todos}, nil
}

func (s *Server) GetTodoById(ctx context.Context, in *TodoId) (*Todo, error) {
	log.Printf("Receive message body from client: %s", in.String())
	var foundTodo *Todo
	for _, todo := range todos {
		if todo.Id == in.Id {
			foundTodo = todo
		}
	}
	return foundTodo, nil
}
