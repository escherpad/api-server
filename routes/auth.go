package routes

import (
	"github.com/escherpad/api-server/handlers"
)

// AuthRoutes defines all routes in auth group.
var AuthRoutes = []Route{
	{
		"register",
		"POST",
		"/api/v1/auth/register",
		handlers.Register,
	},
	{
		"reset-password",
		"POST",
		"/api/v1/auth/reset-password",
		handlers.ResetPassword,
	},
	{
		"set-password",
		"POST",
		"/api/v1/auth/set-password",
		handlers.SetPassword,
	},
	{
		"create-team",
		"POST",
		"/api/v1/auth/create-team",
		handlers.CreateTeam,
	},
	{
		"signup-for-team",
		"POST",
		"/api/v1/auth/signup", // signup-for-team?
		handlers.SignupForTeam,
	},
	{
		"authenicate",
		"POST",
		"/api/v1/auth/authenticate",
		handlers.Authenticate,
	},
	{
		"logout",
		"POST",
		"/api/v1/auth/logout",
		handlers.Logout,
	},
	{
		"test-public",
		"GET",
		"/api/v1/auth/test-public",
		handlers.TestDummy,
	},
	{
		"test-private",
		"GET",
		"/api/v1/auth/test-private",
		handlers.TestDummy,
	},
}
