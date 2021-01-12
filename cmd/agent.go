package cmd

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"ab.ru/bigbrother/data/admin"
	"ab.ru/bigbrother/data/js"
)

var (
	verb     bool
	agentCmd = &cobra.Command{
		Use:   "agent",
		Short: "BigBrother agent",
		Long: `A Fast and Flexible Static Site Generator built with
love by spf13 and friends in Go.
Complete documentation is available at http://hugo.spf13.com`,
		Run: runAgent,
	}
)

func init() {
	agentCmd.PersistentFlags().BoolVarP(&verb, "bb", "b", false, "berbose output")

	rootCmd.AddCommand(agentCmd)
}

func runAgent(cmd *cobra.Command, args []string) {
	http.Handle("/admin/", isRegistered(http.StripPrefix("/admin/", http.FileServer(admin.Assets))))
	http.Handle("/static/", isRegistered(http.StripPrefix("/static/", http.FileServer(js.Assets))))

	http.HandleFunc("/api", agentHandler)

	log.Fatal(http.ListenAndServe(":8081", nil))
}

func isRegistered(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if len(r.URL.Query().Get("siteId")) == 0 {
			http.Error(w, "missing key", http.StatusUnauthorized)
			return // don't call original handler
		}
		h.ServeHTTP(w, r)
	})
}

func agentHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}
