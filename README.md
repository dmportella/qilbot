# Qilbot

Qilbot is a discord bot written in `GOLANG`.

@dmportella

[![Build Status](https://travis-ci.org/dmportella/qilbot.svg?branch=master)](https://travis-ci.org/dmportella/qilbot)[![GoDoc](https://godoc.org/github.com/dmportella/qilbot?status.svg)](https://godoc.org/github.com/dmportella/qilbot)

```
.
├── LICENSE
├── main.go
├── makefile
├── README.md
└── vendor
    ├── github.com
    │   ├── bwmarrin
    │   │   └── discordgo
    │   │       ├── discord.go
    │   │       ├── endpoints.go
    │   │       ├── events.go
    │   │       ├── LICENSE
    │   │       ├── message.go
    │   │       ├── oauth2.go
    │   │       ├── README.md
    │   │       ├── restapi.go
    │   │       ├── state.go
    │   │       ├── structs.go
    │   │       ├── util.go
    │   │       ├── voice.go
    │   │       └── wsapi.go
    │   └── gorilla
    │       └── websocket
    │           ├── AUTHORS
    │           ├── client.go
    │           ├── compression.go
    │           ├── conn.go
    │           ├── conn_read.go
    │           ├── conn_read_legacy.go
    │           ├── doc.go
    │           ├── json.go
    │           ├── LICENSE
    │           ├── README.md
    │           ├── server.go
    │           └── util.go
    ├── golang.org
    │   └── x
    │       └── crypto
    │           ├── LICENSE
    │           ├── nacl
    │           │   └── secretbox
    │           │       └── secretbox.go
    │           ├── PATENTS
    │           ├── poly1305
    │           │   ├── const_amd64.s
    │           │   ├── poly1305_amd64.s
    │           │   ├── poly1305_arm.s
    │           │   ├── poly1305.go
    │           │   ├── sum_amd64.go
    │           │   ├── sum_arm.go
    │           │   └── sum_ref.go
    │           └── salsa20
    │               └── salsa
    │                   ├── hsalsa20.go
    │                   ├── salsa2020_amd64.s
    │                   ├── salsa208.go
    │                   ├── salsa20_amd64.go
    │                   └── salsa20_ref.go
    └── vendor.json

14 directories, 45 files
```