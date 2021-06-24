package main

type brandB struct {
}

func (b *brandB) makeShoe() iShoe {
	return &brandBShoe{
		shoe: shoe{
			logo: "brandB",
			size: 14,
		},
	}
}

func (b *brandB) makeShort() iShort {
	return &brandBShort{
		short: short{
			logo: "brandB",
			size: 14,
		},
	}
}