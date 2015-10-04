Chatterbox is a next-generation IRC client coming soon to Windows, Mac, Linux, and the web in the next couple months. It's built with a modern, minimalistic, visually pleasing interface and support for richly formatted messages with embedded media and Markdown message formatting.

![Screenshot of the first version of the interface](http://i.imgur.com/ZjpBm3w.png)

It will also have a hosted version that keeps subscribers conntected at all times with an uninterrupted experience across multiple devices and the web, even when a user's computer is off.

The desktop and web version is being developed in HTML, CSS, and JavaScript. The Windows, Mac, and Linux versions will use [Electron](http://electron.atom.io) for this. The backend will be written in [Go](https://golang.org) that will run on dedicated servers to power the web version and run locally alongside Electron for the local version. Mobile apps are planned after the initial desktop and web release.

Our development IRC channel is `#chatterbox-irc` on Freenode.

CI Status: [![Build Status](https://img.shields.io/travis/Chatterbox-IRC/Chatterbox.svg?style=flat)](https://travis-ci.org/Chatterbox-IRC/Chatterbox)
[![Code Climate](https://img.shields.io/codeclimate/github/Chatterbox-IRC/Chatterbox.svg?style=flat)](https://codeclimate.com/github/Chatterbox-IRC/Chatterbox)
