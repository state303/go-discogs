// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameStyle = "style"

// Style mapped from table <style>
type Style struct {
	ID   int32  `gorm:"column:id;type:integer;primaryKey;autoIncrement:true;index:pk_style,priority:1" json:"id"`
	Name string `gorm:"column:name;type:character varying(50);not null;uniqueIndex:style_name_key,priority:1" json:"name"`
}

// TableName Style's table name
func (*Style) TableName() string {
	return TableNameStyle
}
