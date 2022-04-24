package commons

import (
	"strconv"
	"time"
)

func GetMongoHost() string {
	return GetEnv("MONGO_HOST", "localhost")
}

func GetMongoUsername() string {
	return GetEnv("MONGO_USER", "")
}

func GetMongoPassword() string {
	return GetEnv("MONGO_PASSWORD", "")
}

func GetMongoPort() string {
	return GetEnv("MONGO_PORT", "27017")
}

func GetMongoDBName() string {
	return GetEnv("MONGO_DB", "development")
}

func GetMongoContainerName() string {
	return GetEnv("MONGO_CONTAINER_NAME", "mongoservice")
}

func GetSWAPIURL() string {
	return GetEnv("SWAPI_URL", "https://swapi.dev/api/")
}

func GetDefaultTimeout() time.Duration {
	const defaultTimeout = 100 * time.Second
	result, err := strconv.Atoi(GetEnv("DEFAULT_TIMEOUT", defaultTimeout.String()))
	if err != nil {
		return defaultTimeout
	}

	return time.Duration(result)
}
