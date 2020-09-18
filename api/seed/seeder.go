package seed

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/yuanyu90221/golang_jwt_api_server/api/models"
)

var users = []models.User{
	models.User{
		Nickname: "Json liang",
		Email:    "json@rplab.ai",
		Passwd:   "1234test",
	},
	models.User{
		Nickname: "Steven wu",
		Email:    "wu@rplab.ai",
		Passwd:   "test",
	},
}

var posts = []models.Post{
	models.Post{
		Title:   "First Tilte",
		Content: "Fist content",
	},
	models.Post{
		Title:   "Second Tilte",
		Content: "Second content",
	},
}

//Load route
func Load(db *gorm.DB) {
	err := db.Debug().DropTableIfExists(&models.Post{}, &models.User{}).Error
	if err != nil {
		log.Fatalf("Cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}, &models.Post{}).Error
	if err != nil {
		log.Fatalf("Cannot migrate table: %v", err)
	}

	err = db.Debug().Model(&models.Post{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}
	for i := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		posts[i].AuthorID = users[i].ID

		err = db.Debug().Model(&models.Post{}).Create(&posts[i]).Error
		if err != nil {
			log.Fatalf("cannot seed posts table: %v", err)
		}
	}
}
