package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/DeRuina/timberjack"
)

func main() {
	// cli flags
	flag.Usage = func() {
		fmt.Printf("Logtee is a tool for writing a stream of log lines\n")
		fmt.Printf("into rolling log files, optionally with compression.\n")
		fmt.Printf("\n")
		fmt.Printf("Usage:\n")
		fmt.Printf("\n")
		fmt.Printf("  logtee [options...]\n")
		fmt.Printf("\n")
		fmt.Printf("Options:\n")
		fmt.Printf("\n")
		flag.PrintDefaults()
		fmt.Printf("\n")
		fmt.Printf("Examples:\n")
		fmt.Printf("\n")
		fmt.Printf("  myprogram | logtee\n")
		fmt.Printf("\n")
		fmt.Printf("    This command runs myprogram, pipes its stdout to logtee, which\n")
		fmt.Printf("    then writes all output to out.log. When out.log grows beyond\n")
		fmt.Printf("    100 MB, it gets 'gzip' compressed and moved to a timestamped\n")
		fmt.Printf("    backup file. It keeps 10 backup files.\n")
		fmt.Printf("\n")
		fmt.Printf("  myprogram | logtee -f my.json -s 20 -n 30 \n")
		fmt.Printf("\n")
		fmt.Printf("    Same as above, but the file is my.json (instead of out.log),\n")
		fmt.Printf("    the max. file size is 20 MB (instead of 100 MB),\n")
		fmt.Printf("    and it keeps 30 backup files (instead of 10).\n")
	}
	filenameFlag := flag.String("f", "out.log", "Filename to write logs to.")
	maxSizeFlag := flag.Int("s", 100, "Max. file size in MB.")
	maxBackupsFlag := flag.Int("n", 10, "Max. backup files file size in MB.")
	compressionFlag := flag.String("c", "gzip", "File compression mode: 'none'|'gzip'|'zstd'.")
	stdoutFlag := flag.Bool("stdout", false, "Write stdout also. Might be useful for debugging.")
	flag.Parse()
	// initialize timerjack logger
	timber := &timberjack.Logger{
		Filename:   *filenameFlag,   // Choose an appropriate path
		MaxSize:    *maxSizeFlag,    // megabytes
		MaxBackups: *maxBackupsFlag, // backups
		// MaxAge:             28,                         // days
		Compression: *compressionFlag, // "none" | "gzip" | "zstd"
		LocalTime:   true,             // default: false (use UTC)
		// RotationInterval:   24 * time.Hour,             // Rotate daily if no other rotation met
		// RotateAtMinutes:    []int{0, 15, 30, 45},       // Also rotate at HH:00, HH:15, HH:30, HH:45
		// RotateAt:           []string{"00:00", "12:00"}, // Also rotate at 00:00 and 12:00 each day
		// BackupTimeFormat: "2006-01-02T150405", // Rotated files will have format <logfilename>-2006-01-02T150405-<reason>.log
		// AppendTimeAfterExt: true,                       // put timestamp after ".log" (foo.log-<timestamp>-<reason>)
		FileMode: 0o644, // Custom permissions for newly created files. If unset or 0, defaults to 640.
	}
	defer timber.Close() // close logger to stop background goroutines
	// read stdin line by line
	sca := bufio.NewScanner(os.Stdin)
	for sca.Scan() {
		line := sca.Bytes()
		// stdout
		if *stdoutFlag {
			os.Stdout.Write(line)
			os.Stdout.Write(newline)
		}
		// file
		timber.Write(line)
		timber.Write(newline)
	}
	err := sca.Err()
	if err != nil && !errors.Is(err, io.EOF) {
		log.Fatal(err)
	}
}

var newline = []byte{10}
