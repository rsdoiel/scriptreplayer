
# Dev notes


Need a terminal emulator (e.g. xterm) for the page:

+ https://www.npmjs.com/package/xterm2
+ https://github.com/131/xterm2

Need something to take a *.timing and a *.log  and turn them
into a JSON object to pass to page.

Need to create a scriptreplay running that will feed the timing and content into xterm2 for playback

+ look at https://github.com/ysangkok/terminal_web_player

Need to have a template that will generate our page based on xterm2, our JSONized 
replay description and our replayer.

