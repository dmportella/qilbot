---
date: 2016-08-30T16:44:25+01:00
documentation: true
draft: false
title: Creating a Discord Bot
---

# Discord Bots

This section desbribes the steps required to create a new discord application so we can associate Qilbot to it.

## Application Creation

1. Browse to [https://discordapp.com/developers/applications](https://discordapp.com/developers/applications)
2. Click `Create New Application`

## Creating an bot

1. Click `My Applications`
2. Give the Application a name (this will be the name of the Bot).
3. Click `Create Application`.

![Step One](http://i.imgur.com/O216xc5.png)

1. This is your application `ID` you will need this for later.
   * Application Id is used to add your bot to a discord server.
2. This is the secret for you application *dont share this with anyone*.
3. Click `Create a Bot user`
   * An dialog will open asking if you are really sure you want todo this just say yes.

![Step Two](http://i.imgur.com/fgT59FI.png)

1. Once the bot is created click the `Reveal Token` link
2. Make note of the token presented to you.
   * This token is what we pass to `qilbot` as the token parameter.

![Step Two](http://i.imgur.com/eD0aUBZ.png)