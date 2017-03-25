//
// This file generated by rdl 1.4.12
//

package hello

import (
	rdl "github.com/ardielle/ardielle-go/rdl"
)

var schema *rdl.Schema

func init() {
	sb := rdl.NewSchemaBuilder("hello")
	sb.Version(1)

	tGreeting := rdl.NewStructTypeBuilder("Struct", "Greeting")
	tGreeting.Field("message", "String", false, nil, "")
	sb.AddType(tGreeting.Build())

	rGetGreeting := rdl.NewResourceBuilder("Greeting", "GET", "/greeting")
	rGetGreeting.Input("name", "String", false, "sender", "", false, "human", "")
	sb.AddResource(rGetGreeting.Build())

	schema = sb.Build()
}

func HelloSchema() *rdl.Schema {
	return schema
}