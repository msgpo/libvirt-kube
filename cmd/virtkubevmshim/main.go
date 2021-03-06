/*
 * This file is part of the libvirt-kube project
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * Copyright 2017 Red Hat, Inc.
 *
 */

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/spf13/pflag"

	"libvirt.org/libvirt-kube/pkg/vmshim"
)

var (
	shimaddr = pflag.String("shim-addr", "/run/virtkubevmshim/shim.sock",
		"UNIX socket path for shim service")
	skipvalidate = pflag.Bool("no-validate", false,
		"Skip validating client container identity")
	connect = pflag.String("connect", "qemu:///system",
		"Libvirt connection URI")
	kubeconfig = pflag.String("kubeconfig", "", "Path to a kube config, if running outside cluster")
	repopath   = pflag.String("repopath", "/srv/images", "Path to image repository mount point")
)

func main() {
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	// Convince glog that we really have parsed CLI
	flag.CommandLine.Parse([]string{})

	svc, err := vmshim.NewShim(*shimaddr, *skipvalidate, *connect, *repopath, *kubeconfig)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = svc.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.Exit(0)
}
