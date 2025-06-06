package cluster

import (
	"github.com/mrgb7/playground/cmd/cluster/plugin"
	"github.com/spf13/cobra"
)

var ClusterCmd = &cobra.Command{
	Use:   "cluster",
	Short: "Manage clusters",
	Long:  `Commands to create, delete, and get information about clusters`,
}

func init() {
	ClusterCmd.AddCommand(plugin.PluginCmd)
	ClusterCmd.AddCommand(createCmd)
	ClusterCmd.AddCommand(deleteCmd)
	ClusterCmd.AddCommand(cleanCmd)
	ClusterCmd.AddCommand(listCmd)
}
