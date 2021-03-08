package repository

import (
	"database/sql"
	"fmt"
	"imserver/lib/db"
	"imserver/models"
	"time"
)

type LoginReq struct {
	Signature   string `db:"signature"`
	Password    string `db:"password"`
	HeadPicture string `db:"head_picture"`
	Name        string `db:"name"`
	UserId      int    `db:"user_id"`
	Username    string `db:"username"`
	QrCode      string `db:"qr_code"`
	Mobile      string `db:"phone"`
}

// 获取用户积分
func GetUserIntegral(userId int) string {
	var balance int
	_ = db.Conn().QueryRow(`SELECT balance FROM im_user_record WHERE user_id=? ORDER BY create_time LIMIT 1`, userId).Scan(&balance)

	return fmt.Sprintf("%.2f", float64(balance)/100)
}

func UpdatePwd(mobile, pwd string) error {
	_, err := db.Conn().Exec(`update im_users set password=? where phone=?`, pwd, mobile)
	return err
}

func ExistByUUID(uuid string) bool {
	var id int
	_ = db.Conn().QueryRow(`select user_id from im_users where username=?`, uuid).Scan(&id)
	return id != 0
}

func CheckPhone(mobile string) bool {
	var id int
	_ = db.Conn().QueryRow(`select user_id from im_users where phone=?`, mobile).Scan(&id)
	return id != 0
}

// 获取登陆信息
func Login(mobile string) (*LoginReq, error) {
	var resp LoginReq
	err := db.Conn().QueryRowx(`SELECT signature,password,head_picture,name,user_id,username,qr_img qr_code FROM im_users WHERE phone=?`, mobile).StructScan(&resp)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return &resp, nil
}

// 注册信息
func Register(model *models.RegisterReq) (int, error) {
	res, err := db.Conn().Exec(`insert into im_users(username,name,head_picture,password,signature,phone,sex,created_time,updated_time) values(?,?,?,?,?,?,?,?,?)`,
		model.Username, model.Name, model.HeadPicture, model.Password, model.Salt, model.Mobile, model.Sex, time.Now(), time.Now())
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), err
}
