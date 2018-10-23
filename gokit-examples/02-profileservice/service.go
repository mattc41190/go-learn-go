package profileservice

import (
	"context"
	"errors"
	"sync"
)

// Service interface dictates our core functionality. This is the services our application is capable of.
type Service interface {
	PostProfile(ctx context.Context, p Profile) error
	GetProfile(ctx context.Context, id string) (Profile, error)
	PutProfile(ctx context.Context, id string, p Profile) error
	PatchProfile(ctx context.Context, id string, p Profile) error
	DeleteProfile(ctx context.Context, id string) error
	GetAddresses(ctx context.Context, profileID string) ([]Address, error)
	GetAddress(ctx context.Context, profileID string, addressID string) (Address, error)
	PostAddress(ctx context.Context, profileID string, a Address) error
	DeleteAddress(ctx context.Context, profileID string, addressID string) error
}

// Profile represents a profile. ID should be unique.
type Profile struct {
	ID        string    `json:id`
	Name      string    `json:"name,omitempty"`
	Addresses []Address `json:"addresses,omitempty"`
}

// Address is a field within profile. It represent a geographical address
type Address struct {
	ID       string `json:id`
	Location string `json:"location,omitempty"`
}

// inmemService represents our datastore
type inmemService struct {
	mtx      sync.RWMutex
	profiles map[string]Profile
}

// Create canned errors for consistent message for caller
var (
	ErrInconsistentIDs = errors.New("inconsistent IDs")
	ErrAlreadyExists   = errors.New("already exists")
	ErrNotFound        = errors.New("not found")
)

// NewInmemService will return a new inmemService.
// It should be noted here that *inmemService implements all of the functions
// required to make an entity a `Service`
// For more information on this, see all functions whose receiver is inmemService
func NewInmemService() Service {
	return &inmemService{
		profiles: map[string]Profile{},
	}
}

// PostProfile will save a Profile to the database
func (svc *inmemService) PostProfile(ctx context.Context, p Profile) error {
	// Get the lock from the inmemService struct
	svc.mtx.Lock()

	// Immediately set up a lock release to occur when the function finishes
	defer svc.mtx.Unlock()

	// Check to see if the profile already exists
	if _, ok := svc.profiles[p.ID]; ok {

		// If it does return an error which informs the caller that profile already exists
		return ErrAlreadyExists
	}

	// Insert the profile to our datastore
	svc.profiles[p.ID] = p

	// Return a nil error
	return nil
}

// GetProfile get a profile using an id provided by the caller
func (svc *inmemService) GetProfile(ctx context.Context, id string) (Profile, error) {
	// Get the Read lock from the inmemService struct
	svc.mtx.RLock()

	// Immediately set up a lock release to occur when the function finishes
	defer svc.mtx.RUnlock()

	// Look for the profile by the `id` function param
	profile, ok := svc.profiles[id]

	// Check if the profile id was not found in the datastore
	if !ok {

		// Return an empty profile and an error informing the caller that the profile was not found
		return Profile{}, ErrNotFound

	}

	// Return the profile to the caller and a nil error
	return profile, nil

}

// PutProfile creates or recreates a profile with a given id
func (svc *inmemService) PutProfile(ctx context.Context, id string, profile Profile) error {

	// Check if the id provided by caller fails to match the ID on the Profile value passed in
	if id != profile.ID {
		// Return a canned error message for mismatched IDs
		return ErrInconsistentIDs
	}

	// Get a Lock on the svc for atomic access to the datastore
	svc.mtx.Lock()

	// Immediately set up an unlock to occur when function exits
	defer svc.mtx.Unlock()

	// Create or replace the profile in our datastore
	svc.profiles[id] = profile

	// Return a nil value for the error
	return nil
}

// PatchProfile will only update a name or addresses field of a profile
func (svc *inmemService) PatchProfile(ctx context.Context, id string, profile Profile) error {

	// Check the profile for an ID field AND check if the
	if profile.ID != "" && id != profile.ID {
		return ErrInconsistentIDs
	}

	// Get a Lock on the svc for atomic access to the datastore
	svc.mtx.Lock()

	// Immediately set up an unlock to occur when function exits
	defer svc.mtx.Unlock()

	// Find profile in the datastore
	existing, ok := svc.profiles[id]

	// If we couldn't find the profile in the datastore
	if !ok {

		// Return a canned error informing the caller that the profile was not found
		return ErrNotFound
	}

	// Check to see if the profile name is not at it's empty value
	if profile.Name != "" {

		// Set the profile name to the existing copy of the name
		existing.Name = profile.Name
	}

	// Check to see if the slice of addresses is not at it's empty value
	if len(profile.Addresses) > 0 {

		// Set the profile addresses to the existing copy of the addresses
		existing.Addresses = profile.Addresses
	}

	// Set the profile in the datastore to our working copy (`existing`)
	svc.profiles[id] = existing

	// Return a nil error value
	return nil
}

