package model

import "time"

type PropertyListing struct {
	ID                int       `xorm:"id" json:"id"`
	ListingID         int       `xorm:"listing_id" json:"listing_id"`
	PropertyType      int       `xorm:"property_type" json:"property_type,omitempty"`
	ProjectID         int       `xorm:"project_id" json:"project_id"`
	ListingType       int       `xorm:"listing_type" json:"listing_type"`
	Status            string    `xorm:"status" json:"status"`
	TotalSpace        float64   `xorm:"total_space" json:"total_space"`
	CoveredSpace      float64   `xorm:"covered_space" json:"covered_space"`
	Bathrooms         int       `xorm:"bathrooms_number" json:"bathrooms_number"`
	Bedrooms          int       `xorm:"bedrooms_number" json:"bedrooms_number"`
	NumFloor          int       `xorm:"num_floor" json:"num_floor"`
	PropertyFloor     int       `xorm:"property_floor" json:"property_floor"`
	MaxGuests         int       `xorm:"max_guests" json:"max_guests"`
	InternalID        string    `xorm:"internal_id" json:"internal_id"`
	MainImage         string    `xorm:"main_image" json:"main_image"`
	YearBuild         int       `xorm:"year_build" json:"year_build"`
	Width             float64   `xorm:"width" json:"width"`
	Length            float64   `xorm:"length" json:"length"`
	MinPrice          int       `xorm:"min_price" json:"min_price"`
	MaxPrice          int       `xorm:"max_price" json:"max_price"`
	LocationID        int       `xorm:"location_id" json:"location_id"`
	OwnerID           int       `xorm:"owner_id" json:"owner_id"`
	Description       string    `xorm:"description" json:"description"`
	Title             string    `xorm:"title" json:"title"`
	Price             int       `xorm:"price" json:"price"`
	ActivatedAt       string    `xorm:"activated_at" json:"activated_at,omitempty"`
	IsActivated       int       `xorm:"is_activated" json:"is_activated"`
	SharingPercentage int       `xorm:"sharing_percentage" json:"sharing_percentage"`
	Private           bool      `xorm:"private" json:"private"`
	AvailableNow      int       `xorm:"available_now" json:"available_now"`
	CreatedAt         time.Time `xorm:"created_at" json:"created_at"`
	CreatedBy         int       `xorm:"created_by" json:"created_by"`
	UpdatedAt         time.Time `xorm:"updated_at" json:"updated_at"`
	UpdatedBy         int       `xorm:"updated_by" json:"updated_by"`
	ListingByAdmin    bool      `xorm:"listing_by_admin" json:"listing_by_admin"`
}

type MinePropertyResult struct {
	ID                int       `xorm:"id" json:"id"`
	ListingID         int       `xorm:"listing_id" json:"listing_id"`
	PropertyType      int       `xorm:"property_type" json:"property_type,omitempty"`
	ProjectID         int       `xorm:"project_id" json:"project_id"`
	ListingType       int       `xorm:"listing_type" json:"listing_type"`
	Status            string    `xorm:"status" json:"status"`
	MainImage         string    `xorm:"main_image" json:"main_image"`
	TotalSpace        float64   `xorm:"total_space" json:"total_space"`
	Distance          float64   `xorm:"distance" json:"distance"`
	Bathrooms         int       `xorm:"bathrooms_number" json:"bathrooms_number"`
	Bedrooms          int       `xorm:"bedrooms_number" json:"bedrooms_number"`
	SharingPercentage int       `xorm:"sharing_percentage" json:"sharing_percentage"`
	Title             string    `xorm:"title" json:"title"`
	Price             int       `xorm:"price" json:"price"`
	CreatedBy         int       `xorm:"created_by" json:"created_by"`
	CreatedAt         time.Time `xorm:"created_at" json:"created_at"`
	ListingByAdmin    bool      `xorm:"listing_by_admin" json:"listing_by_admin"`
}

type PropertyListingDetail struct {
	Listing PropertyListing `json:"listing"`
	Images  []Image         `json:"images"`
}

// IsExists struct
func (m PropertyListing) IsExists() (ok bool) {
	if m.ListingID != 0 {
		ok = true
	}
	return
}
