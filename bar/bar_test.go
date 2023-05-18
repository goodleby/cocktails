package bar

import (
	"reflect"
	"testing"
)

func TestMix(t *testing.T) {
	type args struct {
		name   string
		method string
		drinks []*Drink
	}
	tests := []struct {
		name string
		args args
		want *Drink
	}{
		{
			name: "Should mix drinks",
			args: args{
				name:   "Cosmopolitan",
				method: "Shake",
				drinks: []*Drink{
					{Name: "Vodka", AlcoholContents: 0.4, VolumeOz: 1},
					{Name: "Cointreau", AlcoholContents: 0.4, VolumeOz: 0.5},
					{Name: "Lime Juice", AlcoholContents: 0, VolumeOz: 0.5},
					{Name: "Cranberry Juice", AlcoholContents: 0, VolumeOz: 1},
				},
			},
			want: &Drink{
				Name:            "Cosmopolitan",
				Method:          "Shake",
				VolumeOz:        3,
				AlcoholContents: 0.2,
				Recipe: []*Drink{
					{Name: "Vodka", AlcoholContents: 0.4, VolumeOz: 1},
					{Name: "Cointreau", AlcoholContents: 0.4, VolumeOz: 0.5},
					{Name: "Lime Juice", AlcoholContents: 0, VolumeOz: 0.5},
					{Name: "Cranberry Juice", AlcoholContents: 0, VolumeOz: 1},
				},
			},
		},
		{
			name: "Should mix drinks and ignore nil drinks",
			args: args{
				name:   "Cosmopolitan",
				method: "Shake",
				drinks: []*Drink{
					nil,
					{Name: "Vodka", AlcoholContents: 0.4, VolumeOz: 1},
					{Name: "Cointreau", AlcoholContents: 0.4, VolumeOz: 0.5},
					nil,
					{Name: "Lime Juice", AlcoholContents: 0, VolumeOz: 0.5},
					{Name: "Cranberry Juice", AlcoholContents: 0, VolumeOz: 1},
					nil,
				},
			},
			want: &Drink{
				Name:            "Cosmopolitan",
				Method:          "Shake",
				VolumeOz:        3,
				AlcoholContents: 0.2,
				Recipe: []*Drink{
					{Name: "Vodka", AlcoholContents: 0.4, VolumeOz: 1},
					{Name: "Cointreau", AlcoholContents: 0.4, VolumeOz: 0.5},
					{Name: "Lime Juice", AlcoholContents: 0, VolumeOz: 0.5},
					{Name: "Cranberry Juice", AlcoholContents: 0, VolumeOz: 1},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Mix(tt.args.name, tt.args.method, tt.args.drinks); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Mix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_roundTo(t *testing.T) {
	type args struct {
		num      float64
		decimals int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "should trim to the amount of decimals, if rounding number 4 or lower",
			args: args{
				num:      3.721,
				decimals: 2,
			},
			want: 3.72,
		},
		{
			name: "should round up, if rounding number 5 or higher",
			args: args{
				num:      3.725,
				decimals: 2,
			},
			want: 3.73,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := roundTo(tt.args.num, tt.args.decimals); got != tt.want {
				t.Errorf("roundTo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBar_Use(t *testing.T) {
	type fields struct {
		Ingredients []*Ingredient
	}
	type args struct {
		name string
		oz   float64
	}
	tests := []struct {
		name              string
		fields            fields
		args              args
		want              *Drink
		wantUpdatedFields fields
	}{
		{
			name: "Should return drink from ingredients and record consumption",
			fields: fields{
				Ingredients: []*Ingredient{
					{Name: "Vodka", AlcoholContents: 0.4, ConsumptionOz: 0, UseFrequency: 0},
					{Name: "Cranberry Juice", AlcoholContents: 0, ConsumptionOz: 0, UseFrequency: 0},
					{Name: "Cointreau", AlcoholContents: 0.4, ConsumptionOz: 0, UseFrequency: 0},
					{Name: "Lime Juice", AlcoholContents: 0, ConsumptionOz: 0, UseFrequency: 0},
				},
			},
			args: args{
				name: "Vodka",
				oz:   3,
			},
			want: &Drink{
				Name:            "Vodka",
				VolumeOz:        3,
				AlcoholContents: 0.4,
			},
			wantUpdatedFields: fields{
				Ingredients: []*Ingredient{
					{Name: "Vodka", AlcoholContents: 0.4, ConsumptionOz: 3, UseFrequency: 1},
					{Name: "Cranberry Juice", AlcoholContents: 0, ConsumptionOz: 0, UseFrequency: 0},
					{Name: "Cointreau", AlcoholContents: 0.4, ConsumptionOz: 0, UseFrequency: 0},
					{Name: "Lime Juice", AlcoholContents: 0, ConsumptionOz: 0, UseFrequency: 0},
				},
			},
		},
		{
			name: "Should return nil if ingredient not found",
			fields: fields{
				Ingredients: []*Ingredient{
					{Name: "Vodka", AlcoholContents: 0.4, ConsumptionOz: 0, UseFrequency: 0},
					{Name: "Cranberry Juice", AlcoholContents: 0, ConsumptionOz: 0, UseFrequency: 0},
					{Name: "Cointreau", AlcoholContents: 0.4, ConsumptionOz: 0, UseFrequency: 0},
					{Name: "Lime Juice", AlcoholContents: 0, ConsumptionOz: 0, UseFrequency: 0},
				},
			},
			args: args{
				name: "Tequila",
				oz:   3,
			},
			want: nil,
			wantUpdatedFields: fields{
				Ingredients: []*Ingredient{
					{Name: "Vodka", AlcoholContents: 0.4, ConsumptionOz: 0, UseFrequency: 0},
					{Name: "Cranberry Juice", AlcoholContents: 0, ConsumptionOz: 0, UseFrequency: 0},
					{Name: "Cointreau", AlcoholContents: 0.4, ConsumptionOz: 0, UseFrequency: 0},
					{Name: "Lime Juice", AlcoholContents: 0, ConsumptionOz: 0, UseFrequency: 0},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Bar{
				Ingredients: tt.fields.Ingredients,
			}
			if got := b.Use(tt.args.name, tt.args.oz); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Bar.Use() = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(tt.fields, tt.wantUpdatedFields) {
				t.Errorf("Bar.Use() = %v, want updated fields %v", tt.fields, tt.wantUpdatedFields)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		ingredients []*Ingredient
	}
	tests := []struct {
		name string
		args args
		want *Bar
	}{
		{
			name: "Should create a new bar",
			args: args{
				ingredients: []*Ingredient{
					{Name: "Vodka", AlcoholContents: 0.4},
					{Name: "Cranberry Juice", AlcoholContents: 0},
					{Name: "Cointreau", AlcoholContents: 0.4},
					{Name: "Lime Juice", AlcoholContents: 0},
				},
			},
			want: &Bar{
				Ingredients: []*Ingredient{
					{Name: "Vodka", AlcoholContents: 0.4, ConsumptionOz: 0, UseFrequency: 0},
					{Name: "Cranberry Juice", AlcoholContents: 0, ConsumptionOz: 0, UseFrequency: 0},
					{Name: "Cointreau", AlcoholContents: 0.4, ConsumptionOz: 0, UseFrequency: 0},
					{Name: "Lime Juice", AlcoholContents: 0, ConsumptionOz: 0, UseFrequency: 0},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.ingredients); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
