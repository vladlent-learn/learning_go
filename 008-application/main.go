package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type person struct {
	// tags are unnecessary here since all values map to the same name anyway
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

func main() {
	p := "password228"

	bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(p)
	fmt.Println(string(bs))

	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	fmt.Println(s)

	var fromJSON []person

	err = json.Unmarshal([]byte(s), &fromJSON)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(fromJSON)
}
