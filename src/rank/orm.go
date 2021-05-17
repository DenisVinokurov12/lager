package rank

func All() []Rank {
	records := []Rank{}
	DB.Find(&records)
	return records
}