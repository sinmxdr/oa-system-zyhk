package models

import (
	"fmt"
	"oa-backend/config"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func initDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.SystemConfig.Db.Username,
		config.SystemConfig.Db.Password,
		config.SystemConfig.Db.Host,
		config.SystemConfig.Db.Port,
		config.SystemConfig.Db.Name,
	)

	fmt.Println(dsn)

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			// 使用单数表名
			SingularTable: true,
		},
	})

	if err != nil {
		fmt.Println("Database connection failed,please check arguments:", err)
	}

	db.AutoMigrate(
		&Employee{},
		&Office{},
		&Role{},
		&Permission{},
		&Url{},
		&Region{},
		&Vendor{},

		&Supplier{},
		&ProductType{},
		&ProductAttribute{},
		&Product{},

		&CustomerCompany{},
		&Customer{},
		&Contract{},
		&Task{},
		&Payment{},
		&Invoice{},

		&Expense{},
		&Bidbond{},

		&Predesign{},
		&PredesignTask{},

		&Purchasing{},

		&HistoryEmployee{},
		&HistoryOffice{},
		&HistoryProduct{},

		&System{},
		&OfficeBak{},

		&ProductTrial{},
		&ProductCall{},
	)

	sqlDB, _ := db.DB()
	// 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)
	// 置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second)

}

// func changeDate() {
// 	var tasks []Task
// 	db.Find(&tasks)
// 	for i := range tasks {

// 		productAttribute := ProductAttribute{
// 			IsDelete:         false,
// 			StandardPrice:    tasks[i].StandardPrice,
// 			StandardPriceUSD: tasks[i].StandardPriceUSD,
// 		}

// 		db.Create(&productAttribute)
// 		db.Model(&Task{}).Where("id = ?", tasks[i].ID).Update("product_attribute_id", productAttribute.ID)

// 	}
// }
