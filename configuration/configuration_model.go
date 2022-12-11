package configuration

type Config struct {
	Vadek      Vadek       `mapstructure:"vadek"`
	SQLite3    *SQLite3    `mapstructure:"sqlite3"`
	Mysql      *Mysql      `mapstructure:"mysql"`
	PostgreSQL *PostgreSQL `mapstructure:"postgresql"`
	Log        *Log        `mapstructure:"log"`
}

// app相关
type Vadek struct {
	Mode              string `mapstructure:"mode"`
	WorkDir           string `mapstructure:"work_dir"`
	UploadDir         string
	LogDir            string `mapstructure:"log_dir"`
	TemplateDir       string `mapstructure:"template_dir"`
	ThemeDir          string
	AdminResourcesDir string
}

// 数据库相关
type SQLite3 struct {
	Enable bool `mapstructure:"enable"`
	File   string
}
type Mysql struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	DB       string `mapstructure:"db"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}
type PostgreSQL struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	DB       string `mapstructure:"db"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

// 日志相关
type Log struct {
	FileName string `mapstructure:"filename"`
	Levels   Levels `mapstructure:"level"`
	MaxSize  int    `mapstructure:"maxsize"`
	MaxAge   int    `mapstructure:"maxage"`
	Compress bool   `mapstructure:"compress"`
}
type Levels struct {
	App  string `mapstructure:"app"`
	Gorm string `mapstructure:"gorm"`
}
