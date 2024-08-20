package pcmtoopus

import (
	"gopkg.in/hraban/opus.v2"
	"go.k6.io/k6/js/modules"
)

// RootModule is the module's root type.
type RootModule struct{}

// New creates a new instance of the module, exposing its functionality to the JS runtime.
func (r *RootModule) New() interface{} {
	return &PCMToOpus{}
}

// PCMToOpus type will expose the methods for the plugin.
type PCMToOpus struct{}

// OpusEncoder structure to hold the Opus encoder for each VU.
type OpusEncoder struct {
	encoder *opus.Encoder
}

// NewEncoder initializes a new Opus encoder.
func (p *PCMToOpus) NewEncoder(sampleRate int, channels int, application opus.Application) (*OpusEncoder, error) {
	enc, err := opus.NewEncoder(sampleRate, channels, application)
	if err != nil {
		return nil, err
	}
	return &OpusEncoder{encoder: enc}, nil
}

// Encode encodes a chunk of PCM float32 data to Opus using EncodeFloat32.
func (oe *OpusEncoder) Encode(audioChunk []float32) ([]byte, error) {
	data := make([]byte, 500)  // Allocate a buffer for the encoded data
	n, err := oe.encoder.EncodeFloat32(audioChunk, data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

// Close the Opus encoder.
func (oe *OpusEncoder) Close() {
	oe.encoder = nil
}

func init() {
	modules.Register("k6/x/pcmtoopus", &RootModule{})
}

