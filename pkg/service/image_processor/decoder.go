package image_processor

import (
	"bufio"
	"bytes"
	"errors"
	"image"
	"image/gif"
	_ "image/gif"
	"image/jpeg"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"io"
)

func Decode(reader io.Reader) (decoded image.Image, err error) {
	decoded, _, err = image.Decode(reader)

	return
}

func Encode(img image.Image, format string) (decoded []byte, err error) {
	var decodedBuf bytes.Buffer
	decodedWriter := bufio.NewWriter(&decodedBuf)

	switch format {
	case "jpeg":
	case "jpg":
		{
			if err = jpeg.Encode(decodedWriter, img, nil); err != nil {
				return
			}
			break
		}

	case "png":
		{
			if err = png.Encode(decodedWriter, img); err != nil {
				return
			}

			break
		}

	case "gif":
		{
			if err = gif.Encode(decodedWriter, img, nil); err != nil {
				return
			}

			break
		}

	default:
		err = errors.New("unexpected format")
	}

	decoded = decodedBuf.Bytes()

	return
}
