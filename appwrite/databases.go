package appwrite

import (
	"strings"
)

// Databases service
type Databases struct {
	client Client
}

func NewDatabases(clt Client) *Databases {
	return &Databases{
		client: clt,
	}
}

// ListDocuments get a list of all the user's documents in a given collection.
// You can use the query params to filter your results.
func (srv *Databases) ListDocuments(DatabaseId string, CollectionId string, Queries []interface{}) (*ClientResponse, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId)
	apiPath := r.Replace("/databases/{databaseId}/collections/{collectionId}/documents")

	apiParams := map[string]interface{}{
		"queries": Queries,
	}

	apiHeaders := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", apiPath, apiHeaders, apiParams)
}

// CreateDocument create a new Document. Before using this route, you should
// create a new collection resource using either a [server
// integration](https://appwrite.io/docs/server/databases#databasesCreateCollection)
// API or directly from your database console.
func (srv *Databases) CreateDocument(DatabaseId string, CollectionId string, DocumentId string, Data interface{}, Permissions []interface{}) (*ClientResponse, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId)
	apiPath := r.Replace("/databases/{databaseId}/collections/{collectionId}/documents")

	apiParams := map[string]interface{}{
		"documentId": DocumentId,
		"data": Data,
		"permissions": Permissions,
	}

	apiHeaders := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("POST", apiPath, apiHeaders, apiParams)
}

// GetDocument get a document by its unique ID. This endpoint response returns
// a JSON object with the document data.
func (srv *Databases) GetDocument(DatabaseId string, CollectionId string, DocumentId string, Queries []interface{}) (*ClientResponse, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId, "{documentId}", DocumentId)
	apiPath := r.Replace("/databases/{databaseId}/collections/{collectionId}/documents/{documentId}")

	apiParams := map[string]interface{}{
		"queries": Queries,
	}

	apiHeaders := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", apiPath, apiHeaders, apiParams)
}

// UpdateDocument update a document by its unique ID. Using the patch method
// you can pass only specific fields that will get updated.
func (srv *Databases) UpdateDocument(DatabaseId string, CollectionId string, DocumentId string, Data interface{}, Permissions []interface{}) (*ClientResponse, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId, "{documentId}", DocumentId)
	apiPath := r.Replace("/databases/{databaseId}/collections/{collectionId}/documents/{documentId}")

	apiParams := map[string]interface{}{
		"data": Data,
		"permissions": Permissions,
	}

	apiHeaders := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("PATCH", apiPath, apiHeaders, apiParams)
}

// DeleteDocument delete a document by its unique ID.
func (srv *Databases) DeleteDocument(DatabaseId string, CollectionId string, DocumentId string) (*ClientResponse, error) {
	r := strings.NewReplacer("{databaseId}", DatabaseId, "{collectionId}", CollectionId, "{documentId}", DocumentId)
	apiPath := r.Replace("/databases/{databaseId}/collections/{collectionId}/documents/{documentId}")

	apiParams := map[string]interface{}{
	}

	apiHeaders := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("DELETE", apiPath, apiHeaders, apiParams)
}
