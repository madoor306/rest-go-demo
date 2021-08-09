package dto

// Define UserDTO Object
type UserDTO struct {
	UserId    string          `json:"user_id"`
	FirstName string          `json:"first_name"`
	LastName  string          `json:"last_name"`
	Age       int             `json:"user_age"`
	Address   *UserAddressDTO `json:"address"`
}
