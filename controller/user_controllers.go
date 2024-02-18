package user_controllers

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	db "github.com/sudeep162002/ims-go-backend/db"
	model "github.com/sudeep162002/ims-go-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/mongo"
)

func HelloUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello from user ",
	})
}

func GetUsers(c *gin.Context) {
	client := db.GetClient()
	var collection = client.Database("users").Collection("users")

	var users []model.User // assuming user data is stored as documents in MongoDB

	// Retrieve users from the database
	cursor, err := collection.Find(c.Request.Context(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve users",
		})
		return
	}
	defer cursor.Close(c.Request.Context())

	// Iterate over the cursor and decode documents into the users slice
	for cursor.Next(c.Request.Context()) {
		var user model.User
		if err := cursor.Decode(&user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to decode user",
			})
			return
		}
		users = append(users, user)
	}

	// Check for cursor errors
	if err := cursor.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Cursor error",
		})
		return
	}

	// Return the users as JSON response
	c.JSON(http.StatusOK, users)

}

func GetUsersById(c *gin.Context) {
	client := db.GetClient()
	collection := client.Database("users").Collection("users")

	userID := c.Param("id")

	// Convert userID to the appropriate type if necessary

	filter := bson.M{"userId": userID}

	cursor, err := collection.Find(c.Request.Context(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer cursor.Close(c.Request.Context())

	var familyUsers []model.User
	for cursor.Next(c.Request.Context()) {
		var user model.User
		if err := cursor.Decode(&user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		familyUsers = append(familyUsers, user)
	}

	if len(familyUsers) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No users found in the database."})
		return
	}

	c.JSON(http.StatusOK, familyUsers)
}
func InsertUser(c *gin.Context) {
	var user model.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request body."})
		return
	}

	if user.FullName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "fullName is required."})
		return
	}

	if user.UserID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Family Code is required."})
		return
	}

	if user.RitwickName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "ritwickName is required."})
		return
	}

	client := db.GetClient()
	collection := client.Database("users").Collection("users")

	existingUser := &model.User{}
	err := collection.FindOne(c.Request.Context(), bson.M{"fullName": user.FullName}).Decode(existingUser)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{"message": "User already exists."})
		return
	}

	_, err = collection.InsertOne(c.Request.Context(), user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Error inserting data.", "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data successfully inserted."})
}

func UpdateUser(c *gin.Context) {
	userID := c.Param("id")

	// Bind JSON body to modifiedUser struct
	var modifiedUser model.User
	if err := c.BindJSON(&modifiedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request body"})
		return
	}

	// Ensure that fullName is not empty
	if modifiedUser.FullName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Adding fullName is mandatory"})
		return
	}

	client := db.GetClient()
	collection := client.Database("users").Collection("users")

	// Fetch user corresponding to userID and modifiedUser.FullName
	var dbUser model.User
	filter := bson.M{"userId": userID, "fullName": modifiedUser.FullName}
	err := collection.FindOne(c.Request.Context(), filter).Decode(&dbUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error while getting db user", "details": err.Error()})
		return
	}

	// Update fields of dbUser with modifiedUser's fields
	updateFields(&dbUser, &modifiedUser)

	// Save updated dbUser in the database
	_, err = collection.ReplaceOne(c.Request.Context(), filter, dbUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating data in database", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data successfully updated"})
}

func updateFields(dbUser *model.User, modifiedUser *model.User) {
	// Get reflect.Value of dbUser and modifiedUser
	dbUserValue := reflect.ValueOf(dbUser).Elem()
	modifiedUserValue := reflect.ValueOf(modifiedUser).Elem()

	// Iterate over fields of modifiedUser and update corresponding fields in dbUser
	for i := 0; i < modifiedUserValue.NumField(); i++ {
		fieldName := modifiedUserValue.Type().Field(i).Name
		fieldValue := modifiedUserValue.Field(i)
		if fieldValue.Interface() != reflect.Zero(fieldValue.Type()).Interface() {
			dbUserValue.FieldByName(fieldName).Set(fieldValue)
		}
	}
}
