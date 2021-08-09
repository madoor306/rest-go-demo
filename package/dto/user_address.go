package dto

// Define UserAddressDTO
type UserAddressDTO struct {
	AddressId    string `json:"address_id"`
	AddressLine1 string `json:"address_line1"`
	AddressLine2 string `json:"address_line2"`
	Pincode      int    `json:"pincode"`
}
