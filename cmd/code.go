// Copyright © 2016 NAME HERE <EMAIL ADDRESS>
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
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// codeCmd represents the code command
var codeCmd = &cobra.Command{
	Use:   "code",
	Short: "generate code files from yaml files",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		GenerateCode()
	},
}

func init() {
	RootCmd.AddCommand(codeCmd)

	codeCmd.PersistentFlags().StringP("package", "p", "", "code file (.go) package name")
	codeCmd.PersistentFlags().StringP("input", "i", ".", "directory of yaml files")
	codeCmd.PersistentFlags().StringP("model", "m", "", "model need to generate code")
	viper.BindPFlag("package", codeCmd.PersistentFlags().Lookup("package"))
	viper.BindPFlag("code_input", codeCmd.PersistentFlags().Lookup("input"))
	viper.BindPFlag("code_model", codeCmd.PersistentFlags().Lookup("model"))
}
