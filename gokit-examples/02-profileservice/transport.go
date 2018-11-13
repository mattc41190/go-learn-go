package profileservice

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

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
	r.Method("GET").Path("/profiles/{profileID}").Handler(httptransport.NewServer(
		epoints.GetAddressEndpoint,
		decReqGetProfile,
		encRespGetProfile,
		options...,
	))
	r.Method("PUT").Path("/profiles/{profileID}").Handler(httptransport.NewServer(
		epoints.PutProfileEndpoint,
		decReqPutProfile,
		encRespPutProfile,
		options...,
	))
	r.Method("PATCH").Path("/profiles/{profileID}").Handler(httptransport.NewServer(
		epoint.PatchProfileEndpoint,
		decReqPatchProfile,
		encRespPatchProfile,
		options...,
	))
	r.Method("DELETE").Path("/profiles/{profileID}").Handler(httptransport.NewServer(
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
func decReqPostProfile(ctx context.Context, request interface{}) (interface{}, error) {

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

func decReqGetProfile(ctx context.Context, request interface{}) (interface{}, error) {
	// TODO: Supply full explanation

	// Extract the variables from the URL path
	vars := mux.Vars(request)

	// Run a check to see if `profileID` is one of the URL variables and set the return from that check to a variable called `profileID`
	profileID, ok := vars["profileID"]

	// Check to see if `id` was NOT included in the URL path
	if !ok {

		// Inform the caller that this decoder cannot be called for a route without an id param in the its path
		return nil, ErrBadRouting
	}

	// Return a getProfileRequest literal to the user and a nil error value
	return getProfileRequest{ProfileID: profileID}, nil
}

func decReqPutProfile(ctx context.Context, request interface{}) (interface{}, error) {

	// Create a empty valued placeholder for our return value
	var putProfileReq putProfileRequest

	// Create a `vars` variable that will hold URL variables from the path
	vars := mux.Vars(request)

	// Check the `vars` map variable for a `key` called `id`. If id is found set a variable called `id` to that value.
	profileID, ok := vars["profileId"]

	// Check to see if the map lookup encountered an error.
	if !ok {

		// Return a nil value for the request struct and an error informing the user that the decoder can't work on the provided URL
		return nil, ErrBadRouting
	}

	// Set the ID field to the value found in the URL `vars` map
	putProfileReq.ProfileID = profileID

	// Create an error variable whose value is set by attempting to read the passed request body into putProfileReq
	err := json.NewDecoder(request.Body).Decode(&putProfileReq.Profile)

	// Check to see if there was an error decoding the JSON into the specified struct
	if err != nil {

		// Return a nil request value and the error val
		return nil, err
	}

	// Return the putProfileRequest
	return putProfileReq, nil
}

// Decoders for requests are used on the server side of life :)
func decReqPatchProfile(ctx context.Context, request interface{}) (interface{}, error) {

	var patchProfileReq patchProfileRequest

	vars := mux.Vars(request)
	profileID, ok := vars["profileId"]

	if !ok {
		return nil, ErrBadRouting
	}

	patchProfileReq.ProfileID = profileID
	err := json.NewDecoder(request.Body).Decode(&patchProfileReq.Profile)

	if err != nil {
		return nil, err
	}

	return patchProfileReq, nil
}

func decReqDeleteProfile(ctx context.Context, request interface{}) (interface{}, error) {

	var deleteProfileReq deleteProfileRequest

	vars := mux.Vars(request)
	id, ok := vars["profileID"]

	if !ok {
		// Don't call this decoder if the path doesn't include `profileID`
		return nil, ErrBadRouting
	}

	deleteProfileReq.ProfileID = profileID

	return deleteProfileReq, nil

}

// decReqGetAddresses decodes raw bytes coming into a server into a getAddressesRequest
func decReqGetAddresses(ctx context.Context, request interface{}) (interface{}, error) {

	// Get a map of all URL vars using
	vars := mux.Vars(request)

	// Attempt to extract N ("profileID=N") into a variable called profileID
	profileID, ok := vars["profileID"]

	// Check to make sure that "profileID" exists on the URL.
	if !ok {

		// Return an empty getAddressesRequest value and an error informing the caller that the URL was malformed
		return nil, ErrBadRouting
	}

	// Return getAddressesRequest with the Prpfile field filled in with the profileID found on the URL
	return getAddressesRequest{ProfileID: profileID}, nil
}

func decReqGetAddress(ctx context.Context, request interface{}) (interface{}, error) {
	vars := mux.Vars(request)

	profileID, ok := vars["profileID"]

	if !ok {
		return nil, ErrBadRouting
	}

	addressID, ok := vars["addressID"]

	if !ok {
		return nil, ErrBadRouting
	}

	return getAddressRequest{ProfileID: profileID, AddressID: addressID}, nil
}

func decReqPostAddress(ctx context.Context, request interface{}) (interface{}, error) {

	vars := mux.Vars(request)

	profileID, ok := vars["profileID"]

	if !ok {
		return nil, ErrBadRouting
	}

	var address Address

	err := json.NewDecoder(request.Body).Decode(&address)

	if err != nil {
		return nil, err
	}

	return postAddressRequest{ProfileID: profileID, Address: address}, nil
}

func decReqDeleteAddress(ctx context.Context, request interface{}) (interface{}, error) {
	vars := mux.Vars(request)

	profileID, ok := vars["profileID"]

	if !ok {
		return nil, ErrBadRouting
	}

	addressID, ok := vars["addressID"]

	if !ok {
		return nil, ErrBadRouting
	}

	return deleteAddressRequest{ProfileID, profileID, AddressID, addressID}, nil
}

//
// Client Request Encoders
// When a client wants to send a request the server it will use a request encoder.
// Remember that the baseURL will already be housed on the client, so all that needs to be added is the path
// and in the case of a JSON API, a function that transforms a native GoLang type to a slice of bytes
// The entire methodology around request encoders is to operate on pointers and share memory.
// This means that the encoders won't be returning http.Request, but rather manipulating pointers to http.Request
// GoKit will, behind the scenes use this request with the `Do` method to perform the actual http request.
// All the operator is responsible for is manipulating the memory.
//

func encReqPostProfile(ctx context.Context, req *http.Request, request interface{}) error {
	request.URL.Path = "/profile/"
	return encodeJSONRequest(ctx, request, req)
}

func encReqGetProfile(ctx context.Context, req *http.Request, request interface{}) error {
	r := request.(getProfileRequest)
	profileID := url.QueryEscape(r.profileID)
	request.URL.Path = "/profile/" + profileID
	return encodeJSONRequest(ctx, req, request)
}

func encReqPutProfile(ctx context.Context, req *http.Request, request interface{}) error {
	r := request.(putProfileRequest)
	profileID := url.QueryEscape(r.ID)
	request.URL.Path = "/profile/" + profileID
	return encodeJSONRequest(ctx, req, request)
}

func encReqPatchProfile(ctx context.Context, req *http.Request, request interface{}) error {
	r := request.(patchProfileRequest)
	profileID := url.QueryEscape(r.ProfileID)
	req.URL.Path = "/profile/" + profileID
	return encodeJSONRequest(ctx, req, request)
}

func encReqDeleteProfile(ctx context.Context, req *http.Request, request interface{}) error {
	r := request.(deleteProfileRequest)
	profileID := url.QueryEscape(r.ProfileID)
	req.URL.Path = "/profile/" + profileID
	return encodeJSONRequest(ctx, req, request)
}

func encReqGetAddresses(ctx context.Context, req *http.Request, request interface{}) error {
	r := request.(getAddressesRequest)
	profileID := url.QueryEscape(r.ProfileID)
	req.URL.Path = "/profile/" + profileID + "/address/"
	return encodeJSONRequest(ctx, req, request)
}

func encReqGetAddress(ctx context.Context, req *http.Request, request interface{}) error {
	r := request.(getAddressRequest)
	profileID := url.QueryEscape(r.ProfileID)
	addressID := url.QueryEscape(r.AddressID)
	req.URL.Path = "/profile/" + profileID + "/address/" + addressID
	return encodeJSONRequest(ctx, req, request)
}

func encReqPostAddress(ctx context.Context, req *http.Request, request interface{}) error {
	r := request.(postAddressRequest)
	profileID := url.QueryEscape(r.ProfileID)
	req.URL.Path = "/profile/" + profileID + "address"
	return encodeJSONRequest(ctx, req, request)
}

func encReqDeleteAddress(ctx context.Context, req *http.Request, request interface{}) error {
	r := request.(deleteAddressRequest)
	profileID := url.QueryEscape(r.ProfileID)
	addressID := url.QueryEscape(r.AddressID)
	req.URL.Path = "/profile/" + profileID + "/address/" + addressID
	return encodeJSONRequest(ctx, req, request)
}

// encodeJSONRequest is responsible for writing information into a request via a pointer.
// It writes information relating to the passed in request to a JSON body.
// Useful for HTTP request types which carry a body.
// It accepts a context, a pointer to an httpRequest and an empty interface. It returns an error.
func encodeJSONRequest(ctx context.Context, req *http.Request, request interface{}) error {

	// Create an empty value `bytes.Buffer` called `buf`
	var buf bytes.Buffer

	// Create an `err` var and set it to the error value return by calling Encode an new JSON encoder whose
	err := json.NewEncoder(&buf).Encode(request)
	if err != nil {
		// Return an error informing the caller that the request was not encoded, and thus not sent over the wire.
		return err
	}

	// Set the requests body to the bytes slice / buffer encoded json object
	req.Body = ioutil.NopCloser(&buf)

	// Return a nil error value
	return nil
}
