# powercost

Silly little cli to query an API with electricity prices for Norway.

## Building
There are pre-build binaries for most architectures in the release sections.

To build from source, you need to have go installed. Then run:
```bash
go build -o pwrcost cmd/main.go
```
Move the pwrcost binary to a directory in your PATH.
## Usage

```
powercost

 0.78 ┤                       ╭──╮
 0.72 ┤                     ╭─╯  ╰╮
 0.67 ┤                   ╭─╯     ╰───────╮           ╭──────╮
 0.62 ┤                  ╭╯               ╰───────────╯      ╰─╮
 0.57 ┤                 ╭╯                                     ╰╮
 0.52 ┤                ╭╯                                       ╰╮
 0.47 ┤                │                                         ╰╮
 0.42 ┼─╮             ╭╯                                          ╰──╮
 0.37 ┤ ╰╮          ╭─╯                                              ╰─────╮
 0.32 ┤  ╰───╮   ╭──╯                                                      ╰─╮
 0.27 ┤      ╰───╯                                                           ╰
                                      2022-11-09
       01:00 03:00 05:00 07:00 09:00 11:00 13:00 15:00 17:00 19:00 21:00 23:00 
    00:00 02:00 04:00 06:00 08:00 10:00 12:00 14:00 16:00 18:00 20:00 22:00 
```

You can add the argument "tomorrow" and it'll try to give you the prices for tomorrow.


## Todo
 - Highlight the price for the current hour
 - Tests
 - Probably lots. This is a quick hack.

## License
See LICENSE.md