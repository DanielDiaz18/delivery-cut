package data

import "github.com/kamva/mgm/v3"

type Route struct {
	mgm.DefaultModel `bson:",inline"`
	Origin           Position `bson:"origin" json:"origin"`
	Destination      Position `bson:"destination" json:"destination"`
}
