package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type Hotel struct {
	ID       primitive.ObjectID   `bson:"_id,omitempty" json:"id,omitempty"` // omitempty for dont show id if is empty
	Name     string               `bson:"name" json:"name"`
	Location string               `bson:"location" json:"location"`
	Rooms    []primitive.ObjectID `bson:"rooms" json:"rooms"`
	Rating   int                  `bson:"rating" json:"rating"`
}

type Room struct {
	ID primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"` // omitempty for dont show id if is empty
	// Type RoomType           `bson:"type" json:"type"`
	// small , normal , large
	Size    string             `bson:"size" json:"size"`
	SeaSide bool               `bson:"seadside" json:"seadside"`
	Price   float64            `bson:"price" json:"price"`
	HotelID primitive.ObjectID `bson:"hotelID" json:"hotelID"`
}
