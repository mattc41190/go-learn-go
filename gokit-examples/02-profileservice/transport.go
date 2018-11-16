package profileservice

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/go-kit/kit/log"

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
		encodeResponse,
		options...,
	))
	r.Methods("GET").Path("/profiles/{profileID}").Handler(httptransport.NewServer(
		epoints.GetAddressEndpoint,
		decReqGetProfile,
		encodeResponse,
		options...,
	))
	r.Methods("PUT").Path("/profiles/{profileID}").Handler(httptransport.NewServer(
		epoints.PutProfileEndpoint,
		decReqPutProfile,
		encodeResponse,
		options...,
	))
	r.Methods("PATCH").Path("/profiles/{profileID}").Handler(httptransport.NewServer(
		epoints.PatchProfileEndpoint,
		decReqPatchProfile,
		encodeResponse,
		options...,
	))
	r.Methods("DELETE").Path("/profiles/{profileID}").Handler(httptransport.NewServer(
		epoints.DeleteProfileEndpoint,
		decReqDeleteProfile,
		encodeResponse,
		options...,
	))
	r.Methods("GET").Path("/profiles/{profileID}/address").Handler(httptransport.NewServer(
		epoints.GetAddressesEndpoint,
		decReqGetAddresses,
		encodeResponse,
		options...,
	))
	r.Methods("GET").Path("/profiles/{profileID}/address/{addressID}/").Handler(httptransport.NewServer(
		epoints.GetAddressEndpoint,
		decReqGetAddress,
		encodeResponse,
		options...,
	))
	r.Methods("POST").Path("/profiles/{profileID}/address/").Handler(httptransport.NewServer(
		epoints.PostAddressEndpoint,
		decReqPostAddress,
		encodeResponse,
		options...,
	))
	r.Methods("DELETE").Path("/profiles/{profileID}/address/{addressID}").Handler(httptransport.NewServer(
		epoints.DeleteAddressEndpoint,
		decReqDeleteAddress,
		encodeResponse,
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
func decReqPostProfile(ctx context.Context, request *http.Request) (interface{}, error) {

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

func decReqGetProfile(ctx context.Context, request *http.Request) (interface{}, error) {
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

func decReqPutProfile(ctx context.Context, request *http.Request) (interface{}, error) {

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
func decReqPatchProfile(ctx context.Context, request *http.Request) (interface{}, error) {

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

func decReqDeleteProfile(ctx context.Context, request *http.Request) (interface{}, error) {

	var deleteProfileReq deleteProfileRequest

	vars := mux.Vars(request)
	profileID, ok := vars["profileID"]

	if !ok {
		// Don't call this decoder if the path doesn't include `profileID`
		return nil, ErrBadRouting
	}

	deleteProfileReq.ProfileID = profileID

	return deleteProfileReq, nil

}

// decReqGetAddresses decodes raw bytes coming into a server into a getAddressesRequest
func decReqGetAddresses(ctx context.Context, request *http.Request) (interface{}, error) {

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

func decReqGetAddress(ctx context.Context, request *http.Request) (interface{}, error) {
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

func decReqPostAddress(ctx context.Context, request *http.Request) (interface{}, error) {

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

func decReqDeleteAddress(ctx context.Context, request *http.Request) (interface{}, error) {
	vars := mux.Vars(request)

	profileID, ok := vars["profileID"]

	if !ok {
		return nil, ErrBadRouting
	}

	addressID, ok := vars["addressID"]

	if !ok {
		return nil, ErrBadRouting
	}

	return deleteAddressRequest{ProfileID: profileID, AddressID: addressID}, nil
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
	req.URL.Path = "/profile/"
	return encodeJSONRequest(ctx, req, request)
}

func encReqGetProfile(ctx context.Context, req *http.Request, request interface{}) error {
	r := request.(getProfileRequest)
	profileID := url.QueryEscape(r.ProfileID)
	req.URL.Path = "/profile/" + profileID
	return encodeJSONRequest(ctx, req, request)
}

func encReqPutProfile(ctx context.Context, req *http.Request, request interface{}) error {
	r := request.(putProfileRequest)
	profileID := url.QueryEscape(r.ProfileID)
	req.URL.Path = "/profile/" + profileID
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

//
// Client Response Decoders
// The following decoders will accept a raw http response message (`bytes.Buffer`) and decode it into a known and usable type.
// Basically all of the decoders are the same...
//

// decRespPostProfile accepts a context, and a http.Response value. It returns an interface and error.
// It's return value should be assertable as a postProfileResponse value. This occurs directly before the client's
// endpoint function does a final clean up of the final value to return to the originator.
func decRespPostProfile(ctx context.Context, resp *http.Response) (interface{}, error) {
	var response postProfileResponse
	err := json.NewDecoder(resp.Body).Decode(&response)
	return response, err
}

func decRespGetProfile(ctx context.Context, resp *http.Response) (interface{}, error) {
	var response getProfileResponse
	err := json.NewDecoder(resp.Body).Decode(&response)
	return response, err
}

func decRespPutProfile(ctx context.Context, resp *http.Response) (interface{}, error) {
	var response putProfileResponse
	err := json.NewDecoder(resp.Body).Decode(&response)
	return response, err
}

func decRespPatchProfile(ctx context.Context, resp *http.Response) (interface{}, error) {
	var response patchProfileResponse
	err := json.NewDecoder(resp.Body).Decode(&response)
	return response, err
}

func decRespDeleteProfile(ctx context.Context, resp *http.Response) (interface{}, error) {
	var response deleteProfileResponse
	err := json.NewDecoder(resp.Body).Decode(&response)
	return response, err
}

func decRespGetAddresses(ctx context.Context, resp *http.Response) (interface{}, error) {
	// Create an empty value of type getAddressesResponse
	var response getAddressesResponse

	// Create a `decoder` which uses `resp` as a "source" of information.
	decoder := json.NewDecoder(resp.Body)

	// Create an error variable which is the result of attempting to decode the information stored in the decoder to
	// the `response` variable.
	err := decoder.Decode(&response)

	// Return the response var and any error value created during the decoding process.
	return response, err
}

func decRespGetAddress(ctx context.Context, resp *http.Response) (interface{}, error) {
	var response getAddressResponse
	err := json.NewDecoder(resp.Body).Decode(&response)
	return response, err
}

func decRespPostAddress(ctx context.Context, resp *http.Response) (interface{}, error) {
	var response postAddressResponse
	err := json.NewDecoder(resp.Body).Decode(&response)
	return response, err
}

func decRespDeleteAddress(ctx context.Context, resp *http.Response) (interface{}, error) {
	var response deleteAddressResponse
	err := json.NewDecoder(resp.Body).Decode(&response)
	return response, err

}

//
// Multi-Purpose Encoders
// The following encoders are used either by a server or a client (this should be specified on the function).
// Since the API is largely JSON based once a request or response is "generic enough" to be encoded as JSON there
// is no reason to copy and paste the EXACT same code over and over.
//

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

type errorer interface {
	error() error
}

// encodeResponse, once a request has been fully executed and the response has been created a GoKit server response
// encoder will be invoked (by GoKit). These encoders will write to a a ResponseWriter entity.
func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {

	// Attempt to assert that the response recieved from the server endpoint function implements the errorer interface
	// If it does set `err` equal to that value (that value being the value that has an error func) and `ok` to true.
	err, ok := response.(errorer)

	// Check to see if `ok` is true AND the implementer of errorer has an error func that returns a non-nil value
	if (ok) && (err.error() != nil) {

		// Pass the error value to the encode error function along with the `ResponseWriter`.
		// This will allow encodeError to finish the request lifecycle.
		encodeError(ctx, err.error(), w)

		// Return a nil error value to the caller
		return nil
	}

	// On the response writer create a header value informing the consumer that the message being carried is of type JSON
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	// Return the error value created by attempting to write the response value (interface{} / map[string]interface{})
	// to a JSON I/O encoder.
	return json.NewEncoder(w).Encode(response)
}

// encodeError accepts a value of type Context, an error, and a http.ResponseWriter.
// It has no return value.
func encodeError(ctx context.Context, err error, w http.ResponseWriter) {

	// Check to see if the error passed in was nil
	if err == nil {

		// panic!
		panic("cannot encode a nil error value")
	}

	// Set a header informing the consumer that the response they are about to receive is JSON
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	// Set the status for the error based on the result of passing the `err` value to `codeFrom`
	w.WriteHeader(codeFrom(err))

	// Encode a JSON parsable type of `map[string]interface{}` with a single error field whose value is set to
	// the return value from calling the error interface Error func into the ResponseWriter passed in.
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

// codeFrom is a wrapper around a switch statement that provide HTTP status code based on user created err types.
func codeFrom(err error) int {
	switch err {
	case ErrNotFound:
		return http.StatusNotFound
	case ErrAlreadyExists, ErrInconsistentIDs:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
