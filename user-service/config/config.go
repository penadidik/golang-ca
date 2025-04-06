package config

import "os"

func GetMongoURI() string {
    uri := os.Getenv("MONGODB_URI")
    if uri == "" {
        uri = "mongodb://localhost:27017"
    }
    return uri
}

