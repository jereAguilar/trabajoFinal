package users

import (
	"net/http"
	"time"
	"https://github.com/gin-gonic/gin"
)

type User struct { //agrego formato de tipo json para que se pueda leer en REST
	ID   int64  `json:"id"`
	Nickname string `json:"nickname"`
}

type UserService interface { 
	Add(User) User
	GetUser(*gin.Context)
	UpdateUser(*gin.Context)
}
type uServ struct { // --> implementacion del servicio
	data map[int64]User
}
func (us *uServ) GetUser(c *gin.Context) { 
	c.JSON(http.StatusCreated, us.data)
}

func (s *uServ) Add(u User) User{
	u.ID = time.Now().Unix()
	s.data[u.ID] = user
	return u
}
func (us *uServ) UpdateUser(c *gin.Context) {
	var user User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, exists := us.data[user.ID]
	if exists {
		us.data[u.ID] = user
		c.JSON(http.StatusOK, gin.H{"user": user})
		return
	}

	c.JSON(http.StatusNotFound, gin.H{"status": "Not Found"})
}
func NewService() UserService {  //--> factory method
	m := make(map[int64]User)
	return &uServ{data: m}
}