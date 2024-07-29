package types

type User struct {
	ID        string `json:"id" bson:"_id,omitempty"`
	FirstName string `json:"firstName" bson:"firstName"`
	LastName  string `json:"lastName" bson:"lastName"`
}
