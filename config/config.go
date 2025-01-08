package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var configFile *appConfig

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func GetConfig() appConfig {
	if configFile != nil {
		return *configFile
	}

	appConfig := appConfig{
		DBUser: dbConfig{
			Write: dbConfigObj{
				UserName: os.Getenv("DB_USER_USERNAME"),
				Password: os.Getenv("DB_USER_PASSWORD"),
				Host:     os.Getenv("DB_USER_HOST"),
				Port:     parseUint(os.Getenv("DB_USER_PORT")),
				DBname:   os.Getenv("DB_USER_DBNAME"),
			},
			MaxIdleConns:    parseUint(os.Getenv("DB_USER_MAXIDLECONNS")),
			MaxOpenConns:    parseUint(os.Getenv("DB_USER_MAXOPENCONNS")),
			ConnMaxLifetime: parseUint(os.Getenv("DB_USER_CONNMAXLIFETIME")),
		},
		Redis: redisConfig{
			Host:     os.Getenv("REDIS_HOST"),
			Port:     parseUint(os.Getenv("REDIS_PORT")),
			DBIndex:  parseUint(os.Getenv("REDIS_DBINDEX")),
			UserName: os.Getenv("REDIS_USERNAME"),
			Password: os.Getenv("REDIS_PASSWORD"),
		},
		CodePush: codePush{
			FileLocal: os.Getenv("CODEPUSH_FILELOCAL"),
			Local: localConfig{
				SavePath: os.Getenv("CODEPUSH_LOCAL_SAVEPATH"),
			},
			Aws: awsConfig{
				Endpoint:         os.Getenv("CODEPUSH_AWS_ENDPOINT"),
				Region:           os.Getenv("CODEPUSH_AWS_REGION"),
				S3ForcePathStyle: parseBool(os.Getenv("CODEPUSH_AWS_S3FORCEPATHSTYLE")),
				KeyId:            os.Getenv("CODEPUSH_AWS_KEYID"),
				Secret:           os.Getenv("CODEPUSH_AWS_SECRET"),
				Bucket:           os.Getenv("CODEPUSH_AWS_BUCKET"),
			},
			Ftp: ftpConfig{
				ServerUrl: os.Getenv("CODEPUSH_FTP_SERVERURL"),
				UserName:  os.Getenv("CODEPUSH_FTP_USERNAME"),
				Password:  os.Getenv("CODEPUSH_FTP_PASSWORD"),
			},
		},
		UrlPrefix:       os.Getenv("URL_PREFIX"),
		Port:            os.Getenv("PORT"),
		ResourceUrl:     os.Getenv("RESOURCE_URL"),
		TokenExpireTime: parseInt64(os.Getenv("TOKEN_EXPIRE_TIME")),
	}
	configFile = &appConfig
	return *configFile
}

// Fungsi untuk parsing string ke uint
func parseUint(s string) uint {
	value, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		log.Fatalf("Error parsing uint: %v", err)
	}
	return uint(value)
}

// Fungsi untuk parsing string ke bool
func parseBool(s string) bool {
	value, err := strconv.ParseBool(s)
	if err != nil {
		log.Fatalf("Error parsing bool: %v", err)
	}
	return value
}

// Fungsi untuk parsing string ke int64
func parseInt64(s string) int64 {
	value, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Fatalf("Error parsing int64: %v", err)
	}
	return value
}

type modeConfig struct {
	Mode string
}

type appConfig struct {
	DBUser          dbConfig
	Redis           redisConfig
	CodePush        codePush
	UrlPrefix       string
	Port            string
	ResourceUrl     string
	TokenExpireTime int64
}
type dbConfig struct {
	Write           dbConfigObj
	MaxIdleConns    uint
	MaxOpenConns    uint
	ConnMaxLifetime uint
}
type dbConfigObj struct {
	UserName string
	Password string
	Host     string
	Port     uint
	DBname   string
}
type redisConfig struct {
	Host     string
	Port     uint
	DBIndex  uint
	UserName string
	Password string
}
type codePush struct {
	FileLocal string
	Local     localConfig
	Aws       awsConfig
	Ftp       ftpConfig
}
type awsConfig struct {
	Endpoint         string
	Region           string
	S3ForcePathStyle bool
	KeyId            string
	Secret           string
	Bucket           string
}
type ftpConfig struct {
	ServerUrl string
	UserName  string
	Password  string
}
type localConfig struct {
	SavePath string
}
