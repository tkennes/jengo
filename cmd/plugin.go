package cmd

import (
	"github.com/spf13/cobra"
	src "github.com/tkennes/jengo/src"
)

var PluginRootCmd = &cobra.Command{
	Use:   "plugin",
	Short: "The entrypoint for the plugin interface",
}

var GetPluginCmd = &cobra.Command{
	Use:   "get",
	Short: "Get plugin",
	Run: func(cmd *cobra.Command, args []string) {
		src.YAML(src.GetPlugin(args))
	},
}

var ListPluginsCmd = &cobra.Command{
	Use:   "list",
	Short: "Get plugins",
	Run: func(cmd *cobra.Command, args []string) {
		available, _ := cmd.Flags().GetBool("available")
		updateable, _ := cmd.Flags().GetBool("updateable")
		regex_filter, _ := cmd.Flags().GetString("filter")

		if !available && !updateable {
			// Listing Installed Plugins
			src.Table(src.ListPlugins(regex_filter), src.HeadersPlugins)
		} else if updateable{
			src.Table(src.GetAllUpdateablePlugins(regex_filter), src.HeadersAvailablePlugins)
		} else {
			//src.GetAllAvailablePlugins(regex_filter)
			src.Table(src.GetAllAvailablePlugins(regex_filter), src.HeadersAvailablePlugins)
		}
	},
}

func init() {
	// Toplevel
	RootCmd.AddCommand(PluginRootCmd)

	// Sublevel
	var regex_filter string
	PluginRootCmd.AddCommand(GetPluginCmd)

	PluginRootCmd.AddCommand(ListPluginsCmd)
	ListPluginsCmd.Flags().BoolP("available", "a", false, "Show available plugins")
	ListPluginsCmd.Flags().BoolP("updateable", "u", false, "Show updateable plugins")
	ListPluginsCmd.PersistentFlags().StringVarP(&regex_filter, "filter", "f", "", "Regex filter plugin names")
}
