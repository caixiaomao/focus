// ==========================================================================
// This is auto-generated by gf cli tool. You may not really want to edit it.
// ==========================================================================

package category

import (
	"database/sql"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/os/gtime"
)

// Entity is the golang structure for table gf_category.
type Entity struct {
    Id        uint        `orm:"id,primary" json:"id"`         // 分类ID，自增主键                                        
    Key       string      `orm:"key"        json:"key"`        // 栏目唯一键名，用于程序部分场景硬编码，一般不会用得到    
    ParentId  uint        `orm:"parent_id"  json:"parent_id"`  // 父级分类ID，用于层级管理                                
    UserId    uint        `orm:"user_id"    json:"user_id"`    // 创建的用户ID                                            
    Name      string      `orm:"name"       json:"name"`       // 分类名称                                                
    Sort      uint        `orm:"sort"       json:"sort"`       // 排序，数值越低越靠前，默认为添加时的时间戳，可用于置顶  
    Thumb     string      `orm:"thumb"      json:"thumb"`      // 封面图                                                  
    Brief     string      `orm:"brief"      json:"brief"`      // 简述                                                    
    Content   string      `orm:"content"    json:"content"`    // 详细介绍                                                
    CreatedAt *gtime.Time `orm:"created_at" json:"created_at"` // 创建时间                                                
    UpdatedAt *gtime.Time `orm:"updated_at" json:"updated_at"` // 修改时间                                                
}

// OmitEmpty sets OPTION_OMITEMPTY option for the model, which automatically filers
// the data and where attributes for empty values.
func (r *Entity) OmitEmpty() *arModel {
	return Model.Data(r).OmitEmpty()
}

// Inserts does "INSERT...INTO..." statement for inserting current object into table.
func (r *Entity) Insert() (result sql.Result, err error) {
	return Model.Data(r).Insert()
}

// InsertIgnore does "INSERT IGNORE INTO ..." statement for inserting current object into table.
func (r *Entity) InsertIgnore() (result sql.Result, err error) {
	return Model.Data(r).InsertIgnore()
}

// Replace does "REPLACE...INTO..." statement for inserting current object into table.
// If there's already another same record in the table (it checks using primary key or unique index),
// it deletes it and insert this one.
func (r *Entity) Replace() (result sql.Result, err error) {
	return Model.Data(r).Replace()
}

// Save does "INSERT...INTO..." statement for inserting/updating current object into table.
// It updates the record if there's already another same record in the table
// (it checks using primary key or unique index).
func (r *Entity) Save() (result sql.Result, err error) {
	return Model.Data(r).Save()
}

// Update does "UPDATE...WHERE..." statement for updating current object from table.
// It updates the record if there's already another same record in the table
// (it checks using primary key or unique index).
func (r *Entity) Update() (result sql.Result, err error) {
	return Model.Data(r).Where(gdb.GetWhereConditionOfStruct(r)).Update()
}

// Delete does "DELETE FROM...WHERE..." statement for deleting current object from table.
func (r *Entity) Delete() (result sql.Result, err error) {
	return Model.Where(gdb.GetWhereConditionOfStruct(r)).Delete()
}