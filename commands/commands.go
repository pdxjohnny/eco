package commands

import (
	"github.com/spf13/cobra"

	"github.com/pdxjohnny/eco/call"
	"github.com/pdxjohnny/eco/end"
	"github.com/pdxjohnny/eco/http"
	"github.com/pdxjohnny/eco/ws"
	"github.com/pdxjohnny/s-db/db"
)

// Commands are the commands that can be run
var Commands = []*cobra.Command{
	&cobra.Command{
		Use:   "call",
		Short: "Make a call",
		Run: func(cmd *cobra.Command, args []string) {
			ConfigBindFlags(cmd)
			call.Run()
		},
	},
	&cobra.Command{
		Use:   "end",
		Short: "End a call",
		Run: func(cmd *cobra.Command, args []string) {
			ConfigBindFlags(cmd)
			end.Run()
		},
	},
	&cobra.Command{
		Use:   "ws",
		Short: "Communicate with the call center as a ws",
		Run: func(cmd *cobra.Command, args []string) {
			ConfigBindFlags(cmd)
			ws.Run()
		},
	},
	&cobra.Command{
		Use:   "http",
		Short: "Start the http server",
		Run: func(cmd *cobra.Command, args []string) {
			ConfigBindFlags(cmd)
			http.Run()
		},
	},
	&cobra.Command{
		Use:   "db",
		Short: "Start the db service",
		Run: func(cmd *cobra.Command, args []string) {
			ConfigBindFlags(cmd)
			db.Run()
		},
	},
}

func init() {
	ConfigDefaults(Commands...)
}
