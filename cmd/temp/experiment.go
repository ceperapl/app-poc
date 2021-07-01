package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/ceperapl/app-poc/pkg/models"
	"github.com/ceperapl/app-poc/pkg/repository/memory"
	"github.com/ceperapl/app-poc/pkg/usecase"
)

func main() {
	memoryTaskRepo, err := memory.NewTaskRepo()
	if err != nil {
		log.Fatal(err)
	}
	taskService := usecase.NewTaskService(memoryTaskRepo)

	printTasks(taskService)

	for i := 0; i < 9; i++ {
		if _, err = taskService.CreateTask(&models.Task{Description: "task" + strconv.Itoa(i)}); err != nil {
			log.Fatal(err)
		}
	}
	if _, err = taskService.CreateTask(&models.Task{ID: "f4636a57-2237-468d-ba46-016e9a3c62e6", Description: "task10"}); err != nil {
		log.Fatal(err)
	}
	if _, err = taskService.CreateTask(&models.Task{ID: "f4636a57-2237-468d-ba46-016e9a3c62e7", Description: "task11"}); err != nil {
		log.Fatal(err)
	}

	if _, err = taskService.GetTask("f4636a57-2237-468d-ba46-016e9a3c62e6"); err != nil {
		log.Fatal(err)
	}

	printTasks(taskService)

	_ = taskService.DeleteTask("f4636a57-2237-468d-ba46-016e9a3c62e6")

	printTasks(taskService)

	task11, err := taskService.GetTask("f4636a57-2237-468d-ba46-016e9a3c62e7")
	if err != nil {
		log.Fatal(err)
	}
	task11.Description = "Task11"

	if _, err = taskService.UpdateTask(task11); err != nil {
		log.Fatal(err)
	}
	printTasks(taskService)

	_ = taskService.DeleteTask(task11.ID)

	printTasks(taskService)
}

func printTasks(taskService usecase.TaskService) {
	tasks, count, _ := taskService.ListTasks("", "", 0, 0)
	fmt.Println(count)
	for _, task := range tasks {
		fmt.Printf("%v\n", task)
	}
}
