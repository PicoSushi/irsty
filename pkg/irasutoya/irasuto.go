package irasutoya

// Irasuto describes a Illustration of Irasutoya.
// Irasuto has only and always one Illustration.
type Irasuto struct {
	Title    string `json:"title"`
	ImageURL string `json:"image_url"`
}
