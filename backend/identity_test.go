package main

import (
	"encoding/json"
	"os"
	"strings"
	"testing"
)

func TestIndependentForkIdentityAndMCPDeclarations(t *testing.T) {
	raw, err := os.ReadFile("../manifest.json")
	if err != nil {
		t.Fatal(err)
	}
	var m struct {
		ID, Version, Publisher string
		Contributions          []struct {
			ID, Type string
			Tools    map[string]struct{ Access string }
		}
	}
	if err = json.Unmarshal(raw, &m); err != nil {
		t.Fatal(err)
	}
	if m.ID != pluginID || m.Version != version || m.Publisher != "crayon-hua" {
		t.Fatal("manifest and native identity differ")
	}
	seen := map[string]bool{}
	count := 0
	for _, c := range m.Contributions {
		if !strings.HasPrefix(c.ID, pluginID+".") || seen[c.ID] {
			t.Fatalf("invalid/duplicate contribution %s", c.ID)
		}
		seen[c.ID] = true
		if c.Type != "mcp" {
			continue
		}
		count++
		if len(c.Tools) != len(toolCatalog()) {
			t.Fatal("MCP declarations differ from backend catalog")
		}
		for _, tool := range toolCatalog() {
			decl, ok := c.Tools[tool.Name]
			if !ok || (decl.Access != "read" && decl.Access != "write" && decl.Access != "full") {
				t.Fatalf("missing access for %s", tool.Name)
			}
			if (decl.Access == "read") != tool.Annotations["readOnlyHint"] {
				t.Fatalf("access and read-only annotation disagree for %s", tool.Name)
			}
		}
	}
	if count != 1 {
		t.Fatal("exactly one MCP contribution is required")
	}
	ui, err := os.ReadFile("../src/App.vue")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(ui), pluginID+".log") || strings.Contains(string(ui), "io.dbx.xxljob-console") {
		t.Fatal("log workbench still targets upstream identity")
	}
}
