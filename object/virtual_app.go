/*
Copyright (c) 2014 VMware, Inc. All Rights Reserved.

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

package object

import (
	"fmt"

	"golang.org/x/net/context"

	"github.com/vmware/govmomi/vim25"
	"github.com/vmware/govmomi/vim25/mo"
	"github.com/vmware/govmomi/vim25/types"
)

type VirtualApp struct {
	*ResourcePool
}

func NewVirtualApp(c *vim25.Client, ref types.ManagedObjectReference) *VirtualApp {
	return &VirtualApp{
		ResourcePool: NewResourcePool(c, ref),
	}
}

func (p VirtualApp) String() string {
	if p.InventoryPath == "" {
		return p.Common.String()
	}
	return fmt.Sprintf("%v @ %v", p.Common, p.InventoryPath)
}

func (p VirtualApp) Name(ctx context.Context) (string, error) {
	var o mo.VirtualApp

	err := p.Properties(ctx, p.Reference(), []string{"name"}, &o)
	if err != nil {
		return "", err
	}

	return o.Name, nil
}