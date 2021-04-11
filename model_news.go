package main

import (
  "time"
  "github.com/google/uuid"
  
  "github.com/Lunkov/lib-model/fields"
)

type NewsModel struct {
  ID             uuid.UUID     `db:"id"                         json:"id"            yaml:"id"               gorm:"column:id;type:uuid;primary_key;default:uuid_generate_v4()"`
  CreatedAt      time.Time     `db:"created_at;default: now()"  json:"created_at"    sql:"default: now()"    gorm:"type:timestamp with time zone"`
  PublicAt       time.Time     `db:"public_at;default: null"    json:"public_at"     sql:"default: null"     gorm:"type:timestamp with time zone"`
  DeletedAt     *time.Time     `db:"deleted_at;default: null"   json:"deleted_at"    sql:"default: null"     gorm:"type:timestamp with time zone"`
  
  Image          string        `db:"image"        json:"image"          yaml:"image"`
  
  Language       string        `db:"language"     json:"language"       yaml:"language"`

  Title          string        `db:"title"        json:"title"          yaml:"title"`
  Author         string        `db:"author"       json:"author"         yaml:"author"`
  Article        string        `db:"article"      json:"article"        yaml:"article"`
  
  Keywords       string        `db:"keywords"     json:"keywords"       yaml:"keywords"`
  Description    string        `db:"description"  json:"description"    yaml:"description"`
  
  Images         fields.Files        `db:"images"       json:"images"         yaml:"images"       gorm:"type:jsonb;"`
  Documents      fields.Files        `db:"documents"    json:"documents"      yaml:"documents"    gorm:"type:jsonb;"` 
}
