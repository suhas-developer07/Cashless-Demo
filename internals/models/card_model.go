package models

type Card struct {
	CardId      string  `json:"card_id"`
	RollNumber  string  `json:"roll_number"`
	StudentName string  `json:"strudent_name"`
	Balance     float64 `json:"balance"`
}

// type Vendor struct {
// 	VendorId  string `json:"vendor_id"`
// 	VendorName string  ``
// }