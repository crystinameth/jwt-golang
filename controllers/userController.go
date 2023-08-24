package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/crystinameth/jwt-golang/database"
	helper "github.com/crystinameth/jwt-golang/helpers"
	"github.com/crystinameth/jwt-golang/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection = database.OpenCollection(database.Client, "user")
var validate = validator.New()

func HashPassword(){

}

func VerifyPassword(){

}

func Signup(){

}

func Login(){

}

func GetUsers(){

}

func GetUser(){

}