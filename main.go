package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func proxy(c *gin.Context) {
	remote, err := url.Parse("http://localhost:8082")
	if err != nil {
		c.String(http.StatusInternalServerError, "Erreur de parsing de l'URL du backend")
		return
	}
	proxy := httputil.NewSingleHostReverseProxy(remote)
	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		originalDirector(req)
		proxyPath := c.Param("proxyPath")
		if proxyPath == "" {
			proxyPath = "/"
		}
		req.URL.Path = strings.TrimSuffix(remote.Path, "/") + proxyPath
	}
	proxy.ServeHTTP(c.Writer, c.Request)
}

func main() {
	r := gin.Default()
	r.Any("/proxy/*proxyPath", proxy)
	r.Run(":8081")
}
