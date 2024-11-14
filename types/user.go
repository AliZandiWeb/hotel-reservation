package types

type User struct {
	Id        string `bson:"_id,omitempty" json:"id,omitempty"` // omitempty for dont show id if is empty
	FirstName string `bson:"firstName" json:"firstName"`
	LastName  string `bson:"lastName" json:"lastName"`
}
