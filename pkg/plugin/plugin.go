package plug

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// PluginType 插件类型
type PluginType uint

const (
	TypePreset      PluginType = iota // 预置插件
	TypeOfficialExt                   // 官方扩展插件
	//TYPE_THIRD_EXT                      // 第三方扩展插件
)

type Plugin interface {
	// Name returns the name of the plugin.
	Name() string
	// Version returns the version of the plugin.
	Version() string
	// Type returns the type of the plugin.
	Type() PluginType
}

// AgentPlug agent plugin
type AgentPlug interface {
	Plugin
	// Register initializes the plugin.
	Register(r *gin.Engine, db *gorm.DB, log *zap.Logger) error
}

// ServerPlug Server plugin
type ServerPlug interface {
	Plugin
	// Register initializes the plugin.
	Register(r *gin.Engine, db *gorm.DB, log *zap.Logger) *gin.Engine
}

// EdgePlug edge plugin
//type EdgePlug interface {
//	Plugin
//	Register(client mqtt.Client, log *zap.Logger) error
//}

// RegisterAgentPlugFunc register agent plugin func
type RegisterAgentPlugFunc func(engine *gin.Engine, db *gorm.DB, ownCfg string, log *zap.Logger) (AgentPlug, error)
