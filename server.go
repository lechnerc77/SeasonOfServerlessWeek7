package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

type Recipe struct {
	Country    string `json:"country"`
	RecipeName string `json:"recipeName"`
	RecipeLink string `json:"recipeLink"`
}

var Recipes []Recipe

func recipeHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	lastSegment := path.Base(r.URL.Path)

	searchCountry := strings.Title(lastSegment)

	var foundRecipesForCountry []Recipe

	for _, recipeentry := range Recipes {
		if recipeentry.Country == searchCountry {
			foundRecipesForCountry = append(foundRecipesForCountry, recipeentry)
		}
	}

	urlParameter := r.URL.Query().Get("recipename")

	var foundRecipesByName []Recipe

	if urlParameter != "" {

		for _, recipeentry := range foundRecipesForCountry {
			if recipeentry.RecipeName == urlParameter {
				foundRecipesByName = append(foundRecipesByName, recipeentry)
			}
		}

		if len(foundRecipesByName) > 0 {

			json.NewEncoder(w).Encode(foundRecipesByName)

		} else {

			json.NewEncoder(w).Encode("No entry found")

		}

	} else {

		json.NewEncoder(w).Encode(foundRecipesForCountry)

	}
}

func main() {

	Recipes = []Recipe{
		{Country: "Nigeria", RecipeName: "Jollof Rice", RecipeLink: "https://food52.com/recipes/61557-classic-nigerian-jollof-rice"},
		{Country: "Nigeria", RecipeName: "Egusi Soup", RecipeLink: "https://www.demandafrica.com/food/recipes/nigerian-egusi-soup/"},
		{Country: "Nigeria", RecipeName: "Nkwobi", RecipeLink: "https://www.allnigerianrecipes.com/restaurant/nkwobi/"},
		{Country: "Kenya", RecipeName: "Tilapia Fry", RecipeLink: "https://www.allrecipes.com/recipe/279239/deep-fried-tilapia/"},
		{Country: "Kenya", RecipeName: "Ugali", RecipeLink: "https://cookpad.com/ke/recipes/4727293-how-to-cook-the-perfect-ugali"},
		{Country: "Kenya", RecipeName: "Karanga", RecipeLink: "https://www.allrecipes.com/recipe/255437/karanga-soup/"},
	}

	listenAddr := ":8080"

	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		listenAddr = ":" + val
	}

	http.HandleFunc("/api/recipe/kenya", recipeHandler)
	http.HandleFunc("/api/recipe/nigeria", recipeHandler)
	log.Printf("About to listen on %s. Go to http://127.0.0.1%s", listenAddr, listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}
