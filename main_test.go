package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	kubus         Kubus   = Kubus{12}
	volumeValue   float64 = 1728
	luasValue     float64 = 864
	kelilingValue float64 = 144
)

// Test menggunakan testify
func TestKelilingKubus(t *testing.T) {
	assert.Equal(t, kubus.keliling(), kelilingValue, "Hasil keliling salah!")
}
func TestLuasKubus(t *testing.T) {
	assert.Equal(t, kubus.luas(), luasValue, "Hasil luas salah!")
}
func TestVolumeKubus(t *testing.T) {
	assert.Equal(t, kubus.volume(), volumeValue, "Hasil volume salah!")
}
