// Discordgo - Discord bindings for Go
// Available at https://github.com/bwmarrin/discordgo

// Copyright 2015-2016 Bruce Marriner <bruce@sqls.net>.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file contains functions related to Discord OAuth2 endpoints

package discordgo

// ------------------------------------------------------------------------------------------------
// Code specific to Discord OAuth2 Applications
// ------------------------------------------------------------------------------------------------

// An Application struct stores values for a Discord OAuth2 Application
type Application struct {
	ID           string    `json:"id,omitempty"`
	Name         string    `json:"name"`
	Description  string    `json:"description,omitempty"`
	Icon         string    `json:"icon,omitempty"`
	Secret       string    `json:"secret,omitempty"`
	RedirectURIs *[]string `json:"redirect_uris,omitempty"`
}

// Application returns an Application structure of a specific Application
//   appID : The ID of an Application
func (s *Session) Application(appID string) (st *Application, err error) {

	body, err := s.Request("GET", APPLICATION(appID), nil)
	if err != nil {
		return
	}

	err = unmarshal(body, &st)
	return
}

// Applications returns all applications for the authenticated user
func (s *Session) Applications() (st []*Application, err error) {

	body, err := s.Request("GET", APPLICATIONS, nil)
	if err != nil {
		return
	}

	err = unmarshal(body, &st)
	return
}

// ApplicationCreate creates a new Application
//    name : Name of Application / Bot
//    uris : Redirect URIs (Not required)
func (s *Session) ApplicationCreate(ap *Application) (st *Application, err error) {

	data := struct {
		Name         string    `json:"name"`
		Description  string    `json:"description"`
		RedirectURIs *[]string `json:"redirect_uris,omitempty"`
	}{ap.Name, ap.Description, ap.RedirectURIs}

	body, err := s.Request("POST", APPLICATIONS, data)
	if err != nil {
		return
	}

	err = unmarshal(body, &st)
	return
}

// ApplicationUpdate updates an existing Application
//   var : desc
func (s *Session) ApplicationUpdate(appID string, ap *Application) (st *Application, err error) {

	data := struct {
		Name         string    `json:"name"`
		Description  string    `json:"description"`
		RedirectURIs *[]string `json:"redirect_uris,omitempty"`
	}{ap.Name, ap.Description, ap.RedirectURIs}

	body, err := s.Request("PUT", APPLICATION(appID), data)
	if err != nil {
		return
	}

	err = unmarshal(body, &st)
	return
}

// ApplicationDelete deletes an existing Application
//   appID : The ID of an Application
func (s *Session) ApplicationDelete(appID string) (err error) {

	_, err = s.Request("DELETE", APPLICATION(appID), nil)
	if err != nil {
		return
	}

	return
}

// ------------------------------------------------------------------------------------------------
// Code specific to Discord OAuth2 Application Bots
// ------------------------------------------------------------------------------------------------

// ApplicationBotCreate creates an Application Bot Account
//
//   appID : The ID of an Application
//   token : The authentication Token for a user account to convert into
//           a bot account.  This is optional, if omited a new account
//           is created using the name of the application.
//
// NOTE: func name may change, if I can think up something better.
func (s *Session) ApplicationBotCreate(appID, token string) (st *User, err error) {

	data := struct {
		Token string `json:"token,omitempty"`
	}{token}

	body, err := s.Request("POST", APPLICATIONS_BOT(appID), data)
	if err != nil {
		return
	}

	err = unmarshal(body, &st)
	return
}
