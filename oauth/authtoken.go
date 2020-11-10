//
// Copyright (c) 2020 SSH Communications Security Inc.
//
// All rights reserved.
//

package oauth

import "github.com/SSHcom/privx-sdk-go/restapi"

type tAuthExplicit struct{ string }

/*

WithToken uses explicitly defined JWT to authenticate client
*/
func WithToken(token string) restapi.Authorizer {
	return &tAuthExplicit{token}
}

func (auth *tAuthExplicit) AccessToken() (string, error) {
	return auth.string, nil
}
