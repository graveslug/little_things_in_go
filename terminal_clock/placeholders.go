package main

func placeholder() {
	type placeholder [5]string

	zero := placeholder{
		"▓▓▓",
		"▓ ▓",
		"▓ ▓",
		"▓ ▓",
		"▓▓▓",
	}

	one := placeholder{
		"▓▓ ",
		" ▓ ",
		" ▓ ",
		" ▓ ",
		"▓▓▓",
	}

	two := placeholder{
		"▓▓▓",
		"  ▓",
		"▓▓▓",
		"▓  ",
		"▓▓▓",
	}

	three := placeholder{
		"▓▓▓",
		"  ▓",
		"▓▓▓",
		"  ▓",
		"▓▓▓",
	}

	four := placeholder{
		"▓ ▓",
		"▓ ▓",
		"▓▓▓",
		"  ▓",
		"  ▓",
	}

	five := placeholder{
		"▓▓▓",
		"▓  ",
		"▓▓▓",
		"  ▓",
		"▓▓▓",
	}

	six := placeholder{
		"▓▓▓",
		"▓  ",
		"▓▓▓",
		"▓ ▓",
		"▓▓▓",
	}

	seven := placeholder{
		"▓▓▓",
		"  ▓",
		"  ▓",
		"  ▓",
		"  ▓",
	}

	eight := placeholder{
		"▓▓▓",
		"▓ ▓",
		"▓▓▓",
		"▓ ▓",
		"▓▓▓",
	}

	nine := placeholder{
		"▓▓▓",
		"▓ ▓",
		"▓▓▓",
		"  ▓",
		"▓▓▓",
	}
	colon := placeholder{
		"   ",
		" ▓ ",
		"   ",
		" ▓ ",
		"   ",
	}

	digits := placeholder{
		zero, one, two, three, four, five, six, seven, eight, nine,
	}
}
