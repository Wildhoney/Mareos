package main

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
)

/**
 * @struct Mareos
 */
type Mareos struct { }

/**
 * @struct Collection
 */
type Collection struct {
	Name string `json:"name"`
	Friends []string `json:"friends"`
}

/**
 * @method Map
 * @return void
 */
func (m Mareos) Map() string {
	return "Map"
}

/**
 * @method Group
 * @return void
 */
func (m Mareos) Group() string {
	return "Group"
}

/**
 * @method Reduce
 * @return void
 */
func (m Mareos) Reduce() string {
	return "Reduce"
}

/**
 * @method main
 * @return void
 */
func main() {

    m := Mareos{}
    fmt.Println("Map: ", m.Map())

    // Read and parse the document as JSON.
    data, _   := ioutil.ReadFile("Facebook.json")
    dataModel := make([]Collection, 0)
	json.Unmarshal(data, &dataModel)

	// What do we have, sunshine?
	fmt.Printf("%d People - %#v", len(dataModel), dataModel)

}