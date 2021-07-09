package main

import (
	"context"
	"log"

	api "github.com/ceperapl/app-poc/pkg/delivery/grpc/pb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:9090", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := api.NewRecipesClient(conn)
	ctx := context.Background()

	// List recipes
	err = listRecipes(ctx, client)
	if err != nil {
		log.Fatal(err)
	}

	// Create recipes
	recipes := []api.Recipe{
		{
			Id:   "00000000-0000-0000-0000-000000000001",
			Name: "Baked Salmon",
			Ingredients: []*api.Ingredient{
				{Name: "Salmon", Amount: 1, Measurement: "lb"},
				{Name: "Pine Nuts", Amount: 1, Measurement: "cup"},
				{Name: "Butter Lettuce", Amount: 2, Measurement: "cups"},
				{Name: "Yellow Squash", Amount: 1, Measurement: "med"},
				{Name: "Olive Oil", Amount: 0.5, Measurement: "cup"},
				{Name: "Garlic", Amount: 3, Measurement: "cloves"},
			},
			Steps: []string{
				"Preheat the oven to 350 degrees.",
				"Spread the olive oil around a glass baking dish.",
				"Add the salmon, garlic, and pine nuts to the dish.",
				"Bake for 15 minutes.",
				"Add the yellow squash and put back in the oven for 30 mins.",
				"Remove from oven and let cool for 15 minutes. Add the lettuce and serve.",
			},
		},
		{
			Id:   "00000000-0000-0000-0000-000000000002",
			Name: "Fish Tacos",
			Ingredients: []*api.Ingredient{
				{Name: "Whitefish", Amount: 1, Measurement: "lb"},
				{Name: "Cheese", Amount: 1, Measurement: "cup"},
				{Name: "Iceberg Lettuce", Amount: 2, Measurement: "cups"},
				{Name: "Tomatoes", Amount: 2, Measurement: "large"},
				{Name: "Tortillas", Amount: 3, Measurement: "med"},
			},
			Steps: []string{
				"Cook the fish on the grill until hot.",
				"Place the fish on the 3 tortillas.",
				"Top them with lettuce, tomatoes, and cheese.",
			},
		},
	}

	if err := createRecipes(ctx, client, recipes); err != nil {
		log.Fatal(err)
	}

	// List recipes
	err = listRecipes(ctx, client)
	if err != nil {
		log.Fatal(err)
	}

	if _, err = client.DeleteRecipe(ctx, &api.DeleteRecipeRequest{Id: "00000000-0000-0000-0000-000000000001"}); err != nil {
		log.Fatal(err)
	}

	// List recipes
	err = listRecipes(ctx, client)
	if err != nil {
		log.Fatal(err)
	}
}

func createRecipes(ctx context.Context, client api.RecipesClient, recipes []api.Recipe) error {
	for _, recipe := range recipes {
		_, err := client.CreateRecipe(ctx, &recipe)
		if err != nil {
			return err
		}
	}

	return nil
}

func listRecipes(ctx context.Context, client api.RecipesClient) error {
	response, err := client.ListRecipes(ctx, &api.ListRecipesRequest{})
	if err != nil {
		return err
	}

	recipes := response.Result
	log.Println("List recipes:", len(recipes))
	for _, recipe := range recipes {
		log.Printf("Recipe: Id: %s, Description: %s, CreatedAt: %v, UpdatedAt: %v\n",
			recipe.Id, recipe.Name, recipe.CreatedAt, recipe.UpdatedAt)
	}
	return nil
}
