# wiimmfi-room-watcher

`The tool to re-structure the room statistics from Wiimmfi for readability and provide custom overlays on live streaming`

----

> Wiimmfi is currently unable to provide the exact room statistics (such as joined players and game modes). Due to this, wiimmfi-room-watcher may not work properly.

## Usage

1. Download the latest release, and unzip it anywhere
   * ex. If you're a 64-bit Windows user, you need to download **wiimmfi-room-watcher-win-amd64**.
   >    In the root folder of the program, you can find the **static** directory. It contains all of the available overlays. Those are getting streamed via HTTP-File server.
2. Run wiimmfi-room-watcher
   * Visit the link in your browser (http://localhost:24050), and choose the layout.
   * Add a browser source in OBS

#### More technically information can be seen in readme of the based project [gosumemory](https://github.com/l3lackShark/gosumemory)

## Build

1. If you haven't installed [Go](https://go.dev/) yet, download and install it first
   * Requires v1.22 or newer
2. Clone the repository
3. Execute `go build` in terminal

## Credits

* [gosumemory](https://github.com/l3lackShark/gosumemory) - The original idea for wiimmfi-room-watcher
* [tosu](https://github.com/KotRikD/tosu) - The original idea for the graphical overlay manager
