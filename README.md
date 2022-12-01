# fabula

Yet another of instagram stories downloader. Downloading instagram stories without account. Not working if the instagram target are private.

Thanks to https://storiesig.info/en/ for providing the website and the API.

*Note : There is a rate limit when hitting the API, because it's calling the Instagram official API for getting the UserID based on the username, I'm still figuring out to get UserID without using Instagram official API*

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