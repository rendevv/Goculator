package calculator

import "errors"

func Calculate(a, b int, operation string) (int, error) {
	var result int

	if operation == "tambah" {
		result = a + b
	} else if operation == "kurang" {
		result = a - b
	} else if operation == "kali" {
		result = a * b
	} else if operation == "bagi" {
		if b == 0 {
			return 0, errors.New("tidak boleh di bagi 0")
		}
		result = a / b
	} else if operation == "sisa bagi" {
		if b == 0 {
			return 0, errors.New("tidak boleh di bagi 0")
		}
		result = a % b
	} else {
		return 0, errors.New("operasi tidak dikenali")
	}

	return result, nil
}
