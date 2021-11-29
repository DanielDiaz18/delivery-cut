package data

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Position = [2]float64

type Parner struct {
	mgm.DefaultModel  `bson:",inline"`
	UserId            primitive.ObjectID `bson:"user_id" json:"userId"`
	OrderId           primitive.ObjectID `bson:"order_id" json:"orderId"`
	Status            string             `bson:"status" json:"status"`
	LastKnownPosition Position           `bson:"last_known_position" json:"lastKnownPosition"`
}

type Parners []*Parner

type ParnerRepo struct {
	c *mgm.Collection
}

type ParnerForm struct {
	UserId primitive.ObjectID `json:"userId"`
}

func NewParnerRepo() *ParnerRepo {
	p := &Parner{}
	c := mgm.Coll(p)
	c.Collection.Indexes().CreateOne(mgm.Ctx(), mongo.IndexModel{
		Keys: bson.M{
			"user_id": 1,
		},
		Options: options.Index().SetUnique(true),
	})
	return &ParnerRepo{c: c}
}

func (r *ParnerRepo) Create(p *Parner) error {
	return r.c.Create(p)
}

func (r *ParnerRepo) UpdatePosition(id string, pos Position) error {
	p := &Parner{}
	err := r.c.First(bson.M{"_id": id}, p)
	if err != nil {
		return err
	}
	p.LastKnownPosition = pos
	return r.c.Update(p)
}

func (r *ParnerRepo) GetAll(p Parners) error {
	return r.c.SimpleFind(p, bson.M{})
}
