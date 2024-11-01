package httpmocks

import (
	"encoding/json"
	"fmt"
	"io"
)

var (
	rickSanchez = person{
		Name:         "Rick Sanchez",
		Age:          70,
		CatchPhrases: []string{"IM PICKLE RICK!", "Wubba lubba dub dub!"},
	}

	mortySmith = person{
		Name:         "Morty Smith",
		Age:          14,
		CatchPhrases: []string{"What the hell, Rick?! What the hell?!"},
	}
)

type person struct {
	Name         string
	Age          int
	CatchPhrases []string
}

func (p person) String() string {
	return fmt.Sprintf("[name=%s, age=%d, Catchphrases=%s]", p.Name, p.Age, p.CatchPhrases)
}

// ExampleNewReadCloserMock how to create a ReadCloserMock that returns the expected bytes and error
func ExampleNewReadCloserMock() {
	b, err := json.Marshal(rickSanchez)
	if err != nil {
		panic(err)
	}

	readCloser := NewReadCloserMock(b, nil)

	bOut, err := io.ReadAll(readCloser)
	if err != nil {
		panic(err)
	}

	var output person
	err = json.Unmarshal(bOut, &output)
	if err != nil {
		panic(err)
	}

	fmt.Println(output)

	// Output:
	// [name=Rick Sanchez, age=70, Catchphrases=[IM PICKLE RICK! Wubba lubba dub dub!]]
}

// ExampleNewResponseMock how to create a mocked http.Response that contains a body and a 200 status code
func ExampleNewResponseMock() {
	b, err := json.Marshal(mortySmith)
	if err != nil {
		panic(err)
	}

	body := NewReadCloserMock(b, nil)
	resp := NewResponseMock(body, 200)

	defer body.Close()

	bout, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var actual person
	err = json.Unmarshal(bout, &actual)
	if err != nil {
		panic(err)
	}

	fmt.Println(actual)
	fmt.Println(resp.StatusCode)

	// Output:
	// [name=Morty Smith, age=14, Catchphrases=[What the hell, Rick?! What the hell?!]]
	// 200
}
