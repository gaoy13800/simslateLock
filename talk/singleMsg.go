package talk




type SingleSe struct {

	Cstu string

	Stus string

	Guid string
}



func (this *SingleSe) SetCstu(cstu string){
	this.Cstu = cstu
}

func (this *SingleSe) GetCstu() string{
	return this.Cstu
}


func (this *SingleSe) GetStus() string{
	return this.Stus
}


func NewSingleRe() *SingleSe{
	return &SingleSe{
		Cstu:"3",
		Stus:"80",
	}
}