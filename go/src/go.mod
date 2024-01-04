module modA

go 1.21.3

require (
	go.uber.org/mock v0.3.0
	mock_modA v1.0.0
)

require github.com/golang/mock v1.6.0 // indirect

replace mock_modA v1.0.0 => ../mock-1/
