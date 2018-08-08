package memberlevel

const orderCondition = 1000

type User struct {
	ID     int
	Name   string
	Level  string
	Points int
}

type FreePoints struct {
	Level  string
	Points int
}

func CheckSpending(spending int) bool {
	if spending > orderCondition {
		return true
	}
	return false
}

func (point *FreePoints) GetFreePoints(userLevel string) int {
	if userLevel == "Gold" {
		return 200
	}
	return point.Points
}

func (user *User) UpdateLevel(userID int) string {
	if userID == user.ID && user.Level == "Gold" {
		user.Level = "Platinum"
	}
	return user.Level
}

func (user *User) UpdatePoints(userID int, userLevel string, points int) int {
	if userID == user.ID && userLevel == user.Level {
		user.Points += points
	}
	return user.Points
}

func (user User) CheckUserInRankByID(userID int) bool {
	if userID == user.ID {
		return true
	}
	return false
}

func UplevelProcess(userID int) User {
	user := GetUserByID(userID)
	transactions := GetLastSixMonthByUserID(userID)
	filteredTransactions := FilterTransactionBySpending(transactions)
	countTransactions := len(filteredTransactions)
	if checkCountOrder(countTransaction) &&
		checkUserInRankByID(userID) {
		currentLevel := user.UpdateLevel(userID)
		user.UpdatePoints(userID, currentLevel)
	}
	return user
}
