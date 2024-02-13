package appwrite

import (
	"strings"
)

// Functions service
type Functions struct {
	client Client
}

func NewFunctions(clt Client) *Functions {
	return &Functions{
		client: clt,
	}
}

// ListExecutions get a list of all the current user function execution logs.
// You can use the query params to filter your results.
func (srv *Functions) ListExecutions(FunctionId string, Queries []interface{}, Search string) (*ClientResponse, error) {
	r := strings.NewReplacer("{functionId}", FunctionId)
	apiPath := r.Replace("/functions/{functionId}/executions")

	apiParams := map[string]interface{}{
		"queries": Queries,
		"search": Search,
	}

	apiHeaders := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", apiPath, apiHeaders, apiParams)
}

// CreateExecution trigger a function execution. The returned object will
// return you the current execution status. You can ping the `Get Execution`
// endpoint to get updates on the current execution status. Once this endpoint
// is called, your function execution process will start asynchronously.
func (srv *Functions) CreateExecution(FunctionId string, Body string, Async bool, Path string, Method string, Headers interface{}) (*ClientResponse, error) {
	r := strings.NewReplacer("{functionId}", FunctionId)
	apiPath := r.Replace("/functions/{functionId}/executions")

	apiParams := map[string]interface{}{
		"body": Body,
		"async": Async,
		"path": Path,
		"method": Method,
		"headers": Headers,
	}

	apiHeaders := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("POST", apiPath, apiHeaders, apiParams)
}

// GetExecution get a function execution log by its unique ID.
func (srv *Functions) GetExecution(FunctionId string, ExecutionId string) (*ClientResponse, error) {
	r := strings.NewReplacer("{functionId}", FunctionId, "{executionId}", ExecutionId)
	apiPath := r.Replace("/functions/{functionId}/executions/{executionId}")

	apiParams := map[string]interface{}{
	}

	apiHeaders := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", apiPath, apiHeaders, apiParams)
}
