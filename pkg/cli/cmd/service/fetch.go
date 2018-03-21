//
// Last.Backend LLC CONFIDENTIAL
// __________________
//
// [2014] - [2018] Last.Backend LLC
// All Rights Reserved.
//
// NOTICE:  All information contained herein is, and remains
// the property of Last.Backend LLC and its suppliers,
// if any.  The intellectual and technical concepts contained
// herein are proprietary to Last.Backend LLC
// and its suppliers and may be covered by Russian Federation and Foreign Patents,
// patents in process, and are protected by trade secret or copyright law.
// Dissemination of this information or reproduction of this material
// is strictly forbidden unless prior written permission is obtained
// from Last.Backend LLC.
//

package service

import (
	"fmt"

	"github.com/lastbackend/lastbackend/pkg/cli/envs"
	"github.com/lastbackend/lastbackend/pkg/cli/view"
	"github.com/spf13/cobra"
)

func FetchCmd(cmd *cobra.Command, args []string) {

	if len(args) != 1 {
		cmd.Help()
		return
	}

	namespace, _ := cmd.Flags().GetString("namespace")

	if namespace == "" {
		fmt.Errorf("namesapace parameter not set")
		return
	}

	cli := envs.Get().GetClient()
	response, err := cli.V1().Namespace(namespace).Service(args[0]).Get(envs.Background())
	if err != nil {
		fmt.Println(err)
		return
	}

	ss := view.FromApiServiceView(response)
	ss.Print()
}
