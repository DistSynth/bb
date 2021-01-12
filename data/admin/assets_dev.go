//+build dev

package admin

import "net/http"

// Assets contains project assets.
var Assets http.FileSystem = http.Dir("admin_ui")
