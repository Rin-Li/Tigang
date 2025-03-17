package cache

import "time"



const RestCodePrefix = "rest_code:"

func SetResetCode(email string, code string, expiration time.Duration) error {
	return RedisClient.Set(RestCodePrefix + email, code, expiration).Err()
}

func GetResetCode(email string) (string, error) {
	return RedisClient.Get(RestCodePrefix + email).Result()
}

func DelResetCode(email string) error {
	return RedisClient.Del(RestCodePrefix + email).Err()
}