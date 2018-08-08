package memberlevel

import "testing"

func Test_CheckSpending_Input_1200_Should_Be_True(t *testing.T) {
	expected := true
	spending := 1200

	actualResult := CheckSpending(spending)

	if expected != actualResult {
		t.Errorf("Expected %v but it got %v", expected, actualResult)
	}
}

func Test_CheckSpending_Input_800_Should_Be_False(t *testing.T) {
	expected := false
	spending := 800

	actualResult := CheckSpending(spending)

	if expected != actualResult {
		t.Errorf("Expected %v but it got %v", expected, actualResult)
	}
}

func Test_GetFreePoints_Input_Platinum_Should_Be_300(t *testing.T) {
	expected := 300
	userLevel := "Platinum"

	actualResult := GetFreePoints(userLevel)

	if expected != actualResult {
		t.Errorf("Expected %d but it got %d", expected, actualResult)
	}
}

func Test_UpdateLevel_Input_006_Should_Be_Platinum(t *testing.T) {
	expected := "Platinum"
	userID := 006

	user := User{ID: 006, Level: "Gold"}
	actualResult := user.UpdateLevel(userID)

	if expected != actualResult {
		t.Errorf("Expected %s but it got %s", expected, actualResult)
	}
}

func Test_UpdatePoints_Input_006_Platinum_Should_Be_800(t *testing.T) {
	expected := 800
	userID := 006
	userLevel := "Platinum"

	freePoint := 300
	user := User{ID: 006, Level: "Platinum", Points: 500}
	actualResult := user.UpdatePoints(userID, userLevel, freePoint)

	if expected != actualResult {
		t.Errorf("Expected %d but it got %d", expected, actualResult)
	}
}

func Test_UplevelProcess_Input_006_Should_Be_User_with_Platinum(t *testing.T) {
	expectedUser := User{ID: 006, Level: "Platinum", Points: 300}
	userID := 006

	service := Service{
		Users: []User{User{ID: 006, Level: "Gold"}},
		Transactions: []Transaction{
			Transaction{UserID: 006, Amount: 1000},
			Transaction{UserID: 006, Amount: 1000},
			Transaction{UserID: 006, Amount: 1000},
			Transaction{UserID: 006, Amount: 1000},
			Transaction{UserID: 006, Amount: 1000},
			Transaction{UserID: 006, Amount: 1000},
			Transaction{UserID: 006, Amount: 1000},
			Transaction{UserID: 006, Amount: 1000},
		},
	}
	actualResult := service.UplevelProcess(userID)

	if expectedUser != actualResult {
		t.Errorf("Expected %v but it go %v", expectedUser, actualResult)
	}
}

func Test_GetUserByID_Input_006_Should_Be_User_ID_006(t *testing.T) {
	expectedUser := User{ID: 006, Level: "Gold"}
	userID := 006

	service := Service{Users: []User{User{ID: 006, Level: "Gold"}}}
	actualResult := service.GetUserByID(userID)
	if expectedUser != actualResult {
		t.Errorf("Expected %v but it go %v", expectedUser, actualResult)
	}
}

func Test_GetLastSixMonthByUserID_Input_006_Should_Be_8_Transactions_of_UserID_006(t *testing.T) {
	expectedTransactionsCount := 8
	userID := 006

	service := Service{
		Users: []User{User{ID: 006, Level: "Gold"}},
		Transactions: []Transaction{
			Transaction{},
			Transaction{},
			Transaction{},
			Transaction{},
			Transaction{},
			Transaction{},
			Transaction{},
			Transaction{},
		},
	}
	actualResult := service.GetLastSixMonthByUserID(userID)
	if expectedTransactionsCount != len(actualResult) {
		t.Errorf("Expected %v but it go %v", expectedTransactionsCount, len(actualResult))
	}
}

func Test_FilterTransactionBySpending_Input_8_Transactions_Should_Be_8_filtered_transactions(t *testing.T) {
	expectedTransactionsCount := 8
	transactions := []Transaction{
		Transaction{UserID: 006, Amount: 1000},
		Transaction{UserID: 006, Amount: 1000},
		Transaction{UserID: 006, Amount: 1000},
		Transaction{UserID: 006, Amount: 1000},
		Transaction{UserID: 006, Amount: 1000},
		Transaction{UserID: 006, Amount: 1000},
		Transaction{UserID: 006, Amount: 1000},
		Transaction{UserID: 006, Amount: 1000},
	}

	filteredTransactions := FilterTransactionBySpending(transactions)

	if expectedTransactionsCount != len(filteredTransactions) {
		t.Errorf("Expected %v but it go %v", expectedTransactionsCount, len(filteredTransactions))
	}
}
