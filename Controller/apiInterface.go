package Controller

type entity interface{
	Select()
	SelectOne()
	Update()
	Delete()
	Insert()
}



