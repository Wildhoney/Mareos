package main

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
)

/**
 * @struct Mareos
 */
type Mareos struct {
    length, width int
}

/**
 * @struct Models
 */
type Models struct {
	Collection []Model
}

type Model struct {
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
 * @return 0
 */
func main() {

    m := Mareos{}
    fmt.Println("Map: ", m.Map())

    // Read and parse the document as JSON.
    data, _   := ioutil.ReadFile("Facebook.json")
    dataModel := make([]Model, 0)
	json.Unmarshal(data, &dataModel)

	// What do we have, sunshine?
	fmt.Printf("%#v", dataModel)

}