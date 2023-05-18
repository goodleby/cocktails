package bar

import (
	"log"
	"math"
)

// Bar is a bar.
type Bar struct {
	Ingredients []*Ingredient
}

// New creates a new bar.
func New(ingredients []*Ingredient) *Bar {
	return &Bar{Ingredients: ingredients}
}

// Use uses an ingredient.
func (b *Bar) Use(name string, oz float64) *Drink {
	for _, i := range b.Ingredients {
		if i.Name == name {
			i.ConsumptionOz += oz
			i.UseFrequency += 1

			return &Drink{
				Name:            i.Name,
				AlcoholContents: i.AlcoholContents,
				VolumeOz:        oz,
			}
		}
	}

	log.Printf("Ingredient %s not found", name)

	return nil
}

// Ingredient is a basic ingredient with consumption stats.
type Ingredient struct {
	Name            string  `json:"name"`
	AlcoholContents float64 `json:"alcoholContents"`
	ConsumptionOz   float64 `json:"consumptionOz"`
	UseFrequency    int     `json:"useFrequency"`
}

// Drink is a served drink.
type Drink struct {
	Name            string   `json:"name"`
	Method          string   `json:"method,omitempty"`
	AlcoholContents float64  `json:"alcoholContents"`
	VolumeOz        float64  `json:"oz"`
	Recipe          []*Drink `json:"recipe,omitempty"`
}

// Mix mixes drinks together.
func Mix(name, method string, drinks []*Drink) *Drink {
	drink := &Drink{
		Name:   name,
		Method: method,
	}

	var drinkOz, alcoholOz float64

	for _, d := range drinks {
		if d == nil {
			log.Print("Received nil drink")
			continue
		}

		drinkOz += d.VolumeOz
		alcoholOz += d.VolumeOz * d.AlcoholContents
		drink.Recipe = append(drink.Recipe, d)
	}

	drink.VolumeOz = drinkOz
	drink.AlcoholContents = roundTo(alcoholOz/drinkOz, 2)

	return drink
}

// roundTo rounds a number to the specified number of decimals.
func roundTo(num float64, decimals int) float64 {
	return math.Round(num*math.Pow(10, float64(decimals))) / math.Pow(10, float64(decimals))
}
