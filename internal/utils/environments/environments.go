package environments

import (
	"os"
	"strconv"
)

var DbHost = os.Getenv("DB_HOST")
var DbPort = os.Getenv("DB_PORT")
var DbName = os.Getenv("DB_NAME")
var DbUser = os.Getenv("DB_USER")
var DbPassword = os.Getenv("DB_PASSWORD")

var MongoUrl = os.Getenv("MONGO_URL")

func GetDbMaxConnections() int {
	maxConnections, err := strconv.Atoi(os.Getenv("DB_MAX_CONNECTIONS"))

	if err != nil {
		maxConnections = 5
	}

	return maxConnections
}

func GetMachineId() (uint16, error) {
	var machineId, err = strconv.ParseUint(os.Getenv("MACHINE_ID"), 10, 64)

	if err != nil {
		panic(err)
	}

	return uint16(machineId), err
}

var RepositoryType = os.Getenv("REPOSITORY")

var RedisUrl = os.Getenv("REDIS_URL")

var RedisPassword = os.Getenv("REDIS_PASSWORD")

func IsCacheEnable() bool {
	isCacheEnable, err := strconv.ParseBool(os.Getenv("ENABLE_CACHE"))

	if err != nil {
		return false
	}

	return isCacheEnable
}
