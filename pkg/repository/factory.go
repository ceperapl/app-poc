package repository

// import (
// 	"fmt"

// 	"github.com/ceperapl/app-poc/pkg/repository/sqlite"
// )

// func FactoryTaskRepository(repoType string) (TaskRepository, error) {
// 	switch repoType {
// 	case "sqlite":
// 		return sqlite.NewSQLiteTaskRepo()
// 	default:
// 		return nil, fmt.Errorf("Repository was not created. Invalid type: %s", repoType)
// 	}
// }
