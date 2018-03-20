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
		"CreateResume",
		"POST",
		"/resumes",
		CreateResume,
	},
	Route{
		"GetAccount",
		"GET",
		"/resumes/{accountId}",
		GetResume,
	},
	Route{
		"UpdateBasicResume",
		"PUT",
		"/resumes/{accountId}/basic",
		UpdateBasicResume,
	},
	Route{
		"UpdateWorkResume",
		"PUT",
		"/resumes/{accountId}/work",
		UpdateWorkResume,
	},
	Route{
		"UpdateEducationResume",
		"PUT",
		"/resumes/{accountId}/education",
		UpdateEducationResume,
	},
	Route{
		"UpdateProjectResume",
		"PUT",
		"/resumes/{accountId}/project",
		UpdateProjectResume,
	},
	Route{
		"HealthCheck",
		"GET",
		"/health",
		HealthCheck,
	},
}
