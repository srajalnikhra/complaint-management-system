package dto

// ComplaintEmailData holds details needed to compose a complaint notification email.
type ComplaintEmailData struct {
	UserName       string
	UserEmail      string
	ComplaintTitle string
	Status         string
}
