package main

type Moto struct {
	On    bool
	Ammo  int
	Power int
}

func (m *Moto) Shoot() bool {
	if m.On == true {
		if m.Ammo > 0 {
			m.Ammo--
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func (m *Moto) RideBike() bool {
	if m.On == true {
		if m.Power > 0 {
			m.Power--
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func main() {
	a := Moto{true, 1, 1}
	testStruct := &a
	testStruct.Shoot()
	a.RideBike()
}
