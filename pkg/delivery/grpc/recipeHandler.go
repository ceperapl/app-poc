package grpc

import (
	"context"
	"fmt"
	"time"

	"github.com/ceperapl/app-poc/pkg/delivery/grpc/pb"
	"github.com/ceperapl/app-poc/pkg/models"
	"github.com/ceperapl/app-poc/pkg/usecase"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewRecipeServerGrpc(gserver *grpc.Server, recipeService usecase.RecipeService) {

	server := &recipeServer{
		usecase: recipeService,
	}

	pb.RegisterRecipesServer(gserver, server)
}

type recipeServer struct {
	usecase usecase.RecipeService
}

func (s *recipeServer) transformRecipe(recipe *models.Recipe) (*pb.Recipe, error) {
	var err error

	if recipe == nil {
		return nil, nil
	}

	var pbIngredients []*pb.Ingredient
	for _, ingredient := range recipe.Ingredients {
		pbIngredient := pb.Ingredient{
			Name:        ingredient.Name,
			Amount:      ingredient.Amount,
			Measurement: ingredient.Measurement,
		}
		pbIngredients = append(pbIngredients, &pbIngredient)
	}

	recipePB := &pb.Recipe{
		Id:          recipe.ID,
		Name:        recipe.Name,
		Ingredients: pbIngredients,
		Steps:       recipe.Steps,
	}

	if recipePB.CreatedAt, err = ptypes.TimestampProto(recipe.CreatedAt); err != nil {
		return nil, err
	}
	if recipePB.UpdatedAt, err = ptypes.TimestampProto(recipe.UpdatedAt); err != nil {
		return nil, err
	}
	return recipePB, nil
}

func (s *recipeServer) transformRecipePB(recipePB *pb.Recipe) *models.Recipe {
	createdAt := time.Unix(recipePB.CreatedAt.GetSeconds(), 0)
	updatedAt := time.Unix(recipePB.UpdatedAt.GetSeconds(), 0)

	var ingredients []models.Ingredient
	for _, pbIngredient := range recipePB.Ingredients {
		ingredient := models.Ingredient{
			Name:        pbIngredient.Name,
			Amount:      pbIngredient.Amount,
			Measurement: pbIngredient.Measurement,
		}
		ingredients = append(ingredients, ingredient)
	}

	recipe := &models.Recipe{
		ID:          recipePB.Id,
		Name:        recipePB.Name,
		Ingredients: ingredients,
		Steps:       recipePB.Steps,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}

	return recipe
}

func (s *recipeServer) CreateRecipe(ctx context.Context, req *pb.Recipe) (*pb.Recipe, error) {
	recipe := s.transformRecipePB(req)
	createdRecipe, err := s.usecase.CreateRecipe(recipe)
	result, err := s.transformRecipe(createdRecipe)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *recipeServer) GetRecipe(ctx context.Context, req *pb.GetRecipeRequest) (*pb.Recipe, error) {
	id := req.Id

	if !isValidUUID(id) {
		return nil, status.Error(codes.InvalidArgument, "ERROR: Invalid UUID: "+id)
	}

	recipe, err := s.usecase.GetRecipe(id)
	if err != nil {
		return nil, err
	}

	if recipe == nil {
		return nil, status.Error(codes.NotFound, "ERROR: Recipe is not found: "+id)
	}

	recipePB, err := s.transformRecipe(recipe)
	if err != nil {
		return nil, fmt.Errorf("ERROR: Unable transform recipe to protobuf recipe: %v", err)
	}

	return recipePB, nil
}

func (s *recipeServer) ListRecipes(ctx context.Context, req *pb.ListRecipesRequest) (*pb.ListRecipesResponse, error) {
	recipes, count, err := s.usecase.ListRecipes(req.FilterBy, req.SortBy, int(req.Limit), int(req.Page))
	if err != nil {
		return nil, err
	}
	pbRecipes := []*pb.Recipe{}
	for _, recipe := range recipes {
		recipePB, err := s.transformRecipe(&recipe)
		if err != nil {
			return nil, fmt.Errorf("ERROR: Unable transform recipe to protobuf recipe: %v", err)
		}
		pbRecipes = append(pbRecipes, recipePB)
	}
	return &pb.ListRecipesResponse{Result: pbRecipes, Count: int32(count)}, nil
}

func (s *recipeServer) UpdateRecipe(ctx context.Context, req *pb.Recipe) (*pb.Recipe, error) {
	return nil, nil
}

func (s *recipeServer) DeleteRecipe(ctx context.Context, req *pb.DeleteRecipeRequest) (*empty.Empty, error) {
	err := s.usecase.DeleteRecipe(req.Id)
	if err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}
