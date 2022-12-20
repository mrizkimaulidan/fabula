# fabula

Yet another of instagram stories downloader. Downloading instagram stories without account. Not working if the instagram target are private.

Thanks to https://storiesig.info/en/ for providing the website and the API.

*Note: There is a rate limiter from Instagram website.*

Clone:

```bash
$ git clone https://github.com/mrizkimaulidan/fabula.git
```

*Note: There is a `Makefile` if you need to create the binary file by each operating system.*

Build:
```bash
$ go build
```

Usage:

Show help

```bash
$ ./fabula --help
Usage of ./fabula:
  -username string
        the instagram username
```

Run instagram story downloader

```bash
$ ./fabula -username=john.doe
2022/12/20 18:12:55 main.go:46: found the user with 2 story, downloading now please wait..
2022/12/20 18:12:56 file.go:30: downloading.. 1529430727970[.jpg]
2022/12/20 18:12:56 file.go:30: downloading.. 116560088719[.jpg]
2022/12/20 18:12:56 file.go:36: stories saved on : stories/john.doe
```

Check the `./stories` folder where the story contents has been downloaded.