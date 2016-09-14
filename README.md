## 🍴openmensa

This is a golang wrapper for the [OpenMensa](https://openmensa.org) API. Very WIP.

Use it something like this:

```go
import (
	"fmt"
  
  	"github.com/kiliankoe/openmensa"
)

func main() {
  canteen, _ := openmensa.GetCanteen(79)
  meals, _ := canteen.Today()
  for _, meal := range meals {
    fmt.Println(meal)
  }
}

// Hähnchenkeule mit Geflügelsoße, dazu buntes Gemüse und Petersilienkartoffeln
// Sojageschnetzeltes mit Waldpilzen, dazu Apfelrotkohl und Klöße
// Gefülltes Pizzabrötchen
// ...
```

