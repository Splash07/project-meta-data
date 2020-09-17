package model

// ClientType struct
type ClientType struct {
	ID            int    `xorm:"ClientTypeID"`
	TemplateID    int    `xorm:"ClientTypeTemplateID"`
	Name          string `xorm:"TypeName"`
	Description   string `xorm:"Description"`
	IsInternalUse bool   `xorm:"IsInternalUse"`
	Weight        int    `xorm:"Weight"`
	DefaultUrl    string `xorm:"DefaultUrl"`
	ParentID      int    `xorm:"ParentID"`
}

type ClientTypeChangeRequest struct {
	TemplateID    int    `xorm:"ClientTypeTemplateID" json:"template_id"`
	Name          string `xorm:"TypeName" json:"name"`
	Description   string `xorm:"Description" json:"description"`
	IsInternalUse bool   `xorm:"IsInternalUse" json:"is_internal_use"`
	Weight        int    `xorm:"Weight" json:"weight"`
	DefaultUrl    string `xorm:"DefaultUrl" json:"default_url"`
	ParentID      int    `xorm:"ParentID" json:"parent_id"`
}

// IsExists struct
func (s ClientType) IsExists() (ok bool) {
	if s.ID != 0 {
		ok = true
	}
	return
}
