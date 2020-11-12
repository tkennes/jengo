package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "jengo",
	Short: "JENGO is a golang-based Jenkins API wrapper",
}
