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
	//"github.com/HewlettPackard/oneview-golang/ov"

	"fmt"
	"os"

	"github.com/hjma29/ovcli/ovextra"

	"github.com/spf13/cobra"
)

// connectCmd represents the connect command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add remote enclosure",
	Long:  `add remote enclosure during initial setup`,
	Run:   addenc,
}

func addenc(cmd *cobra.Command, args []string) {

	if ipv6 == "" {
		fmt.Print("Please specify remote enclosure ipv6 address")
		os.Exit(1)
	}

	ovextra.AddRemoteEnc(ipv6)
}

func init() {

	addCmd.Flags().StringVarP(&ipv6, "ipv6", "i", "", "remote enclosure ipv6 address")

}
