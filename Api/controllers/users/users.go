package users

import (
	"log"

	util "github.com/Rajendro1/AccuKnox/Util"
	"github.com/Rajendro1/AccuKnox/config"
	pgdatabase "github.com/Rajendro1/AccuKnox/pgDatabase"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateUsers() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var input = users{}
		if err := c.ShouldBindJSON(&input); err == nil {
			hashPassword, hashPasswordErr := util.HashPassword(input.Password)
			if hashPasswordErr != nil {
				log.Println("password hashing error ", hashPasswordErr.Error())
				c.JSON(200, gin.H{
					"status":  config.INPUTERROR,
					"message": "password hashing error",
					"error":   hashPasswordErr.Error(),
				})
				return
			}
			if input.Email == "" {
				c.JSON(config.INPUTERROR, gin.H{
					"status":  config.INPUTERROR,
					"message": "Please give email id",
				})
				return
			}
			if !util.IsEmailValid(input.Email) {
				c.JSON(config.INPUTERROR, gin.H{
					"status":  config.INPUTERROR,
					"message": "Please give valid email",
				})
				return
			}
			if input.Password == "" {
				c.JSON(config.INPUTERROR, gin.H{
					"status":  config.INPUTERROR,
					"message": "Please give password",
				})
				return
			}
			if input.Name == "" {
				c.JSON(config.INPUTERROR, gin.H{
					"status":  config.INPUTERROR,
					"message": "Please give Name",
				})
				return
			}
			if !pgdatabase.VerifyUsersEmailToDatabase(input.Email) {
				createUsers, userID, createUsersErr := pgdatabase.CreateUsersToDatabase(input.Email, input.Name, hashPassword)
				if createUsers && createUsersErr == nil {
					getUsers, getUsersErr := pgdatabase.GetUsersByIdFromDatabase(userID)
					if getUsersErr != nil {
						c.JSON(config.SUCCESS, gin.H{
							"status":  config.DBERROR,
							"message": "Sorry! we can't create your users",
						})
					} else {
						c.JSON(config.SUCCESS, gin.H{
							"status":  config.CREATED,
							"message": "User created successfully",
							"data":    getUsers,
						})
					}
				} else if createUsersErr != nil {
					c.JSON(200, gin.H{
						"message": "Sorry! we can't create your users",
						"status":  config.DBERROR,
						"error":   createUsersErr.Error(),
					})
				}
			} else {
				c.JSON(config.DUPLICATE, gin.H{
					"status":  config.DUPLICATE,
					"message": "This email already registerd with us",
				})
			}
		} else {
			log.Println("input ERROR ", err.Error())
			c.JSON(config.INPUTERROR, gin.H{
				"status":  config.INPUTERROR,
				"message": "input error",
				"error":   err.Error(),
			})
		}
	}
	return gin.HandlerFunc(fn)
}
func PostLogin() gin.HandlerFunc {
	fn := func(c *gin.Context) {
		var input = users{}
		if err := c.ShouldBindJSON(&input); err == nil {
			if input.Email == "" {
				c.JSON(config.INPUTERROR, gin.H{
					"status":  config.INPUTERROR,
					"message": "Please give email id",
				})
				return
			}
			if !util.IsEmailValid(input.Email) {
				c.JSON(config.INPUTERROR, gin.H{
					"status":  config.INPUTERROR,
					"message": "Please give valid email",
				})
				return
			}
			if input.Password == "" {
				c.JSON(config.INPUTERROR, gin.H{
					"status":  config.INPUTERROR,
					"message": "Please give password",
				})
				return
			}
			if pgdatabase.VerifyUsersEmailToDatabase(input.Email) {
				getUsers, getUsersErr := pgdatabase.GetUsersByEmailFromDatabase(input.Email)
				if getUsersErr != nil {
					c.JSON(config.SUCCESS, gin.H{
						"status":  config.DBERROR,
						"message": "Sorry! we can't create your users",
					})
					return
				}
				dbPassword := getUsers.Password
				if util.CheckPassword(input.Password, dbPassword) {
					sessionID := uuid.New()
					_, updateSessionIDErr := pgdatabase.PatchSessionIDUsingEmail(input.Email, sessionID)
					if updateSessionIDErr != nil {
						c.JSON(config.SUCCESS, gin.H{
							"status":  config.DBERROR,
							"message": "Sorry! we can't updated the session id",
							"error":   updateSessionIDErr.Error(),
						})
					} else {
						c.JSON(config.SUCCESS, gin.H{
							"status":  config.SUCCESS,
							"message": "Log in successfully",
							"sid":     sessionID,
						})
					}
				} else {
					c.JSON(config.UNAUTHORIZE, gin.H{
						"status":  config.UNAUTHORIZE,
						"message": "Sorry! your password did not match with our database",
					})
				}

			} else {
				c.JSON(config.UNAUTHORIZE, gin.H{
					"status":  config.UNAUTHORIZE,
					"message": "Sorry! your email is not present in our database",
				})
				return
			}
		} else {
			log.Println("input ERROR ", err.Error())
			c.JSON(config.INPUTERROR, gin.H{
				"status":  config.INPUTERROR,
				"message": "input error",
				"error":   err.Error(),
			})
		}
	}
	return gin.HandlerFunc(fn)
}
