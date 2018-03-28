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

package route

import (
	"fmt"
	"github.com/lastbackend/lastbackend/pkg/cli/envs"
	"github.com/lastbackend/lastbackend/pkg/cli/view"
	"github.com/spf13/cobra"
)

func ListCmd(cmd *cobra.Command, _ []string) {

	namespace, _ := cmd.Flags().GetString("namespace")

	if namespace == "" {
		fmt.Println("namesapace parameter not set")
		return
	}

	cli := envs.Get().GetClient()
	response, err := cli.V1().Namespace(namespace).Route().List(envs.Background())
	if err != nil {
		fmt.Println(err)
		return
	}

	list := view.FromApiRouteListView(response)
	list.Print()
}