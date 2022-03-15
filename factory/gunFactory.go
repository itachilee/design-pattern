package factory

import "fmt"

func getGun(gunType string) (iGun, error) {

	switch {
	case gunType == "ak47":
		return newAk47(), nil
	case gunType == "maverick":
		return newMaverick(), nil
	}

	return nil, fmt.Errorf("Wrong gun type passed")
}
