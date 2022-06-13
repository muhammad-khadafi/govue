package entity

/*
 * Created by muhammad.khadafi on 27/05/2022
 */

type User struct {
	Id       uint   `json:"id" db:"ID"`
	Username string `json:"username" db:"USERNAME"`
	Password string `json:"password" db:"PASSWORD"`
	Email    string `json:"email" db:"EMAIL"`
}
