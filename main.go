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

package main

import (
	"github.com/hjma29/ovcli/cmd"
)

func main() {

	//pflag.Parse()
	//if commandflag -d is not set, then we should change default logger destination to discard
	//fmt.Println(cmd.Debugmode)
	// if !cmd.Debugmode {
	// 	log.SetOutput(ioutil.Discard)
	// }
	cmd.Execute()
}

// In your project directory (where your main.go file is) you would run the following:
//
// cobra add serve
// cobra add config
// cobra add create -p 'configCmd
