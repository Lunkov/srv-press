package main

import (
  "time"
  "github.com/google/uuid"
)

type GroupModel struct {
  ID             uuid.UUID     `db:"id"                         json:"id"            yaml:"id"               gorm:"column:id;type:uuid;primary_key;default:uuid_generate_v4()"`
  CreatedAt      time.Time     `db:"created_at;default: now()"  json:"created_at"    sql:"default: now()"    gorm:"type:timestamp with time zone"`
  UpdatedAt      time.Time     `db:"updated_at;default: null"   json:"updated_at"    sql:"default: null"     gorm:"type:timestamp with time zone"`
  DeletedAt     *time.Time     `db:"deleted_at;default: null"   json:"deleted_at"    sql:"default: null"     gorm:"type:timestamp with time zone"`
  
  CODE           string        `db:"code"         json:"code"           yaml:"code"`

  Name           string        `db:"name"         json:"name"           yaml:"name"`
  Description    string        `db:"description"  json:"description"    yaml:"description"`
}
