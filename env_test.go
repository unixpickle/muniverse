package muniverse

import (
	"testing"
	"time"
)

func TestEnvs(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping in short mode")
	}
	for _, spec := range EnvSpecs {
		t.Run(spec.Name, func(t *testing.T) {
			env, err := NewEnv(spec)
			if err != nil {
				t.Error(err)
				return
			}
			for i := 0; i < 2; i++ {
				if err := env.Reset(); err != nil {
					t.Error(err)
					break
				}
				if _, _, err := env.Step(time.Millisecond); err != nil {
					t.Error(err)
					break
				}
			}
		})
	}
}
