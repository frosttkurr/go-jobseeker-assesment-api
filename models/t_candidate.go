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
	Last_Salary  string `json:"last_salary"`
}

func GetCandidates(page, page_size int, filters map[string]string) (results []T_Candidate, page_total_data int, total_data int64, err error) {
	offset := (page - 1) * page_size
	query := initializers.DB.Table("t_candidate").Offset(offset).Limit(page_size)

	for filter_name, filter_value := range filters {
		if filter_value != "" {
			query = query.Where(filter_name+" = ?", filter_value)
		}
	}

	db := query.Scan(&results)
	if db.Error != nil {
		err = db.Error
	}

	page_total_data = len(results)
	initializers.DB.Table("t_candidate").Count(&total_data)
	return results, page_total_data, total_data, err
}

func FindCandidate(candidate_id int) (result T_Candidate, err error) {
	db := initializers.DB.Raw("SELECT * FROM t_candidate WHERE candidate_id = ?", candidate_id).Scan(&result)
	if db.Error != nil {
		err = db.Error
	}
	return result, err
}

func inputValidator(data T_Candidate) (err error) {
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
		return errors.New("ERROR: Email already exist")
	}

	db = initializers.DB.Raw("SELECT COUNT(candidate_id) AS id FROM t_candidate WHERE phone_number = ?", data.Phone_Number).Scan(&count_phone_number)
	if db.Error != nil {
		return db.Error
	}

	if count_phone_number > 0 {
		return errors.New("ERROR: Phone number already exist")
	}

	if data.Gender != "M" && data.Gender != "F" {
		return errors.New("ERROR: Gender must be male (M) or female (F)")
	}

	return nil
}

func InsertCandidate(data T_Candidate) (status_code int, err error) {
	if err := inputValidator(data); err != nil {
		return 400, err
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
		return 500, err
	}
	return 200, nil
}

func UpdateCandidate(id int, data T_Candidate) (status_code int, err error) {
	result, _ := FindCandidate(id)
	if result.Candidate_ID == 0 {
		return 404, errors.New("ERROR: Candidate data not found")
	}

	if err := inputValidator(data); err != nil {
		return 400, err
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

	if err = initializers.DB.Model(&T_Candidate{}).Where("candidate_id", id).Updates(&candidate).Error; err != nil {
		return 500, err
	}
	return 200, nil
}

func DeleteCandidate(candidate_id int) (status_code int, err error) {
	result, _ := FindCandidate(candidate_id)
	if result.Candidate_ID == 0 {
		return 404, errors.New("ERROR: Candidate data not found")
	}

	if err = initializers.DB.Delete(&T_Candidate{}, candidate_id).Error; err != nil {
		return 500, err
	}
	return 200, nil
}
