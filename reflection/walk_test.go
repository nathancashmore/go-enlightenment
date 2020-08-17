package reflection

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	type Profile struct {
		Age  int
		City string
	}

	type Person struct {
		Name    string
		Profile Profile
	}

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one string",
			struct {
				Name string
			}{"Nathan"},
			[]string{"Nathan"},
		},
		{
			"Struct with multiple strings",
			struct {
				Name string
				City string
			}{"Nathan", "London"},
			[]string{"Nathan", "London"},
		},
		{
			"Struct with non string field",
			struct {
				Name string
				Age  int
			}{"Nathan", 33},
			[]string{"Nathan"},
		},
		{
			"Nested fields",
			Person{
				"Nathan",
				Profile{33, "London"},
			},
			[]string{"Nathan", "London"},
		},
		{
			"Slices",
			[]Profile{
				{33, "London"},
				{34, "Reykjavík"},
			},
			[]string{"London", "Reykjavík"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string

			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v expected %v", got, test.ExpectedCalls)
			}

		})
	}

}
