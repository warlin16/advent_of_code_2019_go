package main

// GetAbsModuleMass will return absolute module mass
func GetAbsModuleMass(module int) int {
	moduleMass := module/3 - 2
	moduleMassTotalSum := moduleMass
	for moduleMass > 0 {
		moduleMass = moduleMass/3 - 2
		if moduleMass > 0 {
			moduleMassTotalSum += moduleMass
		}
	}
	return moduleMassTotalSum
}
