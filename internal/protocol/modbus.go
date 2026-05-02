// GRID-SAST-010: out-of-bounds slice indexing on attacker-controlled length.
package protocol

import "encoding/binary"

// ParseFrame decodes a Modbus-like ADU frame.
func ParseFrame(buf []byte) (uint16, []byte) {
	// declared length comes from the frame itself; not bounds-checked against len(buf).
	declared := binary.BigEndian.Uint16(buf[4:6])
	// GRID-SAST-010: if `declared` > len(buf), this slice expression panics or reads OOB memory
	// in older Go versions / unsafe variants.
	payload := buf[6 : 6+int(declared)]
	return declared, payload
}
