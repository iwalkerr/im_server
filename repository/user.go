package repository

import (
	"database/sql"
	"imserver/lib/db"
	"imserver/models"
)

// 取消所有默认的地址
func CancelDefaultAddr(uid int) error {
	_, err := db.Conn().Exec(`update im_user_addr set is_default=0 where user_id=?`, uid)
	return err
}

// 修改用户地址
func EditUserAddr(param *models.EditShopAddr) error {
	_, err := db.Conn().Exec(`update im_user_addr set is_default=?,name=?,phone=?,addr=?,detail=? where shop_addr_id=? and user_id=?`,
		param.IsDefault, param.Name, param.Phone, param.Addr, param.Detail, param.ShopUserId, param.Uid)
	return err
}

// 获取收货地址列表
func ShopAddrList(uid int) (list []models.ShopAddr) {
	_ = db.Conn().Select(&list, `SELECT * FROM im_user_addr WHERE user_id=?`, uid)
	return
}

// 添加收货地址
func AddUserShopAddr(param *models.ShopAddr) error {
	_, err := db.Conn().Exec(`insert into im_user_addr(is_default,name,phone,addr,detail,user_id) values(?,?,?,?,?,?)`, param.IsDefault, param.Name, param.Phone, param.Addr, param.Detail, param.Uid)
	return err
}

func UpdateAlipayInfo(uid int, realname, alipayAccount string) error {
	_, err := db.Conn().Exec(`update im_users set realname=?,alipay_account=? where user_id=?`, realname, alipayAccount, uid)
	return err
}

func GetAlipayInfo(uid int) (string, string) {
	var realname, alipayAccount string
	_ = db.Conn().QueryRow(`select realname,alipay_account from im_users where user_id=?`, uid).Scan(&realname, &alipayAccount)
	return realname, alipayAccount
}

// 获取密码
func GetPwdById(uid int) (*LoginReq, error) {
	var resp LoginReq
	err := db.Conn().QueryRow(`select password,signature,phone from im_users where user_id=?`, uid).Scan(&resp.Password, &resp.Signature, &resp.Mobile)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return &resp, nil
}

// 更新头像
func UploadImage(userId int, imgPath string) error {
	_, err := db.Conn().Exec(`update im_users set head_picture=? where user_id=?`, imgPath, userId)
	return err
}

// 获取个人信息
func GetUserInfo(uid int) *models.UserInfo {
	var res models.UserInfo
	err := db.Conn().QueryRowx(`SELECT name,realname,user_addr,sex,email FROM im_users WHERE user_id=?`, uid).StructScan(&res)
	if err != nil && err != sql.ErrNoRows {
		return nil
	}

	return &res
}

// 更新用户信息
func UpdateUser(param *models.UserInfo) error {
	_, err := db.Conn().Exec(`update im_users set name=?,realname=?,user_addr=?,sex=?,email=? where user_id=?`, param.Name, param.RealName, param.UserAddr, param.Sex, param.Email, param.Uid)
	return err
}

// 更新个人二维码
func UpdateQrcode(uid int, qrcodePath string) error {
	_, err := db.Conn().Exec(`update im_users set qr_img=? where user_id=?`, qrcodePath, uid)
	return err
}
