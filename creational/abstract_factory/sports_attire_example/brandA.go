package main

type brandA struct {
}

func (a *brandA) makeShoe() iShoe {
	return &brandAShoe{
		shoe: shoe{
			logo: "brandA",
			size: 14,
		},
	}
}

func (a *brandA) makeShort() iShort {
	return &brandAShort{
		short: short{
			logo: "brandA",
			size: 14,
		},
	}
}