package log

import (
	"github.com/rs/zerolog"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"path"
	"path/filepath"
	"strings"
	"time"
)

var Logger zerolog.Logger

func init() {
	Logger = zerolog.New(zerolog.NewConsoleWriter(func(w *zerolog.ConsoleWriter) {
		w.TimeFormat = time.DateTime
		//w.FormatLevel = func(i interface{}) string {
		//	return strings.ToUpper(fmt.Sprint("[", i, "]"))
		//}
		w.FormatCaller = func(i interface{}) string {
			if strings.Contains(filepath.Base(i.(string)), "main") {
				return "[Main]"
			}
			dir := strings.Split(path.Dir(i.(string)), "/")
			return "[" + cases.Title(language.AmericanEnglish).String(dir[len(dir)-1]) + "]"
		}
	})).Level(zerolog.DebugLevel).With().Timestamp().Caller().Logger()
}
