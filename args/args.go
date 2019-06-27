package args

import (
	"flag"
)

type Args struct {
	Host     string
	Port     int
	Username string
	Password string
	Dbname   string
}

func Parse() *Args {
	host := flag.String("h", "127.0.0.1", "hostname (default=127.0.0.1)")
	port := flag.Int("p", 5432, "port (default=5432)")
	username := flag.String("U", "", "username")
	password := flag.String("W", "", "password")
	dbname := flag.String("d", "", "dbname")

	flag.Parse()

	return &Args{
		*host,
		*port,
		*username,
		*password,
		*dbname,
	}
}
