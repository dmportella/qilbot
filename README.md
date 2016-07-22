# Qilbot

Qilbot is a discord bot written in `GOLANG`.

@dmportella

[![Build Status](https://travis-ci.org/dmportella/qilbot.svg?branch=master)](https://travis-ci.org/dmportella/qilbot)[![GoDoc](https://godoc.org/github.com/dmportella/qilbot?status.svg)](https://godoc.org/github.com/dmportella/qilbot)

## Installing Qilbot

### Create a Discord Application

* Create an Discord Application.
  * At: https://discordapp.com/developers/applications/me
    Make sure you are logged in.
* Create a Bot for your newly created Application.
  * At: Ignore the message about not being reversable we dont care.
* Save the newly created token (you will need to click the reveal link)

### Running Qilbot

* TODO: https://discordapp.com/oauth2/authorize?client_id=appid&scope=bot&permissions=0

```
.
├── bot
│   ├── plugin_func.go
│   ├── plugin.go
│   ├── qilbot_func.go
│   └── qilbot.go
├── LICENSE
├── logging
│   ├── logging_func.go
│   └── logging.go
├── main.go
├── makefile
├── plugins
│   ├── common
│   │   ├── common_func.go
│   │   └── common.go
│   └── edsm
│       ├── distances_func.go
│       ├── distances.go
│       ├── edsmclient_func.go
│       ├── edsmclient.go
│       └── regex_func.go
├── qilbot
├── README.md
├── utilities
│   ├── json.go
│   └── regex.go
└── vendor
    └── vendor.json

12 directories, 21 files
```