package common

import "time"

type Config interface {
	HasPath(path string) bool
	GetBoolean(path string, val ...bool) bool
	GetString(path string, val ...string) string
	GetFloat64(path string, val ...float64) float64
	GetInt32(path string, val ...int32) int32
	GetStringList(path string) []string
	GetTimeDurationInfiniteNotAllowed(path string, val ...time.Duration) time.Duration
	GetTimeDuration(path string, val ...time.Duration) time.Duration
	GetObject(path string) Config
	GetStringMap(path string) map[string]Config
}

func Required[T any](path string, c Config, get func(string, ...T) T) (t T) {
	if c.HasPath(path) {
		return get(path)
	}
	return
}
func Exists[T any](path string, c Config, get func(string, ...T) T, act func(T)) {
	if c.HasPath(path) {
		act(get(path))
	}
}
