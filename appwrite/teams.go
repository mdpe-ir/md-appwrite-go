package appwrite

import (
	"strings"
)

// Teams service
type Teams struct {
	client Client
}

func NewTeams(clt Client) *Teams {
	return &Teams{
		client: clt,
	}
}

// List get a list of all the teams in which the current user is a member. You
// can use the parameters to filter your results.
func (srv *Teams) List(Queries []interface{}, Search string) (*ClientResponse, error) {
	apiPath := "/teams"
	
	apiParams := map[string]interface{}{
		"queries": Queries,
		"search": Search,
	}

	apiHeaders := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", apiPath, apiHeaders, apiParams)
}

// Create create a new team. The user who creates the team will automatically
// be assigned as the owner of the team. Only the users with the owner role
// can invite new members, add new owners and delete or update the team.
func (srv *Teams) Create(TeamId string, Name string, Roles []interface{}) (*ClientResponse, error) {
	apiPath := "/teams"
	
	apiParams := map[string]interface{}{
		"teamId": TeamId,
		"name": Name,
		"roles": Roles,
	}

	apiHeaders := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("POST", apiPath, apiHeaders, apiParams)
}

// Get get a team by its ID. All team members have read access for this
// resource.
func (srv *Teams) Get(TeamId string) (*ClientResponse, error) {
	r := strings.NewReplacer("{teamId}", TeamId)
	apiPath := r.Replace("/teams/{teamId}")

	apiParams := map[string]interface{}{
	}

	apiHeaders := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", apiPath, apiHeaders, apiParams)
}

// UpdateName update the team's name by its unique ID.
func (srv *Teams) UpdateName(TeamId string, Name string) (*ClientResponse, error) {
	r := strings.NewReplacer("{teamId}", TeamId)
	apiPath := r.Replace("/teams/{teamId}")

	apiParams := map[string]interface{}{
		"name": Name,
	}

	apiHeaders := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("PUT", apiPath, apiHeaders, apiParams)
}

// Delete delete a team using its ID. Only team members with the owner role
// can delete the team.
func (srv *Teams) Delete(TeamId string) (*ClientResponse, error) {
	r := strings.NewReplacer("{teamId}", TeamId)
	apiPath := r.Replace("/teams/{teamId}")

	apiParams := map[string]interface{}{
	}

	apiHeaders := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("DELETE", apiPath, apiHeaders, apiParams)
}

// ListMemberships use this endpoint to list a team's members using the team's
// ID. All team members have read access to this endpoint.
func (srv *Teams) ListMemberships(TeamId string, Queries []interface{}, Search string) (*ClientResponse, error) {
	r := strings.NewReplacer("{teamId}", TeamId)
	apiPath := r.Replace("/teams/{teamId}/memberships")

	apiParams := map[string]interface{}{
		"queries": Queries,
		"search": Search,
	}

	apiHeaders := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", apiPath, apiHeaders, apiParams)
}

// CreateMembership invite a new member to join your team. Provide an ID for
// existing users, or invite unregistered users using an email or phone
// number. If initiated from a Client SDK, Appwrite will send an email or sms
// with a link to join the team to the invited user, and an account will be
// created for them if one doesn't exist. If initiated from a Server SDK, the
// new member will be added automatically to the team.
// 
// You only need to provide one of a user ID, email, or phone number. Appwrite
// will prioritize accepting the user ID > email > phone number if you provide
// more than one of these parameters.
// 
// Use the `url` parameter to redirect the user from the invitation email to
// your app. After the user is redirected, use the [Update Team Membership
// Status](https://appwrite.io/docs/references/cloud/client-web/teams#updateMembershipStatus)
// endpoint to allow the user to accept the invitation to the team. 
// 
// Please note that to avoid a [Redirect
// Attack](https://github.com/OWASP/CheatSheetSeries/blob/master/cheatsheets/Unvalidated_Redirects_and_Forwards_Cheat_Sheet.md)
// Appwrite will accept the only redirect URLs under the domains you have
// added as a platform on the Appwrite Console.
// 
func (srv *Teams) CreateMembership(TeamId string, Roles []interface{}, Email string, UserId string, Phone string, Url string, Name string) (*ClientResponse, error) {
	r := strings.NewReplacer("{teamId}", TeamId)
	apiPath := r.Replace("/teams/{teamId}/memberships")

	apiParams := map[string]interface{}{
		"email": Email,
		"userId": UserId,
		"phone": Phone,
		"roles": Roles,
		"url": Url,
		"name": Name,
	}

	apiHeaders := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("POST", apiPath, apiHeaders, apiParams)
}

