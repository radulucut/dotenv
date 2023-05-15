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
			Value: " does not trim  spaces  ",
		},
		{
			Name: "TEST_MULTI_LINE",
			Value: `first line
second line
third line`,
		},
		{
			Name:  "TEST_SINGLE_QUOTE",
			Value: "test single quotes",
		},
		{
			Name:  "TEST_SINGLE_QUOTE_WITH_DOUBLE_QUOTE",
			Value: "\"\"",
		},
	}

	for i := 0; i < len(expectedVars); i++ {
		var actualValue = os.Getenv(expectedVars[i].Name)
		if actualValue != expectedVars[i].Value {
			t.Errorf("Expected var %s to be %v but found %v", expectedVars[i].Name, []byte(expectedVars[i].Value), []byte(actualValue))
		}
	}

}
