package config

const (
	Name          = "GenPassaaS"
	MinCount      = 1
	MinLength     = 6
	MaxCount      = 64
	MaxLength     = 32
	DefaultCount  = 10
	DefaultLength = 16
)

func ClampCount(v int) (rv int) {
	return clamp(v, MinCount, MaxCount)
}

func ClampLength(v int) (rv int) {
	return clamp(v, MinLength, MaxLength)
}

func clamp(v, min, max int) (rv int) {
	switch {
	case v < min:
		return min
	case v > max:
		return max
	}

	return v
}
