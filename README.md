# fabula

## **NOTE!**
## This application are not working anymore until the problem are fixed.
<hr>
Yet another of instagram stories scraper. Downloading instagram stories without account. Not working if the instagram target are private.

This tools using the best scraper and crawler framework : https://github.com/gocolly/colly

And thanks to https://hookgram.com for providing the website.

**Note**: Sometimes there is a bug that the application cannot scraping the stories, i am still figuring out what causing that. Just re-running the application until it shows downloading status on terminal and it should be working.

Clone:

```bash
$ git clone https://github.com/mrizkimaulidan/fabula.git
```

Download the required depedencies:

```bash
$ go mod tidy
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

Tested on Windows 10 Pro 21H2.