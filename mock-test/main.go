package main

import "fmt"
import "github.com/xroger88/hello-world/mock-test/mocks"

func main() {

	mockT := &mocks.Stringer{}

	mockResultFn := func() string {
		return "mockery test"
	}

	mockT.On("String").Return(mockResultFn, nil)

	result := mockT.String()
	/**
	if err != nil {
		panic(err)
	}
	**/

	fmt.Printf("result mocked: %v\n", result)

}
