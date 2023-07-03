package compressor

import (
	"github.com/nfnt/resize"
	"image"
	"photo-gallery-consumer/pkg/service/compressor/types"
)

type compressed struct {
	OutImg image.Image
	types.CompressionLevel
}

var compressionType = []types.CompressionLevel{
	types.Low,
	types.Medium,
	types.High,
}

var compTypeCoefficientMap = map[types.CompressionLevel]float32{
	types.Low:    2,
	types.Medium: 1.25,
	types.High:   1.05,
}

func Compress(img image.Image) []compressed {
	bounds := img.Bounds()
	compressedImages := make([]compressed, len(compressionType))

	for index, compType := range compressionType {
		coefficient := compTypeCoefficientMap[compType]
		resized := resize.Resize(
			uint(float32(bounds.Dx())/coefficient),
			uint(float32(bounds.Dy())/coefficient),
			img,
			resize.Lanczos2,
		)
		compressedImages[index] = compressed{
			OutImg:           resized,
			CompressionLevel: compType,
		}
	}

	return compressedImages
}
