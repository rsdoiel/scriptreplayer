
# scriptreplayer

This is a hack to be able to leverage the output of the Unix "script" and "scriptreplay" commands
in a web browser. It is inspired by https://github.com/ysangkok/terminal_web_player which was
then inspired by others...

## Usage

```
    scriptreplayer [OPTIONS] [TIMING_FILE] [SCRIPT_LOG] TEMPLATE 
```

```shell
    scriptreplayer -t demo.timing -s demo.log page.tmpl > demo.html
```

## Approach

For a script session saved as "demo.timing" and "demo.log" 

1. Merge _demo.timing_ and _demo.log_ into a JSON object (e.g. demo.json), save to be rendered into the page
2. Generate Go templated page (or default template) into HTML markup, embedded JavaScript and JSON playback object

If this works, merge code into [mkpage](https://github.com/caltechlibrary/mkpage) project.

## Getting xterm.js

### Requires

+ node
+ npm

```
    sudo apt install node npm
    npm install xterm
```

### Building

```
    go build scriptreplayer.go
```
