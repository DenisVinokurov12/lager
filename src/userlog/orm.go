package userlog

import(
	"time"
)

func Add(u *UserLog) *UserLog {

	if u.AddTs.IsZero(){
		u.AddTs = time.Now()
	}

	DB.Create(u)
	return u
}


func GetByFilter(parapms map[string]string) []*UserLog{

	res := []*UserLog{}

	if _, ok := parapms["object"]; ok{
		DB.Where("object = ?", parapms["object"])
	}
	if _, ok := parapms["object_id"]; ok{
		DB.Where("object_id = ?", parapms["object_id"])
	}
	if _, ok := parapms["user_id"]; ok{
		DB.Where("user_id = ?", parapms["user_id"])
	}
	if _, ok := parapms["action"]; ok{
		DB.Where("action = ?", parapms["action"])
	}
	if _, ok := parapms["add_ts_p"]; ok{
		DB.Where("add_ts > ?", parapms["add_ts_p"])
	}
	if _, ok := parapms["add_ts_m"]; ok{
		DB.Where("add_ts < ?", parapms["add_ts_m"])
	}
	if _, ok := parapms["add_ts"]; ok{
		DB.Where("add_ts = ?", parapms["add_ts"])
	}
	if _, ok := parapms["limit"]; ok{
		DB.Limit(parapms["limit"])
	}
	if _, ok := parapms["order"]; ok{
		DB.Order(parapms["order"])
	}
	if _, ok := parapms["sort_object"]; ok{
		if parapms["sort_object"] == "desc" {
			DB.Order("object desc")
		}else{
			DB.Order("object")
		}
	}
	if _, ok := parapms["sort_object_id"]; ok{
		if parapms["sort_object_id"] == "desc" {
			DB.Order("object_id desc")
		}else{
			DB.Order("object_id")
		}
	}
	if _, ok := parapms["sort_add_ts"]; ok{
		if parapms["sort_add_ts"] == "desc" {
			DB.Order("add_ts desc")
		}else{
			DB.Order("add_ts")
		}
	}

	DB.Find(&res)

	return res

}