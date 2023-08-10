# wiimmfi-room-watcher

This tool is programmed to aimed to overlay your joined room on the live streaming.

## Usage

1. Currently the tool is not officially released yet. But you can download the executable file from Actions. Only Windows build is available.
    * Unzip it anywhere
    > In the root folder of the program, you can find the **static** directory. It contains all of the available overlays. Those are getting streamed via HTTP-File server.
2. Run wiimmfi-room-watcher
    * Visit the link in your browser (http://localhost:24050), and choose the layout.
    * Add a browser source in OBS

#### More information can be seen in readme of the based project [gosumemory](https://github.com/l3lackShark/gosumemory)

## Build

1. If you haven't installed [Go](https://go.dev/) yet, download and install it first
    * Requires v1.20 or newer to build
2. Clone the repository
3. Execute `go get -u` in terminal
4. Execute `go build` in terminal

## Credits

* [gosumemory](https://github.com/l3lackShark/gosumemory) - The original idea for wiimmfi-room-watcher
