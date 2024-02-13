package appwrite

import (
	"strings"
)

// Storage service
type Storage struct {
	client Client
}

func NewStorage(clt Client) *Storage {
	return &Storage{
		client: clt,
	}
}

// ListFiles get a list of all the user files. You can use the query params to
// filter your results.
func (srv *Storage) ListFiles(BucketId string, Queries []interface{}, Search string) (*ClientResponse, error) {
	r := strings.NewReplacer("{bucketId}", BucketId)
	apiPath := r.Replace("/storage/buckets/{bucketId}/files")

	apiParams := map[string]interface{}{
		"queries": Queries,
		"search": Search,
	}

	apiHeaders := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", apiPath, apiHeaders, apiParams)
}

// CreateFile create a new file. Before using this route, you should create a
// new bucket resource using either a [server
// integration](https://appwrite.io/docs/server/storage#storageCreateBucket)
// API or directly from your Appwrite console.
// 
// Larger files should be uploaded using multiple requests with the
// [content-range](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Range)
// header to send a partial request with a maximum supported chunk of `5MB`.
// The `content-range` header values should always be in bytes.
// 
// When the first request is sent, the server will return the **File** object,
// and the subsequent part request must include the file's **id** in
// `x-appwrite-id` header to allow the server to know that the partial upload
// is for the existing file and not for a new one.
// 
// If you're creating a new file using one of the Appwrite SDKs, all the
// chunking logic will be managed by the SDK internally.
// 
func (srv *Storage) CreateFile(BucketId string, FileId string, File string, Permissions []interface{}) (*ClientResponse, error) {
	r := strings.NewReplacer("{bucketId}", BucketId)
	apiPath := r.Replace("/storage/buckets/{bucketId}/files")

	apiParams := map[string]interface{}{
		"fileId": FileId,
		"file": File,
		"permissions": Permissions,
	}

	apiHeaders := map[string]interface{}{
		"content-type": "multipart/form-data",
	}
	return srv.client.Call("POST", apiPath, apiHeaders, apiParams)
}

// GetFile get a file by its unique ID. This endpoint response returns a JSON
// object with the file metadata.
func (srv *Storage) GetFile(BucketId string, FileId string) (*ClientResponse, error) {
	r := strings.NewReplacer("{bucketId}", BucketId, "{fileId}", FileId)
	apiPath := r.Replace("/storage/buckets/{bucketId}/files/{fileId}")

	apiParams := map[string]interface{}{
	}

	apiHeaders := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", apiPath, apiHeaders, apiParams)
}

// UpdateFile update a file by its unique ID. Only users with write
// permissions have access to update this resource.
func (srv *Storage) UpdateFile(BucketId string, FileId string, Name string, Permissions []interface{}) (*ClientResponse, error) {
	r := strings.NewReplacer("{bucketId}", BucketId, "{fileId}", FileId)
	apiPath := r.Replace("/storage/buckets/{bucketId}/files/{fileId}")

	apiParams := map[string]interface{}{
		"name": Name,
		"permissions": Permissions,
	}

	apiHeaders := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("PUT", apiPath, apiHeaders, apiParams)
}

// DeleteFile delete a file by its unique ID. Only users with write
// permissions have access to delete this resource.
func (srv *Storage) DeleteFile(BucketId string, FileId string) (*ClientResponse, error) {
	r := strings.NewReplacer("{bucketId}", BucketId, "{fileId}", FileId)
	apiPath := r.Replace("/storage/buckets/{bucketId}/files/{fileId}")

	apiParams := map[string]interface{}{
	}

	apiHeaders := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("DELETE", apiPath, apiHeaders, apiParams)
}

// GetFileDownload get a file content by its unique ID. The endpoint response
// return with a 'Content-Disposition: attachment' header that tells the
// browser to start downloading the file to user downloads directory.
func (srv *Storage) GetFileDownload(BucketId string, FileId string) (*ClientResponse, error) {
	r := strings.NewReplacer("{bucketId}", BucketId, "{fileId}", FileId)
	apiPath := r.Replace("/storage/buckets/{bucketId}/files/{fileId}/download")

	apiParams := map[string]interface{}{
	}

	apiHeaders := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", apiPath, apiHeaders, apiParams)
}

// GetFilePreview get a file preview image. Currently, this method supports
// preview for image files (jpg, png, and gif), other supported formats, like
// pdf, docs, slides, and spreadsheets, will return the file icon image. You
// can also pass query string arguments for cutting and resizing your preview
// image. Preview is supported only for image files smaller than 10MB.
func (srv *Storage) GetFilePreview(BucketId string, FileId string, Width int, Height int, Gravity string, Quality int, BorderWidth int, BorderColor string, BorderRadius int, Opacity float64, Rotation int, Background string, Output string) (*ClientResponse, error) {
	r := strings.NewReplacer("{bucketId}", BucketId, "{fileId}", FileId)
	apiPath := r.Replace("/storage/buckets/{bucketId}/files/{fileId}/preview")

	apiParams := map[string]interface{}{
		"width": Width,
		"height": Height,
		"gravity": Gravity,
		"quality": Quality,
		"borderWidth": BorderWidth,
		"borderColor": BorderColor,
		"borderRadius": BorderRadius,
		"opacity": Opacity,
		"rotation": Rotation,
		"background": Background,
		"output": Output,
	}

	apiHeaders := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", apiPath, apiHeaders, apiParams)
}

// GetFileView get a file content by its unique ID. This endpoint is similar
// to the download method but returns with no  'Content-Disposition:
// attachment' header.
func (srv *Storage) GetFileView(BucketId string, FileId string) (*ClientResponse, error) {
	r := strings.NewReplacer("{bucketId}", BucketId, "{fileId}", FileId)
	apiPath := r.Replace("/storage/buckets/{bucketId}/files/{fileId}/view")

	apiParams := map[string]interface{}{
	}

	apiHeaders := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", apiPath, apiHeaders, apiParams)
}
