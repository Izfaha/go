package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	// gin/go-lang = this is our main model name in go.mod
	"gin/go-lang/models"
)

// dummy data for testing 
var users = []models.User{
	{ID: 1, Name: "Faiz Maulana Habibi"},
	{ID: 2, Name: "Anas"},
}