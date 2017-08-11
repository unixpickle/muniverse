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
			defer env.Close()
			for i := 0; i < 2; i++ {
				if err := env.Reset(); err != nil {
					t.Error(err)
					t.Log(env.Log())
					break
				}
				if _, _, err := env.Step(time.Millisecond); err != nil {
					t.Error(err)
					t.Log(env.Log())
					break
				}
			}
		})
	}
}

func BenchmarkEnvObserve(b *testing.B) {
	for _, envName := range []string{"PopUp-v0", "FlappyBird-v0"} {
		for _, compression := range []bool{false, true} {
			name := envName
			if compression {
				name += "-Compressed"
			}
			b.Run(name, func(b *testing.B) {
				opts := &Options{}
				if compression {
					opts.Compression = true
					opts.CompressionQuality = 100
				}
				env, err := NewEnvOptions(SpecForName(envName), opts)
				if err != nil {
					b.Fatal(err)
				}
				defer env.Close()
				if err := env.Reset(); err != nil {
					b.Fatal(err)
				}
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					if _, err := env.Observe(); err != nil {
						b.Fatal(err)
					}
				}
			})
		}
	}
}
