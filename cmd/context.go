package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	src "github.com/tkennes/jengo/src"
)

var ConfigRootCmd = &cobra.Command{
	Use:   "config",
	Short: "The entrypoint for the config interface",
}

var ContextRootCmd = &cobra.Command{
	Use:   "context",
	Short: "The entrypoint for the config interface",
}

var AddContextCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new context",
	Run: func(cmd *cobra.Command, args []string) {
		context_name := parse_argument(cmd, "context_name")
		token := parse_argument(cmd, "token")
		url := parse_argument(cmd, "url")
		username := parse_argument(cmd, "username")
		src.DoAddContext("", context_name, token, url, username)
	},
}

var RemoveContextCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove context",
	Run: func(cmd *cobra.Command, args []string) {
		context_name := parse_argument(cmd, "context_name")
		src.DoRemoveContext("", context_name)
	},
}

var UpdateContextCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a context",
	Run: func(cmd *cobra.Command, args []string) {
		context_name := parse_argument(cmd, "context_name")
		new_context_name := parse_argument(cmd, "new_context_name")
		token := parse_argument(cmd, "token")
		url := parse_argument(cmd, "url")
		username := parse_argument(cmd, "username")
		manual := parse_bool_argument(cmd, "manual")
		fmt.Println(token)
		if manual {
			src.DoUpdateManualContext("", context_name)
		} else {
			src.DoUpdateContext("", context_name, new_context_name, token, url, username)
		}
	},
}

var GetContextCmd = &cobra.Command{
	Use:   "get",
	Short: "Get context info",
	Run: func(cmd *cobra.Command, args []string) {
		for _, context := range args {
			src.YAML(src.GetAndShowContext("", context))
		}
	},
}

var GetCurrentContextCmd = &cobra.Command{
	Use:   "current",
	Short: "Get current context info",
	Run: func(cmd *cobra.Command, args []string) {
		src.YAML(src.GetAndShowCurrentContext(""))
	},
}

var SetCurrentContextCmd = &cobra.Command{
	Use:   "set",
	Short: "Set current context",
	Run: func(cmd *cobra.Command, args []string) {
		context_name := parse_argument(cmd, "context_name")
		if context_name == "" {
			src.ErrorLog(errors.New(fmt.Sprintf("Context name \"%s\" does not exist", context_name)))
		}
		src.DoUpdateCurrentContext("", context_name)
	},
}

var ListContextsCmd = &cobra.Command{
	Use:   "list",
	Short: "Show all contexts",
	Run: func(cmd *cobra.Command, args []string) {
		src.YAML(src.GetAndShowContexts(""))
	},
}


func init() {
	// Toplevel
	RootCmd.AddCommand(ContextRootCmd)

	// Context handling- Sublevel
	var (
		context_name string
		manual bool
		new_context_name string
		token string
		url string
		username string
	)
	// CRU(D)
	ContextRootCmd.AddCommand(AddContextCmd)
	AddContextCmd.PersistentFlags().StringVarP(&context_name, "context_name", "n", "", "The name of the context")
	AddContextCmd.PersistentFlags().StringVarP(&token, "token", "t", "", "Token of the context")
	AddContextCmd.PersistentFlags().StringVarP(&url, "url", "", "", "The URL of the context")
	AddContextCmd.PersistentFlags().StringVarP(&username, "username", "u", "", "The username of the context")

	ContextRootCmd.AddCommand(RemoveContextCmd)
	RemoveContextCmd.PersistentFlags().StringVarP(&context_name, "context_name", "n", "", "The name of the context")

	ContextRootCmd.AddCommand(UpdateContextCmd)
	UpdateContextCmd.PersistentFlags().StringVarP(&context_name, "context_name", "n", "", "The name of the context")
	UpdateContextCmd.PersistentFlags().StringVarP(&new_context_name, "new_context_name", "a", "", "The new name of the context")
	UpdateContextCmd.PersistentFlags().StringVarP(&token, "token", "t", "", "Token of the context")
	UpdateContextCmd.PersistentFlags().StringVarP(&url, "url", "", "", "The URL of the context")
	UpdateContextCmd.PersistentFlags().StringVarP(&username, "username", "u", "", "The username of the context")
	UpdateContextCmd.PersistentFlags().BoolVarP(&manual, "manual", "m", false, "Manual update")

	// Getters
	ContextRootCmd.AddCommand(GetContextCmd)
	ContextRootCmd.AddCommand(GetCurrentContextCmd)
	ContextRootCmd.AddCommand(ListContextsCmd)

	// Setters
	ContextRootCmd.AddCommand(SetCurrentContextCmd)
	SetCurrentContextCmd.PersistentFlags().StringVarP(&context_name, "context_name", "n", "", "The name of the context")

}
////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Helpers
////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func parse_argument(cmd *cobra.Command, argument_name string) string {
	context_name, err := cmd.Flags().GetString(argument_name)
	if err != nil {
		src.ErrorLog(err)
	}
	return context_name
}

func parse_bool_argument(cmd *cobra.Command, argument_name string) bool {
	context_name, err := cmd.Flags().GetBool(argument_name)
	if err != nil {
		src.ErrorLog(err)
	}
	return context_name
}
