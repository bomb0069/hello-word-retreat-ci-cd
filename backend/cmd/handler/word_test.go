package handler

import "testing"

func Test_GetWord_With_1_Should_Be_Hello_World_In_Thai(t *testing.T) {
	expected := `{text:"สวัสดีชาวโลก"}`

	actual := GetWord(1)

	if actual != expected {
		t.Errorf("actual %s is not equals to expected %s", actual, expected)
	}

}
