package appwrite

import (
	"strings"
)

// Account service
type Account struct {
	client Client
}

func NewAccount(clt Client) *Account {
	return &Account{
		client: clt,
	}
}

// Get get the currently logged in user.
func (srv *Account) Get() (*ClientResponse, error) {
	apiPath := "/account"
	
	apiParams := map[string]interface{}{
	}

	apiHeaders := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", apiPath, apiHeaders, apiParams)
}

// Create use this endpoint to allow a new user to register a new account in
// your project. After the user registration completes successfully, you can
// use the
// [/account/verfication](https://appwrite.io/docs/references/cloud/client-web/account#createVerification)
// route to start verifying the user email address. To allow the new user to
// login to their new account, you need to create a new [account
// session](https://appwrite.io/docs/references/cloud/client-web/account#createEmailSession).
func (srv *Account) Create(UserId string, Email string, Password string, Name string) (*ClientResponse, error) {
	apiPath := "/account"
	
	apiParams := map[string]interface{}{
		"userId": UserId,
		"email": Email,
		"password": Password,
		"name": Name,
	}

	apiHeaders := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("POST", apiPath, apiHeaders, apiParams)
}

// UpdateEmail update currently logged in user account email address. After
// changing user address, the user confirmation status will get reset. A new
// confirmation email is not sent automatically however you can use the send
// confirmation email endpoint again to send the confirmation email. For
// security measures, user password is required to complete this request.
// This endpoint can also be used to convert an anonymous account to a normal
// one, by passing an email address and a new password.
// 
func (srv *Account) UpdateEmail(Email string, Password string) (*ClientResponse, error) {
	apiPath := "/account/email"
	
	apiParams := map[string]interface{}{
		"email": Email,
		"password": Password,
	}

	apiHeaders := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("PATCH", apiPath, apiHeaders, apiParams)
}

// ListIdentities get the list of identities for the currently logged in user.
func (srv *Account) ListIdentities(Queries string) (*ClientResponse, error) {
	apiPath := "/account/identities"
	
	apiParams := map[string]interface{}{
		"queries": Queries,
	}

	apiHeaders := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", apiPath, apiHeaders, apiParams)
}

// DeleteIdentity delete an identity by its unique ID.
func (srv *Account) DeleteIdentity(IdentityId string) (*ClientResponse, error) {
	r := strings.NewReplacer("{identityId}", IdentityId)
	apiPath := r.Replace("/account/identities/{identityId}")

	apiParams := map[string]interface{}{
	}

	apiHeaders := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("DELETE", apiPath, apiHeaders, apiParams)
}

// CreateJWT use this endpoint to create a JSON Web Token. You can use the
// resulting JWT to authenticate on behalf of the current user when working
// with the Appwrite server-side API and SDKs. The JWT secret is valid for 15
// minutes from its creation and will be invalid if the user will logout in
// that time frame.
func (srv *Account) CreateJWT() (*ClientResponse, error) {
	apiPath := "/account/jwt"
	
	apiParams := map[string]interface{}{
	}

	apiHeaders := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("POST", apiPath, apiHeaders, apiParams)
}

// ListLogs get the list of latest security activity logs for the currently
// logged in user. Each log returns user IP address, location and date and
// time of log.
func (srv *Account) ListLogs(Queries []interface{}) (*ClientResponse, error) {
	apiPath := "/account/logs"
	
	apiParams := map[string]interface{}{
		"queries": Queries,
	}

	apiHeaders := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", apiPath, apiHeaders, apiParams)
}

// UpdateName update currently logged in user account name.
func (srv *Account) UpdateName(Name string) (*ClientResponse, error) {
	apiPath := "/account/name"
	
	apiParams := map[string]interface{}{
		"name": Name,
	}

	apiHeaders := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("PATCH", apiPath, apiHeaders, apiParams)
}

// UpdatePassword update currently logged in user password. For validation,
// user is required to pass in the new password, and the old password. For
// users created with OAuth, Team Invites and Magic URL, oldPassword is
// optional.
func (srv *Account) UpdatePassword(Password string, OldPassword string) (*ClientResponse, error) {
	apiPath := "/account/password"
	
	apiParams := map[string]interface{}{
		"password": Password,
		"oldPassword": OldPassword,
	}

	apiHeaders := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("PATCH", apiPath, apiHeaders, apiParams)
}

// UpdatePhone update the currently logged in user's phone number. After
// updating the phone number, the phone verification status will be reset. A
// confirmation SMS is not sent automatically, however you can use the [POST
// /account/verification/phone](https://appwrite.io/docs/references/cloud/client-web/account#createPhoneVerification)
// endpoint to send a confirmation SMS.
func (srv *Account) UpdatePhone(Phone string, Password string) (*ClientResponse, error) {
	apiPath := "/account/phone"
	
	apiParams := map[string]interface{}{
		"phone": Phone,
		"password": Password,
	}

	apiHeaders := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("PATCH", apiPath, apiHeaders, apiParams)
}

