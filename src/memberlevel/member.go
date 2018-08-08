package memberlevel

const orderCondition = 1000
const priceCondition = 8

type Transaction struct {
	Amount int
	UserID int
}

type Service struct {
	Users        []User
	Transactions []Transaction
}

func (s Service) GetUserByID(userID int) User {
	for _, user := range s.Users {
		if user.ID == userID {
			return user
		}
	}
	return User{}
}
func (s Service) GetLastSixMonthByUserID(userID int) []Transaction {
	return s.Transactions
}

func CheckSpending(spending int) bool {
	if spending >= orderCondition {
		return true
	}
	return false
}

func FilterTransactionBySpending(transactions []Transaction) []Transaction {
	filtered := []Transaction{}
	for _, transaction := range transactions {
		if CheckSpending(transaction.Amount) {
			filtered = append(filtered, transaction)
		}
	}
	return filtered
}

func (s *Service) UplevelProcess(userID int) User {
	user := s.GetUserByID(userID)
	transactions := s.GetLastSixMonthByUserID(userID)
	filteredTransactions := FilterTransactionBySpending(transactions)
	countTransactions := len(filteredTransactions)
	if checkCountOrder(countTransactions) &&
		user.CheckUserInRankByID(userID) {
		currentLevel := user.UpdateLevel(userID)
		freePoint := GetFreePoints(currentLevel)
		user.UpdatePoints(userID, currentLevel, freePoint)
	}
	return user
}

func checkCountOrder(count int) bool {
	return count >= priceCondition
}
