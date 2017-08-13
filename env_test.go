package muniverse

import (
	"math"
	"testing"
	"time"
)

var BenchmarkEnvs = []string{"PopUp-v0", "FlappyBird-v0"}

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

func TestEnvObserve(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping in short mode")
	}
	for _, spec := range EnvSpecs {
		if !spec.AllCanvas {
			continue
		}
		t.Run(spec.Name, func(t *testing.T) {
			env, err := NewEnv(spec)
			if err != nil {
				t.Error(err)
				return
			}
			defer env.Close()
			if err := env.Reset(); err != nil {
				t.Log(env.Log())
				t.Fatal(err)
			}
			if _, _, err := env.Step(time.Millisecond * 30); err != nil {
				t.Log(env.Log())
				t.Fatal(err)
			}
			actualObs, err := env.Observe()
			if err != nil {
				t.Fatal(err)
			}
			env.(*rawEnv).spec.AllCanvas = false
			expectedObs, err := env.Observe()
			if err != nil {
				t.Fatal(err)
			}

			actual, actualWidth, actualHeight, err := RGB(actualObs)
			if err != nil {
				t.Fatal(err)
			}
			expected, expectedWidth, expectedHeight, err := RGB(expectedObs)
			if err != nil {
				t.Fatal(err)
			}
			if actualWidth != expectedWidth || actualHeight != expectedHeight {
				t.Fatalf("dimensions should be %dx%d but got %dx%d",
					expectedWidth, expectedHeight, actualWidth, actualHeight)
			}
			for i, a := range actual {
				x := expected[i]
				diff := math.Abs(float64(x) - float64(a))
				if math.Abs(diff) > 10 {
					t.Fatalf("pixel %d: got %d, expected %d", i, a, x)
				}
			}
		})
	}
}

func BenchmarkEnvObserve(b *testing.B) {
	for _, envName := range BenchmarkEnvs {
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
				defer b.StopTimer()
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

func BenchmarkEnvReset(b *testing.B) {
	for _, envName := range BenchmarkEnvs {
		name := envName
		b.Run(name, func(b *testing.B) {
			env, err := NewEnv(SpecForName(envName))
			if err != nil {
				b.Fatal(err)
			}
			defer env.Close()
			defer b.StopTimer()
			if err := env.Reset(); err != nil {
				b.Fatal(err)
			}
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				if err := env.Reset(); err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}
