package data

import (
	"fmt"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string `json:"name" bson:"name"`
	LastName         string `json:"lastName" bson:"last_name"`
	Password         string `json:"password" bson:"password"`
	Email            string `json:"email" bson:"email"`
	PhoneNumber      string `json:"phoneNumber" bson:"phone_number"`
	Kind             string `json:"kind" bson:"kind"`
}

type LoginForm struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

type SignUpForm struct {
	Name        string `json:"name" validate:"required"`
	LastName    string `json:"lastName" validate:"required"`
	Password    string `json:"password" validate:"required,min=6"`
	Email       string `json:"email" validate:"required,email"`
	PhoneNumber string `json:"phoneNumber" validate:"required,min=10"`
	Kind        string `json:"kind" validate:"required,oneof=admin user parner"`
}

type UsersRepo struct {
	c *mgm.Collection
}

func NewUsersRepo() *UsersRepo {
	u := &User{}
	c := mgm.Coll(u)
	c.Collection.Indexes().CreateOne(mgm.Ctx(), mongo.IndexModel{
		Keys: bson.M{
			"email": 1,
		},
		Options: options.Index().SetUnique(true),
	})
	return &UsersRepo{c}
}

func (r *UsersRepo) Login(email, password string, u *User) error {
	if err := r.c.First(bson.M{"email": email}, u); err != nil {
		return fmt.Errorf("usuario o contaseña incorrectos")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)); err != nil {
		return fmt.Errorf("usuario o contaseña incorrectos")
	}
	return nil
}

func (r *UsersRepo) Create(u *User) error {
	return r.c.Create(u)
}

func (r *UsersRepo) Get(id string) (*User, error) {
	u := &User{}
	err := r.c.FindByID(id, u)
	return u, err
}
