package main
//Executable commands must always use package main.

//An import path is a string that uniquely identifies a package. A package's import path corresponds to its location inside a workspace or in a remote repository 
import (
    "encoding/json"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)


/*
 * 1.使用 Gorilla Web Toolkit 的 mux router.
 * 2.自定資料結構 (無DB)
 * 3.參考 https://www.codementor.io/codehakase/building-a-restful-api-with-golang-a6yivzqdo
 * 4. URL 127.0.0.1:8080/people
 * 5. POST 數據範例{"firstname":"Jack","lastname":"Cheng", "address":{"city":"台中","state":"台灣"}}
 *	> 須自行指定 /{id}
 * 
 * 6. $ go build && ./rest-api
 */


// The person Type (more like an object)
type Person struct {
    ID        string   `json:"id,omitempty"`
    Firstname string   `json:"firstname,omitempty"`
    Lastname  string   `json:"lastname,omitempty"`
    Address   *Address `json:"address,omitempty"`
}
type Address struct {
    City  string `json:"city,omitempty"`
    State string `json:"state,omitempty"`
}

var people []Person


/*
 *
 */
// Display all from the people var
func GetPeople(w http.ResponseWriter, r *http.Request) {
    json.NewEncoder(w).Encode(people)
}

// Display a single data
func GetPerson(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for _, item := range people {
        if item.ID == params["id"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    json.NewEncoder(w).Encode(&Person{})
}

// create a new item
func CreatePerson(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var person Person
    _ = json.NewDecoder(r.Body).Decode(&person)
    person.ID = params["id"]
    people = append(people, person)
    json.NewEncoder(w).Encode(people)
}

// Delete an item
func DeletePerson(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for index, item := range people {
        if item.ID == params["id"] {
            people = append(people[:index], people[index+1:]...)
            break
        }
    }
    //原程式在for內, 會輸出多次! (我改移到 for外)
    json.NewEncoder(w).Encode(people)
}


// main function to boot up everything
func main() {
    people = append(people, Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &Address{City: "City X", State: "State X"}})
    people = append(people, Person{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &Address{City: "City Z", State: "State Y"}})
    people = append(people, Person{ID: "3", Firstname: "Francis", Lastname: "Sunday"})

    //
    router := mux.NewRouter()
    router.HandleFunc("/people", GetPeople).Methods("GET")
    router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
    router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
    router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")
    log.Fatal(http.ListenAndServe(":8080", router))
}




