package server

import (
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"pangud.io/pangud/apps/pdcenter/internal/biz"
)

// ReverseService ftp用户业务接口
type ReverseService struct {
	log       *zap.Logger
	agentRepo biz.EndpointReadRepository
}

// NewReverseProxy new reverseProxy proxy
func NewReverseProxy(agentRepo biz.EndpointReadRepository, log *zap.Logger) *ReverseService {
	return &ReverseService{
		log:       log,
		agentRepo: agentRepo,
	}
}

// ListFtpUsers 列出ftp用户 	// api/v1/agent/:agent/proxy/:plugin
func (r *ReverseService) reverseProxy(ctx *gin.Context) {
	//check param
	agentStr := ctx.Param("agent")
	agentId, err := strconv.Atoi(agentStr)
	if err != nil {
		ctx.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	plugin := ctx.Param("plugin")
	if plugin == "" {
		ctx.Writer.WriteHeader(http.StatusNotFound)
		return
	}

	path := ctx.Param("path")
	if path == "" {
		ctx.Writer.WriteHeader(http.StatusNotFound)
		return
	}

	// get agent
	agent, err := r.agentRepo.FindOne(ctx, uint32(agentId))
	if err != nil {
		ctx.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	//todo api/v1 can be config
	url, err := url.JoinPath(agent.Addr, plugin, "api/v1", path)

	if err != nil {
		ctx.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	req, err := http.NewRequest(ctx.Request.Method, url, ctx.Request.Body)

	req.Header.Set("X-API-KEY", agent.Token)

	if err != nil {
		ctx.Writer.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		r.log.Error("reverseProxy", zap.Error(err))
		ctx.Writer.WriteHeader(http.StatusBadGateway)
		return
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		ctx.Writer.WriteHeader(http.StatusBadRequest)
		return
	}
	if res.StatusCode >= 300 {
		ctx.Writer.WriteHeader(res.StatusCode)
		return
	}
	ctx.Writer.Write(data)
}

// RegisterServices 注册反向代理服务
func (r *ReverseService) RegisterServices(router *gin.Engine) {
	// 实现代理请求到agent
	//api/v1/agent/:agent/:plugin/path
	router.Any("api/v1/proxy/agent/:agent/:plugin/*path", r.reverseProxy)
}
