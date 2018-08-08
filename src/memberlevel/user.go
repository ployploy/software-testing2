package memberlevel

type User struct {
	ID     int
	Name   string
	Level  string
	Points int
}

func GetFreePoints(userLevel string) int {
	if userLevel == "Gold" {
		return 200
	}
	if userLevel == "Platinum" {
		return 300
	}
	return 0
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
