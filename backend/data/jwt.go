package data

import (
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CustomClaims struct {
	jwt.StandardClaims
	Id   primitive.ObjectID `json:"id"`
	Kind string             `json:"kind"`
}
