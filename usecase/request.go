package usecase

type PatchUser struct {
	FirstName   *string `json:"first_name"`
	LastName    *string `json:"last_name"`
	Email       *string `json:"email"        binding:"omitempty,email"`
	PhoneNumber *string `json:"phone_number" binding:"omitempty,min=9,max=15"`
}
