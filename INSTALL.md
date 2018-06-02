
# Installation

*scriptreplayer* is a set of command line programs run from a shell like Bash. It will
build an HTML page and other assets for playing back a "script" transcript and time log.
This lets you see a "scriptreplay" like version of the commands from the transcript.

This is generalized instructions for a release.  For deployment suggestions see NOTES.md

## Compiled version

Compiled versions are available for Mac OS X (amd64 processor, macosx-amd64), Linux (amd64 process, linux-amd64), 
Windows (amd64 processor, windows-amd64) and Rapsberry Pi (arm7 processor, raspbian-arm7)

VERSION_NUMBER is a [symantic version number](http://semver.org/) (e.g. v0.1.2)


For all the released version go to the project page on Github and click latest release

>    https://github.com/rsdoiel/scriptreplayer/releases/latest


| Platform    | Zip Filename                             |
|-------------|------------------------------------------|
| Windows     | scriptreplayer-VERSION_NUMBER-windows-amd64.zip |
| Mac OS X    | scriptreplayer-VERSION_NUMBER-macosx-amd64.zip  |
| Linux/Intel | scriptreplayer-VERSION_NUMBER-linux-amd64.zip   |
| Raspbery Pi | scriptreplayer-VERSION_NUMBER-raspbian-arm7.zip |


## The basic recipe

+ Find the Zip file listed matching the architecture you're running and download it
    + (e.g. if you're on a Windows 10 laptop/Surface with a amd64 style CPU you'd choose the Zip file with "windows-amd64" in the name).
+ Download the zip file and unzip the file.
+ Copy the contents of the folder named "bin" to a folder that is in your path 
    + (e.g. "$HOME/bin" is common).
+ Adjust your PATH if needed
    + (e.g. `export PATH="$HOME/bin:$PATH"`)
+ Test


### Mac OS X

1. Download the zip file
2. Unzip the zip file
3. Copy the executables to $HOME/bin (or a folder in your path)
4. Make sure the new location in in our path
5. Test

Here's an example of the commands run in the Terminal App after downloading the 
zip file.

```shell
    cd Downloads/
    unzip scriptreplayer-*-macosx-amd64.zip
    mkdir -p $HOME/bin
    cp -v bin/* $HOME/bin/
    export PATH=$HOME/bin:$PATH
    scriptreplayer -version
```

### Windows

1. Download the zip file
2. Unzip the zip file
3. Copy the executables to $HOME/bin (or a folder in your path)
4. Test

Here's an example of the commands run in from the Bash shell on Windows 10 after
downloading the zip file.

```shell
    cd Downloads/
    unzip scriptreplayer-*-windows-amd64.zip
    mkdir -p $HOME/bin
    cp -v bin/* $HOME/bin/
    export PATH=$HOME/bin:$PATH
    scriptreplayer -version
```


### Linux 

1. Download the zip file
2. Unzip the zip file
3. Copy the executables to $HOME/bin (or a folder in your path)
4. Test

Here's an example of the commands run in from the Bash shell after
downloading the zip file.

```shell
    cd Downloads/
    unzip scriptreplayer-*-linux-amd64.zip
    mkdir -p $HOME/bin
    cp -v bin/* $HOME/bin/
    export PATH=$HOME/bin:$PATH
    scriptreplayer -version
```


### Raspberry Pi

Released version is for a Raspberry Pi 2 or later use (i.e. requires ARM 7 support).

1. Download the zip file
2. Unzip the zip file
3. Copy the executables to $HOME/bin (or a folder in your path)
4. Test

Here's an example of the commands run in from the Bash shell after
downloading the zip file.

```shell
    cd Downloads/
    unzip scriptreplayer-*-raspbian-arm7.zip
    mkdir -p $HOME/bin
    cp -v bin/* $HOME/bin/
    export PATH=$HOME/bin:$PATH
    scriptreplayer -version
```


## Compiling from source

_scriptreplayer_ is "go gettable".  Use the "go get" command to download the dependant packages
as well as _scriptreplayer_'s source code. 


```shell
    go get -u github.com/rsdoiel/scriptreplayer/...
```

Or clone the repstory and then compile

```shell
    cd
    git clone https://github.com/rsdoiel/scriptreplayer src/github.com/rsdoiel/scriptreplayer
    cd src/github.com/rsdoiel/scriptreplayer
    make
    make test
    make install
```

Compilation assumes [go](https://github.com/golang/go) v1.10, node v8, npm v3.5

