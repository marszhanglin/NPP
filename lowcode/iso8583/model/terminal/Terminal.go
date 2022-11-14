package terminal

import (
	"NPP/dao"
	"time"
)

//
//// 删除错误数据
//func DeleteErrorData() {
//	dao.Db.Exec(" DELETE FROM t_zhu_msg_h WHERE zhu_id NOT IN ( SELECT id FROM t_zhu ) ")
//}
//
//// 删除5天数据
//func DeleteHistoryData() {
//	dao.Db.Exec(" DELETE FROM t_zhu_msg_h WHERE DATE_SUB(CURDATE(), INTERVAL 5 DAY) > date(cre_time)    ")
//}
//
//func SaveZhu(po *ZhuPO) error {
//	po.Cre_time = time.Now()
//	//tx := db.Begin()
//	//flow.TransflowId = getTimeUUID()
//	if err := dao.Db.Table("t_zhu").Create(&po).Error; err != nil {
//		logUtils.GlogInfo(err.Error())
//		//tx.Rollback()
//		return err
//	}
//	//tx.Commit()
//	return nil
//}
//
//func SaveZhuMsg(po *ZhuMsgPO) error {
//	po.Cre_time = time.Now()
//	//tx := db.Begin()
//	//flow.TransflowId = getTimeUUID()
//	if err := dao.Db.Table("t_zhu_msg").Create(&po).Error; err != nil {
//		logUtils.GlogInfo(err.Error())
//		//tx.Rollback()
//		return err
//	}
//	//tx.Commit()
//	return nil
//}
//
//func UpdateZhu(po *ZhuPO) error {
//	//po.Upd_time = time.Now()
//	tx := dao.Db.Begin()
//	if err := tx.Table("t_zhu").Where("zhu_sn=?", po.ZhuSn).Update(&po).Error; err != nil {
//		logUtils.GlogInfo(err.Error())
//		tx.Rollback()
//		return err
//	}
//	tx.Commit()
//	return nil
//}
//
//func UpdateZhuMsg(po *ZhuMsgPO) error {
//	//po.Upd_time = time.Now()
//	tx := dao.Db.Begin()
//	if err := tx.Table("t_zhu_msg").Where("zhu_id=?", po.ZhuId).Update(&po).Error; err != nil {
//		logUtils.GlogInfo(err.Error())
//		tx.Rollback()
//		return err
//	}
//	tx.Commit()
//	return nil
//}
//
//// 根据DeviceSn跟FlowNo查询订单是否重复
//func ZhuIsExit(po *ZhuPO) (isExit bool, err error, zhuId int64) {
//	// 查询redis看是否有这个zhu_sn  并返回zhuID
//	id, _ := noSqlUtils.GetNoSqlStr(noSqlUtils.RT_ZhuID + po.ZhuSn)
//	//logUtils.GlogInfo(id+"<-----------")
//	if len(id) > 0 {
//		//logUtils.GlogInfo("?<-----------")
//		idInt, _ := strconv.ParseInt(id, 10, 64)
//		return true, nil, int64(idInt)
//	}
//	//logUtils.GlogInfo("??<-----------")
//	//
//	zhuPos := []ZhuPO{}
//	queryErr := dao.Db.Table("t_zhu").Select(" id ").Where(" zhu_sn = ? ", po.ZhuSn).Find(&zhuPos).Error
//	if len(zhuPos) > 0 {
//		return true, queryErr, zhuPos[0].Id
//	} else {
//		return false, queryErr, 0
//	}
//}
//
//func GetAllZhuIdSnMap() (isExit bool, err error, zhuId []ZhuPO) {
//	zhuPos := []ZhuPO{}
//	queryErr := dao.Db.Table("t_zhu").Select(" id,zhu_sn").Where(" 1 = 1 ").Find(&zhuPos).Error
//	if len(zhuPos) > 0 {
//		return true, queryErr, zhuPos
//	} else {
//		return false, queryErr, nil
//	}
//}
//
//type ZhuSnVideo struct {
//	ZhuSn       string
//	VideoCamera string
//}
//
//func GetAllZhuSnVideoMap() (isExit bool, zhuSVs []ZhuSnVideo) {
//	zhuSnVideos := []ZhuSnVideo{}
//	rows, _ := dao.Db.Raw(" SELECT tz.zhu_sn  ,tzm.Video_camera  FROM t_zhu tz LEFT JOIN t_zhu_msg tzm ON tzm.zhu_id = tz.id ").Rows()
//	defer rows.Close()
//	for rows.Next() {
//		var zhuSnVideo ZhuSnVideo
//		var ZhuSn string
//		var VideoCamera string
//		rows.Scan(&ZhuSn, &VideoCamera)
//		zhuSnVideo.VideoCamera = VideoCamera
//		zhuSnVideo.ZhuSn = ZhuSn
//		zhuSnVideos = append(zhuSnVideos, zhuSnVideo)
//	}
//	if len(zhuSnVideos) > 0 {
//		return true, zhuSnVideos
//	} else {
//		return false, nil
//	}
//}
//
//func String2VO(str string) (voData ZhuVO) {
//	unmarshalErr := json.Unmarshal([]byte(str), &voData)
//	if nil != unmarshalErr {
//		logUtils.GlogInfo("Zhu:" + unmarshalErr.Error())
//	}
//	return voData
//}
//
//func String2PO(str string) (poData ZhuPO) {
//	unmarshalErr := json.Unmarshal([]byte(str), &poData)
//	if nil != unmarshalErr {
//		logUtils.GlogInfo("Zhu:" + unmarshalErr.Error())
//	}
//	return poData
//}
//
//func String2ZhuMsgPO(str string) (poData ZhuMsgPO) {
//	unmarshalErr := json.Unmarshal([]byte(str), &poData)
//	if nil != unmarshalErr {
//		logUtils.GlogInfo("ZhuMsg:" + unmarshalErr.Error())
//	}
//	return poData
//}
//
//func VO2POAndSaveOrUpdate(voData ZhuVO) (isSucc bool, errMsg string) {
//
//	for i := 0; i < int(voData.Count); i++ {
//		zhu := &ZhuPO{}
//		if len(nc.NcMap) < 1 {
//			return
//		}
//		ncId := nc.NcMap[voData.NcId]
//		// TODO 临时调整
//		if voData.Name == "A0" || voData.Name == "B0" {
//			ncId = 5
//		}
//		if ncId == 0 {
//			return
//		}
//		zhu.NcId = ncId
//		zhu.PName = voData.Name
//		zhu.Input = voData.Input
//		index := ""
//		if (int(voData.Index) + i) >= 10 {
//			index = "-" + strconv.Itoa((int(voData.Index) + i))
//		} else {
//			index = "-0" + strconv.Itoa((int(voData.Index) + i))
//		}
//
//		zhu.ZhuSn = voData.DrId + index
//		if voData.Output == 0 {
//			zhu.Output = 1 << 16
//		} else {
//			zhu.Output = voData.Output
//		}
//		zhu.Mods = voData.Mods
//
//		zhuMsg := &ZhuMsgPO{}
//		zhuMsg.NcId = ncId
//		zhuMsg.GrowingTemperature = strconv.FormatFloat(voData.Stemp, 'f', -1, 32)
//		zhuMsg.TopTemperature = strconv.FormatFloat(voData.Ttemp, 'f', -1, 32)
//		zhuMsg.SoilTemperature = strconv.FormatFloat(voData.Etemp, 'f', -1, 32)
//		zhuMsg.GrowingHumidity = strconv.FormatFloat(voData.Shumi, 'f', -1, 32)
//		zhuMsg.TopHumidity = strconv.FormatFloat(voData.Thumi, 'f', -1, 32)
//		zhuMsg.SoilHumidity = strconv.FormatFloat(voData.Ehumi, 'f', -1, 32)
//		zhuMsg.GrowingIlluminance = strconv.FormatFloat(voData.Sillum, 'f', -1, 32)
//		zhuMsg.TopIlluminance = strconv.FormatFloat(voData.Tillum, 'f', -1, 32)
//		zhuMsg.SoilEC = strconv.Itoa(int(voData.EC))
//		zhuMsg.CO2_concentratione = strconv.Itoa(int(voData.CO2))
//		zhuMsg.NH3_concentratione = strconv.Itoa(int(voData.NH3))
//		//zhuMsg.Video_camera = voData.Cam1
//
//		// 1.zhu是否存在
//		isZhuExits, _, dbZhuId := ZhuIsExit(zhu)
//		// 2.保存更新zhu
//		if isZhuExits {
//
//			zhu.Id = dbZhuId
//			go updateZhuCache(zhu, dbZhuId)
//			//go UpdateZhu(zhu)
//		} else {
//
//			//  第一次保存时名称预制一个
//			zhu.ZName = voData.Name + index //"-0" + strconv.Itoa(int(voData.Index)+i)
//			SaveZhu(zhu)
//			_, _, dbZhuId = ZhuIsExit(zhu)
//		}
//
//		// 3.编辑zhumsg的zhuid
//		zhuMsg.ZhuId = dbZhuId
//
//		// 4.保存更新zhumsg
//		if isZhuExits {
//			go updateZhuMsgCache(zhuMsg, zhu)
//			//go UpdateZhuMsg(zhuMsg)
//		} else {
//			SaveZhuMsg(zhuMsg)
//		}
//	}
//	return true, ""
//}
//
//// TODO 演示后,将数据保存至redis
//
//var CasheExpire int64 = 5 * 60
//
//func updateZhuMsgCache(zhuMsg *ZhuMsgPO, zhu *ZhuPO) {
//	zhuMsg.Upd_time = time.Now()
//	zhuMsgJsonBytes, _ := json.Marshal(zhuMsg)
//	zhuMsgStr := string(zhuMsgJsonBytes)
//	noSqlUtils.SetNoSqlStrExpire(noSqlUtils.RT_ZhuMSG+strconv.FormatInt(zhu.Id, 10), CasheExpire, zhuMsgStr)
//}
//
//// TODO 演示后,将数据保存至redis
//
//func updateZhuCache(zhu *ZhuPO, dbZhuId int64) {
//	zhu.Upd_time = time.Now()
//	zhuJsonBytes, _ := json.Marshal(zhu)
//	zhuStr := string(zhuJsonBytes)
//	noSqlUtils.SetNoSqlStrExpire(noSqlUtils.RT_Zhu+zhu.ZhuSn, CasheExpire, zhuStr)
//	noSqlUtils.SetNoSqlStrExpire(noSqlUtils.RT_ZhuID+zhu.ZhuSn, CasheExpire, strconv.FormatInt(dbZhuId, 10))
//}

func QueryTerminalBySn(terminalId string) (isExit bool, err error, terminal []TerminalPO) {
	terminals := []TerminalPO{}
	queryErr := dao.Db.Table("t_terminal").Select(" * ").Where(" terminal_id = ? ", terminalId).Find(&terminals).Error
	if len(terminals) > 0 {
		return true, queryErr, terminals
	} else {
		terminal = nil
		return false, queryErr, nil
	}
}

type TerminalPO struct {
	Id            int64     `gorm:"column:id;primary_key" `
	DeviceSn      string    `gorm:"column:device_sn" `
	TerminalId    string    `gorm:"column:terminal_id" `
	MerchantId    string    `gorm:"column:merchant_id" `
	SecretkeyType int64     `gorm:"column:secretkey_type" `
	MasterKey     string    `gorm:"column:master_key" `
	PinKey        string    `gorm:"column:pin_key" `
	MacKey        string    `gorm:"column:mac_key" `
	TrackKey      string    `gorm:"column:track_key" `
	Batch         int64     `gorm:"column:batch" `
	CreTime       time.Time `gorm:"column:cre_time" `
	UpdTime       time.Time `gorm:"column:upd_time" `
	Remark        string    `gorm:"column:remark" `
}
