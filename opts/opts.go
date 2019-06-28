package opts

import (
	"os"

	flags "github.com/jessevdk/go-flags"
)

type Options struct {
	Host     string `long:"host" short:"h" description:"hostname" default:"127.0.0.1"`
	Port     int    `long:"port" short:"p" description:"port" default:"5432"`
	Username string `long:"username" short:"U" description:"username" required:"true"`
	Password string `long:"password" short:"W" description:"password" required:"true"`
	Dbname   string `long:"dbname" short:"d" description:"dbname" required:"true"`
}

func Parse() *Options {
	var options Options
	parser := flags.NewParser(&options, flags.Default)
	if _, err := parser.Parse(); err != nil {
		parser.WriteHelp(os.Stderr)
		os.Exit(2)
	}

	return &options
}
