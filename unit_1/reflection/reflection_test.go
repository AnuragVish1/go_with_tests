package reflection

import (
	"reflect"
	"slices"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Gender string
	Age    int
}

func TestWalk(t *testing.T) {

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Testing number of expected calls meet the actual result but for one string only",
			struct {
				Name string
			}{"Anurag"},
			[]string{"Anurag"},
		},
		{
			"Testing for multiple input for string",
			struct {
				Name   string
				Gender string
			}{"Anurag", "Male"},
			[]string{"Anurag", "Male"},
		},
		{
			"Testing if feild is not string error is trhrown",
			struct {
				Name     string
				PhoneNum int
			}{"Anurag", 123456},
			[]string{"Anurag"},
		},
		{
			"Checking if nested struct work",
			Person{"Anurag", Profile{"Male", 21}},
			[]string{"Anurag", "Male"},
		},
		{
			"Checking if pointer struct work",
			&Person{"Anurag", Profile{"Male", 21}},
			[]string{"Anurag", "Male"},
		},
		{
			"slices boy",
			[]Person{
				{
					"Anurag",
					Profile{"Male", 20},
				},
				{
					"Anu",
					Profile{"Kid", 16},
				},
			},
			[]string{"Anurag", "Male", "Anu", "Kid"},
		},
		{
			"array boy",
			[2]Person{
				{
					"Anurag",
					Profile{"Male", 20},
				},
				{
					"Anu",
					Profile{"Kid", 16},
				},
			},
			[]string{"Anurag", "Male", "Anu", "Kid"},
		},
	}

	for _, test := range cases {

		t.Run(test.Name, func(t *testing.T) {
			got := []string{}
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("Got %v want %v", got, test.ExpectedCalls)
			}
		})

	}

	t.Run("maps boy", func(t *testing.T) {
		test := map[string]string{
			"Hello": "Hey",
			"HE":    "MAN",
		}

		expected := []string{"Hey", "MAN"}

		got := []string{}

		walk(test, func(input string) {
			got = append(got, input)
		})

		correct(got, expected, t)

	})

	t.Run("Handling channels", func(t *testing.T) {
		channels := make(chan Profile)

		go func() {
			channels <- Profile{Age: 20, Gender: "Male"}
			channels <- Profile{Age: 21, Gender: "Male"}
			close(channels)
		}()

		got := []string{}

		expected := []string{"Male", "Male"}

		walk(channels, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("Got %v want %v", got, expected)
		}

	})

	t.Run("Handling functions", func(t *testing.T) {
		testFunc := func() (Profile, Person) {
			return Profile{Age: 20, Gender: "Male"}, Person{Name: "Anurag", Profile: Profile{Age: 10, Gender: "Male"}}
		}

		got := []string{}
		expected := []string{"Male", "Anurag", "Male"}

		walk(testFunc, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("Got %v want %v", got, expected)
		}
	})

}

func correct(got []string, expected []string, t *testing.T) {
	if len(got) != len(expected) {
		t.Fatalf("Got %v want %v", got, expected)
	}

	for _, value := range got {
		contains := slices.Contains(expected, value)
		if !contains {
			t.Errorf("Got %v want %v", got, expected)
		}
	}
}
