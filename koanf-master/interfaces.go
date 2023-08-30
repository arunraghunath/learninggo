package koanf

type Provider interface {
	ReadBytes() ([]byte, error)
}
type Parser interface {
	Unmarshal([]byte) (map[string]interface{}, error)
}
