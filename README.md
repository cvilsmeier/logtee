# logtee

[![GoDoc Reference](https://pkg.go.dev/badge/github.com/cvilsmeier/logtee)](http://godoc.org/github.com/cvilsmeier/logtee)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/license/mit)

Logtee is a tool for writing a stream of log lines into rolling log files, optionally with compression.

## Installation

```bash
go install github.com/cvilsmeier/logtee@latest
```

## Usage

~~~
Usage:

  logtee [options...]

Options:

  -c string
    	File compression mode: 'none'|'gzip'|'zstd'. (default "gzip")
  -f string
    	Filename to write logs to. (default "out.log")
  -n int
    	Max. backup files file size in MB. (default 10)
  -s int
    	Max. file size in MB. (default 100)
  -stdout
    	Write stdout also. Might be useful for debugging.

Examples:

  myprogram | logtee

    This command runs myprogram, pipes its stdout to logtee, which
    then writes all output to out.log. When out.log grows beyond
    100 MB, it gets 'gzip' compressed and moved to a timestamped
    backup file. The default settings keep 10 backup files, old
    files are deleted.

  myprogram | logtee -f my.log -s 20 -n 30 

    Same as above, but the file is my.log (instead of out.log), the
    max. file size is 20 MB (instead of 100 MB) and the max. number
    of backup files is 30 (instead of 10).
~~~


## License

The MIT License (MIT)

Copyright (c) 2026 C.Vilsmeier

See [LICENSE](LICENSE)
