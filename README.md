# unseenmail
Cli application to get the number of unseen mails (atm only imap is supported)
I personally use it for adding a mail notification count to polybar. 

## Configuration
The configfile is read from $XDG_CONFIG_HOME/unseenmail/config.toml. If
XDG_CONFIG_HOME isn't set its read from $HOME/.config/unseenmail/config.toml.

### Configfile
You can add as much Imap servers as you want.

#### Example
```
[imap]
[imap.testserver]
Server = "imap.testserver.com"
Port = 993
Username = "YourUsername"
Password = "YourPassword"                                       
```

## Usage
Just run unseenmail. 

## Installation
Get it from your distros package manager. If it isn't in your distro package it
yourself. Every other way to get it to your system sucks.

### Gentoo/Funtoo
The ebuild will be in my gentoo overlay "oboeverlay"
