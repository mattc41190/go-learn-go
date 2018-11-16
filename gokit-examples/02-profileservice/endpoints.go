package profileservice

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
)

// Endpoints is a struct which contains the names of our service's functionality which at the time of usage will be "wrapped" in an Endpoint
type Endpoints struct {
	PostProfileEndpoint   endpoint.Endpoint
	GetProfileEndpoint    endpoint.Endpoint
	PutProfileEndpoint    endpoint.Endpoint
	PatchProfileEndpoint  endpoint.Endpoint
	DeleteProfileEndpoint endpoint.Endpoint
	GetAddressesEndpoint  endpoint.Endpoint
	GetAddressEndpoint    endpoint.Endpoint
	PostAddressEndpoint   endpoint.Endpoint
	DeleteAddressEndpoint endpoint.Endpoint
}

// MakeServerEndpoints accepts an implementer of the Service interface and returns an Endpoints struct.
// Each field within the Endpoints struct is and Endpoint. An Endpoint is a value of type function.
// Essentially what an Endpoint is a single function interface. Since an Endpoint has a set definition GoKit has a means by which it can
// handle usages of it.
func MakeServerEndpoints(s Service) Endpoints {
	return Endpoints{
		PostProfileEndpoint:   MakePostProfileEndpoint(s),
		GetProfileEndpoint:    MakeGetProfileEndpoint(s),
		PutProfileEndpoint:    MakePutProfileEndpoint(s),
		PatchProfileEndpoint:  MakePatchProfileEndpoint(s),
		DeleteProfileEndpoint: MakeDeleteProfileEndpoint(s),
		GetAddressesEndpoint:  MakeGetAddressesEndpoint(s),
		GetAddressEndpoint:    MakeGetAddressEndpoint(s),
		PostAddressEndpoint:   MakePostAddressEndpoint(s),
		DeleteAddressEndpoint: MakeDeleteAddressEndpoint(s),
	}
}

// MakeClientEndpoints is responsbile for making the Endpoints a client will use.
// Think of Client Endpoints as more a consumer of service than a service in itself.
// The core idea here is that our Client will allow us to interact with our service in a predictable fashion
// We create a consumer of our service as we build our service and
// our Client service will, if we follow the correct patterns, implement all our Service interface thus allowing
// consumers of our MicroService to interact with our service as if it were a dependency that resided as a dependency in the same codebase
func MakeClientEndpoints(instance string) (Endpoints, error) {

	// Check to see if the url passed is NOT prefixed with http
	if !strings.HasPrefix(instance, "http") {

		// Prepend http to the passed in URL
		instance = fmt.Sprintf("http://%s", instance)
	}

	// Create a new value of type URL based on the instance string passed.
	tgt, err := url.Parse(instance)

	// Check to see if err is nil
	if err != nil {
		// Panic!
		panic(err)
	}

	// Clear the URL (tgt) Path field. Warning: I am not sure why we do this...
	tgt.Path = ""

	options := []httptransport.ClientOption{}

	// Create an Endpoionts struct literal.
	// This is done by creating a new GoKit client for each Endpoint. Then using the Endpoint method
	return Endpoints{
		PostProfileEndpoint:   httptransport.NewClient("POST", tgt, encReqPostProfile, decRespPostProfile, options...).Endpoint(),
		GetProfileEndpoint:    httptransport.NewClient("GET", tgt, encReqGetProfile, decRespGetProfile, options...).Endpoint(),
		PutProfileEndpoint:    httptransport.NewClient("PUT", tgt, encReqPutProfile, decRespPutProfile, options...).Endpoint(),
		PatchProfileEndpoint:  httptransport.NewClient("PATCH", tgt, encReqPatchProfile, decRespPatchProfile, options...).Endpoint(),
		DeleteProfileEndpoint: httptransport.NewClient("DELETE", tgt, encReqDeleteProfile, decRespDeleteProfile, options...).Endpoint(),
		GetAddressesEndpoint:  httptransport.NewClient("GET", tgt, encReqGetAddresses, decRespGetAddresses, options...).Endpoint(),
		GetAddressEndpoint:    httptransport.NewClient("GET", tgt, encReqGetAddress, decRespGetAddress, options...).Endpoint(),
		PostAddressEndpoint:   httptransport.NewClient("POST", tgt, encReqPostAddress, decRespPostAddress, options...).Endpoint(),
		DeleteAddressEndpoint: httptransport.NewClient("DELETE", tgt, encReqDeleteAddress, decRespDeleteAddress, options...).Endpoint(),
	}, nil
}

