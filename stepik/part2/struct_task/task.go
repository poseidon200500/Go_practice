package main

type Gun struct {
	On    bool
	Ammo  int
	Power bool
}

func (ex *Gun) Shoot() bool {
	var result bool = true
	if !ex.On || ex.Ammo == 0 {
		result = false
	} else {
		ex.Ammo--
		result = true
	}

	return result
}

func (ex *Gun) RideBike() bool {
	var result bool = true
	if !ex.Power || ex.Ammo == 0 {
		result = false
	} else {
		ex.Ammo--
		result = true
	}

	return result
}

func main() {

}
