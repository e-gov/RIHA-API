package riha

import (
	"fmt"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
	Role        string
}

func (r Route) String() string {
	return fmt.Sprintf("Route name:%s method:%s pattern:%s role:%s", r.Name, r.Method, r.Pattern, r.Role)
}

type Routes []Route

// The last element of the route is internal role name that is required for access
// * indicates an un-restricted route.
var routes = Routes{
	Route{
		"GetAll",
		"GET",
		"/systems.json",
		List,
		"*",
	},
	Route{
		"GetOne",
		"GET",
		"/{shortname}",
		GetOne,
		"*",
	},
}
