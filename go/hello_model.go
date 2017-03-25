//
// This file generated by rdl 1.4.12
//

package hello

import (
	"encoding/json"
	"fmt"
	rdl "github.com/ardielle/ardielle-go/rdl"
)

var _ = rdl.Version
var _ = json.Marshal
var _ = fmt.Printf

//
// Greeting -
//
type Greeting struct {
	Message string `json:"message"`
}

//
// NewGreeting - creates an initialized Greeting instance, returns a pointer to it
//
func NewGreeting(init ...*Greeting) *Greeting {
	var o *Greeting
	if len(init) == 1 {
		o = init[0]
	} else {
		o = new(Greeting)
	}
	return o
}

type rawGreeting Greeting

//
// UnmarshalJSON is defined for proper JSON decoding of a Greeting
//
func (self *Greeting) UnmarshalJSON(b []byte) error {
	var r rawGreeting
	err := json.Unmarshal(b, &r)
	if err == nil {
		o := Greeting(r)
		*self = o
		err = self.Validate()
	}
	return err
}

//
// Validate - checks for missing required fields, etc
//
func (self *Greeting) Validate() error {
	if self.Message == "" {
		return fmt.Errorf("Greeting.message is missing but is a required field")
	} else {
		val := rdl.Validate(HelloSchema(), "String", self.Message)
		if !val.Valid {
			return fmt.Errorf("Greeting.message does not contain a valid String (%v)", val.Error)
		}
	}
	return nil
}