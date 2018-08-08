package memberlevel

import "testing"

func Test_CheckSpending_Input_1200_Should_Be_True(t *testing.T) {
	expected := true
	spending := 1200

	actualResult := CheckSpending(spending)

	if expected != actualResult {
		t.Errorf("Expected %v but got it %v", expected, actualResult)
	}
}

func Test_CheckSpending_Input_800_Should_Be_False(t *testing.T) {
	expected := false
	spending := 800

	actualResult := CheckSpending(spending)

	if expected != actualResult {
		t.Errorf("Expected %v but got it %v", expected, actualResult)
	}
}

func Test_GetFreePoints_Input_Platinum_Should_Be_300(t *testing.T) {
	expected := 300
	userLevel := "Platinum"

	point := FreePoints{Level: "Platinum", Points: 300}
	actualResult := point.GetFreePoints(userLevel)

	if expected != actualResult {
		t.Errorf("Expected %d but got it %d", expected, actualResult)
	}
}

func Test_UpdateLevel_Input_006_Should_Be_Platinum(t *testing.T) {
	expected := "Platinum"
	userID := 006

	user := User{ID: 006, Level: "Gold"}
	actualResult := user.UpdateLevel(userID)

	if expected != actualResult {
		t.Errorf("Expected %s but got it %s", expected, actualResult)
	}
}

func Test_UpdatePoints_Input_006_Platinum_Should_Be_800(t *testing.T) {
	expected := 800
	userID := 006
	userLevel := "Platinum"

	freePoint := FreePoints{Level: "Platinum", Points: 300}
	user := User{ID: 006, Level: "Platinum", Points: 500}
	actualResult := user.UpdatePoints(userID, userLevel, freePoint.Points)

	if expected != actualResult {
		t.Errorf("Expected %d but got it %d", expected, actualResult)
	}
}
