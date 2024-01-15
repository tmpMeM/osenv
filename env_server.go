package osenv

import (
	"os"
)

// server port :port :8080
// SERVER_PORT
func GetServerPort() string {
	return os.Getenv("SERVER_PORT")
}

// server db redis address
// SERVER_DB_REDIS_ADDRESS 127.0.0.1:6379
func GetServerDbRedisAddress() string {
	return os.Getenv("SERVER_DB_REDIS_ADDRESS")
}

// 数据库类型, mysql,sqlite,postgres 创建的mysql数据库数据集建议使用utf8mb4
// SERVER_DB_TYPE
func GetServerDbType() string {
	return os.Getenv("SERVER_DB_TYPE")
}

// 数据库地址，如本地为 127.0.0.1
// SERVER_DB_HOST
func GetServerDbHost() string {
	return os.Getenv("SERVER_DB_HOST")
}

// 数据库端口，mysql默认为3306 postgres默认端口为5432
// SERVER_DB_PORT
func GetServerDbPort() string {
	return os.Getenv("SERVER_DB_PORT")
}

// 连接地址，形式如 数据库域名:端口号
// SERVER_DB_ADDRESS
func GetServerDbAddress() string {
	return os.Getenv("SERVER_DB_ADDRESS")
}

// 数据库用户名
// SERVER_DB_USERNAME
func GetServerDbUsername() string {
	return os.Getenv("SERVER_DB_USERNAME")
}

// 数据库用户名密码
// SERVER_DB_PASSWORD
func GetServerDbPassword() string {
	return os.Getenv("SERVER_DB_PASSWORD")
}

// 数据库名称，在mysql或psql中创建的数据库名称
// SERVER_DB_NAME
func GetServerDbName() string {
	return os.Getenv("SERVER_DB_NAME")
}
