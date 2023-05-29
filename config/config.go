package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"strconv"
	"time"
)

func LoadConfig(path string) IConfig {
	envMap, err := godotenv.Read(path)
	if err != nil {
		log.Fatalf("load dotenv failed: %v", err)
	}

	parseInt := func(key string) int {
		value, err := strconv.Atoi(envMap[key])
		if err != nil {
			log.Fatalf("load %s failed: %v", key, err)
		}
		return value
	}

	parseDuration := func(key string) time.Duration {
		value := parseInt(key)
		return time.Duration(value) * time.Second
	}

	return &config{
		app: &app{
			host:         envMap["APP_HOST"],
			port:         parseInt("APP_PORT"),
			name:         envMap["APP_NAME"],
			version:      envMap["APP_VERSION"],
			readTimeout:  parseDuration("APP_READ_TIMEOUT"),
			writeTimeout: parseDuration("APP_WRITE_TIMEOUT"),
			bodyLimit:    parseInt("APP_BODY_LIMIT"),
			fileLimit:    parseInt("APP_FILE_LIMIT"),
			gcpbucket:    envMap["APP_GCP_BUCKET"],
		},
		db: &db{
			host:           envMap["DB_HOST"],
			port:           parseInt("DB_PORT"),
			protocol:       envMap["DB_PROTOCOL"],
			username:       envMap["DB_USERNAME"],
			password:       envMap["DB_PASSWORD"],
			database:       envMap["DB_DATABASE"],
			sslMode:        envMap["DB_SSL_MODE"],
			maxConnections: parseInt("DB_MAX_CONNECTIONS"),
		},
		jwt: &jwt{
			adminKey:         envMap["JWT_SECRET_KEY"],
			secretKey:        envMap["JWT_ADMIN_KEY"],
			apiKey:           envMap["JWT_API_KEY"],
			accessExpiresAt:  parseInt("JWT_ACCESS_EXPIRES"),
			refreshExpiresAt: parseInt("JWT_REFRESH_EXPIRES"),
		},
	}
}

// 2
type config struct {
	app *app
	db  *db
	jwt *jwt
}

// IConfig 6
type IConfig interface {
	App() IAppConfig
	Db() IDbConfig
	Jwt() IJwtConfig
}

// IAppConfig 7
type IAppConfig interface {
	//host:port
	Url() string
	Name() string
	Version() string
	ReadTimeout() time.Duration
	WriteTimeout() time.Duration
	BodyLimit() int
	FileLimit() int
	GCPBucket() string
}

// 3
type app struct {
	host         string
	port         int
	name         string
	version      string
	readTimeout  time.Duration
	writeTimeout time.Duration
	bodyLimit    int //bytes
	fileLimit    int //bytes
	gcpbucket    string
}

// App 10
func (c *config) App() IAppConfig {
	return c.app
}

// Url 11
func (a *app) Url() string                 { return fmt.Sprintf("%s:%d", a.host, a.port) }
func (a *app) Name() string                { return a.name }
func (a *app) Version() string             { return a.version }
func (a *app) ReadTimeout() time.Duration  { return a.readTimeout }
func (a *app) WriteTimeout() time.Duration { return a.writeTimeout }
func (a *app) BodyLimit() int              { return a.bodyLimit }
func (a *app) FileLimit() int              { return a.fileLimit }
func (a *app) GCPBucket() string           { return a.gcpbucket }

// 8
type IDbConfig interface {
	Url() string
	MaxOpenConns() int
}

// 4
type db struct {
	host           string
	port           int
	protocol       string
	username       string
	password       string
	database       string
	sslMode        string
	maxConnections int
}

func (c *config) Db() IDbConfig {
	return c.db
}

func (d *db) Url() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", d.host, d.port, d.username, d.password, d.database, d.sslMode)
}

func (d *db) MaxOpenConns() int {
	return d.maxConnections
}

// 9
type IJwtConfig interface {
	SecretKey() []byte
	AdminKey() []byte
	ApiKey() []byte
	AccessExpiresAt() int
	RefreshExpiresAt() int
	SetJwtAccessExpires(t int)
	SetJwtRefreshExpires(t int)
}

// 5
type jwt struct {
	adminKey         string
	secretKey        string
	apiKey           string
	accessExpiresAt  int //seconds
	refreshExpiresAt int //seconds
}

func (c *config) Jwt() IJwtConfig {
	return c.jwt
}

func (j *jwt) SecretKey() []byte          { return []byte(j.secretKey) }
func (j *jwt) AdminKey() []byte           { return []byte(j.adminKey) }
func (j *jwt) ApiKey() []byte             { return []byte(j.apiKey) }
func (j *jwt) AccessExpiresAt() int       { return j.accessExpiresAt }
func (j *jwt) RefreshExpiresAt() int      { return j.refreshExpiresAt }
func (j *jwt) SetJwtAccessExpires(t int)  { j.accessExpiresAt = t }
func (j *jwt) SetJwtRefreshExpires(t int) { j.refreshExpiresAt = t }