// GetPrefs get the preferences as a key-value object for the currently logged
// in user.
func (srv *Account) GetPrefs() (*ClientResponse, error) {
	apiPath := "/account/prefs"
	
	apiParams := map[string]interface{}{
	}

	apiHeaders := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", apiPath, apiHeaders, apiParams)
}

// UpdatePrefs update currently logged in user account preferences. The object
// you pass is stored as is, and replaces any previous value. The maximum
// allowed prefs size is 64kB and throws error if exceeded.
func (srv *Account) UpdatePrefs(Prefs interface{}) (*ClientResponse, error) {
	apiPath := "/account/prefs"
	
	apiParams := map[string]interface{}{
		"prefs": Prefs,
	}

	apiHeaders := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("PATCH", apiPath, apiHeaders, apiParams)
}

// CreateRecovery sends the user an email with a temporary secret key for
// password reset. When the user clicks the confirmation link he is redirected
// back to your app password reset URL with the secret key and email address
// values attached to the URL query string. Use the query string params to
// submit a request to the [PUT
// /account/recovery](https://appwrite.io/docs/references/cloud/client-web/account#updateRecovery)
// endpoint to complete the process. The verification link sent to the user's
// email address is valid for 1 hour.
func (srv *Account) CreateRecovery(Email string, Url string) (*ClientResponse, error) {
	apiPath := "/account/recovery"
	
	apiParams := map[string]interface{}{
		"email": Email,
		"url": Url,
	}

	apiHeaders := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("POST", apiPath, apiHeaders, apiParams)
}

// UpdateRecovery use this endpoint to complete the user account password
// reset. Both the **userId** and **secret** arguments will be passed as query
// parameters to the redirect URL you have provided when sending your request
// to the [POST
// /account/recovery](https://appwrite.io/docs/references/cloud/client-web/account#createRecovery)
// endpoint.
// 
// Please note that in order to avoid a [Redirect
// Attack](https://github.com/OWASP/CheatSheetSeries/blob/master/cheatsheets/Unvalidated_Redirects_and_Forwards_Cheat_Sheet.md)
// the only valid redirect URLs are the ones from domains you have set when
// adding your platforms in the console interface.
func (srv *Account) UpdateRecovery(UserId string, Secret string, Password string, PasswordAgain string) (*ClientResponse, error) {
	apiPath := "/account/recovery"
	
	apiParams := map[string]interface{}{
		"userId": UserId,
		"secret": Secret,
		"password": Password,
		"passwordAgain": PasswordAgain,
	}

	apiHeaders := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("PUT", apiPath, apiHeaders, apiParams)
}

// ListSessions get the list of active sessions across different devices for
// the currently logged in user.
func (srv *Account) ListSessions() (*ClientResponse, error) {
	apiPath := "/account/sessions"
	
	apiParams := map[string]interface{}{
	}

	apiHeaders := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", apiPath, apiHeaders, apiParams)
}

// DeleteSessions delete all sessions from the user account and remove any
// sessions cookies from the end client.
func (srv *Account) DeleteSessions() (*ClientResponse, error) {
	apiPath := "/account/sessions"
	
	apiParams := map[string]interface{}{
	}

	apiHeaders := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("DELETE", apiPath, apiHeaders, apiParams)
}

// CreateAnonymousSession use this endpoint to allow a new user to register an
// anonymous account in your project. This route will also create a new
// session for the user. To allow the new user to convert an anonymous account
// to a normal account, you need to update its [email and
// password](https://appwrite.io/docs/references/cloud/client-web/account#updateEmail)
// or create an [OAuth2
// session](https://appwrite.io/docs/references/cloud/client-web/account#CreateOAuth2Session).
func (srv *Account) CreateAnonymousSession() (*ClientResponse, error) {
	apiPath := "/account/sessions/anonymous"
	
	apiParams := map[string]interface{}{
	}

	apiHeaders := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("POST", apiPath, apiHeaders, apiParams)
}

// CreateEmailSession allow the user to login into their account by providing
// a valid email and password combination. This route will create a new
// session for the user.
// 
// A user is limited to 10 active sessions at a time by default. [Learn more
// about session
// limits](https://appwrite.io/docs/authentication-security#limits).
func (srv *Account) CreateEmailSession(Email string, Password string) (*ClientResponse, error) {
	apiPath := "/account/sessions/email"
	
	apiParams := map[string]interface{}{
		"email": Email,
		"password": Password,
	}

	apiHeaders := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("POST", apiPath, apiHeaders, apiParams)
}

