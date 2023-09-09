package main

import (
	"fmt"
	"net/http"

	"github.com/goccy/go-json"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/testutil"
)

func main() {
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("query")
		result := graphql.Do(graphql.Params{
			Schema:        testutil.StarWarsSchema,
			RequestString: query,
		})
		json.NewEncoder(w).Encode(result)
	})
	fmt.Println("Now server is running on port 8080")
	fmt.Println("Test with Get      : curl -g 'http://localhost:8080/graphql?query={hero{name}}'")
	http.ListenAndServe(":8080", nil)
}
