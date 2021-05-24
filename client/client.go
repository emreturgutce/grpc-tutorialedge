package main

import (
	"fmt"
	"log"
	"math/rand"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/emreturgutce/grpc-tutorialedge/chat"
)

func CreateTodo(c chat.TodoServiceClient) (*chat.Message, error) {
	id := rand.Int31()
	return c.AddTodo(context.Background(), &chat.Todo{
		Id:          id,
		Title:       fmt.Sprintf("Todo %d", id),
		Description: fmt.Sprintf("Description %d", id),
	})
}

func GetTodos(c chat.TodoServiceClient) (*chat.Todos, error) {
	return c.GetTodos(context.Background(), &chat.Message{})
}

func GetTodoById(id int32, c chat.TodoServiceClient) (*chat.Todo, error) {
	return c.GetTodoById(context.Background(), &chat.TodoId{Id: id})
}

func main() {
	var conn *grpc.ClientConn

	conn, err := grpc.Dial(":9000", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}

	defer conn.Close()

	c := chat.NewTodoServiceClient(conn)

	response, err := GetTodoById(1298498081, c)

	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}

	log.Printf("Response from server: %s", response.String())
}
