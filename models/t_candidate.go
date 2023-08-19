package models

import "go-jobseeker-assesment-api/initializers"

type T_Candidate struct {
	Candidate_ID uint   `gorm:"primary_key" json:"candidate_id"`
	Email        string `gorm:"type:varchar(255);unique_index;not null" json:"email"`
	Phone_Number string `gorm:"type:varchar(20);unique_index" json:"phone_number"`
	Full_Name    string `gorm:"type:varchar(255);not null" json:"full_name"`
	DOB          string `gorm:"type:varchar(10)" json:"dob"`
	POB          string `gorm:"type:varchar(255)" json:"pob"`
	Gender       string `gorm:"type:char(1);check:gender IN ('F', 'M')" json:"gender"`
	Year_Exp     string `gorm:"type:varchar(10);not null" json:"year_exp"`
	Last_Salary  string `gorm:"type:varchar(255)" json:"last_salary"`
}

func GetCandidates() (results []T_Candidate, err error) {
	db := initializers.DB.Raw("SELECT * FROM t_candidate").Scan(&results)
	if db.Error != nil {
		err = db.Error
	}
	return results, err
}

func FindCandidate(candidate_id string) (result T_Candidate, err error) {
	db := initializers.DB.Raw("SELECT * FROM t_candidate WHERE candidate_id = ?", candidate_id).Scan(&result)
	if db.Error != nil {
		err = db.Error
	}
	return result, err
}

func InsertCandidate(data T_Candidate) (err error) {
	candidate := T_Candidate{
		Email:        data.Email,
		Phone_Number: data.Phone_Number,
		Full_Name:    data.Full_Name,
		DOB:          data.DOB,
		POB:          data.POB,
		Gender:       data.Gender,
		Year_Exp:     data.Year_Exp,
		Last_Salary:  data.Last_Salary,
	}

	if err = initializers.DB.Create(&candidate).Error; err != nil {
		return err
	}
	return nil
}
