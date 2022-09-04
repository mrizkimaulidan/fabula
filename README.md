# fabula

Yet another of instagram stories scraper. Downloading instagram stories without account. Not working if the instagram target are private.

This tools using https://github.com/PuerkitoBio/goquery for parsing the HTML.

And thanks to https://igpanda.com/ for providing the website and the API.

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
$ ./fabula.exe --help
Usage of fabula:
  -username string
        the instagram username
```

Run instagram story downloader

```bash
$ ./fabula.exe -username=john.doe
2022/06/27 14:14:26 starting to scrape story from john.doe
2022/06/27 14:14:27 downloading, please wait.
2022/06/27 14:14:29 downloading, please wait..
2022/06/27 14:14:30 downloading, please wait...
2022/06/27 14:14:32 downloading, please wait....
2022/06/27 14:14:34 downloading, please wait.....
2022/06/27 14:14:34 downloading, please wait......
2022/06/27 14:14:35 6 story downloaded from john.doe
```

Check the `./stories` folder where the story contents has been downloaded.