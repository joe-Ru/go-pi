package mongo_service

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"

	//"fmt"
	"io/ioutil"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoVars struct {
	UserName string
	Database string
	Host string
}

type Articles struct {
	ID    primitive.ObjectID `bson:"_id"`
	Title string `bson:"title"`
	Link  string `bson:"link"`
}


// Read the host from the JSON file
func MongoGetHostFromJson() string{

	readContent, err := ioutil.ReadFile("./mongo_vars.json")

	if err != nil {
		log.Fatal("Error opening file: ", err)
	}
	var mongoVars MongoVars

	err = json.Unmarshal(readContent, &mongoVars)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	var mongoHost = mongoVars.Host
	return mongoHost

}

func GetArticlesFromCollection(collectionDocument string)  []Articles {

	log.Println("collection document passed:", collectionDocument)

	var connectionHost = MongoGetHostFromJson()
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(connectionHost).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	coll := client.Database("local").Collection(collectionDocument)
	cursor, err := coll.Find(context.TODO(), bson.M{})

	var articles []Articles
	var article Articles
	for cursor.Next(context.TODO()) {
		if err := cursor.Decode(&article); err != nil {
			log.Fatal(err)
		}
		fmt.Println(article.ID)
		articles = append(articles, article)
	}
	return articles
}

func GetOneArticleFromCollection(collectionDocument string, id string)  Articles {

	var connectionHost = MongoGetHostFromJson()
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(connectionHost).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()


	var article Articles

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.D{{"_id", objectId}}
	coll := client.Database("local").Collection(collectionDocument)
	err = coll.FindOne(context.TODO(), filter).Decode(&article)
	if err != nil {
		panic(err)
	}

	return article

}

func DeleteOneArticleFromCollection(collectionDocument string, id string) {

	var connectionHost = MongoGetHostFromJson()
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(connectionHost).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	objectId, _ := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.D{{"_id", objectId}}
	coll := client.Database("local").Collection(collectionDocument)
	result, err := coll.DeleteOne(context.TODO(), filter)
	log.Println(result)
	if err != nil {
		panic(err)
	}

}

func AddArticleFromCollection(collectionDocument string, title string, link string) {

	var connectionHost = MongoGetHostFromJson()
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(connectionHost).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	//objectId, err := primitive.ObjectIDFromHex(id)
	//if err != nil {
	//	log.Fatal(err)
	//}

	filter := bson.D{
		{"title", title},
		{"link", link},
	}

	coll := client.Database("local").Collection(collectionDocument)
	result, err := coll.InsertOne(context.TODO(), filter)
	log.Println(result)
	if err != nil {
		panic(err)
	}

}
