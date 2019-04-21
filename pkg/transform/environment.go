package transform

import (
	"fmt"
	"os"
)

type BuildEnv string

const (
	Development BuildEnv = "development"
	Production  BuildEnv = "production"
	Hosted      BuildEnv = "hosted"
)

type Environment struct {
	Type     BuildEnv
	BasePath string
}

func GetEnvironment() (Environment, error) {
	t, ok := os.LookupEnv("KS_BUILD_ENVIRONMENT")
	if !ok {
		t = "development"
	}
	switch BuildEnv(t) {
	case Development, Production, Hosted:
	default:
		return Environment{}, fmt.Errorf("unknown build environment: %s, must be one of development, production, hosted", t)
	}
	p, ok := os.LookupEnv("KS_BUILD_BASEPATH")
	if !ok {
		return Environment{}, fmt.Errorf("no build environment specified, KS_BUILD_BASEPATH must be set")
	}
	return Environment{
		Type:     BuildEnv(t),
		BasePath: p,
	}, nil
}
