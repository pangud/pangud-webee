package biz

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewImageUsecase(t *testing.T) {

}
func TestImageUsecase_Search(t *testing.T) {
	imgs, err := imageUsecase.SearchByReference(context.Background(), "", "")
	assert.NoError(t, err)
	assert.NotEmpty(t, imgs)
	fmt.Printf("%+v\n", imgs)
	imgs2, err2 := imageUsecase.SearchByReference(context.Background(), "redis", "")
	assert.NoError(t, err2)
	assert.NotEmpty(t, imgs2)
	fmt.Printf("%+v\n", imgs2)
	imgs3, err3 := imageUsecase.SearchByReference(context.Background(), "ysql", "5")
	assert.NoError(t, err3)
	assert.NotEmpty(t, imgs3)
	fmt.Printf("%+v\n", len(imgs3))
}
