package user

import (
	"github.com/jinzhu/gorm"
	"time"
	"course_day"
	"course"
	"crypto/sha1"
	"encoding/hex"
	"issue_day"
	"user_course"
	"user_comment"
)

var DB *gorm.DB

var UserLogCallbackAdd = func(string , int, int){}
var UserLogCallbackDelete = func(string , int, int){}
var UserLogCallbackEdit = func(string , int, int){}

const(
	ADMIN = 2
	CURATOR = 1
	STUDENT = 0
)

type User struct {
	Id             	int    `gorm:"primary_key" json:"id"`
	Login          	string `json:"login"`
	Password       	string `json:"-"`
	FirstName       string `json:"first_name"`
	LastName       	string `json:"last_name"`
	Phone       	string `json:"phone"`
	Rule       		int `json:"rule"`
	AddTs       	time.Time `json:"add_ts"`
	RankId       	int `json:"rank_id"`
	ApiToken       string `json:"-"`
	Email       	string `json:"email"`
	Avatar       	string `json:"avatar"`
	OrderId       	string `json:"order_id"`
	Gender       	string `json:"gender"`
}


func CompressPass(pass string) string {
	h := sha1.New()
	h.Write([]byte(pass))
	return hex.EncodeToString(h.Sum(nil))
}


func (u User) IsAccessCourse(course_ *course.Course) bool {
	return true
}

func (u User) IsAccessUpdateRankUser(user_ *User) bool {
	if u.Rule == ADMIN || u.Rule == CURATOR {
		return true
	}
	return true
}


func (u User) IsAccessViewCourseUser(user_ *User) bool {
	if u.Rule == ADMIN || u.Rule == CURATOR {
		return true
	}
	if user_.Id == u.Id {
		return true
	}
	return true
}

func (u User) IsAccessAddCourse() bool {
	if u.Rule == ADMIN{
		return true
	}
	return false
}

func (u User) IsAccessNagarda() bool {
	if u.Rule == ADMIN || u.Rule == CURATOR{
		return true
	}
	return false
}

func (u User) IsAccessUpdateCourse(course_ *course.Course) bool {
	if u.Rule == ADMIN{
		return true
	}
	return false
}

func (u User) IsAccessDeleteCourse(course_ *course.Course) bool {
	if u.Rule == ADMIN{
		return true
	}
	return false
}


func (u User) IsAccessCourseDay(course_day_ *course_day.CourseDay) bool {
	if u.Rule == ADMIN{
		return true
	}
	// if course_day_.UserId == u.Id {
	// 	return true
	// }
	return true
}

func (u User) IsAccessAddIssueDay() bool {
	if u.Rule == ADMIN{
		return true
	}
	return false
}

func (u User) IsAccessUpdateIssueDay(issue_day_ *issue_day.IssueDay) bool {
	if u.Rule == ADMIN{
		return true
	}
	return false
}
func (u User) IsAccessDeleteIssueDay(issue_day_ *issue_day.IssueDay) bool {
	if u.Rule == ADMIN{
		return true
	}
	return false
}

func (u User) IsAccessAddUserCourse() bool {
	if u.Rule == ADMIN{
		return true
	}
	return false
}

func (u User) IsAccessUpdateUserCourse(user_course_ *user_course.UserCourse) bool {
	if u.Rule == ADMIN || u.Rule == CURATOR {
		return true
	}
	return false
}

func (u User) IsAccessUpdateUser(user_ *User) bool {
	if u.Rule == ADMIN || u.Rule == CURATOR {
		return true
	}
	if u.Id == user_.Id {
		return true
	}
	return false
}

func (u User) IsAccessViewUser(user_ *User) bool {
	if u.Rule == ADMIN || u.Rule == CURATOR {
		return true
	}
	return false
}

func (u User) IsAccessDeleteUserComment(user_comment_ *user_comment.UserComment) bool {
	if user_comment_.UserId == u.Id {
		return true
	}
	return false
}

func (u User) IsAccessViewUserCourse(user_course_ *user_course.UserCourse) bool {
	if u.Rule == ADMIN || u.Rule == CURATOR {
		return true
	}
	if user_course_.UserId == u.Id {
		return true
	}
	return false
}

func (u User) IsAccessViewUserCommentByUserId(user_ *User) bool {
	if u.Rule == ADMIN || u.Rule == CURATOR {
		return true
	}
	if u.Id == user_.Id {
		return true
	}
	return false
}