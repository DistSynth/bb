//+build dev

package js

import "net/http"

// Assets contains project assets.
var Assets http.FileSystem = http.Dir("static")