// PostProfile is a function on the Endpoints which takes an implementer of Context
// and a value of type Profile and returns an error. (Only useful when used from Client)
func (e Endpoints) PostProfile(ctx context.Context, p Profile) error {

	// Initialize a variable called request to equal a struct literal of type postProfileRequest
	// whose Profile field is set to `p` the Profile passed to the function
	request := postProfileRequest{Profile: p}

	// Initialize two variables, one called `response` whose value is set to the result of calling a field on
	// `e` (the Endpoints value for which the current function is a receiver) called PostProfileEndpoint.
	// PostProfileEndpoint IS A CALLABLE ENDPOINT! This is confusing I know, but stick with it.
	// When PostProfileEndpoint is called it returns a response and error value. Response is of type interface{}
	// and error is an implementer of error.
	response, err := e.PostProfileEndpoint(ctx, request)

	// Check to see if calling the PostProfileEndpoint returned a non-nil error value
	if err != nil {

		// Return the error to the caller
		return err
	}

	// Initialize a variable called resp which is a the value inside of the response value asserted as a value of type postProfileResponse
	resp := response.(postProfileResponse)

	// Return the Err field from the resp variable (Remember resp is a value of type postProfileResponse)
	return resp.Err
}

// GetProfile is a function which takes a context and an ID string and returns a value of type Profile and an implementer of error
func (e Endpoints) GetProfile(ctx context.Context, ID string) (Profile, error) {

	// Initialize a variable of type GetProfileRequest and set its ID field to the ID value passed to the function
	request := getProfileRequest{ProfileID: ID}

	// Initialize two variable made by calling `e`'s GetProfileEndpoint field value.
	// This function will return a response value of type interface{} and an error
	response, err := e.GetProfileEndpoint(ctx, request)

	// Check to see if the endpoint call returned a non-nil error value.
	if err != nil {

		// Return an empty value Profile and the error from the endpoint call
		return Profile{}, err
	}

	// Initialize a variable called `resp` which will be the result of asserting the response (`response` interface{})
	//from GetProfileEndpoint as a GetProfileResponse value
	resp := response.(getProfileResponse)

	// Return the Profile
	return resp.Profile, resp.Err

}

// PutProfile is a function which takes a Context, an ID, and a Profile and returns an error.
// It will call a server endpoint which "puts" or replaces an exsiting profile
func (e Endpoints) PutProfile(ctx context.Context, profileID string, p Profile) error {

	// Create a request value based on the arguements passed in by the caller. Note, that this value
	// once sent over the wire will be decoded as the the required arguments to the service, so it may be a good idea to
	// make the request a type which is capable of being parsed as the sort of value which matches the Service interface function's
	// required signature ¯\_(ツ)_/¯. Another option is to have the server side decoder inspect and conditionally decorate the request
	// so that it maybe decoded as the Service expected type.
	request := putProfileRequest{ProfileID: profileID, Profile: p}

	// On the endpoints struct (the same one as the receiver for this function, no less) call the PutProfileEndpoint function.
	// The function will make an HTTP request to a running instance of this service and use the same GoLang native values on both Server and
	// Client side. The Endpoint function will return a response value of type interface{}, and an err value
	response, err := e.PutProfileEndpoint(ctx, request)

	// Check to see if the the PutProfile request returned a non-nil error value
	if err != nil {

		// Return the Endpoint call level error to the caller.
		return err
	}

	// Assert that the response server hold a concrete value of `putProfileResponse`
	resp := response.(putProfileResponse)

	// Return the error value held on the resp Err field
	return resp.Err
}

// PatchProfile is a function which takes a Context, an ID, and a Profile and returns an error
// PatchProfile differs slightly in implementation from PutProfile by selectively replacing data on the existing strcuture.
// It will call a server endpoint which "patches" or updates a profile in the datastore.
func (e Endpoints) PatchProfile(ctx context.Context, profileID string, p Profile) error {

	// TODO: Create detailed ref spec
	request := patchProfileRequest{ProfileID: profileID, Profile: p}

	response, err := e.PatchProfileEndpoint(ctx, request)

	if err != nil {
		return err
	}

	resp := response.(patchProfileResponse)

	return resp.Err
}

