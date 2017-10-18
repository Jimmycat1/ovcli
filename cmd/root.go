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
	"fmt"
	"log"
	"os"

	"github.com/hashicorp/logutils"
	"github.com/hjma29/ovcli/oneview"
	"github.com/spf13/cobra"
)

var cfgFile string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "ovcli",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {

	cobra.OnInitialize(initConfig)

	client := oneview.NewCLIOVClient()
	RootCmd.AddCommand(NewShowCmd(client))
	RootCmd.AddCommand(deleteCmd)
	RootCmd.AddCommand(createCmd)
	RootCmd.AddCommand(connectCmd)
	RootCmd.AddCommand(addCmd)

	RootCmd.PersistentFlags().BoolVarP(&Debugmode, "debug", "d", false, "Debug:true,false")

	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

}

//NewShowCmd creates a cobra command with desired output destination
func NewShowCmd(client oneview.Client) *cobra.Command {

	var showCmd = &cobra.Command{
		Use:   "show",
		Short: "shows different OneView resources",
		Long:  `shows different OneView resources`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	showCmd.AddCommand(showLIGCmd)
	showCmd.AddCommand(showLICmd)
	showCmd.AddCommand(showICCmd)
	showCmd.AddCommand(showUplinkSetCmd)
	showCmd.AddCommand(showEncCmd)
	showCmd.AddCommand(showNetworkCmd)
	showCmd.AddCommand(showEGCmd)
	showCmd.AddCommand(showSPCmd)
	showCmd.AddCommand(NewShowSPTemplateCmd(client))

	return showCmd
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {

	if Debugmode {
		//log.SetDebug(true)
		filter := &logutils.LevelFilter{
			Levels:   []logutils.LogLevel{"DEBUG", "WARN", "ERROR"},
			MinLevel: logutils.LogLevel("DEBUG"),
			Writer:   os.Stderr,
		}
		log.SetOutput(filter)
		//fmt.Println("debug mode")

	} else {
		filter := &logutils.LevelFilter{
			Levels:   []logutils.LogLevel{"DEBUG", "WARN", "ERROR"},
			MinLevel: logutils.LogLevel("WARN"),
			Writer:   os.Stderr,
		}
		log.SetOutput(filter)
		//fmt.Println("non-debug mode")

	}

}
