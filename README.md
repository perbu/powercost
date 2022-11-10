# powercost

Silly little cli to query an API with electricity prices for Norway.

## Building
There are pre-build binaries for most architectures in the release sections.

To build from source, you need to have go installed. Then run:
```bash
go build -o pwrcost main.go
```
Move the `pwrcost` binary to a directory in your PATH.
## Usage

```
$ pwrcost

 0.45 ┤                       ╭─╮
 0.41 ┤                      ╭╯ ╰─╮
 0.37 ┤                    ╭─╯    ╰────╮
 0.34 ┤                   ╭╯           ╰───────────╮
 0.30 ┼─╮               ╭─╯                        ╰────────╮
 0.26 ┤ ╰────────╮   ╭──╯                                   ╰─╮
 0.22 ┤          ╰───╯                                        ╰─╮
 0.19 ┤                                                         ╰───╮
 0.15 ┤                                                             ╰────╮
 0.11 ┤                                                                  ╰──╮
 0.08 ┤                                                                     ╰─
                             Prices for 2022-11-10 in NO1
     00 01 02 03 04 05 06 07 08 09 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24
```

You can add the flag "-tomorrow" and it'll try to give you the prices for tomorrow. If the prices
are not published yet, it will print an error message to stderr.
```bash
pwrcost -tomorrow
```
You can also add the flag "-zone" to display a different price zone.
```bash
pwrcost -zone=NO2
```

See also `pwrcost -h` for more usage information.

## Todo
 - Highlight the price for the current hour
 - Tests
 - Probably lots. This is a quick hack.

## License
See LICENSE.md