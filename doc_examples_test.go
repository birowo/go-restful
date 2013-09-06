package restful

import "net/http"

func ExampleOPTIONSFilter() {
	// Install the OPTIONS filter on the default Container
	Filter(OPTIONSFilter())
}
func ExampleContainer_OPTIONSFilter() {
	// Install the OPTIONS filter on a Container
	myContainer := new(Container)
	myContainer.Filter(myContainer.OPTIONSFilter)
}

func ExampleContainer() {
	// The Default container of go-restful uses the http.DefaultServeMux.
	// You can create your own Container using restful.NewContainer() and create a new http.Server for that particular container

	ws := new(WebService)
	wsContainer := NewContainer()
	wsContainer.Add(ws)
	server := &http.Server{Addr: ":8080", Handler: wsContainer}
	server.ListenAndServe()
}

func ExampleCrossOriginResourceSharing() {
	// To install this filter on the Default Container use:
	cors := CrossOriginResourceSharing{ExposeHeaders: []string{"X-My-Header"}, CookiesAllowed: false, Container: DefaultContainer}
	Filter(cors.Filter)
}
