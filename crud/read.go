package crud

import (
	"context"
    "fmt"
    "go.mongodb.org/mongo-driver/bson"
    "kanolibrary/mongo"
    "encoding/json"
	"kanolibrary/util"
    "kanolibrary/models"
    
)

 func FindOne(collection string, query map[string]interface{}) ([]byte, error) {
 	ctx := context.Background()
     db, err := mongo.Connect()
     util.ErrorChecker(err)

     csr, err := db.Collection(collection).Find(ctx, query)
     util.ErrorChecker(err)
     defer csr.Close(ctx)

     result := make([]models.Books, 0)
     for csr.Next(ctx) {
         var row models.Books
         err := csr.Decode(&row)
         util.ErrorChecker(err)
 
         result = append(result, row)
 
     }
     data, err := json.Marshal(result)
     util.ErrorChecker(err)
     return data, err
 }
func FindAll(collection string) ([]byte, error) {

	ctx := context.Background()
    db, err := mongo.Connect()
    util.ErrorChecker(err)

    csr, err := db.Collection(collection).Find(ctx, bson.M{})
    util.ErrorChecker(err)
    defer csr.Close(ctx)

    result := make([]models.Books, 0)
    for csr.Next(ctx) {
        var row models.Books
        err := csr.Decode(&row)
        util.ErrorChecker(err)

        result = append(result, row)
    }
    fmt.Println(result)
    data, err := json.Marshal(result)
    util.ErrorChecker(err)
    return data, err
    
}

    

