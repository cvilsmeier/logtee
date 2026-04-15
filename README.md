# logtee

Logtee is a tool for writing a stream of log lines into rolling log files, optionally with compression.

Technically, logtee is just a small CLI wrapper for the fantastic https://github.com/DeRuina/timberjack library.


## Installation

Have [Go](https://go.dev/dl/) installed, then: 

```bash
go install github.com/cvilsmeier/logtee@latest
```

## Usage

~~~
$ logtee --help
Logtee is a tool for writing a stream of log lines
into rolling log files, optionally with compression.

Usage:

  logtee [options...]

Options:

  -file string
      Filename to write logs to (default "out.log")
  -size int
      Max. file size in MB (default 100)
  -backups int
      Max. backup files to keep (default 10)
  -compress string
      File compression mode: "none", "gzip" or "zstd" (default "gzip")
  -stdout
      Write additionally to stdout
  -version
      Print version
  -h, --help
      Print help and exit

Examples:

  myprogram | logtee

    This command runs myprogram, pipes its stdout to logtee, which
    then writes all output to out.log. When out.log grows beyond
    100 MB, it gets 'gzip' compressed and moved to a timestamped
    backup file. It keeps 10 backup files.

  myprogram | logtee -file my.json -size 20 -backups 30 

    Same as above, but the file is my.json (instead of out.log),
    the max. file size is 20 MB (instead of 100 MB),
    and it keeps 30 backup files (instead of 10).
~~~


