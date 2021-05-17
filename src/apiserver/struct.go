package apiserver


type (
	ApiConfig struct {
		Port        int    `json:"port"`
		Path        string `json:"ws_path"`
		SessionKey  string `json:"sesion_key"`
		SessionLive int    `json:"sesion_live"`
		RegistrationHost     string         `json:"registration_host"`
		FilesFolder     string         `json:"files_folder"`
	}
	UserCourse struct{
		UserId 			int    		`json:"day_id"`
		IssueDayId 		int    		`json:"issue_day_id"`
	}
	LoginPost struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	CourseDayPost struct{
		DayId int `json:"day_id"`
		CourseId int `json:"course_id"`
		Title          	string `json:"title"`
		Description    	string `json:"description"`
	}
	LoginGet struct {
		Token    string `json:"token"`
		Username string `json:"username"`
		Email    string `json:"email"`
		Access   string `json:"access"`
	}
	Logout struct {
		Token string `json:""token`
	}
	UserSession struct{
		Id int
		Rules map[string]map[int][]string
	}
	Order struct {
		OrderId string `json:"order_id"`
		Quantity int `json:"quantity"`
	}
	Error struct {
		Error string `json:"error"`
	}
	CourseEdit struct{
		Title          	string `json:"title"`
		Description    	string `json:"description"`
		Image    		string `json:"image"`
	}
	RegistrationPost struct {
		FirstName string `json:"first_name"`
		LastName string `json:"last_name"`
		Phone string `json:"phone"`
		Login string `json:"login"`
		Email string `json:"email"`
		Password string `json:"password"`
		CourseId  int `json:"course_id"`
		OrderId  string `json:"order_id"`
		Rule  int `json:"rule"`
		Gender  string `json:"gender"`
	}
)
