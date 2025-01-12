package sfp

import (
	"bytes"
	"testing"
)

//ChatGPT Generated Code to test the Encode and Decode functions

func TestEncodeDecodeCustomPacket(t *testing.T) {
	// Define a sample packet
	originalPacket := SimpleFinancePackage{
		Key:       [8]byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H'},
		Operation: 1,
		Asset:     2,
		Price:     1234.56,
		Volume:    7890.12,
	}

	// Encode the packet
	encoded, err := EncodeSimpleFinanacePackage(originalPacket)
	if err != nil {
		t.Fatalf("Encoding failed: %v", err)
	}

	// Check the encoded length
	if len(encoded) != 28 {
		t.Fatalf("Encoded data has incorrect length: got %d, want 28", len(encoded))
	}

	// Decode the packet
	decodedPacket, err := ParseSimpleFinanacePackage(encoded)
	if err != nil {
		t.Fatalf("Decoding failed: %v", err)
	}

	// Compare the original and decoded packets
	if !bytes.Equal(originalPacket.Key[:], decodedPacket.Key[:]) {
		t.Errorf("Key mismatch: got %v, want %v", decodedPacket.Key, originalPacket.Key)
	}
	if originalPacket.Operation != decodedPacket.Operation {
		t.Errorf("Operation mismatch: got %d, want %d", decodedPacket.Operation, originalPacket.Operation)
	}
	if originalPacket.Asset != decodedPacket.Asset {
		t.Errorf("Asset mismatch: got %d, want %d", decodedPacket.Asset, originalPacket.Asset)
	}
	if originalPacket.Price != decodedPacket.Price {
		t.Errorf("Price mismatch: got %f, want %f", decodedPacket.Price, originalPacket.Price)
	}
	if originalPacket.Volume != decodedPacket.Volume {
		t.Errorf("Volume mismatch: got %f, want %f", decodedPacket.Volume, originalPacket.Volume)
	}
}

func TestInvalidPacketSizes(t *testing.T) {
	// Test with too small a packet
	smallPacket := make([]byte, 10)
	_, err := ParseSimpleFinanacePackage(smallPacket)
	if err == nil {
		t.Fatal("Expected error for small packet, got nil")
	}

	// Test with too large a packet
	largePacket := make([]byte, 30)
	_, err = ParseSimpleFinanacePackage(largePacket)
	if err == nil {
		t.Fatal("Expected error for large packet, got nil")
	}
}

func TestInvalidPacketValues(t *testing.T) {
	// Create a valid packet and modify its fields
	validPacket := SimpleFinancePackage{
		Key:       [8]byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H'},
		Operation: 1,
		Asset:     2,
		Price:     1234.56,
		Volume:    7890.12,
	}
	encoded, _ := EncodeSimpleFinanacePackage(validPacket)

	// Modify encoded data to introduce errors
	encoded[9] = 255 // Invalid operation value

	_, err := ParseSimpleFinanacePackage(encoded)
	if err == nil {
		t.Fatal("Expected error for invalid operation value, got nil")
	}

	// Modify price to a negative value
	copy(encoded[10:], []byte{255, 255, 255, 255, 255, 255, 255, 255}) // Represents -1 in float64 (Big-Endian)
	_, err = ParseSimpleFinanacePackage(encoded)
	if err == nil {
		t.Fatal("Expected error for negative price, got nil")
	}
}
