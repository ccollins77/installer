//
// Copyright 2017 The Maru OS Project
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package main

import (
	"fmt"
	"io/ioutil"

	"github.com/BurntSushi/toml"
)

type device struct {
	Common_name  string
	Product_name string

	Nhos_file string
	Nhos_url  string

	Nhfs_file string
	Nhfs_url  string

	Gapps_file string
	Gapps_url  string

	Twrp_file string
	Twrp_url  string

	Extra_file string
	Extra_url  string
}

type devices struct {
	Device []device
}

func readDevicesConfig() devices {
	b, err := ioutil.ReadFile("devices.toml")
	if err != nil {
		fmt.Print(err)
	}

	var nhConfig devices
	if _, err := toml.Decode(string(b), &nhConfig); err != nil {
		eEcho("ERROR READING TOML")
	}
	return nhConfig
}

func findDeviceConfig(nhDevices devices, deviceProductName string) device {
	for _, d := range nhDevices.Device {
		if d.Product_name == deviceProductName {
			return d
		}
	}
	return device{}
}
