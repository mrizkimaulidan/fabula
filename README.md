# fabula

Fabula is yet another Instagram stories and highlight stories downloader that allows you to download Instagram stories and highlight stories without an account. Please note that it will not work if the Instagram target is private.

Special thanks to https://storiesig.info/en/ for providing the website and API.

**Note**: You need Go installed on your machine!

Clone the repository:

Clone:

```bash
$ git clone https://github.com/mrizkimaulidan/fabula.git
```

Build the project:

```bash
$ go build
```

**Note**: The name of the executable may vary depending on the operating system. Make sure you are aware of this when running the commands!

The development work was done in WSL Ubuntu 22.04.4 LTS and tested well.

Usage:

Show help:

```bash
$ ./fabula --help
Usage of ./fabula:
  -option string
        the parsing option 'story' or 'highlight'
  -username string
        the Instagram username
```

Run the Instagram story downloader:

```bash
$ ./fabula -username=john.doe -option=story
=======================================
= Name          : John Doe(@john.doe)
= Followers     : 1000
= Followings    : 500
= Public Email  : john.doe@mail.com
=======================================
Found the user with 6 stories
Downloading.. 1677518969229567834.mp4
Downloading.. 1677518969229567434.mp4
Downloading.. 1677518969229566634.jpg
Downloading.. 1677518969229567334.jpg
Downloading.. 1677518969229567634.jpg
Downloading.. 1677518969229567734.jpg
Downloaded.. 1677518969229567634.jpg
Downloaded.. 1677518969229567334.jpg
Downloaded.. 1677518969229567734.jpg
Downloaded.. 1677518969229566634.jpg
Downloaded.. 1677518969229567834.mp4
Downloaded.. 1677518969229567434.mp4
All stories have been downloaded!
```

Run the Instagram story highlight downloader:

```bash
$ ./fabula -username=john.doe -option=highlight
=======================================
= Found the users with 2 story highlights
= 1. Florida
= 2. Los Angeles
=======================================
Which highlight you want to download (in number): 1
Downloading... 1726077594588483045.mp4
Downloading... 1726077594613101645.mp4
Downloading... 1726077594590142745.mp4
Downloading... 1726077594606143245.mp4
Downloaded... 1726077594606143245.mp4
Downloaded... 1726077594588483045.mp4
Downloaded... 1726077594613101645.mp4
Downloaded... 1726077594590142745.mp4
All highlight story have been downloaded!
```

Check the `./stories` or `./highlights` folder where the story or the highlight story contents have been downloaded.
