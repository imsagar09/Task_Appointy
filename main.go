package main

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
    "context"
    
   "time"

   "go.mongodb.org/mongo-driver/bson"
   "go.mongodb.org/mongo-driver/mongo"
   "go.mongodb.org/mongo-driver/mongo/options"
   "go.mongodb.org/mongo-driver/mongo/readpref"
    
)

type Participants struct {
    Name     string `json:"Name"`
    Email string `json:"Email"`
    RSVP string `json:"Rsvp"`
}
type Article struct {
    Id      string `json:"Id"`
    Title string `json:"Title"`
    
    Start Time string `json:"Stt"`
    End Time string `json:"Edt"`
    Creation Time Stamp string `json:"Crt"`
    details Participants
}



var Articles []Article
func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    
}

func handleRequests() {
    myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/", homePage)
    myRouter.HandleFunc("/articles", returnAllArticles)
    myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
    // add our new DELETE endpoint here
    myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
    myRouter.HandleFunc("/article/{id}", returnSingleArticle)
    log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
    client, err := mongo.NewClient(options.Client().ApplyURI(“<<MongoDB Connection URI>>))
    if err != nil {
        log.Fatal(err)
    }
    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
    err = client.Connect(ctx)
    if err != nil {
            log.Fatal(err)
    }
    defer client.Disconnect(ctx)

    Articles = []Article{
        Article{Id: "1", Title: "Hello", Stt: "12:12:12", Edt: "13:13:13",Crt: "19-10-20",Name:"Sagar",Email:"gup@gmail.com"},
        Article{Id: "2", Title: "Hello1", Stt: "14:14:14", Edt: "15:15:15",Crt: "19-10-20",Name:"Sagar Gupta",Email:"gupta@gmail.com"},
    }
    handleRequests()
}

func returnAllArticles(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: returnAllArticles")
    json.NewEncoder(w).Encode(Articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    key := vars["id"]

    // Loop over all of our Articles
    // if the article.Id equals the key we pass in
    // return the article encoded as JSON
    for _, article := range Articles {
        if article.Id == key {
            json.NewEncoder(w).Encode(article)
        }
    }
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
    // get the body of our POST request
    // unmarshal this into a new Article struct
    // append this to our Articles array.    
    reqBody, _ := ioutil.ReadAll(r.Body)
    var article Article 
    json.Unmarshal(reqBody, &article)
    // update our global Articles array to include
    // our new Article
    Articles = append(Articles, article)

    json.NewEncoder(w).Encode(article)
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
    // once again, we will need to parse the path parameters
    vars := mux.Vars(r)
    // we will need to extract the `id` of the article we
    // wish to delete
    id := vars["id"]

    // we then need to loop through all our articles
    for index, article := range Articles {
        // if our id path parameter matches one of our
        // articles
        if article.Id == id {
            // updates our Articles array to remove the 
            // article
            Articles = append(Articles[:index], Articles[index+1:]...)
        }
    }

}
func InsertPost(title string, body string) {

post := Post{title, body}

collection := client.Database(“my_database”).Collection(“posts”)

insertResult, err := collection.InsertOne(context.TODO(), post)



if err != nil {

log.Fatal(err)

}



fmt.Println(“Inserted post with ID:”, insertResult.InsertedID)

}
func GetPost(id bson.ObjectId) {

    collection := client.Database(“my_database”).Collection(“posts”)
    
    filter := bson.D
    
    var post Post
    
    
    
    err := collection.FindOne(context.TODO(), filter).Decode(&post)
    
    if err != nil {
    
    log.Fatal(err)
    
    }
    
    
    
    fmt.Println(“Found post with title “, post.Title)
    
    }
    


