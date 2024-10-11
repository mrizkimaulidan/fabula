# fabula

Fabula is yet another Instagram stories and highlight stories downloader that allows you to download Instagram stories and highlight stories without an account. Please note that it will not work if the Instagram target is private.

Special thanks to https://storiesig.info/en/ for providing the website and API.

## Prerequisites

Ensure that Go is installed on your machine before proceeding. You can download and install it from the official [Go website](https://golang.org/dl/).

## Installation

1. Clone the repository using the following command:

```bash
$ git clone https://github.com/mrizkimaulidan/fabula.git
```

2. Build the Project

```bash
$ go build
```

**Note**: The executable file will vary based on your operating system (e.g., `fabula.exe` for Windows, `fabula` for Unix-based systems).

## Usage

Run the program with the --help flag to see usage instructions:

```bash
$ ./fabula --help
```

Output:

```bash
Usage of ./fabula:
  -option string
        The parsing option 'story' or 'highlight'
  -username string
        The Instagram username
```

### Download Instagram Stories

To download all available public stories from a user, use the `story` option:

```bash
$ ./fabula -username=john.doe -option=story
```

Sample output:

```bash
=======================================
= Name          : John Doe (@john.doe)
= Followers     : 1000
= Followings    : 500
= Public Email  : john.doe@mail.com
=======================================
Found 6 stories for the user.
Downloading... 1677518969229567834.mp4
Downloading... 1677518969229567434.mp4
Downloading... 1677518969229566634.jpg
...
All stories have been downloaded!
```

The stories will be saved in the `./stories/<username>` folder.

### Download Instagram Highlights

To download highlight stories from a public Instagram user, use the `highlight` option:

```bash
$ ./fabula -username=john.doe -option=highlight
```

Sample output:

```bash
=======================================
= Found 2 story highlights for the user.
= 1. Florida
= 2. Los Angeles
=======================================
Which highlight do you want to download (enter number): 1
Downloading... 1726077594588483045.mp4
Downloading... 1726077594613101645.mp4
...
All highlight stories have been downloaded!
```

The selected highlight stories will be saved in the `./highlights/<username>/<highlight_number>` folder..

## Using the Makefile (Developers Only)

Alternatively, you can use the `Makefile` to build for multiple platforms (Windows, MacOS, Linux).

The provided `Makefile` simplifies building the project and running specific tasks.

### Build for Multiple Platforms

```bash
$ make build
```

This will build the project for the following platforms:

- Windows (32-bit and 64-bit)
- MacOS (64-bit)
- Linux (64-bit)

The resulting binaries will be placed in the `bin/` directory.

### Run for Stories

To download stories for a specific username:

```bash
$ make story USERNAME=john.doe
```

### Run for Highlights

To download highlights for a specific username:

```bash
$ make highlight USERNAME=john.doe
```

### Clean Up

To clean up build files and downloaded data:

```bash
$ make clean
```

This will remove the `bin/` directory and the downloaded stories and highlights.

## Platform Support

This project has been developed and tested on **WSL Ubuntu 22.04.4 LTS**, but it should work on other platforms that support Go as well.

## Directory Structure

After downloading, the files will be stored in the following directory structure:

```bash
fabula/
│
├── stories/
│   └── <username>/
│       └── <story_files>
│
└── highlights/
    └── <username>/
        └── <highlight_number>/
            └── <highlight_files>
```

## Important Note

- **Instagram Privacy**: Fabula cannot download content from private Instagram accounts.
- **File Naming**: The downloaded files retain their original Instagram filenames (e.g., `1677518969229567834.mp4`).
