package alamatController

import (
	"encoding/json"
	"errors"
	"gocrud/config"
	"gocrud/helper"

	// "gocrud/helper"
	"gocrud/model"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var alamats []model.Alamat
	var alamatResponse []model.AlamatResponse

	if err := config.DB.Joins("User").Find(&alamats).Find(&alamatResponse).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}
	helper.Response(w, 200, "List alamat", alamatResponse)
}

func Create(w http.ResponseWriter, r *http.Request) {
	var alamat model.Alamat
	if err := json.NewDecoder(r.Body).Decode(&alamat); err != nil{
		helper.Response(w, 500, err.Error(), nil)
		return
	}
	defer r.Body.Close()

	//cek user
	var user model.User
	if err := config.DB.First(&user, alamat.UserID).Error; err!=nil{
		if errors.Is(err, gorm.ErrRecordNotFound){
			helper.Response(w, 404, "User not Found", nil)
		return
		}

		helper.Response(w, 500, err.Error(), nil)
		return
	}

	if err := config.DB.Create(&alamat).Error; err != nil{
		helper.Response(w, 500, err.Error(), nil)
		return
	}
	helper.Response(w, 201, "Success create alamat", nil)
}

func Detail(w http.ResponseWriter, r *http.Request) {
	idParams := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParams)

	var alamat model.Alamat
	var alamatResponse model.AlamatResponse

	if err := config.DB.Joins("User").Where("is_deleted = 0").First(&alamat, id).First(&alamatResponse).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Response(w, 404, "Alamat not found", nil)
			return
		}
		helper.Response(w, 500, err.Error(), nil)
		return
	}
	helper.Response(w, 200, "Detail Alamat", alamatResponse)
}

func Update(w http.ResponseWriter, r *http.Request) {
	idParams := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParams)

	var alamat model.Alamat


	if err := config.DB.First(&alamat, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound){
		helper.Response(w, 404, "Alamat not found", nil)
		return
	}

	helper.Response(w, 500, err.Error(), nil)
	return
}

	var alamatPayload model.Alamat
	if err := json.NewDecoder(r.Body).Decode(&alamatPayload); err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return	
	}
	defer r.Body.Close()

	var user model.User
	if alamatPayload.UserID != 0 {
		if err:= config.DB.First(&user, alamatPayload.UserID).Error; err != nil{
			if errors.Is(err, gorm.ErrRecordNotFound){
				helper.Response(w, 404, "User not found", nil)
				return
			}
		
			helper.Response(w, 500, err.Error(), nil)
			return

		}
	}

	if err := config.DB.Where("id = ?", id).Updates(&alamatPayload).Error; err != nil{
		helper.Response(w, 500, err.Error(), nil)
		return
	}
	helper.Response(w, 201, "Success Update Alamat", nil)
}


func Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(params)

	var alamat model.Alamat

	res := config.DB.Delete(&alamat, id)

	if res.Error != nil {
		helper.Response(w, 500, res.Error.Error(), nil)
		return
	}

	if res.RowsAffected == 0{
		helper.Response(w, 404, "Alamat not Found", nil)
		return
	}
	helper.Response(w, 200, "Success delete alamat", nil)
	return
}
