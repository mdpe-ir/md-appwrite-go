package appwrite

import (
)

// Graphql service
type Graphql struct {
	client Client
}

func NewGraphql(clt Client) *Graphql {
	return &Graphql{
		client: clt,
	}
}

// Query execute a GraphQL mutation.
func (srv *Graphql) Query(Query interface{}) (*ClientResponse, error) {
	apiPath := "/graphql"
	
	apiParams := map[string]interface{}{
		"query": Query,
	}

	apiHeaders := map[string]interface{}{
		"x-sdk-graphql": "true",
		"content-type": "application/json",
	}
	return srv.client.Call("POST", apiPath, apiHeaders, apiParams)
}

// Mutation execute a GraphQL mutation.
func (srv *Graphql) Mutation(Query interface{}) (*ClientResponse, error) {
	apiPath := "/graphql/mutation"
	
	apiParams := map[string]interface{}{
		"query": Query,
	}

	apiHeaders := map[string]interface{}{
		"x-sdk-graphql": "true",
		"content-type": "application/json",
	}
	return srv.client.Call("POST", apiPath, apiHeaders, apiParams)
}
