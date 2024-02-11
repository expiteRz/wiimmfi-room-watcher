# wiimmfi-room-watcher

This tool is programmed to aimed to overlay your joined room on the live streaming.
> NOTE: Currently, wiimmfi-room-watcher cannot retrieve the room statistics correctly because Wiimmfi is now unstable for it. It is not a issue on the tool. Please don't ask about it to the author.

## Usage

1. The tool is on the state of work-in-progress, and not officially released yet. But you can download the executable file from [Actions](https://github.com/expiteRz/wiimmfi-room-watcher/actions). Currently Windows and Linux build is available.
    * Download the latest build, and unzip it anywhere
      * ex. If you're a 64-bit Windows user, you need to download **wiimmfi-room-watcher-win-amd64**.
    > In the root folder of the program, you can find the **static** directory. It contains all of the available overlays. Those are getting streamed via HTTP-File server.
2. Run wiimmfi-room-watcher
    * Visit the link in your browser (http://localhost:24050), and choose the layout.
    * Add a browser source in OBS

#### More technically information can be seen in readme of the based project [gosumemory](https://github.com/l3lackShark/gosumemory)

## Build

1. If you haven't installed [Go](https://go.dev/) yet, download and install it first
    * Requires v1.20 or newer to build
2. Clone the repository
3. Execute `go get -u` in terminal
4. Execute `go build` in terminal

## Credits

* [gosumemory](https://github.com/l3lackShark/gosumemory) - The original idea for wiimmfi-room-watcher
