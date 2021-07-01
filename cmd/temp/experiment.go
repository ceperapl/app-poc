package main

import (
	"context"
	"log"

	api "github.com/ceperapl/app-poc/pkg/delivery/grpc/pb"
	"google.golang.org/grpc"
)

const (
	serverAddr = "localhost:9090"
)

func main() {
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := api.NewTasksClient(conn)
	ctx := context.Background()

	// List tasks
	err = listTasks(ctx, client)
	if err != nil {
		log.Fatal(err)
	}

	// Create tasks
	tasks := []api.Task{
		{
			Id:          "00000000-0000-0000-0000-000000000001",
			Description: "task1",
		},
		{
			Id:          "00000000-0000-0000-0000-000000000002",
			Description: "task2",
		},
		{
			Id:          "00000000-0000-0000-0000-000000000003",
			Description: "task3",
		},
	}

	if err := createTasks(ctx, client, tasks); err != nil {
		log.Fatal(err)
	}

	// List tasks
	err = listTasks(ctx, client)
	if err != nil {
		log.Fatal(err)
	}

	if _, err = client.DeleteTask(ctx, &api.DeleteTaskRequest{Id: "00000000-0000-0000-0000-000000000001"}); err != nil {
		log.Fatal(err)
	}

	// List tasks
	err = listTasks(ctx, client)
	if err != nil {
		log.Fatal(err)
	}
}

func createTasks(ctx context.Context, client api.TasksClient, tasks []api.Task) error {
	for _, task := range tasks {
		_, err := client.CreateTask(ctx, &task)
		if err != nil {
			return err
		}
	}

	return nil
}

func listTasks(ctx context.Context, client api.TasksClient) error {
	response, err := client.ListTasks(ctx, &api.ListTasksRequest{})
	if err != nil {
		return err
	}

	tasks := response.Result
	log.Println("List tasks:", len(tasks))
	for _, task := range tasks {
		log.Printf("Task: Id: %s, Description: %s, CreatedAt: %v, UpdatedAt: %v\n",
			task.Id, task.Description, task.CreatedAt, task.UpdatedAt)
	}
	return nil
}
