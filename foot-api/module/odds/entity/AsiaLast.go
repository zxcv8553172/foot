package entity

import (
	"tesou.io/platform/foot-parent/foot-api/common/base/entity"
)

type AsiaLast struct {
	//博彩公司id
	CompId string `xorm:"unique(CompId_MatchId)"`
	//比赛id
	MatchId string `xorm:"unique(CompId_MatchId)"`

	/**
	初上下盘赔率
	*/
	Sp3 float64	`xorm:" comment('Sp3') index"`
	Sp0 float64	`xorm:" comment('Sp0') index"`
	//让球
	SLetBall string `xorm:" comment('s让球') index"`

	/**
	即时上下盘赔率
	*/
	Ep3 float64	`xorm:" comment('Ep3') index"`
	Ep0 float64	`xorm:" comment('Ep0') index"`
	//让球
	ELetBall string `xorm:" comment('e让球') index"`

	//数据时间
	OddDate string	`xorm:" comment('数据时间') index"`

	entity.Base `xorm:"extends"`
}