package models

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

type User struct {
	ID       uint64 `gorm:"primaryKey"`
	Password string
	Username string
	Email    string
}

type HashTags struct {
	Tags []string `bson:"tags"`
}

type Quest struct {
	ID      uint64   `bson:"id"`
	UserID  uint64   `bson:"user_id"`
	Title 	string	 `bson:"title"`
	Content string   `bson:"content"`
	HashTag HashTags `bson:"hash_tag"`
}

type Offer struct{
	AnswerQuest Quest `bson:"answer_quest"`
}