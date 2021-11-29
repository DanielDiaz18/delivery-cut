package data

import (
	"time"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Order struct {
	mgm.DefaultModel `bson:",inline"`
	Route            primitive.ObjectID `bson:"route" json:"route"`
	User             primitive.ObjectID `bson:"user" json:"user"`
	Parner           primitive.ObjectID `bson:"parner" json:"parner"`
	Status           string             `bson:"status" json:"status"`
	startedAt        time.Time          `bson:"stated_at" json:"startedAt"`
	finishedAt       time.Time          `bson:"finished_at" json:"finishedAt"`
}

type OrderForm struct {
	Route       primitive.ObjectID `bson:"route" json:"route" validate:"required"`
	Parner      primitive.ObjectID `bson:"parner" json:"parner" validate:"required"`
	Origin      Position           `bson:"origin" json:"origin" validate:"required"`
	Destination Position           `bson:"destination" json:"destination" validate:"required"`
}

type Orders []*Order

type OrderRepo struct {
	c *mgm.Collection
}

func NewOrderRepo() *OrderRepo {
	o := &Order{}
	c := mgm.Coll(o)
	c.Collection.Indexes().CreateOne(mgm.Ctx(), mongo.IndexModel{
		Keys: bson.M{
			"user_id": 1,
		},
		Options: options.Index().SetUnique(true),
	})
	return &OrderRepo{c: c}
}

func (r *OrderRepo) Create(o *Order) error {
	return r.c.Create(o)
}

func (r *OrderRepo) GetOrders(o Orders) error {
	return r.c.SimpleFind(o, bson.M{})
}
