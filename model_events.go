package main

import (
  "time"
  "github.com/google/uuid"
  
  "github.com/Lunkov/lib-model/fields"
)

type EventsModel struct {
  ID             uuid.UUID     `db:"id"                         json:"id"            yaml:"id"               gorm:"column:id;type:uuid;primary_key;default:uuid_generate_v4()"`
  CreatedAt      time.Time     `db:"created_at;default: now()"  json:"created_at"    sql:"default: now()"    gorm:"type:timestamp with time zone"`
  PublicAt       time.Time     `db:"public_at;default: null"    json:"public_at"     sql:"default: null"     gorm:"type:timestamp with time zone"`
  DeletedAt     *time.Time     `db:"deleted_at;default: null"   json:"deleted_at"    sql:"default: null"     gorm:"type:timestamp with time zone"`
  
  Image          string        `db:"image"        json:"image"          yaml:"image"`

  Title          string        `db:"title"        json:"title"          yaml:"title"`
  Author         string        `db:"author"       json:"author"         yaml:"author"`
  Description    string        `db:"description"  json:"description"    yaml:"description"`
  
  DateStartAt    *time.Time     `db:"date_start;default: now()"    json:"date_start"     sql:"default: now()"     gorm:"type:timestamp with time zone"`
  DateFinishAt   *time.Time     `db:"date_finish;default: now()"   json:"date_finish"    sql:"default: now()"     gorm:"type:timestamp with time zone"`
  Address        string         `db:"address"      json:"address"        yaml:"address"`
  
  Images         fields.Files        `db:"images"       json:"images"         yaml:"images"       gorm:"type:jsonb;"`
  Documents      fields.Files        `db:"documents"    json:"documents"      yaml:"documents"    gorm:"type:jsonb;"` 
}
