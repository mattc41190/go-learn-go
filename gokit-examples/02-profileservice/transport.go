package profileservice

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

// Create a `var`s block
var (

	// Create an error for trying to call a bad route + method combination
	ErrBadRouting = errors.New("inconsistent mapping between route and handler (programmer error)")
)

// MakeHTTPHandler will accept an implementer of the Service interface AKA a Service and a logger
// and it will return an http.Handler. Note that an http.Handler is an interface which maintains a single method called
// ServeHTTP. ServeHTTP accepts a http.ResponseWriter (interface) and a a pointer to a http.Request (struct).
func MakeHTTPHandler(s Service, logger log.Logger) http.Handler {

	// Create a variable router and set it the return value of a new route multiplexer
	r := mux.NewRouter()

	// Create a var called endpoints and set it to an Endpoints struct with a server side configuration.
	epoints := MakeServerEndpoints(s)

	// Create a slice of server options
	options := []httptransport.ServerOption{
		httptransport.ServerErrorLogger(logger),
		httptransport.ServerErrorEncoder(encodeError),
	}

	// Create Gorilla Mux Handlers For Each Endpoint
	// TODO: Decompose one of these guys method by method
	r.Methods("POST").Path("/profiles/").Handler(httptransport.NewServer(
		epoints.PostAddressEndpoint,
		decReqPostProfile,
		encRespPostProfile,
		options...,
	))
	r.Method("GET").Path("/profiles/{id}").Handler(httptransport.NewServer(
		epoints.GetAddressEndpoint,
		decReqGetProfile,
		encRespGetProfile,
		options...,
	))
	r.Method("PUT").Path("/profiles/{id}").Handler(httptransport.NewServer(
		epoints.PutProfileEndpoint,
		decReqPutProfile,
		encRespPutProfile,
		options...,
	))
	r.Method("PATCH").Path("/profiles/{id}").Handler(httptransport.NewServer(
		epoint.PatchProfileEndpoint,
		decReqPatchProfile,
		encRespPatchProfile,
		options...,
	))
	r.Method("DELETE").Path("/profiles/{id}").Handler(httptransport.NewServer(
		epoint.DeleteProfileEndpoint,
		decReqDeleteProfile,
		encRespDeleteProfile,
		options...,
	))
	r.Method("GET").Path("/profiles/{profileID}/address").Handler(httptransport.NewServer(
		epoints.GetAddressesEndpoint,
		decReqGetAddresses,
		encRespGetAddresses,
		options...,
	))
	r.Method("GET").Path("/profiles/{profileID}/address/{addressID}/").Handler(httptransport.NewServer(
		epoints.GetAddressEndpoint,
		decReqGetAddress,
		encRespGetAddress,
		options...,
	))
	r.Method("POST").Path("/profiles/{profileID}/address/").Handler(httptransport.NewServer(
		epoints.PostAddressEndpoint,
		decReqPostAddress,
		encRespPostAddress,
		options...,
	))
	r.Method("DELETE").Path("/profiles/{profileID}/address/{addressID}").Handler(httptransport.NewServer(
		epoints.DeleteAddressEndpoint,
		deqReqDeleteAddress,
		encRespDeleteAddress,
		options...,
	))

	// Return the route multiplexer.
	return r

}

// --------
// DECODERS
// --------

// decReqPostProfile will be the first place a request goes to once it enters the handler.
// This flow is determined by the httptransport Server's `ServeHTTP` method.
// The `ServeHTTP` method calls a collection of option functions(pre or post).
// Then calls a request decoder which is responsible for decoding the request content from it's RAW byte form into something usable by GoLang.
// TODO: Create full decomposition of GoKit server
func decReqPostProfile(ctx context.Context, request interface{}) (postProfileRequest, error) {

	// Create an empty valued holding place for the desired value
	var req postProfileRequest

	// Decode the value in the the request body into the `Profile` field of the `req` var.
	// Capture any errors that occurred as `err`
	err := json.NewDecoder(request.Body).Decode(&req.Profile)

	// Check to see if any error occurred during the decoding process
	if err != nil {

		// Return a nil value for the decoded request
		return nil, err
	}

	// Return the value of type postProfileRequest (populated with the decoded JSON from the HTTP request), and `nil` error value to the caller.
	return req, nil
}

func decReqGetProfile(ctx context.Context, request interface{}) (getProfileRequest, error) {
	// TODO: Supply full explanation

	// Extract the variables from the URL path
	vars := mux.Vars()

	// Run a check to see if `id` is one of the URL variables and set the return from that check to a variable called `profileID`
	profileID, ok := vars["id"]

	// Check to see if `id` was NOT included in the URL path
	if !ok {

		// Inform the caller that this decoder cannot be called for a route without an id param in the its path
		return nil, ErrBadRouting
	}

	// Return a getProfileRequest literal to the user and a nil error value
	return getProfileRequest{ID: profileID}, nil

}
