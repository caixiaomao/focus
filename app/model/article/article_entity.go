// ==========================================================================
// This is auto-generated by gf cli tool. You may not really want to edit it.
// ==========================================================================

package article

import (
	"database/sql"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/os/gtime"
)

// Entity is the golang structure for table gf_article.
type Entity struct {
    Id         uint        `orm:"id,primary"  json:"id"`          // 自增ID                                                  
    Key        string      `orm:"key"         json:"key"`         // 文章唯一键名，用于程序硬编码使用，一般用不到            
    CategoryId uint        `orm:"category_id" json:"category_id"` // 栏目ID                                                  
    UserId     uint        `orm:"user_id"     json:"user_id"`     // 创建文章的用户ID                                        
    Title      string      `orm:"title"       json:"title"`       // 文章标题                                                
    Content    string      `orm:"content"     json:"content"`     // 文章内容                                                
    Sort       uint        `orm:"sort"        json:"sort"`        // 排序，数值越低越靠前，默认为添加时的时间戳，可用于置顶  
    Brief      string      `orm:"brief"       json:"brief"`       // 摘要                                                    
    Thumb      string      `orm:"thumb"       json:"thumb"`       // 缩略图                                                  
    Tags       string      `orm:"tags"        json:"tags"`        // 标签名称列表，以JSON存储                                
    Referer    string      `orm:"referer"     json:"referer"`     // 内容来源                                                
    Status     uint        `orm:"status"      json:"status"`      // 状态 0: 正常, 1: 禁用                                   
    ZanCount   uint        `orm:"zan_count"   json:"zan_count"`   // 赞                                                      
    CaiCount   uint        `orm:"cai_count"   json:"cai_count"`   // 踩                                                      
    CreatedAt  *gtime.Time `orm:"created_at"  json:"created_at"`  // 创建时间                                                
    UpdatedAt  *gtime.Time `orm:"updated_at"  json:"updated_at"`  // 修改时间                                                
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