// CreateMagicURLSession sends the user an email with a secret key for
// creating a session. If the provided user ID has not been registered, a new
// user will be created. When the user clicks the link in the email, the user
// is redirected back to the URL you provided with the secret key and userId
// values attached to the URL query string. Use the query string parameters to
// submit a request to the [PUT
// /account/sessions/magic-url](https://appwrite.io/docs/references/cloud/client-web/account#updateMagicURLSession)
// endpoint to complete the login process. The link sent to the user's email
// address is valid for 1 hour. If you are on a mobile device you can leave
// the URL parameter empty, so that the login completion will be handled by
// your Appwrite instance by default.
// 
// A user is limited to 10 active sessions at a time by default. [Learn more
// about session
// limits](https://appwrite.io/docs/authentication-security#limits).
// 
func (srv *Account) CreateMagicURLSession(UserId string, Email string, Url string) (*ClientResponse, error) {
	apiPath := "/account/sessions/magic-url"
	
	apiParams := map[string]interface{}{
		"userId": UserId,
		"email": Email,
		"url": Url,
	}

	apiHeaders := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("POST", apiPath, apiHeaders, apiParams)
}

// UpdateMagicURLSession use this endpoint to complete creating the session
// with the Magic URL. Both the **userId** and **secret** arguments will be
// passed as query parameters to the redirect URL you have provided when
// sending your request to the [POST
// /account/sessions/magic-url](https://appwrite.io/docs/references/cloud/client-web/account#createMagicURLSession)
// endpoint.
// 
// Please note that in order to avoid a [Redirect
// Attack](https://github.com/OWASP/CheatSheetSeries/blob/master/cheatsheets/Unvalidated_Redirects_and_Forwards_Cheat_Sheet.md)
// the only valid redirect URLs are the ones from domains you have set when
// adding your platforms in the console interface.
func (srv *Account) UpdateMagicURLSession(UserId string, Secret string) (*ClientResponse, error) {
	apiPath := "/account/sessions/magic-url"
	
	apiParams := map[string]interface{}{
		"userId": UserId,
		"secret": Secret,
	}

	apiHeaders := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("PUT", apiPath, apiHeaders, apiParams)
}

// CreateOAuth2Session allow the user to login to their account using the
// OAuth2 provider of their choice. Each OAuth2 provider should be enabled
// from the Appwrite console first. Use the success and failure arguments to
// provide a redirect URL's back to your app when login is completed.
// 
// If there is already an active session, the new session will be attached to
// the logged-in account. If there are no active sessions, the server will
// attempt to look for a user with the same email address as the email
// received from the OAuth2 provider and attach the new session to the
// existing user. If no matching user is found - the server will create a new
// user.
// 
// A user is limited to 10 active sessions at a time by default. [Learn more
// about session
// limits](https://appwrite.io/docs/authentication-security#limits).
// 
func (srv *Account) CreateOAuth2Session(Provider string, Success string, Failure string, Scopes []interface{}) (*ClientResponse, error) {
	r := strings.NewReplacer("{provider}", Provider)
	apiPath := r.Replace("/account/sessions/oauth2/{provider}")

	apiParams := map[string]interface{}{
		"success": Success,
		"failure": Failure,
		"scopes": Scopes,
	}

	apiHeaders := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", apiPath, apiHeaders, apiParams)
}

// CreatePhoneSession sends the user an SMS with a secret key for creating a
// session. If the provided user ID has not be registered, a new user will be
// created. Use the returned user ID and secret and submit a request to the
// [PUT
// /account/sessions/phone](https://appwrite.io/docs/references/cloud/client-web/account#updatePhoneSession)
// endpoint to complete the login process. The secret sent to the user's phone
// is valid for 15 minutes.
// 
// A user is limited to 10 active sessions at a time by default. [Learn more
// about session
// limits](https://appwrite.io/docs/authentication-security#limits).
func (srv *Account) CreatePhoneSession(UserId string, Phone string) (*ClientResponse, error) {
	apiPath := "/account/sessions/phone"
	
	apiParams := map[string]interface{}{
		"userId": UserId,
		"phone": Phone,
	}

	apiHeaders := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("POST", apiPath, apiHeaders, apiParams)
}

// UpdatePhoneSession use this endpoint to complete creating a session with
// SMS. Use the **userId** from the
// [createPhoneSession](https://appwrite.io/docs/references/cloud/client-web/account#createPhoneSession)
// endpoint and the **secret** received via SMS to successfully update and
// confirm the phone session.
func (srv *Account) UpdatePhoneSession(UserId string, Secret string) (*ClientResponse, error) {
	apiPath := "/account/sessions/phone"
	
	apiParams := map[string]interface{}{
		"userId": UserId,
		"secret": Secret,
	}

	apiHeaders := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("PUT", apiPath, apiHeaders, apiParams)
}

