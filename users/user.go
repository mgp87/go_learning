package users

import (
	"fmt"
	"time"

	"github.com/mgp87/go_learning/models"
)

func NewUser() {
	u := new(models.User)
	u.AddUser(1, "Miguel", time.Now(), true)
	fmt.Println(u)
}
