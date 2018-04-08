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

package config

import (
	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"
)

type Config struct {
	File string
	Data Data
}

type Data struct {
	Imap map[string]Imap
}

type Imap struct {
	Server   string
	Port     int
	Username string
	Password string
}

func New(file string) (Config, error) {
	var d Data
	_, err := toml.DecodeFile(file, &d)
	if err != nil {
		return Config{}, errors.Wrap(err, "Can't read configfile")
	}
	return Config{file, d}, nil
}
