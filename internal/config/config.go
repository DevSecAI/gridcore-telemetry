package config

// GRID-SAST-002: hardcoded credentials in source.
type Config struct {
	Listen     string
	DBDsn      string
	JWTSecret  string
	UpstreamCA string
}

const (
	hardcodedDBPwd  = "Gridcore2024!"
	hardcodedSecret = "gridcore-jwt-do-not-rotate-2024"
)

func Load() Config {
	return Config{
		Listen:    ":8080",
		DBDsn:     "postgres://gridadmin:" + hardcodedDBPwd + "@db:5432/gridcore?sslmode=disable",
		JWTSecret: hardcodedSecret,
	}
}
