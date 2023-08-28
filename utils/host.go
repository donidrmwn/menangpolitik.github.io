package utils

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type HostSetting struct {
	FrontEndUrl string
	BackendUrl  string
	Company     string
	ViewFolder  string
	CompanyID   int
}

func Host(c *fiber.Ctx) (out HostSetting) {
	protocol := "http"
	if c.Protocol() == "https" {
		protocol = "https"
	}
	host := c.Hostname()
	url := fmt.Sprintf("%s://%s", protocol, host)
	var data = map[string]HostSetting{

		"http://localhost:3000": {
			FrontEndUrl: "http://localhost:3000",
			BackendUrl:  "http://apidev.politikpintar.id",
			Company:     "1",
			ViewFolder:  "menangpol",
		},
	}
	out = data[url]
	if out.FrontEndUrl == "" {
		out = data[""]
	}
	return
}
