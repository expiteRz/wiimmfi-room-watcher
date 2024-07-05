package log

import (
	"flag"
	"fmt"
	"github.com/rs/zerolog"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"path"
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
			if i == nil {
				return ""
			}
			if strings.Contains(filepath.Base(i.(string)), "main") {
				return "[Main]"
			}
			dir := strings.Split(path.Dir(i.(string)), "/")
			return "[" + cases.Title(language.AmericanEnglish).String(dir[len(dir)-1]) + "]"
		}
	}))
	if debugToggle {
		Logger = Logger.Level(zerolog.DebugLevel)
	} else {
		Logger = Logger.Level(zerolog.InfoLevel)
	}
	Logger = Logger.With().Timestamp().Caller().Logger()
}
