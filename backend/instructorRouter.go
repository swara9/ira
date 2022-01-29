package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// instructor registration
func instructorRegister(c *gin.Context) {
	// Parse input request
	type Req struct {
		Email    string `json: "email"`
		Password string `json: "password"`
	}
	req := Req{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "incorrect parameters",
		})
		return
	}
	// Check if the student with email already exists
	existingInstructor := Instructor{}
	result := DB.Where("email = ?", req.Email).First(&existingInstructor)
	if result.RowsAffected == 1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "user with email already exists",
		})
		return
	}
	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "interval server error",
		})
		return
	}
	// Insert into database
	newInstructor := Instructor{
		Email:    req.Email,
		Password: string(hashedPassword),
	}
	result = DB.Create(&newInstructor)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}
	c.JSON(http.StatusOK, newInstructor)
}

// Student login
func instructorLogin(c *gin.Context) {
	type Req struct {
		Email    string `json: "email"`
		Password string `json: "password"`
	}
	req := Req{}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "incorrect parameters",
		})
		return
	}
	// Check if the student exists
	instructor := Instructor{}
	res := DB.Where("email = ?", req.Email).First(&instructor)
	if res.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "user not found",
		})
		return
	}
	// Check if the password match
	err = bcrypt.CompareHashAndPassword([]byte(instructor.Password), []byte(req.Password))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "wrong password",
		})
		return
	}
	c.JSON(http.StatusOK, instructor)
}
