package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Source struct {
	ID            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	DeveloperName string             `json:"devname,omitempty" bson:"devname,omitempty"`
	Email         string             `json:"email,omitempty" bson:"email,omitempty"`
	SourceLink    string             `json:"sourcelink,omitempty" bson:"sourcelink,omitempty"`
	Ticket        string             `json:"ticket,omitempty" bson:"ticket,omitempty"`
	Timestamp     time.Time          `json:"timestamp,omitempty" bson:"timestamp,omitempty"`
}
type Deployment struct {
	SourceLink      string `json:"sourcelink,omitempty" bson:"sourcelink,omitempty"`
	DestinationLink string `json:"destinationlink,omitempty" bson:"destinationlink,omitempty"`
}
