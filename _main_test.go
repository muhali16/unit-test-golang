package main

import "testing"

var (
	kubus         Kubus   = Kubus{12}
	volumeValue   float64 = 1728
	luasValue     float64 = 864
	kelilingValue float64 = 144
)

func TestKelilingKubus(t *testing.T) {
	t.Log("Hasil dari keliling kubus:", kubus.keliling())
	if kubus.keliling() != kelilingValue {
		t.Errorf("Hasil keliling tidak sesuai yang seharusnya %.2f", kelilingValue)
	}
}

func TestLuasKubus(t *testing.T) {
	t.Log("Hasil dari luas kubus:", kubus.luas())
	if kubus.luas() != luasValue {
		t.Errorf("Hasil luas tidak sesuai yang seharusnya %.2f", luasValue)
	}
}

func TestVolumeKubus(t *testing.T) {
	t.Log("Hasil dari volume kubus:", kubus.volume())
	if kubus.volume() != volumeValue {
		t.Errorf("Hasil volume tidak sesuai yang seharusnya %.2f", volumeValue)
	}
}
