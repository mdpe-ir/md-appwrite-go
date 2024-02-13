package appwrite

import (
	"strings"
)

// Avatars service
type Avatars struct {
	client Client
}

func NewAvatars(clt Client) *Avatars {
	return &Avatars{
		client: clt,
	}
}

// GetBrowser you can use this endpoint to show different browser icons to
// your users. The code argument receives the browser code as it appears in
// your user [GET
// /account/sessions](https://appwrite.io/docs/references/cloud/client-web/account#getSessions)
// endpoint. Use width, height and quality arguments to change the output
// settings.
// 
// When one dimension is specified and the other is 0, the image is scaled
// with preserved aspect ratio. If both dimensions are 0, the API provides an
// image at source quality. If dimensions are not specified, the default size
// of image returned is 100x100px.
func (srv *Avatars) GetBrowser(Code string, Width int, Height int, Quality int) (*ClientResponse, error) {
	r := strings.NewReplacer("{code}", Code)
	apiPath := r.Replace("/avatars/browsers/{code}")

	apiParams := map[string]interface{}{
		"width": Width,
		"height": Height,
		"quality": Quality,
	}

	apiHeaders := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", apiPath, apiHeaders, apiParams)
}

// GetCreditCard the credit card endpoint will return you the icon of the
// credit card provider you need. Use width, height and quality arguments to
// change the output settings.
// 
// When one dimension is specified and the other is 0, the image is scaled
// with preserved aspect ratio. If both dimensions are 0, the API provides an
// image at source quality. If dimensions are not specified, the default size
// of image returned is 100x100px.
// 
func (srv *Avatars) GetCreditCard(Code string, Width int, Height int, Quality int) (*ClientResponse, error) {
	r := strings.NewReplacer("{code}", Code)
	apiPath := r.Replace("/avatars/credit-cards/{code}")

	apiParams := map[string]interface{}{
		"width": Width,
		"height": Height,
		"quality": Quality,
	}

	apiHeaders := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", apiPath, apiHeaders, apiParams)
}

// GetFavicon use this endpoint to fetch the favorite icon (AKA favicon) of
// any remote website URL.
// 
func (srv *Avatars) GetFavicon(Url string) (*ClientResponse, error) {
	apiPath := "/avatars/favicon"
	
	apiParams := map[string]interface{}{
		"url": Url,
	}

	apiHeaders := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", apiPath, apiHeaders, apiParams)
}

// GetFlag you can use this endpoint to show different country flags icons to
// your users. The code argument receives the 2 letter country code. Use
// width, height and quality arguments to change the output settings. Country
// codes follow the [ISO 3166-1](https://en.wikipedia.org/wiki/ISO_3166-1)
// standard.
// 
// When one dimension is specified and the other is 0, the image is scaled
// with preserved aspect ratio. If both dimensions are 0, the API provides an
// image at source quality. If dimensions are not specified, the default size
// of image returned is 100x100px.
// 
func (srv *Avatars) GetFlag(Code string, Width int, Height int, Quality int) (*ClientResponse, error) {
	r := strings.NewReplacer("{code}", Code)
	apiPath := r.Replace("/avatars/flags/{code}")

	apiParams := map[string]interface{}{
		"width": Width,
		"height": Height,
		"quality": Quality,
	}

	apiHeaders := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", apiPath, apiHeaders, apiParams)
}

// GetImage use this endpoint to fetch a remote image URL and crop it to any
// image size you want. This endpoint is very useful if you need to crop and
// display remote images in your app or in case you want to make sure a 3rd
// party image is properly served using a TLS protocol.
// 
// When one dimension is specified and the other is 0, the image is scaled
// with preserved aspect ratio. If both dimensions are 0, the API provides an
// image at source quality. If dimensions are not specified, the default size
// of image returned is 400x400px.
// 
func (srv *Avatars) GetImage(Url string, Width int, Height int) (*ClientResponse, error) {
	apiPath := "/avatars/image"
	
	apiParams := map[string]interface{}{
		"url": Url,
		"width": Width,
		"height": Height,
	}

	apiHeaders := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", apiPath, apiHeaders, apiParams)
}

// GetInitials use this endpoint to show your user initials avatar icon on
// your website or app. By default, this route will try to print your
// logged-in user name or email initials. You can also overwrite the user name
// if you pass the 'name' parameter. If no name is given and no user is
// logged, an empty avatar will be returned.
// 
// You can use the color and background params to change the avatar colors. By
// default, a random theme will be selected. The random theme will persist for
// the user's initials when reloading the same theme will always return for
// the same initials.
// 
// When one dimension is specified and the other is 0, the image is scaled
// with preserved aspect ratio. If both dimensions are 0, the API provides an
// image at source quality. If dimensions are not specified, the default size
// of image returned is 100x100px.
// 
func (srv *Avatars) GetInitials(Name string, Width int, Height int, Background string) (*ClientResponse, error) {
	apiPath := "/avatars/initials"
	
	apiParams := map[string]interface{}{
		"name": Name,
		"width": Width,
		"height": Height,
		"background": Background,
	}

	apiHeaders := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", apiPath, apiHeaders, apiParams)
}

// GetQR converts a given plain text to a QR code image. You can use the query
// parameters to change the size and style of the resulting image.
// 
func (srv *Avatars) GetQR(Text string, Size int, Margin int, Download bool) (*ClientResponse, error) {
	apiPath := "/avatars/qr"
	
	apiParams := map[string]interface{}{
		"text": Text,
		"size": Size,
		"margin": Margin,
		"download": Download,
	}

	apiHeaders := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", apiPath, apiHeaders, apiParams)
}
