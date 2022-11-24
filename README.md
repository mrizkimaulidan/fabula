# fabula

Yet another of instagram stories scraper. Downloading instagram stories without account. Not working if the instagram target are private.

Thanks to https://storiesig.info/en/ for providing the website and the API.

Clone:

```bash
$ git clone https://github.com/mrizkimaulidan/fabula.git
```

Download the required depedencies:

```bash
$ go mod download
```

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
```

Check the `./stories` folder where the story contents has been downloaded.