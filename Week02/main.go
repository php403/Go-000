/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"errors"
	_ "errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID         	uint32 `gorm:"primary_key" json:"id"`
	OpenId 		string `json:"open_id"`
	IsDel      	uint8  `json:"is_del"`
}

type UserName struct {
	ID         	uint32 `gorm:"primary_key" json:"id"`
	Uid 		uint32 `json:"uid"`
	Name      	string  `json:"name"`
}

type UserInfo struct {
	ID 		uint32
	OpenId  string
	Name	string
}

func main() {
	userInfo,err :=GetService("123")
	if err != nil {
		fmt.Print(err)
	}
	fmt.Print(userInfo)
}


//service
func GetService(open_id string) (UserInfo, error){
	dsn := "root:pg719&1996@tcp(sh-cynosdbmysql-grp-emlvkwaa.sql.tencentcdb.com:24930)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	userInfo := UserInfo{}
	if err != nil{
		return userInfo,newErr
	}
	userData, err := User{OpenId:open_id}.Get(db)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound){
		newErr:=fmt.Errorf("userData为空: %v", err)
		return userInfo,newErr
	}
	userInfoData, err := UserName{ID:userData.ID}.Get(db)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound){
		newErr:=fmt.Errorf("userInfo为空: %v", err)
		return userInfo,newErr
	}
	userInfo = UserInfo{userData.ID,userData.OpenId,userInfoData.Name}
	return userInfo,nil
}

//model
func (u User) Get(db *gorm.DB) (User, error) {
	var user User
	db = db.Where("open_id = ? AND is_del = ?", u.OpenId, 0)
	err := db.First(&user).Error
	return user, err
}

//model
func (u UserName) Get(db *gorm.DB) (UserName, error) {
	var UserName UserName
	db = db.Where("uid = ? ", u.Uid)
	err := db.First(&UserName).Error
	return UserName, err
}
