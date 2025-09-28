package encode

import (
	"encoding/base64"
	"fmt"
	"os"
)

// ToDataURL encodes an image file to a base64 data URL
func ToDataURL(imgPath string) (string, error) {
	bin, err := os.ReadFile(imgPath)
	if err != nil {
		return "", err
	}
	b64 := base64.StdEncoding.EncodeToString(bin)
	return fmt.Sprintf("data:image/png;base64,%s", b64), nil
}
