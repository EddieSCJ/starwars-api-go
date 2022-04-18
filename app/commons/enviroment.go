package commons

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
