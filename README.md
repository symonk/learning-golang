#### Learning go

-----

A simple repository for documenting and learning the basic functionality of `go`.  The repository is simple,
each module contains a `main.go` and the entrypoint to that module is the `Run` method.  Creating multiple
`main` functions in each package is not viable.  The `main.go` in the root directory synchronously calls
all the examples `Run` functions sequentially in line with the table of contents below.

All modules are exposed under the `-module` flag.  To run a particular exercise, take the top level folder name
and run it like so:

```console
# To run the `closingchannels` examples.
go run . -module closingchannels
```

-----


### Table Of Contents:

* [01 - Hello World](helloworld/main.go)
* [02 - Values](values/main.go)
* [03 - Variables](variables/main.go)
* [04 - Constants](constants/main.go)
* [05 - For](forloop/main.go)
* [06 - If, else if and else](ifelse/main.go)
* [07 - Switch](switches/main.go)
* [08 - Arrays](arrays/main.go)
* [09 - Slices](slices/main.go)
* [10 - Maps](maps/main.go)
* [11 - Range](ranges/main.go)
* [12 - Functions](functions/main.go)
* [13 - Variadic Functions](variadicfunctions/main.go)
* [14 - Closures](closures/main.go)
* [15 - Recursion](recursion/main.go)
* [16 - Pointers](pointers/main.go)
* [17 - Strings & Runes](stringsrunes/main.go)
* [18 - Structs](structs/main.go)
* [19 - Methods](methods/main.go)
* [20 - Interfaces](interfaces/main.go)
* [21 - Struct Embedding](structembedding/main.go)
* [22 - Generics](generics/main.go)
* [23 - Errors](errors/main.go)
* [24 - Goroutines](goroutines/main.go)
* [25 - Channels](channels/main.go)
* [26 - Buffered Channels](bufferedchannels/main.go)
* [27 - Channel Synchronisation](channelsynchronisation/main.go)
* [28 - Channel Directions](channeldirections/main.go)
* [29 - Selects](selects/main.go)
* [30 - Timeouts](timeouts/main.go)
* [31 - Non Blocking Channel Operations](nonblockingchannelops/main.go)
* [32 - Closing Channels](closingchannels/main.go)
* [33 - Range Over Channels](rangeoverchannels/main.go)
* [34 - Timers](timers/main.go)
* [35 - Tickers](tickers/main.go)
* [36 - Worker Pools](workerpools/main.go)
* [37 - Waitgroups](waitgroups.go)
* [38 - Rate Limiting](ratelimiting/main.go)
* [39 - Atomic Counters](atomiccounters/main.go)
* [40 - Mutexes](mutexes/main.go)
* [41 - Stateful Goroutines](statefulgoroutines/main.go)
* [42 - Sorting](sorting/main.go)
* [43 - Custom Comparator Functions](sortingbyfunctions/main.go)
* [44 - Panic](panics/main.go)
* [45 - Defer](defers/main.go)
* [46 - Recover](recover/main.go)
* [47 - String Functions](stringfunctions/main.go)
* [48 - String Formatting](stringformatting/main.go)
* [49 - Text Templating](texttemplates/main.go)
* [50 - Regexp](regex/main.go)
* [51 - Json](json/main.go)
* [52 - Xml](xml/main.go)
* [53 - Time](time/main.go)
* [54 - Epoch](epoch/main.go)
* [55 - Time Formatting](timeformatting/main.go)
* [56 - Random Numbers](randomnumbers/main.go)
* [57 - Number Parsing](numberparsing/main.go)
* [58 - URL Parsing](urlparsing/main.go)
* [59 - SHA256 Parsing](sha256/main.go)
* [60 - Base64 Encoding](base64/main.go)
* [61 - Reading Files](readingfiles/main.go)
* [62 - Writing Files](writingfiles/main.go)
* [63 - Line Filters](linefilters/main.go)
* [64 - File Paths](filepaths/main.go)
* [65 - Directories](directories/main.go)
* [66 - Temporary Files & Directories](temporaryfiledirs/main.go)
* [67 - Embed Directive](embedding/main.go)
* [68 - Testing & Benchmarking](testingbenchmarking/main.go)
* [69 - Command Line Arguments](cli/main.go)
* [70 - Command Line Flags](clf/main.go)
* [71 - Command Line Subcommands](cls/main.go)
* [72 - Environment Vars](environment/main.go)
* [73 - Logging](logging/main.go)
* [74 - HTTP Client](httpclient/main.go)
* [75 - HTTP Server](httpserver/main.go)
* [76 - Contexts](contexts/main.go)
* [77 - Spawning Processes](spawningprocesses/main.go)
* [78 - Execing Processes](execingprocesses/main.go)
* [79 - Signals](signals/main.go)
* [80 - Exit](exit/main.go)
