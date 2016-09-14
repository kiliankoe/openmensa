package openmensa

// A Meal object
type Meal struct {
	ID       int         `json:"id"`
	Name     string      `json:"name"`
	Category string      `json:"category"`
	Prices   *MealPrices `json:"prices"`
	Notes    []string    `json:"notes"`
}

func (m Meal) String() string {
	return m.Name
}

// A MealPrices object
type MealPrices struct {
	Students  *float32 `json:"students"`
	Employees *float32 `json:"employees"`
	Pupils    *float32 `json:"pupils"`
	Others    *float32 `json:"others"`
}
