// Copyright (c) 2018 ChrisOboe
//
// This file is part of unseenmail
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"fmt"
	"github.com/ChrisOboe/unseenmail/backends"
	"github.com/ChrisOboe/unseenmail/config"
	"os"
)

func getConfigdir(appname string) string {
	if os.Getenv("XDG_CONFIG_HOME") != "" {
		return os.Getenv("XDG_CONFIG_HOME") + "/" + appname
	} else {
		return os.Getenv("HOME") + "/.config/" + appname
	}
}

func main() {
	configfile := getConfigdir("unseenmail") + "/config.toml"
	c, err := config.New(configfile)
	if err != nil {
		fmt.Println(err)
	}

	var unread uint = 0

	// handle imap
	for _, imapConfig := range c.Data.Imap {
		i := backends.NewImap(imapConfig)
		tmpUnread, err := i.GetUnread()
		if err != nil {
			fmt.Println(err)
		}
		unread += tmpUnread
	}

	fmt.Println(unread)
	if unread == 0 {
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}
