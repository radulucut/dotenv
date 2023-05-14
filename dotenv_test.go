package dotenv

import (
	"os"
	"testing"
)

type Var struct {
	Name  string
	Value string
}

func Test_Load_ExpectVarsLoadedSuccessfully(t *testing.T) {
	os.Clearenv()

	err := Load(".env.test")

	if err != nil {
		t.Errorf("Expected no error parsing but found %s", err)
	}

	var expectedVars = []Var{
		{
			Name:  "TEST_VAR",
			Value: "test-value",
		},
		{
			Name:  "test_var_quotes",
			Value: "test-value-quotes",
		},
		{
			Name:  "TEST_VAR_SPACES",
			Value: " three   spaces  ",
		},
		{
			Name:  "TEST_MULTI_LINE",
			Value: "first line\nsecond line\nthird line",
		},
		{
			Name:  "TEST_ESCAPE",
			Value: "\"\"",
		},
		{
			Name:  "TEST_ESCAPE_INSIDE_QUOTES",
			Value: "\"\\",
		},
	}

	for i := 0; i < len(expectedVars); i++ {
		var actualValue = os.Getenv(expectedVars[i].Name)
		if actualValue != expectedVars[i].Value {
			t.Errorf("Expected var %s to be %v but found %v", expectedVars[i].Name, []byte(expectedVars[i].Value), []byte(actualValue))
		}
	}

}
