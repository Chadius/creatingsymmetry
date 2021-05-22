package wave

import (
	"wallpaper/entities/utility"
)

// WallpaperFormulaMarshalled can be marshalled into Wave Packet formulas
type WallpaperFormulaMarshalled struct {
	WavePackets []*FormulaMarshalable			`json:"wave_packets" yaml:"wave_packets"`
	Multiplier utility.ComplexNumberForMarshal	`json:"multiplier" yaml:"multiplier"`
}
