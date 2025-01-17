// +build !fast

package screenshotr_test

import (
	"encoding/binary"
	"github.com/danielpaulus/go-ios/ios"
	"github.com/danielpaulus/go-ios/ios/screenshotr"
	"github.com/stretchr/testify/assert"
	"testing"
)

const png uint32 = 0x89504E47

func TestScreenshot(t *testing.T) {
	device, err := ios.GetDevice("")
	if err != nil {
		t.Error(err)
		return
	}
	screenshotrService, err := screenshotr.New(device)
	if err != nil {
		t.Error(err)
		return
	}
	imageBytes, err := screenshotrService.TakeScreenshot()
	if assert.NoError(t, err) {
		assert.Equal(t, png, binary.BigEndian.Uint32(imageBytes[0:4]),
			"check that response starts with png magic bytes")
	}
}
