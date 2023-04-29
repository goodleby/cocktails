package main

import (
	"encoding/json"
	"log"
	"sort"

	"github.com/goodleby/cocktails/bar"
)

func main() {
	b := bar.New([]*bar.Ingredient{
		// Spirits
		{Name: "Tequila", AlcoholContents: 0.4},
		{Name: "Gin", AlcoholContents: 0.4},
		{Name: "Rum", AlcoholContents: 0.4},
		{Name: "Whiskey", AlcoholContents: 0.4},

		// Liqueurs
		{Name: "Cointreau", AlcoholContents: 0.4},
		{Name: "Campari", AlcoholContents: 0.28},
		{Name: "Sweet Vermouth", AlcoholContents: 0.2},

		// Bitters
		{Name: "Angostura Bitters", AlcoholContents: 0.44},

		// Juices
		{Name: "Lemon Juice", AlcoholContents: 0},
		{Name: "Lime Juice", AlcoholContents: 0},
		{Name: "Cranberry Juice", AlcoholContents: 0},
		{Name: "Orange Juice", AlcoholContents: 0},
		{Name: "Pineapple Juice", AlcoholContents: 0},
		{Name: "Pomegranate Juice", AlcoholContents: 0},
		{Name: "Blackcurrant Juice", AlcoholContents: 0},
		{Name: "Coconut Cream", AlcoholContents: 0},

		// Syrups
		{Name: "Simple Syrup", AlcoholContents: 0},
		{Name: "Grenadine Syrup", AlcoholContents: 0},

		// Sodas
		{Name: "Soda", AlcoholContents: 0},
		{Name: "Coke", AlcoholContents: 0},
		{Name: "Tonic", AlcoholContents: 0},
		{Name: "Grapefruit Soda", AlcoholContents: 0},

		// Garnishes
		{Name: "Egg White", AlcoholContents: 0},
		{Name: "Mint", AlcoholContents: 0},
		{Name: "Lime Wedge", AlcoholContents: 0},
		{Name: "Orange Wedge", AlcoholContents: 0},
		{Name: "Pineapple Wedge", AlcoholContents: 0},
		{Name: "Orange Peel", AlcoholContents: 0},
		{Name: "Cherry", AlcoholContents: 0},
	})

	drinks := []*bar.Drink{
		// Gin
		bar.Stir("Gin and Tonic", []*bar.Drink{
			b.Use("Gin", 1.5),
			b.Use("Tonic", 4),
			b.Use("Lime Wedge", 0),
		}),
		bar.Shake("Gin Fizz", []*bar.Drink{
			b.Use("Gin", 2),
			b.Use("Lemon Juice", 0.75),
			b.Use("Simple Syrup", 0.5),
			b.Use("Soda", 2),
			b.Use("Lime Wedge", 0),
		}),
		bar.Shake("Negroni", []*bar.Drink{
			b.Use("Gin", 1),
			b.Use("Sweet Vermouth", 1),
			b.Use("Campari", 1),
			b.Use("Orange Peel", 0),
		}),
		bar.Shake("Blackcurrant Gimlet", []*bar.Drink{
			b.Use("Gin", 1.5),
			b.Use("Blackcurrant Juice", 2),
			b.Use("Lime Juice", 0.75),
			b.Use("Simple Syrup", 0.5),
			b.Use("Lime Wedge", 0),
		}),

		// Tequila
		bar.Shake("Pineapple Margarita", []*bar.Drink{
			b.Use("Tequila", 1),
			b.Use("Lime Juice", 1),
			b.Use("Cointreau", 0.5),
			b.Use("Pineapple Juice", 2),
			b.Use("Lime Wedge", 0),
		}),
		bar.Stir("Paloma", []*bar.Drink{
			b.Use("Tequila", 1.5),
			b.Use("Grapefruit Soda", 3),
			b.Use("Lime Juice", 0.5),
			b.Use("Lime Wedge", 0),
		}),
		bar.Stir("Tequila Sunrise", []*bar.Drink{
			b.Use("Tequila", 1.5),
			b.Use("Orange Juice", 3),
			b.Use("Grenadine Syrup", 0.5),
			b.Use("Orange Wedge", 0),
		}),

		// Rum
		bar.Shake("Pomegranate Daiquiri", []*bar.Drink{
			b.Use("Rum", 1.5),
			b.Use("Lime Juice", 0.75),
			b.Use("Simple Syrup", 0.5),
			b.Use("Pomegranate Juice", 2),
			b.Use("Lime Wedge", 0),
		}),
		bar.Shake("Mojito", []*bar.Drink{
			b.Use("Rum", 1.5),
			b.Use("Mint", 0),
			b.Use("Simple Syrup", 1),
			b.Use("Lime Juice", 0.75),
			b.Use("Soda", 2),
			b.Use("Lime Wedge", 0),
		}),
		bar.Stir("Cuba Libre", []*bar.Drink{
			b.Use("Rum", 1.5),
			b.Use("Coke", 3),
			b.Use("Lime Juice", 0.5),
			b.Use("Lime Wedge", 0),
		}),
		bar.Shake("Pina Colada", []*bar.Drink{
			b.Use("Rum", 1),
			b.Use("Pineapple Juice", 2),
			b.Use("Coconut Cream", 1.5),
			b.Use("Pineapple Wedge", 0),
		}),
		bar.Shake("Cosmopolitan", []*bar.Drink{
			b.Use("Rum", 1),
			b.Use("Cranberry Juice", 1),
			b.Use("Cointreau", 0.5),
			b.Use("Lime Juice", 0.5),
			b.Use("Lime Wedge", 0),
		}),

		// Whiskey
		bar.Shake("Whiskey Sour", []*bar.Drink{
			b.Use("Whiskey", 2),
			b.Use("Lemon Juice", 0.75),
			b.Use("Simple Syrup", 0.5),
			b.Use("Egg White", 0.1),
			b.Use("Orange Wedge", 0),
		}),
		bar.Shake("Manhattan", []*bar.Drink{
			b.Use("Whiskey", 2),
			b.Use("Sweet Vermouth", 1),
			b.Use("Angostura Bitters", 0.1),
			b.Use("Cherry", 0),
		}),
		bar.Shake("Boulevardier", []*bar.Drink{
			b.Use("Whiskey", 1.5),
			b.Use("Sweet Vermouth", 1),
			b.Use("Campari", 1),
			b.Use("Orange Wedge", 0),
		}),
	}

	// Sort in ascending order of alcohol content
	sort.Slice(drinks, func(i, j int) bool {
		return drinks[i].AlcoholContents < drinks[j].AlcoholContents
	})

	drinksJSON, err := json.Marshal(drinks)
	if err != nil {
		log.Fatal(err)
	}

	log.Print(string(drinksJSON))
}
