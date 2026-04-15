package opslevel_test

import (
	"encoding/json"
	"testing"

	opslevel "github.com/opslevel/opslevel-go/v2026"
)

func TestComponentTypeInput_CategoryOmittedWhenNil(t *testing.T) {
	input := opslevel.ComponentTypeInput{
		Name:  opslevel.NewNullableFrom("My Service"),
		Alias: opslevel.NewNullableFrom("my_service"),
	}

	data, err := json.Marshal(input)
	if err != nil {
		t.Fatalf("unexpected marshal error: %s", err)
	}

	var raw map[string]interface{}
	if err := json.Unmarshal(data, &raw); err != nil {
		t.Fatalf("unexpected unmarshal error: %s", err)
	}

	if _, exists := raw["category"]; exists {
		t.Errorf("category should be omitted from JSON when nil, got: %s", string(data))
	}
}

func TestComponentTypeInput_CategoryIncludedWhenSet(t *testing.T) {
	tests := []struct {
		name     string
		category string
	}{
		{"default category", "default"},
		{"infrastructure category", "infrastructure"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := opslevel.ComponentTypeInput{
				Name:     opslevel.NewNullableFrom("My Type"),
				Alias:    opslevel.NewNullableFrom("my_type"),
				Category: opslevel.NewNullableFrom(tt.category),
			}

			data, err := json.Marshal(input)
			if err != nil {
				t.Fatalf("unexpected marshal error: %s", err)
			}

			var raw map[string]interface{}
			if err := json.Unmarshal(data, &raw); err != nil {
				t.Fatalf("unexpected unmarshal error: %s", err)
			}

			got, exists := raw["category"]
			if !exists {
				t.Fatalf("category should be present in JSON when set, got: %s", string(data))
			}
			if got != tt.category {
				t.Errorf("expected category %q, got %q", tt.category, got)
			}
		})
	}
}

func TestComponentTypeInput_CategoryNullWhenExplicitlyNulled(t *testing.T) {
	input := opslevel.ComponentTypeInput{
		Name:     opslevel.NewNullableFrom("My Type"),
		Category: opslevel.NewNullOf[string](),
	}

	data, err := json.Marshal(input)
	if err != nil {
		t.Fatalf("unexpected marshal error: %s", err)
	}

	var raw map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		t.Fatalf("unexpected unmarshal error: %s", err)
	}

	catRaw, exists := raw["category"]
	if !exists {
		t.Fatalf("category should be present in JSON when explicitly nulled, got: %s", string(data))
	}
	if string(catRaw) != "null" {
		t.Errorf("expected category to be null, got: %s", string(catRaw))
	}
}

func TestComponentType_CategoryParsedFromJSON(t *testing.T) {
	tests := []struct {
		name     string
		json     string
		expected string
	}{
		{
			"default category",
			`{"category":"default","name":"Service","description":"","href":"","isDefault":true}`,
			"default",
		},
		{
			"infrastructure category",
			`{"category":"infrastructure","name":"Redis Cloud","description":"","href":"","isDefault":false}`,
			"infrastructure",
		},
		{
			"empty category",
			`{"category":"","name":"Legacy","description":"","href":"","isDefault":false}`,
			"",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var ct opslevel.ComponentType
			if err := json.Unmarshal([]byte(tt.json), &ct); err != nil {
				t.Fatalf("unexpected unmarshal error: %s", err)
			}
			if ct.Category != tt.expected {
				t.Errorf("expected category %q, got %q", tt.expected, ct.Category)
			}
		})
	}
}
