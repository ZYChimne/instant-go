package model

import "time"

type Friend struct{
	FirstID int `json:"firstid"`;
	SecondID int `json:"secondid"`;
	CreateTime time.Time `json:"createtime"`
	UpdateTime time.Time `json:"updatetime"`
}