// DeleteProfile is a function which takes a Context and an ID string and returns an error.
// It will call a server endpoint which deletes a Profile from the datastore.
func (e Endpoints) DeleteProfile(ctx context.Context, profileID string) error {

	// TODO: Create detailed ref spec
	request := deleteProfileRequest{ProfileID: profileID}

	response, err := e.DeleteProfileEndpoint(ctx, request)

	if err != nil {
		return err
	}

	resp := response.(deleteProfileResponse)

	return resp.Err
}

// GetAddresses is a function which accepts a Context and a profileID string and returns a slice of Address and an error.
// It will call a server endpoint which returns a slice of addresses from the datastore. Client func.
func (e Endpoints) GetAddresses(ctx context.Context, profileID string) ([]Address, error) {

	// TODO: Create detailed ref spec
	request := getAddressesRequest{ProfileID: profileID}

	response, err := e.GetAddressesEndpoint(ctx, request)

	if err != nil {
		return []Address{}, nil
	}

	resp := response.(getAddressesResponse)

	return resp.Addresses, resp.Err

}

// GetAddress is a function which accepts a Context and a profileID string and an addressID string
// It will make a request to a server endpoint that returns an address associated with a specified profile
func (e Endpoints) GetAddress(ctx context.Context, profileID string, addressID string) (Address, error) {

	// TODO: Create detailed ref spec
	request := getAddressRequest{ProfileID: profileID, AddressID: addressID}

	response, err := e.GetAddressEndpoint(ctx, request)

	if err != nil {
		return Address{}, err
	}

	resp := response.(getAddressResponse)

	return resp.Address, resp.Err
}

// PostAddress is function which accepts a Context, a profileID, an Address.
// It will call an endpoint function which makes a request to an endpoint server. Client func.
func (e Endpoints) PostAddress(ctx context.Context, profileID string, a Address) error {

	// TODO: Create detailed ref spec
	request := postAddressRequest{ProfileID: profileID, Address: a}

	response, err := e.PostAddressEndpoint(ctx, request)

	if err != nil {
		return err
	}

	resp := response.(postAddressResponse)

	return resp.Err
}

// DeleteAddress is a function that accepts a Context, a profileID string, and an addressID string.
// It will call a server endpoint which returns an error if the address is not successfully deleted.
func (e Endpoints) DeleteAddress(ctx context.Context, profileID string, addressID string) error {
	request := deleteAddressRequest{ProfileID: profileID, AddressID: addressID}

	response, err := e.DeleteAddressEndpoint(ctx, request)

	if err != nil {
		return err
	}

	resp := response.(deleteAddressResponse)

	return resp.Err
}

// MakePostProfileEndpoint accepts a service and returns an Endpoint which is intended to be attached to a server
func MakePostProfileEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		// Assert / convert the value residing inside of the request to a postProfileRequest
		req := request.(postProfileRequest)

		// Initialize a variable e (short for error because GoLang).
		// It is set by calling the passed in Service's PostProfile function
		// Keeping in line with Service interface the PostProfile method accepts the context and request that were passed in.
		e := s.PostProfile(ctx, req.Profile)

		// Return a struct literal of type postProfileResponse whose Err field is set to `e`, the error we got back from the Service call.
		return postProfileResponse{Err: e}, nil
	}
}

// MakeGetProfileEndpoint accepts a service and returns an Endpoint function which should be associated to a server route path.
func MakeGetProfileEndpoint(s Service) endpoint.Endpoint {
	// TODO: Create detailed ref spec
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getProfileRequest)
		p, e := s.GetProfile(ctx, req.ProfileID)
		return getProfileResponse{Profile: p, Err: e}, nil
	}
}

// MakePutProfileEndpoint accepts a service and returns an Endpoint function which should be associated to a server route.
func MakePutProfileEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		// TODO: Create detailed ref spec
		req := request.(putProfileRequest)
		e := s.PutProfile(ctx, req.ProfileID, req.Profile)
		return putProfileResponse{Err: e}, nil
	}
}

// MakePatchProfileEndpoint accepts a service and returns an Endpoint function which should be associated to a server route
func MakePatchProfileEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		// TODO: Create detailed ref spec
		req := request.(patchProfileRequest)
		e := s.PatchProfile(ctx, req.ProfileID, req.Profile)
		return patchProfileResponse{Err: e}, nil
	}
}

