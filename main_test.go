package main

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	kubus         Kubus   = Kubus{12}
	volumeValue   float64 = 1728
	luasValue     float64 = 864
	kelilingValue float64 = 144
)

// Before and After Test
func TestMain(m *testing.M) {
	// Kode sebelum test
	fmt.Println("BEFORE TEST, START!")

	// run all test
	m.Run()

	// Kode setelah test
	fmt.Println("AFTER TEST, DONE!")
}

// benchmarking
func BenchmarkHitungKubus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		kubus.keliling()
		kubus.volume()
		kubus.luas()
	}
}

// func TestKelilingKubus(t *testing.T) {
// 	t.Log("Hasil dari keliling kubus:", kubus.keliling())
// 	if kubus.keliling() != kelilingValue {
// 		t.Errorf("Hasil keliling tidak sesuai yang seharusnya %.2f", kelilingValue)
// 	}
// }

// func TestLuasKubus(t *testing.T) {
// 	t.Log("Hasil dari luas kubus:", kubus.luas())
// 	if kubus.luas() != luasValue {
// 		t.Errorf("Hasil luas tidak sesuai yang seharusnya %.2f", luasValue)
// 	}
// }

// func TestVolumeKubus(t *testing.T) {
// 	t.Log("Hasil dari volume kubus:", kubus.volume())
// 	if kubus.volume() != volumeValue {
// 		t.Errorf("Hasil volume tidak sesuai yang seharusnya %.2f", volumeValue)
// 	}
// }

// Test menggunakan testify
// func TestKelilingKubus(t *testing.T) {
// 	assert.Equal(t, kubus.keliling(), kelilingValue, "Hasil keliling salah!")
// }
// func TestLuasKubus(t *testing.T) {
// 	assert.Equal(t, kubus.luas(), luasValue, "Hasil luas salah!")
// }
// func TestVolumeKubus(t *testing.T) {
// 	assert.Equal(t, kubus.volume(), volumeValue, "Hasil volume salah!")
// }

// subtest
func TestReturnValueType(t *testing.T) {
	t.Run("Keliling", func(t *testing.T) {
		keliling := kubus.keliling()
		typeOf := reflect.ValueOf(keliling).Kind()
		require.Equal(t, "float64", typeOf.String())
	})
	t.Run("Luas", func(t *testing.T) {
		luas := kubus.luas()
		typeOf := reflect.ValueOf(luas).Kind()
		require.Equal(t, "float64", typeOf.String())
	})
	t.Run("Volume", func(t *testing.T) {
		volume := kubus.volume()
		typeOf := reflect.ValueOf(volume).Kind()
		require.Equal(t, "float64", typeOf.String())
	})
}

// test table
func TestTableKubus(t *testing.T) {
	cubes := []struct {
		cube             Kubus
		name             string
		kelilingExpected float64
		luasExpected     float64
		volumeExpected   float64
	}{
		{
			cube:             Kubus{2},
			name:             "sisi_2",
			kelilingExpected: 24,
			luasExpected:     24,
			volumeExpected:   8,
		}, {
			cube:             Kubus{4},
			name:             "sisi_4",
			kelilingExpected: 48,
			luasExpected:     96,
			volumeExpected:   64,
		},
	}

	// keliling test
	for _, cube := range cubes {
		t.Run(cube.name, func(t *testing.T) {
			t.Logf("Hasil Keliling: %.2f", cube.cube.keliling())
			if cube.cube.keliling() != cube.kelilingExpected {
				t.Errorf("Hasil tidak sesuai! hasil yang diharapkan: %.2f", cube.kelilingExpected)
			}
		})
	}

	// luas test
	for _, cube := range cubes {
		t.Run(cube.name, func(t *testing.T) {
			t.Logf("Hasil Luas: %.2f", cube.cube.luas())
			if cube.cube.luas() != cube.luasExpected {
				t.Errorf("Hasil tidak sesuai! hasil yang diharapkan: %.2f", cube.luasExpected)
			}
		})
	}

	// volume test
	for _, cube := range cubes {
		t.Run(cube.name, func(t *testing.T) {
			t.Logf("Hasil Volume: %.2f", cube.cube.volume())
			if cube.cube.volume() != cube.volumeExpected {
				t.Errorf("Hasil tidak sesuai! hasil yang diharapkan: %.2f", cube.volumeExpected)
			}
		})
	}
}
