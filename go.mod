module github.com/ldez/analysistestrun

go 1.20

require golang.org/x/tools v0.10.0

require (
	golang.org/x/mod v0.11.0 // indirect
	golang.org/x/sys v0.9.0 // indirect
)

// use the fork on golang.org/x/tools
replace golang.org/x/tools => github.com/ldez/tools v0.0.0-20230708095651-0aee1d89cb58