// MakeDeleteProfileEndpoint accepts an implementer of Service and returns an Endpoint function which should be attached to a Server route.
// It is responsible for calling the Delete on the service from the request.
// Endpoints are the glue which attach SERVER routes to SERVICE behaviors.
// When an Endpoint is called the code has already been decoded into known Golang types via the GoKit Server side decoder
// The result is sent into the Endpoint function where it is cast as a type that the SERVICE will expect.
// At this point the Service is called
// When the result comes back it is set to a value which our response encoder will know what to do with.
// GoKit will then pass this value to our transport encoder at which point it will send the results back to the caller.
func MakeDeleteProfileEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteProfileRequest)
		e := s.DeleteProfile(ctx, req.ProfileID)
		return deleteProfileResponse{Err: e}, nil
	}
}

// MakeGetAddressesEndpoint accepts a Service and returns an Endpoint function which should be attached to a Server route in the form of a Server.
func MakeGetAddressesEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		// At this point in the flow the current function has access to the decoded value from our request decoder attached to the GoKit Server.
		// This decoded value will be asserted / cast as a Service consumable getAddressesRequest called `req`
		req := request.(getAddressesRequest)
		a, e := s.GetAddresses(ctx, req.ProfileID)
		return getAddressesResponse{Addresses: a, Err: e}, nil
	}
}

// MakeGetAddressEndpoint accepts an implemeter of Service and returns and endpoint function which will be attached to a Server route
// via the GoKit NewServer function which sandwiches an Endpoint function between a request decoder / receiver and a response encoder / sender
func MakeGetAddressEndpoint(s Service) endpoint.Endpoint {
	// TODO: Create detailed ref spec
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getAddressRequest)
		addr, e := s.GetAddress(ctx, req.ProfileID, req.AddressID)
		return getAddressResponse{Err: e, Address: addr}, nil
	}
}

// MakePostAddressEndpoint is a function which takes a service and returns an Endpoint function
// which will be attached to a GoKit server which will be attached to an http.Server's route.
func MakePostAddressEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		// Assert the request interface passed in as a postProfileRequest value
		// req will be a value of type postAddressRequest with the values properly mapped from the dynamic value passed in
		req := request.(postAddressRequest)

		// Create an error based on the return from the Service's PostAddress function
		e := s.PostAddress(ctx, req.ProfileID, req.Address)

		return postAddressResponse{Err: e}, nil

	}
}

// MakeDeleteAddressEndpoint accepts a service and returns an Endpoint function which should be attached to a server
func MakeDeleteAddressEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteAddressRequest)
		e := s.DeleteAddress(ctx, req.ProfileID, req.AddressID)
		return deleteAddressResponse{Err: e}, nil
	}
}

type postProfileRequest struct {
	Profile Profile
}
type postProfileResponse struct {
	Err error `json:"err,omitempty"`
}

type getProfileRequest struct {
	ProfileID string
}

type getProfileResponse struct {
	Err     error   `json:"err,omitempty"`
	Profile Profile `json:"profile,omitempty"`
}

type putProfileRequest struct {
	ProfileID string
	Profile   Profile
}

type putProfileResponse struct {
	Err error `json:"err,omitempty"`
}

type patchProfileRequest struct {
	ProfileID string
	Profile   Profile
}

type patchProfileResponse struct {
	Err error `json:"err,omitempty"`
}

type deleteProfileRequest struct {
	ProfileID string
}

type deleteProfileResponse struct {
	Err error `json:"err,omitempty"`
}

type getAddressesRequest struct {
	ProfileID string
}

type getAddressesResponse struct {
	Err       error     `json:"err,omitempty"`
	Addresses []Address `json:"addresses,omitempty"`
}

type getAddressRequest struct {
	ProfileID string
	AddressID string
}

type getAddressResponse struct {
	Err     error   `json:"err,omitempty"`
	Address Address `json:"address,omitempty"`
}

type postAddressRequest struct {
	Address   Address
	ProfileID string
}

type postAddressResponse struct {
	Err error `json:"err,omitempty"`
}

type deleteAddressRequest struct {
	AddressID string
	ProfileID string
}

type deleteAddressResponse struct {
	Err error `json:"err,omitempty"`
}
