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

package backends

import (
	"github.com/ChrisOboe/unseenmail/config"
	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"github.com/pkg/errors"
	"strconv"
)

type Imap struct {
	config config.Imap
}

func NewImap(config config.Imap) Imap {
	return Imap{config}
}

func (i Imap) GetUnread() (uint, error) {
	var unseen uint

	// connect to server
	imapClient, err := client.DialTLS(i.config.Server+":"+strconv.Itoa(i.config.Port), nil)
	if err != nil {
		return 0, errors.Wrap(err, "Couldn't connect to server")
	}
	defer imapClient.Logout()

	// login
	err = imapClient.Login(i.config.Username, i.config.Password)
	if err != nil {
		return 0, errors.Wrap(err, "Couldn't log in")
	}

	// getting mailboxes
	mbchan := make(chan *imap.MailboxInfo)
	go func() {
		imapClient.List("", "*", mbchan)
	}()

	var mailboxes []string
	for m := range mbchan {
		mailboxes = append(mailboxes, m.Name)
	}

	// iterate through mailboxes
	for _, mailbox := range mailboxes {
		status, err := imapClient.Status(mailbox, []imap.StatusItem{"UNSEEN"})
		if err != nil {
			return 0, errors.Wrap(err, "Couldn't get status of mailbox")
		}
		unseen += uint(status.Unseen)
	}

	return unseen, nil
}