// GetSession use this endpoint to get a logged in user's session using a
// Session ID. Inputting 'current' will return the current session being used.
func (srv *Account) GetSession(SessionId string) (*ClientResponse, error) {
	r := strings.NewReplacer("{sessionId}", SessionId)
	apiPath := r.Replace("/account/sessions/{sessionId}")

	apiParams := map[string]interface{}{
	}

	apiHeaders := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", apiPath, apiHeaders, apiParams)
}

// UpdateSession access tokens have limited lifespan and expire to mitigate
// security risks. If session was created using an OAuth provider, this route
// can be used to "refresh" the access token.
func (srv *Account) UpdateSession(SessionId string) (*ClientResponse, error) {
	r := strings.NewReplacer("{sessionId}", SessionId)
	apiPath := r.Replace("/account/sessions/{sessionId}")

	apiParams := map[string]interface{}{
	}

	apiHeaders := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("PATCH", apiPath, apiHeaders, apiParams)
}

// DeleteSession logout the user. Use 'current' as the session ID to logout on
// this device, use a session ID to logout on another device. If you're
// looking to logout the user on all devices, use [Delete
// Sessions](https://appwrite.io/docs/references/cloud/client-web/account#deleteSessions)
// instead.
func (srv *Account) DeleteSession(SessionId string) (*ClientResponse, error) {
	r := strings.NewReplacer("{sessionId}", SessionId)
	apiPath := r.Replace("/account/sessions/{sessionId}")

	apiParams := map[string]interface{}{
	}

	apiHeaders := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("DELETE", apiPath, apiHeaders, apiParams)
}

// UpdateStatus block the currently logged in user account. Behind the scene,
// the user record is not deleted but permanently blocked from any access. To
// completely delete a user, use the Users API instead.
func (srv *Account) UpdateStatus() (*ClientResponse, error) {
	apiPath := "/account/status"
	
	apiParams := map[string]interface{}{
	}

	apiHeaders := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("PATCH", apiPath, apiHeaders, apiParams)
}

// CreateVerification use this endpoint to send a verification message to your
// user email address to confirm they are the valid owners of that address.
// Both the **userId** and **secret** arguments will be passed as query
// parameters to the URL you have provided to be attached to the verification
// email. The provided URL should redirect the user back to your app and allow
// you to complete the verification process by verifying both the **userId**
// and **secret** parameters. Learn more about how to [complete the
// verification
// process](https://appwrite.io/docs/references/cloud/client-web/account#updateVerification).
// The verification link sent to the user's email address is valid for 7 days.
// 
// Please note that in order to avoid a [Redirect
// Attack](https://github.com/OWASP/CheatSheetSeries/blob/master/cheatsheets/Unvalidated_Redirects_and_Forwards_Cheat_Sheet.md),
// the only valid redirect URLs are the ones from domains you have set when
// adding your platforms in the console interface.
// 
func (srv *Account) CreateVerification(Url string) (*ClientResponse, error) {
	apiPath := "/account/verification"
	
	apiParams := map[string]interface{}{
		"url": Url,
	}

	apiHeaders := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("POST", apiPath, apiHeaders, apiParams)
}

// UpdateVerification use this endpoint to complete the user email
// verification process. Use both the **userId** and **secret** parameters
// that were attached to your app URL to verify the user email ownership. If
// confirmed this route will return a 200 status code.
func (srv *Account) UpdateVerification(UserId string, Secret string) (*ClientResponse, error) {
	apiPath := "/account/verification"
	
	apiParams := map[string]interface{}{
		"userId": UserId,
		"secret": Secret,
	}

	apiHeaders := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("PUT", apiPath, apiHeaders, apiParams)
}

// CreatePhoneVerification use this endpoint to send a verification SMS to the
// currently logged in user. This endpoint is meant for use after updating a
// user's phone number using the
// [accountUpdatePhone](https://appwrite.io/docs/references/cloud/client-web/account#updatePhone)
// endpoint. Learn more about how to [complete the verification
// process](https://appwrite.io/docs/references/cloud/client-web/account#updatePhoneVerification).
// The verification code sent to the user's phone number is valid for 15
// minutes.
func (srv *Account) CreatePhoneVerification() (*ClientResponse, error) {
	apiPath := "/account/verification/phone"
	
	apiParams := map[string]interface{}{
	}

	apiHeaders := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("POST", apiPath, apiHeaders, apiParams)
}

// UpdatePhoneVerification use this endpoint to complete the user phone
// verification process. Use the **userId** and **secret** that were sent to
// your user's phone number to verify the user email ownership. If confirmed
// this route will return a 200 status code.
func (srv *Account) UpdatePhoneVerification(UserId string, Secret string) (*ClientResponse, error) {
	apiPath := "/account/verification/phone"
	
	apiParams := map[string]interface{}{
		"userId": UserId,
		"secret": Secret,
	}

	apiHeaders := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("PUT", apiPath, apiHeaders, apiParams)
}
