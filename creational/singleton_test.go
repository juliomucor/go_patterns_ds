package creational

import "testing"

func TestGetInstance(t *testing.T) {
	counter1 := GetInstance()

	if counter1 == nil {
		// test of acceptance criteria 1 failed
		t.Error("expected pointer to Singleton after calling")
	}

	expectedCounter := counter1

	currentCount := counter1.AddOne()

	if currentCount != 1 {
		t.Errorf("after calling for the first time,the count must be 1 but it is %d\n", currentCount)
	}

	counter2 := GetInstance()

	if counter2 != expectedCounter {
		t.Error("Expected same instance in counter2 but it got a different")
	}

	currentCount = counter2.AddOne()
	if currentCount != 2 {
		t.Errorf("After calling a second time to AddOne, expect value is 2 but got is %d\n", currentCount)
	}
}
