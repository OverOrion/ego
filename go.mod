module github.com/OverOrion/ego

go 1.17

require (
	github.com/edgelesssys/ego v1.0.1
	github.com/stretchr/testify v1.8.1
	gopkg.in/square/go-jose.v2 v2.6.0
)

require (
	github.com/OverOrion/marblerun v0.6.2-0.20230104095344-b6a5c7bb2bb4 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/crypto v0.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/edgelesssys/ego v1.0.1 => github.com/OverOrion/ego v1.1.1-0.20221215133934-dba1b1b0c429

replace github.com/edgelesssys/marblerun => github.com/OverOrion/marblerun v0.6.2-0.20230104090900-f869552668aa
