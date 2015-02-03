package main 

import (
  "fmt"
  "net/http"
  "encoding/json"
  "time"
  "os"
  
  "gopkg.in/mgo.v2"
  "gopkg.in/mgo.v2/bson"
  "github.com/gorilla/mux"
)

type Form struct {
  Id bson.ObjectId    `bson:"_id"`
  Name string         `bson:"name", omitempty`
  Fields []Field      `bson:"fields"`
  CreatedAt time.Time `bson:"created_at"`
}

type Field struct {
  Id bson.ObjectId `bson:"_id"`
  Label string     `bson:"label"`
  Name string      `bson:"name"`
  Required bool    `bson:"required"`
  Pattern string   `bson:"pattern"`
  Type string      `bson:"type"`
}

var connectionString string = os.Getenv("MONGO_CONNECTIONSTRING")
var mongoSession *mgo.Session

func main(){
  
  r := mux.NewRouter()
  mongoSession, _ = mgo.Dial(connectionString)
  defer mongoSession.Close()

  /* GET */
  r.HandleFunc("/", Get).Methods("GET")
  r.HandleFunc("/{id}", GetOne).Methods("GET")
  

  // POST.
  r.HandleFunc("/", Post).Methods("POST")

  /* PUT 
  r.HandleFunc("/{key}/", PUT).Methods("PUT")
  */
  
  /* DELETE 
  r.HandleFunc("/{key}/", DELETE).Methods("DELETE")
  */
  
  http.Handle("/", r)
  
  /* Listen on port 3000 */
  http.ListenAndServe(":3000", nil)
}

func Get(w http.ResponseWriter, r *http.Request){
  
  forms := []Form{}
  collection := mongoSession.DB("formbuilder").C("forms")
  iterator := collection.Find(bson.M{"name": bson.M{"$ne":""}}).Limit(100).Iter()
  err := iterator.All(&forms)
  
  if err != nil {
    fmt.Println(w, err)
  }
  
  data, _ := json.Marshal(forms)
  
  /* Should be done in middleware, but Im not too familiar with creating middleware yet */
  w.Header().Set("Access-Control-Allow-Origin", "*")
  w.Write(data)
}

func GetOne(w http.ResponseWriter, r *http.Request){
  
  vars := mux.Vars(r)
  id := vars["id"]
  collection := mongoSession.DB("formbuilder").C("forms")
  form := Form{}
  err := collection.FindId(bson.ObjectIdHex(id)).One(&form)
  
  if err != nil {
    fmt.Println(w, err)
  }
  
  data, _ := json.Marshal(form)
  
  /* Should be done in middleware, but Im not too familiar with creating middleware yet */
  w.Header().Set("Access-Control-Allow-Origin", "*")
  w.Write(data)
}

func Post(w http.ResponseWriter, r *http.Request){
  form := new(Form)
  form.Id = bson.NewObjectId()
  form.CreatedAt = time.Now()
  json.NewDecoder(r.Body).Decode(form)
  
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusCreated)
  
  data, _ := json.Marshal(form)
  
  collection := mongoSession.DB("formbuilder").C("forms")
  err := collection.Insert(form)
  if err != nil {
    fmt.Printf("Can't insert document: %v\n", err)
  }

  w.Header().Set("Access-Control-Allow-Origin", "*")
  w.Write(data)
}