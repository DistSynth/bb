//go:generate go run -tags=dev data/assets_generate.go

package main

import (
	//"ab.ru/bigbrother/cmd"

	"ab.ru/bigbrother/cmd"
	_ "github.com/mailru/go-clickhouse"
)

// type Base struct {
// 	Date      time.Time   `json:"ts,omitempty" db:"country_code"`
// 	EventName string      `json:"et,omitempty" db:"et"`
// 	Payload   interface{} `json:"pl,omitempty" db:"-"`
// }

// const jsonStr = `{
// 	"ts":"2018-09-22T12:42:31Z",
// 	"et":"SomeEvent",
// 	"pl":{
// 		"url":"http://ab.ru",
// 		"session":"aabbccdd-aabbccdd",
// 		"cnt":1
// 	}
// }
// `

func main() {
	cmd.Execute()
	// var b Base
	// json.Unmarshal([]byte(jsonStr), &b)
	// fmt.Printf("%b\n", b)
	// switch v := b.Payload.(type) {
	// case int:
	// 	// v is an int here, so e.g. v + 1 is possible.
	// 	fmt.Println("Integer: ", v)
	// case float64:
	// 	// v is a float64 here, so e.g. v + 1.0 is possible.
	// 	fmt.Println("Float64: ", v)
	// case string:
	// 	// v is a string here, so e.g. v + " Yeah!" is possible.
	// 	fmt.Println("String: ", v)
	// case map[string]interface{}:

	// 	var fnames []string
	// 	var fvals []interface{}

	// 	for key, val := range v {
	// 		fmt.Println("\tKey: ", key)
	// 		switch f := val.(type) {
	// 		case int64:
	// 			// v is an int here, so e.g. v + 1 is possible.
	// 			fnames = append(fnames, key)
	// 			fvals = append(fvals, f)
	// 			fmt.Println("\tInteger: ", f)
	// 		case float64:
	// 			fnames = append(fnames, key)
	// 			fvals = append(fvals, f)
	// 			// v is a float64 here, so e.g. v + 1.0 is possible.
	// 			fmt.Println("\tFloat64: ", f)
	// 		case string:
	// 			fnames = append(fnames, key)
	// 			fvals = append(fvals, f)
	// 			// v is a string here, so e.g. v + " Yeah!" is possible.
	// 			fmt.Println("\tString: ", f)
	// 		default:
	// 			// And here I'm feeling dumb. ;)
	// 			fmt.Println("\tI don't know, ask stackoverflow.")
	// 		}
	// 	}
	// 	fmt.Println("Keys: ", fnames)
	// 	fmt.Println("Vals: ", fvals)
	// default:
	// 	// And here I'm feeling dumb. ;)
	// 	fmt.Println("I don't know, ask stackoverflow.")
	// }
}
