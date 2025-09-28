package decode

import (
	"encoding/base64"
	"fmt"
	"os"
	"strings"
)

// FromDataURL decodes a base64 data URL and writes the image to a file
func FromDataURL(dataURL, outPath string) error {
	parts := strings.SplitN(dataURL, ",", 2)
	if len(parts) != 2 {
		return fmt.Errorf("invalid data URL")
	}
	bin, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		return err
	}
	return os.WriteFile(outPath, bin, 0644)
}
