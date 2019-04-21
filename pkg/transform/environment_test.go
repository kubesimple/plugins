package transform

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEnvironment(t *testing.T) {
	tcs := []struct {
		Name      string
		Env       map[string]string
		Expect    Environment
		ShouldErr bool
	}{
		{Name: "ok", Env: map[string]string{"KS_BUILD_BASEPATH": "/test", "KS_BUILD_ENVIRONMENT": "hosted"}, Expect: Environment{Hosted, "/test"}, ShouldErr: false},
		{Name: "no basepath", Env: map[string]string{"KS_BUILD_ENVIRONMENT": "hosted"}, Expect: Environment{}, ShouldErr: true},
		{Name: "unk environment", Env: map[string]string{"KS_BUILD_ENVIRONMENT": "7-11"}, Expect: Environment{}, ShouldErr: true},
		{Name: "default dev", Env: map[string]string{"KS_BUILD_BASEPATH": "/test"}, Expect: Environment{Development, "/test"}, ShouldErr: false},
	}

	for _, tc := range tcs {
		t.Run(tc.Name, func(t *testing.T) {
			for _, key := range []string{"KS_BUILD_BASEPATH", "KS_BUILD_ENVIRONMENT"} {
				os.Unsetenv(key)
			}
			for k, v := range tc.Env {
				os.Setenv(k, v)
			}
			e, err := GetEnvironment()
			switch tc.ShouldErr {
			case true:
				assert.Error(t, err)
			default:
				assert.NoError(t, err)
				assert.Equal(t, tc.Expect, e)
			}
		})
	}
}
