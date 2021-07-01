package grpc

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/ceperapl/app-poc/pkg/delivery/grpc/pb"
	"github.com/ceperapl/app-poc/pkg/models"
	"github.com/ceperapl/app-poc/pkg/usecase"
	"github.com/gofrs/uuid"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var NOT_FOUND_ERROR = errors.New("Entity not found")

func NewTaskServerGrpc(gserver *grpc.Server, taskService usecase.TaskService) {

	taskServer := &server{
		usecase: taskService,
	}

	pb.RegisterTasksServer(gserver, taskServer)
}

type server struct {
	usecase usecase.TaskService
}

func (s *server) transformTask(t *models.Task) (*pb.Task, error) {
	var err error

	if t == nil {
		return nil, nil
	}

	taskPB := &pb.Task{
		Id:          t.ID,
		Description: t.Description,
		Completed:   t.Completed,
	}
	if taskPB.Deadline, err = ptypes.TimestampProto(t.Deadline); err != nil {
		return nil, err
	}
	if taskPB.CreatedAt, err = ptypes.TimestampProto(t.CreatedAt); err != nil {
		return nil, err
	}
	if taskPB.UpdatedAt, err = ptypes.TimestampProto(t.UpdatedAt); err != nil {
		return nil, err
	}
	return taskPB, nil
}

func (s *server) transformTaskPB(taskPB *pb.Task) *models.Task {
	deadline := time.Unix(taskPB.Deadline.GetSeconds(), 0)
	createdAt := time.Unix(taskPB.CreatedAt.GetSeconds(), 0)
	updatedAt := time.Unix(taskPB.UpdatedAt.GetSeconds(), 0)

	task := &models.Task{
		ID:          taskPB.Id,
		Description: taskPB.Description,
		Deadline:    deadline,
		Completed:   taskPB.Completed,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}

	return task
}

func (s *server) CreateTask(ctx context.Context, req *pb.Task) (*pb.Task, error) {
	task := s.transformTaskPB(req)
	createdTask, err := s.usecase.CreateTask(task)
	result, err := s.transformTask(createdTask)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *server) GetTask(ctx context.Context, req *pb.GetTaskRequest) (*pb.Task, error) {
	id := req.Id

	if !isValidUUID(id) {
		return nil, status.Error(codes.InvalidArgument, "ERROR: Invalid UUID: "+id)
	}

	task, err := s.usecase.GetTask(id)
	if err != nil {
		return nil, err
	}

	if task == nil {
		return nil, status.Error(codes.NotFound, "ERROR: Task is not found: "+id)
	}

	taskPB, err := s.transformTask(task)
	if err != nil {
		return nil, fmt.Errorf("ERROR: Unable transform task to protobuf task: %v", err)
	}

	return taskPB, nil
}

func (s *server) ListTasks(ctx context.Context, req *pb.ListTasksRequest) (*pb.ListTasksResponse, error) {
	tasks, count, err := s.usecase.ListTasks(req.FilterBy, req.SortBy, int(req.Limit), int(req.Page))
	if err != nil {
		return nil, err
	}
	pbTasks := []*pb.Task{}
	for _, task := range tasks {
		taskPB, err := s.transformTask(&task)
		if err != nil {
			return nil, fmt.Errorf("ERROR: Unable transform task to protobuf task: %v", err)
		}
		pbTasks = append(pbTasks, taskPB)
	}
	return &pb.ListTasksResponse{Result: pbTasks, Count: int32(count)}, nil
}

func (s *server) UpdateTask(ctx context.Context, req *pb.Task) (*pb.Task, error) {
	return nil, nil
}

func (s *server) DeleteTask(ctx context.Context, req *pb.DeleteTaskRequest) (*empty.Empty, error) {
	return nil, nil
}

func isValidUUID(u string) bool {
	_, err := uuid.FromString(u)
	return err == nil
}
