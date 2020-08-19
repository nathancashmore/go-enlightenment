package reflection

import (
	"github.com/stretchr/testify/assert"
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
		{
			"Arrays",
			[2]Profile{
				{33, "London"},
				{34, "Reykjavík"},
			},
			[]string{"London", "Reykjavík"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string

			Walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v expected %v", got, test.ExpectedCalls)
			}

		})
	}

	t.Run("with maps", func(t *testing.T) {
		// This is here because you cannot rely on the order of maps !

		aMap := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz"}

		var got []string
		Walk(aMap, func(input string) {
			got = append(got, input)
		})

		assert.Contains(t, got, "Bar")
		assert.Contains(t, got, "Boz")
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{33, "Berlin"}
			aChannel <- Profile{34, "London"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Berlin", "London"}

		Walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("wanted %v, got %v", want, got)
		}
	})

	t.Run("with function", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{55, "Russia"}, Profile{22, "London"}
		}

		var got []string
		want := []string{"Russia", "London"}

		Walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("wanted %q, got %q", want, got)
		}
	})

}
