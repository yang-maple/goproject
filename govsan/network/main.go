package main

import (
	"context"
	"fmt"
	examples "govsan"

	"github.com/vmware/govmomi/view"
	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/mo"
)

func main() {
	examples.Run(func(ctx context.Context, c *vim25.Client) error {
		// Create a view of Network types
		m := view.NewManager(c)

		v, err := m.CreateContainerView(ctx, c.ServiceContent.RootFolder, []string{"Network"}, true)
		if err != nil {
			return err
		}

		defer v.Destroy(ctx)

		// Reference: http://pubs.vmware.com/vsphere-60/topic/com.vmware.wssdk.apiref.doc/vim.Network.html
		var networks []mo.Network
		err = v.Retrieve(ctx, []string{"Network"}, nil, &networks)
		if err != nil {
			return err
		}

		for _, net := range networks {
			fmt.Printf("%s: %s, %s\n", net.Name, net.Reference(), net.ConfigStatus)
		}

		return nil
	})
}
