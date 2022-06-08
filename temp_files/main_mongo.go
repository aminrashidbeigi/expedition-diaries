package mainmongo

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type stop struct {
	Country string `bson:"country"`
	City    string `bson:"city"`
	Year    string `bson:"year"`
}

type travel struct {
	ID        primitive.ObjectID `bson:"_id"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
	Traveler  string             `bson:"traveler"`
	Stops     []stop             `bson:"stops"`
	Resources []string           `bson:"resources"`
}

var collection *mongo.Collection
var ctx = context.TODO()

func getTravels(c *gin.Context) {
	var travels []travel
	cursor, err := collection.Find(ctx, bson.D{{}})
	if err != nil {
		panic(err)
	}
	if err = cursor.All(ctx, &travels); err != nil {
		panic(err)
	}
	fmt.Println(travels)
	c.IndentedJSON(http.StatusOK, travels)
}

func addTravel(c *gin.Context) {
	var newTravel travel
	if err := c.BindJSON(&newTravel); err != nil {
		return
	}

	newTravel.ID = primitive.NewObjectID()
	newTravel.CreatedAt = time.Now()
	newTravel.UpdatedAt = time.Now()

	_, err := collection.InsertOne(ctx, newTravel)
	if err != nil {
		log.Fatal(err)
	}
	c.IndentedJSON(http.StatusCreated, newTravel)
}

func initDB() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database("history-travels").Collection("travels")
}

func main() {
	initDB()
	router := gin.Default()
	router.GET("/travels", getTravels)
	router.POST("/add-travel", addTravel)
	router.Run("localhost:8080")
}
