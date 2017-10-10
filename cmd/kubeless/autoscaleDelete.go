/*
Copyright 2016 Skippbox, Ltd.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"github.com/Sirupsen/logrus"
	"github.com/kubeless/kubeless/pkg/utils"
	"github.com/spf13/cobra"
)

var autoscaleDeleteCmd = &cobra.Command{
	Use:   "delete <name>",
	Short: "delete an autoscale from Kubeless",
	Long:  `delete an autoscale from Kubeless`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			logrus.Fatal("Need exactly one argument - autoscale name")
		}
		asName := args[0]

		ns, err := cmd.Flags().GetString("namespace")
		if err != nil {
			logrus.Fatal(err)
		}

		client := utils.GetClientOutOfCluster()

		err = utils.DeleteAutoscale(client, asName, ns)
		if err != nil {
			logrus.Fatal(err)
		}

		err = utils.DeleteServiceMonitor(asName, ns)
		if err != nil {
			logrus.Fatal(err)
		}
	},
}
