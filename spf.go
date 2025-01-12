package sfp

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
)

type SimpleFinancePackage struct {
	Key       [8]byte // 8 bytes for the key
	Operation byte    // 1 byte for operation
	Asset     byte    // 1 byte for asset
	Price     float64 // 8 bytes for price (float64 for precision)
	Volume    float64 // 8 bytes for volume
}

func ParseSimpleFinanacePackage(data []byte) (*SimpleFinancePackage, error) {
	if len(data) != 26 {
		return nil, errors.New("invalid packet size, expected 26 bytes")
	}

	reader := bytes.NewReader(data)
	var packet SimpleFinancePackage

	// Read binary data into the struct
	if err := binary.Read(reader, binary.BigEndian, &packet); err != nil {
		return nil, fmt.Errorf("failed to parse packet: %w", err)
	}

	// Validate fields (example rules, customize as needed)
	if packet.Price < 0 || packet.Volume < 0 {
		return nil, errors.New("price and volume must be non-negative")
	}
	if packet.Operation > 127 { // Example: Valid operation values are 0-127
		return nil, errors.New("invalid operation value")
	}

	return &packet, nil
}

func EncodeSimpleFinanacePackage(packet SimpleFinancePackage) ([]byte, error) {
	buf := new(bytes.Buffer)

	// Write binary data to the buffer in Big-Endian order
	if err := binary.Write(buf, binary.BigEndian, packet.Key); err != nil {
		return nil, fmt.Errorf("failed to encode key: %w", err)
	}
	if err := binary.Write(buf, binary.BigEndian, packet.Operation); err != nil {
		return nil, fmt.Errorf("failed to encode operation: %w", err)
	}
	if err := binary.Write(buf, binary.BigEndian, packet.Asset); err != nil {
		return nil, fmt.Errorf("failed to encode asset: %w", err)
	}
	if err := binary.Write(buf, binary.BigEndian, packet.Price); err != nil {
		return nil, fmt.Errorf("failed to encode price: %w", err)
	}
	if err := binary.Write(buf, binary.BigEndian, packet.Volume); err != nil {
		return nil, fmt.Errorf("failed to encode volume: %w", err)
	}

	// Return the resulting byte slice
	return buf.Bytes(), nil
}