// GetMembership get a team member by the membership unique id. All team
// members have read access for this resource.
func (srv *Teams) GetMembership(TeamId string, MembershipId string) (*ClientResponse, error) {
	r := strings.NewReplacer("{teamId}", TeamId, "{membershipId}", MembershipId)
	apiPath := r.Replace("/teams/{teamId}/memberships/{membershipId}")

	apiParams := map[string]interface{}{
	}

	apiHeaders := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", apiPath, apiHeaders, apiParams)
}

// UpdateMembership modify the roles of a team member. Only team members with
// the owner role have access to this endpoint. Learn more about [roles and
// permissions](https://appwrite.io/docs/permissions).
// 
func (srv *Teams) UpdateMembership(TeamId string, MembershipId string, Roles []interface{}) (*ClientResponse, error) {
	r := strings.NewReplacer("{teamId}", TeamId, "{membershipId}", MembershipId)
	apiPath := r.Replace("/teams/{teamId}/memberships/{membershipId}")

	apiParams := map[string]interface{}{
		"roles": Roles,
	}

	apiHeaders := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("PATCH", apiPath, apiHeaders, apiParams)
}

// DeleteMembership this endpoint allows a user to leave a team or for a team
// owner to delete the membership of any other team member. You can also use
// this endpoint to delete a user membership even if it is not accepted.
func (srv *Teams) DeleteMembership(TeamId string, MembershipId string) (*ClientResponse, error) {
	r := strings.NewReplacer("{teamId}", TeamId, "{membershipId}", MembershipId)
	apiPath := r.Replace("/teams/{teamId}/memberships/{membershipId}")

	apiParams := map[string]interface{}{
	}

	apiHeaders := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("DELETE", apiPath, apiHeaders, apiParams)
}

// UpdateMembershipStatus use this endpoint to allow a user to accept an
// invitation to join a team after being redirected back to your app from the
// invitation email received by the user.
// 
// If the request is successful, a session for the user is automatically
// created.
// 
func (srv *Teams) UpdateMembershipStatus(TeamId string, MembershipId string, UserId string, Secret string) (*ClientResponse, error) {
	r := strings.NewReplacer("{teamId}", TeamId, "{membershipId}", MembershipId)
	apiPath := r.Replace("/teams/{teamId}/memberships/{membershipId}/status")

	apiParams := map[string]interface{}{
		"userId": UserId,
		"secret": Secret,
	}

	apiHeaders := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("PATCH", apiPath, apiHeaders, apiParams)
}

// GetPrefs get the team's shared preferences by its unique ID. If a
// preference doesn't need to be shared by all team members, prefer storing
// them in [user
// preferences](https://appwrite.io/docs/references/cloud/client-web/account#getPrefs).
func (srv *Teams) GetPrefs(TeamId string) (*ClientResponse, error) {
	r := strings.NewReplacer("{teamId}", TeamId)
	apiPath := r.Replace("/teams/{teamId}/prefs")

	apiParams := map[string]interface{}{
	}

	apiHeaders := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("GET", apiPath, apiHeaders, apiParams)
}

// UpdatePrefs update the team's preferences by its unique ID. The object you
// pass is stored as is and replaces any previous value. The maximum allowed
// prefs size is 64kB and throws an error if exceeded.
func (srv *Teams) UpdatePrefs(TeamId string, Prefs interface{}) (*ClientResponse, error) {
	r := strings.NewReplacer("{teamId}", TeamId)
	apiPath := r.Replace("/teams/{teamId}/prefs")

	apiParams := map[string]interface{}{
		"prefs": Prefs,
	}

	apiHeaders := map[string]interface{}{
		"content-type": "application/json",
	}
	return srv.client.Call("PUT", apiPath, apiHeaders, apiParams)
}
