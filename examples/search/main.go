package main

import (
	"context"
	"fmt"

	"github.com/restayway/osm"
)

func main() {
	ctx := context.Background()
	results, err := osm.Search(ctx, "Central Park")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	for _, result := range results {
		fmt.Println("Search Result:", result)
	}
}
