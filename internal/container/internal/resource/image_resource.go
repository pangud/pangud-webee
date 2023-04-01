package resource

import (
	"go.uber.org/zap"
)

type ImageResouce struct {
	log *zap.Logger
}

func NewImageResource(logger *zap.Logger) *ImageResouce {
	return &ImageResouce{
		log: logger,
	}
}

func (r *ImageResouce) ListImages() {
	r.log.Sugar().Debug("ab")
}
