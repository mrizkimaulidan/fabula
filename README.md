# fabula

Fabula is yet another Instagram stories downloader that allows you to download Instagram stories without an account. Please note that it will not work if the Instagram target is private.

Special thanks to https://storiesig.info/en/ for providing the website and API.

Note: You need Go installed on your machine!

Clone the repository:

Clone:

```bash
$ git clone https://github.com/mrizkimaulidan/fabula.git
```

Build the project:
```bash
$ go build
```

Usage:

Show help:

```bash
$ ./fabula --help
Usage of ./fabula:
  -username string
        the Instagram username
```

Run the Instagram story downloader:

```bash
$ ./fabula -username=john.doe
2023/02/28 01:29:29 main.go:36: =======================================
2023/02/28 01:29:29 main.go:37: = Name          : John Doe(@john.doe)
2023/02/28 01:29:29 main.go:38: = Followers     : 1000
2023/02/28 01:29:29 main.go:39: = Followings    : 500
2023/02/28 01:29:29 main.go:40: = Public Email  : john.doe@mail.com
2023/02/28 01:29:29 main.go:41: =======================================
2023/02/28 01:29:29 main.go:42: Found the user with 6 stories
2023/02/28 01:29:29 main.go:49: Downloading.. 1677518969229567834.mp4
2023/02/28 01:29:29 main.go:49: Downloading.. 1677518969229567434.mp4
2023/02/28 01:29:29 main.go:49: Downloading.. 1677518969229566634.jpg
2023/02/28 01:29:29 main.go:49: Downloading.. 1677518969229567334.jpg
2023/02/28 01:29:29 main.go:49: Downloading.. 1677518969229567634.jpg
2023/02/28 01:29:29 main.go:49: Downloading.. 1677518969229567734.jpg
2023/02/28 01:29:31 main.go:63: Downloaded.. 1677518969229567634.jpg
2023/02/28 01:29:32 main.go:63: Downloaded.. 1677518969229567334.jpg
2023/02/28 01:29:36 main.go:63: Downloaded.. 1677518969229567734.jpg
2023/02/28 01:29:36 main.go:63: Downloaded.. 1677518969229566634.jpg
2023/02/28 01:29:39 main.go:63: Downloaded.. 1677518969229567834.mp4
2023/02/28 01:29:42 main.go:63: Downloaded.. 1677518969229567434.mp4
2023/02/28 01:29:42 main.go:68: All stories have been downloaded!
```

Check the ./stories folder where the story contents have been downloaded.