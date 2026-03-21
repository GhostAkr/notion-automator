package configprovider

import (
	"testing"
)

func TestGetConfigReadsValidConfig(t *testing.T) {
	// Arrange
	jsonConfigProvider := JsonConfigProvider{}

	// Act
	var config ConfigMain = jsonConfigProvider.GetConfig("./../../../test/config.json")

	// Assert
	if config.InternalToken != "api token" {
		t.Errorf("InternalToken == %s, want \"api token\"", config.InternalToken)
	}
	if len(config.SetProperties) != 2 {
		t.Errorf("len(SetProperties) == %d, want 2", len(config.SetProperties))
	}

	if config.SetProperties[0].PageId != "123" {
		t.Errorf("SetProperties[0].PageId == %s, want \"123\"", config.SetProperties[0].PageId)
	}
	if config.SetProperties[0].PropertyType != "number" {
		t.Errorf("SetProperties[0].PropertyType == %s, want \"number\"", config.SetProperties[0].PropertyType)
	}
	if config.SetProperties[0].NewValue != "123456" {
		t.Errorf("SetProperties[0].NewValue == %s, want \"123456\"", config.SetProperties[0].NewValue)
	}

	if config.SetProperties[1].PageId != "456" {
		t.Errorf("SetProperties[1].PageId == %s, want \"456\"", config.SetProperties[1].PageId)
	}
	if config.SetProperties[1].PropertyType != "number" {
		t.Errorf("SetProperties[1].PropertyType == %s, want \"number\"", config.SetProperties[1].PropertyType)
	}
	if config.SetProperties[1].NewValue != "7575.46" {
		t.Errorf("SetProperties[1].NewValue == %s, want \"7575.46\"", config.SetProperties[1].NewValue)
	}
}

func TestGetConfigPanicsWhenInvalidPathGiven(t *testing.T) {
	// Arrange & Assert
	jsonConfigProvider := JsonConfigProvider{}
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("GetConfig is expected to panic but it didn't")
		}
	}()

	// Act
	jsonConfigProvider.GetConfig("invalid_path.json")
}

func TestGetConfigPanicsWhenJsonIsMalformed(t *testing.T) {
	// Arrange & Assert
	jsonConfigProvider := JsonConfigProvider{}
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("GetConfig is expected to panic but it didn't")
		}
	}()

	jsonConfigProvider.GetConfig("./../../../test/config_invalid.json")
}
