package hackernews

import (
	"context"
	"fmt"
	"log"
	"time"
	"with_coffee/lib/config"
	db "with_coffee/lib/mongo"

	"github.com/go-resty/resty/v2"
	"go.mongodb.org/mongo-driver/bson"
)

// Retrieve top stories ids from the Api
func GetTopStories() TopIds {
	cnf, _ := config.LoadConfig()
	url := fmt.Sprintf("%s/topstories.json", cnf.HackerNews.BaseUrl)

	var ids TopIds
	client := resty.New()

	client.R().SetResult(&ids).Get(url)

	return ids[:cnf.HackerNews.Limit]
}

// Retrieve a specific story from api based on story id
func GetStoryFromId(id int) Story {
	cnf, _ := config.LoadConfig()
	url := fmt.Sprintf("%s/item/%d.json", cnf.HackerNews.BaseUrl, id)

	var story Story
	client := resty.New()

	client.R().SetResult(&story).Get(url)

	return story
}

// Save a single story to mongo
func SaveStoryToMongo(story Story) {
	var existingStory Story

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	collection := db.MongoCollection(ctx, "hackernews")
	collection.FindOne(ctx, bson.D{{"id", story.ID}}).Decode(&existingStory)

	if existingStory.ID != 0 {
		log.Printf("Story with %d is already present on mongo. Skipping", existingStory.ID)
		return
	}

	_, err := collection.InsertOne(ctx, story)

	if err != nil {
		log.Printf("Could not save story with id %d to mongo. Error: %v\n", story.ID, err)

	}

	log.Default().Printf("Story with id %d has been saved to db\n", story.ID)
}

// Parent function to fetch stories from the api and save them to mongo
func ImportStories() []Story {
	log.Println("Getting the latest top stories from hackernews")
	Ids := GetTopStories()

	var stories []Story

	for _, id := range Ids {
		story := GetStoryFromId(id)

		SaveStoryToMongo(story)

		stories = append(stories, story)

	}
	return stories
}
