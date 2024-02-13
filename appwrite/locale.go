package appwrite

import (
)

// Locale service
type Locale struct {
	client Client
}

func NewLocale(clt Client) *Locale {
	return &Locale{
		client: clt,
	}
}

// Get get the current user location based on IP. Returns an object with user
// country code, country name, continent name, continent code, ip address and
// suggested currency. You can use the locale header to get the data in a
// supported language.
// 
// ([IP Geolocation by DB-IP](https://db-ip.com))
func (srv *Locale) Get() (*ClientResponse, error) {
	apiPath := "/locale"
	
	apiParams := map[string]interface{}{
	}

	apiHeaders := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", apiPath, apiHeaders, apiParams)
}

// ListCodes list of all locale codes in [ISO
// 639-1](https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes).
func (srv *Locale) ListCodes() (*ClientResponse, error) {
	apiPath := "/locale/codes"
	
	apiParams := map[string]interface{}{
	}

	apiHeaders := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", apiPath, apiHeaders, apiParams)
}

// ListContinents list of all continents. You can use the locale header to get
// the data in a supported language.
func (srv *Locale) ListContinents() (*ClientResponse, error) {
	apiPath := "/locale/continents"
	
	apiParams := map[string]interface{}{
	}

	apiHeaders := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", apiPath, apiHeaders, apiParams)
}

// ListCountries list of all countries. You can use the locale header to get
// the data in a supported language.
func (srv *Locale) ListCountries() (*ClientResponse, error) {
	apiPath := "/locale/countries"
	
	apiParams := map[string]interface{}{
	}

	apiHeaders := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", apiPath, apiHeaders, apiParams)
}

// ListCountriesEU list of all countries that are currently members of the EU.
// You can use the locale header to get the data in a supported language.
func (srv *Locale) ListCountriesEU() (*ClientResponse, error) {
	apiPath := "/locale/countries/eu"
	
	apiParams := map[string]interface{}{
	}

	apiHeaders := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", apiPath, apiHeaders, apiParams)
}

// ListCountriesPhones list of all countries phone codes. You can use the
// locale header to get the data in a supported language.
func (srv *Locale) ListCountriesPhones() (*ClientResponse, error) {
	apiPath := "/locale/countries/phones"
	
	apiParams := map[string]interface{}{
	}

	apiHeaders := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", apiPath, apiHeaders, apiParams)
}

// ListCurrencies list of all currencies, including currency symbol, name,
// plural, and decimal digits for all major and minor currencies. You can use
// the locale header to get the data in a supported language.
func (srv *Locale) ListCurrencies() (*ClientResponse, error) {
	apiPath := "/locale/currencies"
	
	apiParams := map[string]interface{}{
	}

	apiHeaders := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", apiPath, apiHeaders, apiParams)
}

// ListLanguages list of all languages classified by ISO 639-1 including
// 2-letter code, name in English, and name in the respective language.
func (srv *Locale) ListLanguages() (*ClientResponse, error) {
	apiPath := "/locale/languages"
	
	apiParams := map[string]interface{}{
	}

	apiHeaders := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", apiPath, apiHeaders, apiParams)
}
