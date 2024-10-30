package main

func discount(price int) string {
	switch {
	case price >= 100000 && price < 500000:
		return "10%"
	case price >= 500000 && price < 1000000:
		return "20%"
	case price >= 1000000:
		return "30%"
	}

	return "0%"
}