func (svc *inmemService) DeleteProfile(ctx context.Context, id string) error {

	// Get a Lock on the svc for atomic access to the datastore
	svc.mtx.Lock()

	// Immediately set up an unlock to occur when function exits
	defer svc.mtx.Unlock()

	// Check to see if the profile is in our datastore.
	if _, ok := svc.profiles[id]; !ok {

		// Return canned error informing caller that the profile was not found
		return ErrNotFound
	}

	// From the profiles datastore, delete the profile with the passed in ID
	delete(svc.profiles, id)

	// Return a nil value for error
	return nil
}

// GetAddresses will return a slice of addresses associated with a specified Profile. Profile is specified by ID.
func (svc *inmemService) GetAddresses(ctx context.Context, profileID string) ([]Address, error) {

	// Get a Read Lock on the svc for atomic read access  to the datastore
	svc.mtx.RLock()

	// Immediately set up a lock release to occur when the function finishes
	defer svc.mtx.RUnlock()

	// Check to make sure there is a profile that corresponds to the passed in profile and save the found profile to a profile variable
	profile, ok := svc.profiles[profileID]

	// If no profile was found for the passed in ID
	if !ok {

		// Return error informing the caller that the profile to which the addresses should have been associated was not found
		return nil, ErrNotFound
	}

	// Return all addresses associated with the profile that was passed in and a nil error value
	return profile.Addresses, nil
}

// Get address will return a single address associated to a single profile by it's ID
func (svc *inmemService) GetAddress(ctx context.Context, profileID string, addressID string) (Address, error) {

	// Get a Read Lock on the svc for atomic read access  to the datastore
	svc.mtx.RLock()

	// Immediately set up a lock release to occur when the function finishes
	defer svc.mtx.RUnlock()

	// Check the data store to make sure the requested profile exists and set
	profile, ok := svc.profiles[profileID]

	// If no entry for the profile was fund in the datastore
	if !ok {

		// Return an empty valued Address and an error informing the caller that no profile was found with the provided ID.
		return Address{}, ErrNotFound
	}

	// Loop through each address attached to the found profile
	for _, address := range profile.Addresses {

		// Check to see if the current address's ID matches the addressID passed in
		if address.ID == addressID {

			// Return that address and a nil error for a value
			return address, nil
		}
	}

	// Return an empty Address value and a not found error since we were unable to find the specified address.
	return Address{}, ErrNotFound
}

// PostAddress will save a passed in Address to a specified Profile.
func (svc *inmemService) PostAddress(ctx context.Context, profileID string, address Address) error {

	// Get a Lock on the svc for atomic access to the datastore
	svc.mtx.Lock()

	// Immediately set up a lock release to occur when the function finishes
	defer svc.mtx.Unlock()

	// Check the datastore for the specified profile and set it to a profile variable when found
	profile, ok := svc.profiles[profileID]

	// If the profile was not found in the datastore
	if !ok {

		// Return an error to the caller informing them that the specified profile could not be found
		return ErrNotFound
	}

	// Iterate over the profile's existent addresses
	for _, existingAddress := range profile.Addresses {

		// Check to see if the current existent address has an ID that matches the ID of the address passed in
		if existingAddress.ID == address.ID {

			// Return an error informing the caller that address they wanted to add is already present
			return ErrAlreadyExists
		}
	}

	// Add the new address to the profile Address collection
	profile.Addresses = append(profile.Addresses, address)

	// Replace the profile with the one that has the new Address added to the Address collection
	svc.profiles[profileID] = profile

	// Return a nil error value
	return nil
}

func (svc *inmemService) DeleteAddress(ctx context.Context, profileID string, addressID string) error {

	// Get a Lock on the svc for atomic access to the datastore
	svc.mtx.Lock()

	// Immediately set up a lock release to occur when the function finishes
	defer svc.mtx.Unlock()

	// Check the datastore for the specified Profile and set it to a variable (`profile`) when found.
	profile, ok := svc.profiles[profileID]

	// If no profile was found in the datastore
	if !ok {

		// Return an error informing the caller that the specified profileID does not exist in the datastore
		return ErrNotFound
	}

	// Create a new Address collection. This Address collection will be saved to datastore on the Profile.
	newAddresses := make([]Address, 0, len(profile.Addresses))

	// Loop over each address in the profile address collection
	for _, address := range profile.Addresses {

		// Check to see if the current address's ID matches the addressID passed in
		if address.ID == addressID {
			// Delete the address!
			// NOTE: By continuing on to next iteration without adding the address to the
			// - new Address collection we are EFFECTIVELY deleting the entity.
			continue
		}

		// Store the current address to new Address collection
		newAddresses = append(newAddresses, address)
	}

	// Check to see if the two address collections are the same size
	if len(profile.Addresses) == len(newAddresses) {

		// Return an error informing the user that no element was deleted because the specifed addressID was not found
		return ErrNotFound
	}

	// Replace the profile's current Address collection with the new Address collection (`newAddresses`)
	profile.Addresses = newAddresses

	// Save the profile to the datastore
	svc.profiles[profileID] = profile

	// Return a nil error value
	return nil
}
