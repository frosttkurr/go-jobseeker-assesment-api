package models

import (
	"errors"
	"go-jobseeker-assesment-api/initializers"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type T_Candidate struct {
	Candidate_ID uint   `json:"candidate_id"`
	Email        string `json:"email" validate:"required,email"`
	Phone_Number string `json:"phone_number"`
	Full_Name    string `json:"full_name" validate:"required"`
	DOB          string `json:"dob" validate:"required"`
	POB          string `json:"pob" validate:"required"`
	Gender       string `json:"gender" validate:"required"`
	Year_Exp     string `json:"year_exp" validate:"required"`
	Last_Salary  string `json:"last_salary" validate:"required"`
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

func dataValidator(data T_Candidate) (err error) {
	var db *gorm.DB
	var count_email, count_phone_number int64

	validate := validator.New()
	if err := validate.Struct(data); err != nil {
		return err.(validator.ValidationErrors)
	}

	db = initializers.DB.Raw("SELECT COUNT(candidate_id) AS id FROM t_candidate WHERE email = ?", data.Email).Scan(&count_email)
	if db.Error != nil {
		return db.Error
	}

	if count_email > 0 {
		return errors.New("ERROR: Email must be unique")
	}

	db = initializers.DB.Raw("SELECT COUNT(candidate_id) AS id FROM t_candidate WHERE phone_number = ?", data.Phone_Number).Scan(&count_phone_number)
	if db.Error != nil {
		return db.Error
	}

	if count_phone_number > 0 {
		return errors.New("ERROR: Phone number must be unique")
	}

	if data.Gender != "M" && data.Gender != "F" {
		return errors.New("ERROR: Gender must be male (M) or female (F)")
	}

	return nil
}

func InsertCandidate(data T_Candidate) (err error) {
	if err := dataValidator(data); err != nil {
		return err
	}

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
