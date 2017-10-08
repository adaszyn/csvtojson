package csvtojson
import "testing"
import "github.com/stretchr/testify/assert"

func TestCsvStringToArray (t *testing.T) {
	const testedString =
		`name,surname
		1,1
		1,2`
	got := CsvStringToArray(testedString, ",")
	shouldBe := [][]string {
		{ "name", "surname" },
		{ "1", "1" },
		{ "1", "2" },
	}
	assert.Equal(t, got, shouldBe)
}

func TestArrayToMap(t *testing.T) {
	var constantArray = [] [] string {
		{"a", "b"},
		{"1", "2"},
		{"11", "22"},
	}
	should := [] map [string] string {
		{"a": "1", "b": "2"},
		{"a": "11", "b": "22"},
	}
	got := ArrayToArrayOfMaps(constantArray)
	assert.Equal(t, should, got)
}