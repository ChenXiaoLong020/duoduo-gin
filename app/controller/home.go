package controller

import (
	"duoduo-gin/app/respone"
	"duoduo-gin/model"
	"duoduo-gin/service"
	"strconv"
	"sync"

	"github.com/gin-gonic/gin"
)

type Home struct {
	wg sync.WaitGroup
}

func (h Home) Index(c *gin.Context) {
	respone.WithData(c, service.User{}.Find())
}

func (h Home) Add(c *gin.Context) {

	for i := 0; i < 10; i++ {
		h.wg.Add(1)
		go func() {
			service.User{}.Create(model.Users{
				Username: "xiaobai",
				Password: strconv.Itoa(i),
			})
			h.wg.Done()
		}()
	}

	h.wg.Wait()
	respone.Success(c)
}
