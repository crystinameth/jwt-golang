package helpers

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/crystinameth/jwt-golang/database"
	jwt "github.com/dgrijalva/jwt-go"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SignedDetails struct{
	Email			string
	First_Name		string
	Last_Name		string
	Uid				string
	User_Type		string
	jwt.StandardClaims
}

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")

// can define the secret key in .env for major projects
var SECRET_KEY string = os.Getenv("SECRET_KEY")

func GenerateAllTokens(email string, firstname string, lastname string, usertype string, uid string ) (signedToken string, signedRefreshToken string, error) {
	claims := &SignedDetails{
		Email : email,
		First_Name: firstname,
		Last_Name: lastname,
		Uid: uid,
		User_Type: usertype,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	refreshClaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(SECRET_KEY))

	if err != nil {
		log.Panic(err)
		return
	}
	return token, refreshToken, err
}

func UpdateAllTokens