package userController

import (
	"encoding/json"
	"errors"
	"gocrud/config"
	"gocrud/helper"
	"gocrud/model"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// func Index(w http.ResponseWriter, r *http.Request) {
// 	var users []model.User

// 	if err := config.DB.Where("is_deleted = 0").Find(&users).Error; err != nil{
// 		helper.Response(w, 500, err.Error(), nil)
// 	}

// 	helper.Response(w, 200, "List User", users)
// }

func Index(w http.ResponseWriter, r *http.Request) {
	var users []model.User
	var userResponse []model.UserResponse

	if err := config.DB.Joins("Alamat").Where("is_deleted = 0").Find(&users).Find(&userResponse).Error; err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return
	}
	helper.Response(w, 200, "List alamat", userResponse)
}

func Create(w http.ResponseWriter, r *http.Request) {
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil{
		helper.Response(w, 500, err.Error(), nil)
		return
	}
	defer r.Body.Close()

	if err := config.DB.Create(&user).Error; err != nil{
		helper.Response(w, 500, err.Error(), nil)
		return
	}
	helper.Response(w, 200, "Success create user", nil)
	return
}

func Detail(w http.ResponseWriter, r *http.Request) {
	idParams := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParams)

	var user model.User

	if err := config.DB.Where("is_deleted = 0").First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			helper.Response(w, 404, "User not found", nil)
			return
		}
		helper.Response(w, 500, err.Error(), nil)
		return
	}
	helper.Response(w, 200, "User", user)
	return
}

func Update(w http.ResponseWriter, r *http.Request) {
	idParams := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(idParams)

	var user model.User

	if err := config.DB.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound){
		helper.Response(w, 404, "User not found", nil)
		return
	}

	helper.Response(w, 500, err.Error(), nil)
	return
}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		helper.Response(w, 500, err.Error(), nil)
		return	
	}
	defer r.Body.Close()

	if err := config.DB.Where("id = ?", id).Updates(&user).Error; err != nil{
		helper.Response(w, 500, err.Error(), nil)
		return
	}
	helper.Response(w, 201, "Success Update User", nil)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)["id"]
	id, _ := strconv.Atoi(params)

	var user model.User

	res := config.DB.Delete(&user, id)

	if res.Error != nil {
		helper.Response(w, 500, res.Error.Error(), nil)
		return
	}

	if res.RowsAffected == 0{
		helper.Response(w, 404, "User not Found", nil)
		return
	}
	helper.Response(w, 200, "Success delete user", nil)
	return
}