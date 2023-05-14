package conf

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

// Bootstrap is the bootstrap config
// type Bootstrap struct {
// 	Server  *Server
// 	Data    *Data
// 	Logger  *Logger
// 	Workdir string
// 	Docker  *Docker
// }

type Application struct {
	IsMaster bool   `yaml:"is_master"`
	Workdir  string `yaml:"workdir"`
}

// Server is the server config
type Server struct {
	Addr string
}

type Data struct {
	Database *Database `yaml:"database"`
	Redis    *Redis    `yaml:"redis"`
}

// Database is the database config
type Database struct {
	DSN string `yaml:"dsn"`
}

type Redis struct {
	Network      string
	Addr         string
	Password     string
	Database     int32
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	MaxIdle      int32
	MaxActive    int32
	IdleTimeout  int32
}

type Logger struct {
	File    *FileLogger
	Console *ConsoleLogger
	Gorm    *GormLogger
}

type FileLogger struct {
	Name       string
	Path       string
	Level      string
	MaxSize    int
	MaxAge     int
	MaxBackups int
	LocalTime  bool
	Compress   bool
}

type ConsoleLogger struct {
	Level  string
	Format string
	Color  bool
}

type GormLogger struct {
	Level                     string        `yaml:"level"`
	SlowThreshold             time.Duration `yaml:"slow_threshold"`
	Caller                    bool          `yaml:"caller"`
	IgnoreRecordNotFoundError bool          `yaml:"ignore_record_not_found_error"`
}

type Docker struct {
	Host string
}

func (g *GormLogger) UnmarshalYAML(unmarshal func(interface{}) error) error {
	type alias struct {
		Level                     string `yaml:"level"`
		SlowThreshold             string `yaml:"slow_threshold"`
		Caller                    bool   `yaml:"caller"`
		IgnoreRecordNotFoundError bool   `yaml:"ignore_record_not_found_error"`
	}

	var tmp alias
	if err := unmarshal(&tmp); err != nil {
		return err
	}

	t, err := time.ParseDuration(tmp.SlowThreshold)

	if err != nil {
		return fmt.Errorf("failed to parse '%s' to time.Duration: %v", tmp.SlowThreshold, err)
	}
	g.Level = tmp.Level
	g.IgnoreRecordNotFoundError = tmp.IgnoreRecordNotFoundError
	g.Caller = tmp.Caller
	g.SlowThreshold = t
	return nil
}

func Load(file string, obj interface{}) error {

	content, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	// viper.SetConfigType("yaml")

	// viper.ReadConfig(bytes.NewBuffer(content))

	err = yaml.Unmarshal(content, obj)

	//var bootstrap Bootstrap
	if err != nil {
		return err
	}
	return nil
}
func Test() {
	fmt.Println("test")
}
