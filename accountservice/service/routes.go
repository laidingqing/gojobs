package service

import "net/http"

//Route defines a single route, e.g. a human readable name, HTTP method, pattern the function that will execute when the route is called.
type Route struct {
	Name        string // Name
	Method      string // HTTP method
	Pattern     string // Route pattern
	HandlerFunc http.HandlerFunc
}

// Routes defines the type Routes which is just an array (slice) of Route structs.
type Routes []Route

// Initialize our routes
var routes = Routes{
	Route{
		"SeedAccounts",
		"POST",
		"/accounts/seed",
		SeedAccounts,
	},
	Route{
		"GetAccount",
		"GET",
		"/accounts/{accountId}",
		GetAccount,
	},
	Route{
		"CreateAccount",
		"POST",
		"/accounts",
		CreateAccount,
	},
	Route{
		"HealthCheck",
		"GET",
		"/health",
		HealthCheck,
	},
}
