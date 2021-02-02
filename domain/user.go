/*
	Package domain はドメインのエンティティです。
	対象とするドメインの実態を反映させます。
*/
package domain

//go:generate mockgen -source=$GOFILE -destination=../mock/$GOPACKAGE/mock_$GOFILE -package=mock_$GOPACKAGE

type User struct {
	ID          int    `json:"id"`
	FirstName   string `json:"first_name" binding:"required"`
	LastName    string `json:"last_name"  binding:"required"`
	Email       string `json:"email"      binding:"required,email"`
	PhoneNumber string `json:"phone_number" binding:"required,min=9,max=15"`
}
