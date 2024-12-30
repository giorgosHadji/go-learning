module helloworld

go 1.23.3

require github.com/gh/calculator v0.0.0

require (
	golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c // indirect
	rsc.io/quote v1.5.2 // indirect
	rsc.io/sampler v1.3.0 // indirect
)

replace github.com/gh/calculator => ../calculator
