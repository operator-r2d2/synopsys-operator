// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
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

package synopsysctl

import (
	"fmt"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "List Synopsys Resources in your cluster",
	Args: func(cmd *cobra.Command, args []string) error {
		num_args := 1
		if len(args) != num_args {
			return fmt.Errorf("Must pass Resource Type")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		resource_name := args[0]
		out, err := RunKubeCmd("get", resource_name)
		if err != nil {
			fmt.Printf("Error getting %s with KubeCmd: %s", resource_name, err)
		}
		fmt.Printf("%+v", out)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
