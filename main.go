package main

import (
	"runtime"

	"github.com/spf13/cobra"

	"github.com/pdxjohnny/eco/commands"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var rootCmd = &cobra.Command{Use: "eco"}
	rootCmd.AddCommand(commands.Commands...)
	rootCmd.Execute()
}
