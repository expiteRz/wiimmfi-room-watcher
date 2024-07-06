//go:build debug

package log

import (
	"flag"
	"fmt"
	"github.com/rs/zerolog"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var debugToggle bool
var Logger zerolog.Logger

func init() {
	flag.BoolVar(&debugToggle, "debug", false, "Toggle for debug mode to print the current state of room stat parsing")
	flag.Parse()

	Logger = zerolog.New(zerolog.NewConsoleWriter(func(w *zerolog.ConsoleWriter) {
		w.TimeFormat = time.DateTime
		w.FormatLevel = func(i interface{}) string {
			level, err := zerolog.ParseLevel(i.(string))
			if err != nil {
				return "[UNKNOWN]"
			}
			clr := zerolog.LevelColors[level]
			if level == zerolog.DebugLevel {
				clr = 35
			}
			levelStr := fmt.Sprint("[", strings.ToUpper(i.(string)), "]")
			return fmt.Sprintf("\x1b[%dm%v\x1b[0m", clr, levelStr)
		}
		w.FormatCaller = func(i interface{}) string {
			var c string
			if cc, ok := i.(string); ok {
				c = cc
			}
			if len(c) > 0 {
				if cwd, err := os.Getwd(); err == nil {
					if rel, err := filepath.Rel(cwd, c); err == nil {
						c = rel
					}
				}
				c = fmt.Sprintf("\u001B[%dm%v\u001B[0m", 1, c)
			}
			return c
		}
	}))
	if debugToggle {
		Logger = Logger.Level(zerolog.DebugLevel)
	} else {
		Logger = Logger.Level(zerolog.InfoLevel)
	}
	Logger = Logger.With().Timestamp().Caller().Logger()
}
