/*
Copyright 2019 The Kubernetes Authors.

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

package cmd

import (
	"fmt"
	"github.com/kubernetes-sigs/ingress-controller-conformance/internal/pkg/checks"
	"github.com/spf13/cobra"
	"golang.org/x/text/feature/plural"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"time"
)

func init() {
	verifyCmd.Flags().StringVarP(&checkName, "check", "c", "", "Verify only this specified check name.")

	rootCmd.AddCommand(verifyCmd)

	_ = message.Set(language.English, "%d success",
		plural.Selectf(1, "%d",
			"=0", "No checks passed...",
			"=1", "1 check passed,",
			"other", "%d checks passed!",
		),
	)
	_ = message.Set(language.English, "%d failure",
		plural.Selectf(1, "%d",
			"=0", "No failures!",
			"=1", "1 failure",
			"other", "%d failures!",
		),
	)
}

var (
	checkName = ""
)

var verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "Run Ingress verifications for conformance",
	Long:  "Run Ingress verifications for conformance",
	Run: func(cmd *cobra.Command, args []string) {
		start := time.Now()

		config := checks.Config{}
		successCount, failureCount, err := checks.Checks.Verify(checkName, config)
		if err != nil {
			fmt.Printf(err.Error())
		}

		elapsed := time.Since(start)

		p := message.NewPrinter(language.English)
		fmt.Printf("--- Verification completed ---\n%s %s\nin %s\n",
			p.Sprintf("%d success", successCount),
			p.Sprintf("%d failure", failureCount),
			elapsed)
	},
}
