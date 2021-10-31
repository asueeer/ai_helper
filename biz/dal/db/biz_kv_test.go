package db

import (
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/spf13/cast"
)

type BizKV struct {
	ID      *int64 `gorm:"column:id"`
	Key     string `gorm:"column:key"`
	Val     string `gorm:"column:val"`
	Creator int64  `gorm:"column:creator"`
	Updater int64  `gorm:"column:updater"`
	Deleter int64  `gorm:"column:deleter"`
	gorm.Model
}

func (BizKV) TableName() string {
	return "biz_kv"
}

// 生成若干随机头像
func TestGenerateRandomHeadUrl(t *testing.T) {
	sql := DB.Model(&BizKV{}).Debug().LogMode(true)
	headUrls := []string{
		"https://images.gitee.com/uploads/images/2021/0824/010002_5e01d1ac_7809561.jpeg",
		"https://images.gitee.com/uploads/images/2021/0824/003738_19a51f87_7809561.jpeg",
		"https://images.gitee.com/uploads/images/2021/0824/003710_dc2563c6_7809561.jpeg",
		"https://images.gitee.com/uploads/images/2021/0824/003619_a9c360fe_7809561.jpeg",
		"https://images.gitee.com/uploads/images/2021/0824/003538_3335a24b_7809561.jpeg",
		"https://images.gitee.com/uploads/images/2021/0824/003502_c084b700_7809561.png",
		"https://images.gitee.com/uploads/images/2021/0824/003255_64352af3_7809561.jpeg",
		"https://images.gitee.com/uploads/images/2021/0824/002616_e99f4166_7809561.jpeg",
		"https://images.gitee.com/uploads/images/2021/0731/134131_a864d20c_7809561.jpeg",
	}
	for i := 0; i < len(headUrls); i++ {
		kv := BizKV{
			Key:     "default_head_url_",
			Val:     headUrls[i],
			Creator: 305088049,
			Updater: 305088049,
		}
		kv.Key = kv.Key + cast.ToString(i)
		err := sql.Omit("id", "deleter").Create(&kv).Error
		t.Log(err)
	}
}

// 生成若干随机昵称
func TestGenerateLuckyNames(t *testing.T) {
	sql := DB.Model(&BizKV{}).Debug().LogMode(true)
	luckNames := []string{
		"Estella",
		"Beatrice",
		"Zora",
		"Eleanor",
		"Laila",
		"Lisbeth",
		"Hazel",
		"Serendipity",
		"Natsukashii",
		"Epoch",
		"Hiraeth",
		"Ineffable",
		"Petrichor",
		"Mellifluous",
		"Limerence",
		"Ethereal",
		"Iridescent",
		"Laurel",
		"Nostalgia",
		"Palpitate",
		"Taxol",
		"Ethanol",
		"Ephemeral",
		"Shimmer",
		"Mamihlapinatapai",
	}
	for i := 0; i < len(luckNames); i++ {
		kv := BizKV{
			Key:     "lucky_name_",
			Val:     luckNames[i],
			Creator: 305088049,
			Updater: 305088049,
		}
		kv.Key = kv.Key + cast.ToString(i)
		err := sql.Omit("id", "deleter").Create(&kv).Error
		t.Log(err)
	}
}

// 生成若干随机幸运数字
func TestGenerateLuckyNumbers(t *testing.T) {
	sql := DB.Model(&BizKV{}).Debug().LogMode(true)
	luckNumbers := []string{
		"B612",
		"246",
		"001",
		"007",
		"0826",
		"0827",
		"1014",
		"0727",
	}
	for i := 0; i < len(luckNumbers); i++ {
		kv := BizKV{
			Key:     "lucky_number_",
			Val:     luckNumbers[i],
			Creator: 305088049,
			Updater: 305088049,
		}
		kv.Key = kv.Key + cast.ToString(i)
		err := sql.Omit("id", "deleter").Create(&kv).Error
		t.Log(err)
	}
}

// 生成若干随机的分隔符
func TestGenerateRandomNumbers(t *testing.T) {
	sql := DB.Model(&BizKV{}).Debug().LogMode(true)
	delimiters := []string{
		"-",
		"_",
		"~",
	}
	for i := 0; i < len(delimiters); i++ {
		kv := BizKV{
			Key:     "random_delimiter_",
			Val:     delimiters[i],
			Creator: 305088049,
			Updater: 305088049,
		}
		kv.Key = kv.Key + cast.ToString(i)
		err := sql.Omit("id", "deleter").Create(&kv).Error
		t.Log(err)
	}
}

// 存放小程序AppID和小程序App密钥
func TestSaveMiniProgramConfig(t *testing.T) {
	InitDB()
	sql := DB.Model(&BizKV{}).Debug().LogMode(true)
	kv := BizKV{
		Key:     "mini_program_conf",
		Val:     "wx3103aa7e787ceb5e,95e8eeba52f096882a941507f04cd512",
		Creator: 305088049,
		Updater: 305088049,
	}
	err := sql.Omit("id").Create(&kv).Error
	t.Log(err)
}

// 存放阿里云AccessID和AccessSecret
func TestSaveOSSConfig(t *testing.T) {
	InitDB()
	sql := DB.Model(&BizKV{}).Debug().LogMode(true)
	kv := BizKV{
		Key:     "oss_conf",
		Val:     "LTAI5tHByRNJYanYvoQmKM4W,UluaRdsyJgKconcFD4x04Zm9wcgGrT",
		Creator: 305088049,
		Updater: 305088049,
	}
	err := sql.Omit("id").Create(&kv).Error
	t.Log(err)
}
