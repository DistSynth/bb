package cmd

import (
	"database/sql"
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	printSQL bool

	initCmd = &cobra.Command{
		Use:   "init",
		Short: "Init clikhouse scheme for BigBrother.",
		Long: `A Fast and Flexible web analytics server.
Complete documentation is available at http://...`,
		Run: runInit,
	}
)

const bbSQL = `
CREATE TABLE IF NOT EXISTS events (
	et String,
	s_n Array(String), 
	s_v Array(String),
	f_n Array(String), 
	f_v Array(Float64),
	i_n Array(String), 
    i_v Array(Int64),
	ts UInt64,
	d Date MATERIALIZED toDate(round(ts/1000)),
	dt DateTime MATERIALIZED toDateTime(round(ts/1000))
) ENGINE = MergeTree(d, et, 8192)
`

func init() {
	initCmd.PersistentFlags().BoolVarP(&printSQL, "print", "p", false, "print sql to stdout")

	rootCmd.AddCommand(initCmd)
}

func runInit(cmd *cobra.Command, args []string) {
	if printSQL {
		fmt.Println(bbSQL)
	} else {
		log.Info(viper.GetString("ch.url"))
		connect, err := sql.Open("clickhouse", viper.GetString("ch.url"))
		if err != nil {
			log.Fatal(err)
		}
		if err := connect.Ping(); err != nil {
			log.Fatal(err)
		}

		_, err = connect.Exec(bbSQL)
		if err != nil {
			log.Fatal(err)
		}
		log.Info("Tables created")

		_, err = connect.Exec("INSERT INTO events (et, ts, f_n, f_v,s_n,s_v,i_n,i_v)  VALUES ('cpu', 1509232010254, ['load','temp'], [0.85, 0.68],[],[],[],[])")
		if err != nil {
			log.Fatal(err)
		}
	}
}
