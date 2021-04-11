package main

import (
  "time"
  "github.com/google/uuid"
  
  "github.com/Lunkov/lib-model/fields"
)

type BookModel struct {
  ID             uuid.UUID     `db:"id"                         json:"id"            yaml:"id"               gorm:"column:id;type:uuid;primary_key;default:uuid_generate_v4()"`
  CreatedAt      time.Time     `db:"created_at;default: now()"  json:"created_at"    sql:"default: now()"    gorm:"type:timestamp with time zone"`
  PublicAt       time.Time     `db:"public_at;default: null"    json:"public_at"     sql:"default: null"     gorm:"type:timestamp with time zone"`
  DeletedAt     *time.Time     `db:"deleted_at;default: null"   json:"deleted_at"    sql:"default: null"     gorm:"type:timestamp with time zone"`
  
  Group          string        `db:"code_group"   json:"code_group"     yaml:"code_group"`
  Image          string        `db:"image"        json:"image"          yaml:"image"`

  Author         string        `db:"author"       json:"author"         yaml:"author"`
  Title          string        `db:"title"        json:"title"          yaml:"title"`
  Description    string        `db:"description"  json:"description"    yaml:"description"`
  
  Document       fields.File   `db:"document"    json:"document"       yaml:"document"` 
}
