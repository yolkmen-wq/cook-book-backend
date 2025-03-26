package models

type Menu struct {
	// 菜单类型（0代表菜单、1代表iframe、2代表外链、3代表按钮）
	ID              int64  `json:"id" gorm:"column:id;primary_key"`
	MenuType        *int   `json:"menuType,omitempty" gorm:"column:menu_type"`
	ParentId        int64  `json:"parentId,omitempty" gorm:"column:parent_id"`
	Title           string `json:"title,omitempty" gorm:"column:title"`
	Name            string `json:"name,omitempty" gorm:"column:name"`
	Path            string `json:"path,omitempty" gorm:"column:path"`
	Component       string `json:"component,omitempty" gorm:"column:component"`
	Rank            int    `json:"rank,omitempty" gorm:"column:rank"`
	Redirect        string `json:"redirect,omitempty" gorm:"column:redirect"`
	Icon            string `json:"icon,omitempty" gorm:"column:icon"`
	ExtraIcon       string `json:"extraIcon,omitempty" gorm:"column:extra_icon"`
	EnterTransition string `json:"enterTransition,omitempty" gorm:"column:enter_transition"`
	LeaveTransition string `json:"leaveTransition,omitempty" gorm:"column:leave_transition"`
	ActivePath      string `json:"activePath,omitempty" gorm:"column:active_path"`
	Auths           string `json:"auths,omitempty" gorm:"column:auths"`
	FrameSrc        string `json:"frameSrc,omitempty" gorm:"column:frame_src"`
	FrameLoading    *bool  `json:"frameLoading,omitempty" gorm:"column:frame_loading"`
	KeepAlive       *bool  `json:"keepAlive,omitempty" gorm:"column:keep_alive;default:0"`
	HiddenTag       *bool  `json:"hiddenTag,omitempty" gorm:"column:hidden_tag;default:0"`
	FixedTag        *bool  `json:"fixedTag,omitempty" gorm:"column:fixed_tag;default:0"`
	ShowLink        *bool  `json:"showLink,omitempty" gorm:"column:show_link;default:0"`
	ShowParent      *bool  `json:"showParent,omitempty" gorm:"column:show_parent;default:0"`
}

type GetMenuRequest struct {
	Name     string `form:"name"`
	PageNum  int    `form:"page_num"`
	PageSize int    `form:"page_size"`
}
