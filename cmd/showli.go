// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"os"
	"text/tabwriter"
	"text/template"

	"github.com/hjma29/ovcli/ovextra"
	"github.com/spf13/cobra"
)

var showLICmd = &cobra.Command{
	Use:   "li",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: showLI,
}

const (
	liShowFormat = "" +
		"Name\tConsistency\tStacking\tLIG\n" +
		//"----\t-----\n" +
		"{{range .}}" +
		"{{.Name}}\t{{.ConsistencyStatus}}\t{{.StackingHealth}}\t{{.LIGName}}\n" +
		"{{end}}"
)

func showLI(cmd *cobra.Command, args []string) {

	liMap := ovextra.GetLI()

	tw := tabwriter.NewWriter(os.Stdout, 5, 1, 3, ' ', 0)
	defer tw.Flush()

	t := template.Must(template.New("").Parse(liShowFormat))
	t.Execute(tw, liMap)

}

func init() {

	// showICCmd.AddCommand(showICPortCmd)
	// showICPortCmd.AddCommand(showSFPCmd)

	//showICPortCmd.Flags().StringVarP(&porttype, "type", "t", "all", "Port Type:uplink,downlink,interconnect,all")
	//&porttype = showICPortCmd.Flags().String("type", "", "Port Type:uplink,downlink,interconnect")

	// createNetworkTypePtr = createNetworkCmd.PersistentFlags().String("type", "", "Network Type")
	// createNetworkPurposePtr = createNetworkCmd.PersistentFlags().String("purpose", "", "General or Management etc")
	// createNetworkVlanIDPtr = createNetworkCmd.PersistentFlags().Int("vlan", 777, "General or Management etc")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// networkCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// networkCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
