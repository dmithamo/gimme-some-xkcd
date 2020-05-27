# gimme-some-xkcd

## Why/Motivation

Why not?
Also, I kinda like learning, so, yea.

### Description

This tenee-tiny app is an attempt learning a bit of:

- command line args parsing with go
- file io with go
- making n/work requests with go

### How to Run

```bash
#!/usr/bin/bash
git clone <this-repo>
cd <this-repo>
go run client/cli.go [-n] [-s]
# -n : flag to request a specifi comic. Defaults to latest
# -s : flag to specify whether to save comic. Defaults to false
```

If all went well:

- a json response (or an error) should be printed on the cli
- the requested comic number should be downloaded into the root dir of this project

That's it.

### Feature Roadmap

- Allow saving to another directory
- Add a web client, complete with a cool ui, notifcations, the works.
- Make this run every morning (because who doesn't like to start their day with a good xkcd comic?)
- See what else can be done with the xkcd api

### Resources

- [The XKCD API](http://xkcd.com)
- [The Go Programming Language](http://golang.org)
