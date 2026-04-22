package dto

import (
	"encoding/json"
	"reflect"
	"testing"

	"pgregory.net/rapid"
)

// Feature: auto-import-menu, Property 1: MenuPermissionIO JSON 序列化往返
// Validates: Requirements 3.1

func genApiIO(t *rapid.T) ApiIO {
	return ApiIO{
		Method: rapid.String().Draw(t, "method"),
		URL:    rapid.String().Draw(t, "url"),
	}
}

func genPermissionIO(t *rapid.T) PermissionIO {
	n := rapid.IntRange(0, 5).Draw(t, "numApis")
	apis := make([]ApiIO, n)
	for i := range apis {
		apis[i] = genApiIO(t)
	}
	return PermissionIO{
		Code: rapid.String().Draw(t, "code"),
		Name: rapid.String().Draw(t, "name"),
		Type: rapid.String().Draw(t, "type"),
		Apis: apis,
	}
}

func genMenuIO(t *rapid.T, depth int) MenuIO {
	var children []MenuIO
	if depth > 0 {
		n := rapid.IntRange(0, 3).Draw(t, "numChildren")
		children = make([]MenuIO, n)
		for i := range children {
			children[i] = genMenuIO(t, depth-1)
		}
	} else {
		children = []MenuIO{}
	}
	return MenuIO{
		MenuType:       rapid.String().Draw(t, "menuType"),
		Path:           rapid.String().Draw(t, "path"),
		Component:      rapid.String().Draw(t, "component"),
		Perm:           rapid.String().Draw(t, "perm"),
		MenuName:       rapid.String().Draw(t, "menuName"),
		Title:          rapid.String().Draw(t, "title"),
		PermissionCode: rapid.String().Draw(t, "permissionCode"),
		Icon:           rapid.String().Draw(t, "icon"),
		SortValue:      rapid.Int().Draw(t, "sortValue"),
		IsExternal:     rapid.Bool().Draw(t, "isExternal"),
		ExternalLink:   rapid.String().Draw(t, "externalLink"),
		TextBadge:      rapid.String().Draw(t, "textBadge"),
		ActivePath:     rapid.String().Draw(t, "activePath"),
		Status:         rapid.String().Draw(t, "status"),
		KeepAlive:      rapid.Bool().Draw(t, "keepAlive"),
		IsHide:         rapid.Bool().Draw(t, "isHide"),
		IsIframe:       rapid.Bool().Draw(t, "isIframe"),
		ShowBadge:      rapid.Bool().Draw(t, "showBadge"),
		FixedTab:       rapid.Bool().Draw(t, "fixedTab"),
		IsHideTab:      rapid.Bool().Draw(t, "isHideTab"),
		IsFullPage:     rapid.Bool().Draw(t, "isFullPage"),
		Children:       children,
	}
}

func genMenuPermissionIO(t *rapid.T) MenuPermissionIO {
	numMenus := rapid.IntRange(0, 5).Draw(t, "numMenus")
	menus := make([]MenuIO, numMenus)
	maxDepth := rapid.IntRange(0, 3).Draw(t, "maxDepth")
	for i := range menus {
		menus[i] = genMenuIO(t, maxDepth)
	}

	numPerms := rapid.IntRange(0, 5).Draw(t, "numPermissions")
	perms := make([]PermissionIO, numPerms)
	for i := range perms {
		perms[i] = genPermissionIO(t)
	}

	return MenuPermissionIO{
		Menus:       menus,
		Permissions: perms,
	}
}

// normalizeMenuPermissionIO replaces nil slices with empty slices so that
// reflect.DeepEqual works correctly after a JSON round-trip (json.Unmarshal
// produces nil slices for JSON null / missing fields).
func normalizeMenuPermissionIO(m *MenuPermissionIO) {
	if m.Menus == nil {
		m.Menus = []MenuIO{}
	}
	if m.Permissions == nil {
		m.Permissions = []PermissionIO{}
	}
	for i := range m.Menus {
		normalizeMenuIO(&m.Menus[i])
	}
	for i := range m.Permissions {
		normalizePermissionIO(&m.Permissions[i])
	}
}

func normalizeMenuIO(menu *MenuIO) {
	if menu.Children == nil {
		menu.Children = []MenuIO{}
	}
	for i := range menu.Children {
		normalizeMenuIO(&menu.Children[i])
	}
}

func normalizePermissionIO(perm *PermissionIO) {
	if perm.Apis == nil {
		perm.Apis = []ApiIO{}
	}
}

func TestMenuPermissionIORoundTrip(t *testing.T) {
	// Feature: auto-import-menu, Property 1: MenuPermissionIO JSON 序列化往返
	// Validates: Requirements 3.1
	rapid.Check(t, func(t *rapid.T) {
		original := genMenuPermissionIO(t)

		data, err := json.Marshal(original)
		if err != nil {
			t.Fatalf("json.Marshal failed: %v", err)
		}

		var decoded MenuPermissionIO
		err = json.Unmarshal(data, &decoded)
		if err != nil {
			t.Fatalf("json.Unmarshal failed: %v", err)
		}

		// Normalize both sides to handle nil vs empty slice differences
		normalizeMenuPermissionIO(&original)
		normalizeMenuPermissionIO(&decoded)

		if !reflect.DeepEqual(original, decoded) {
			t.Fatalf("round-trip mismatch:\noriginal: %+v\ndecoded:  %+v", original, decoded)
		}
	})
}

// Feature: auto-import-menu, Property 2: JSON 结构验证正确性
// Validates: Requirements 1.2, 3.1

func TestJSONStructureValidation(t *testing.T) {
	// Feature: auto-import-menu, Property 2: JSON 结构验证正确性
	// **Validates: Requirements 1.2**

	t.Run("valid_structures_parse_correctly", func(t *testing.T) {
		// Generate random valid MenuPermissionIO structs, marshal to JSON,
		// unmarshal back - should succeed with correct values.
		rapid.Check(t, func(t *rapid.T) {
			original := genMenuPermissionIO(t)

			data, err := json.Marshal(original)
			if err != nil {
				t.Fatalf("json.Marshal failed: %v", err)
			}

			var result MenuPermissionIO
			err = json.Unmarshal(data, &result)
			if err != nil {
				t.Fatalf("valid JSON structure should parse without error: %v", err)
			}

			// A valid structure with menus array and permissions array should parse
			// and produce non-zero values when the original had non-zero values.
			normalizeMenuPermissionIO(&original)
			normalizeMenuPermissionIO(&result)

			if !reflect.DeepEqual(original, result) {
				t.Fatalf("valid structure did not parse correctly:\noriginal: %+v\nresult:   %+v", original, result)
			}
		})
	})

	t.Run("arbitrary_bytes_do_not_panic", func(t *testing.T) {
		// Generate random byte slices - unmarshal should not panic.
		// If unmarshal fails, the result should be zero-valued.
		rapid.Check(t, func(t *rapid.T) {
			data := rapid.SliceOf(rapid.Byte()).Draw(t, "randomBytes")

			var result MenuPermissionIO
			func() {
				defer func() {
					if r := recover(); r != nil {
						t.Fatalf("json.Unmarshal panicked on arbitrary bytes: %v", r)
					}
				}()
				err := json.Unmarshal(data, &result)
				if err != nil {
					// Unmarshal failed - result should be zero-valued
					if result.Menus != nil {
						t.Fatalf("expected nil Menus after failed unmarshal, got: %+v", result.Menus)
					}
					if result.Permissions != nil {
						t.Fatalf("expected nil Permissions after failed unmarshal, got: %+v", result.Permissions)
					}
				}
			}()
		})
	})

	t.Run("invalid_structure_fields_are_zero", func(t *testing.T) {
		// Generate JSON objects with missing or wrong-typed fields.
		// Unmarshal should not panic. For missing fields the result should be zero-valued.
		// Note: Go's json.Unmarshal does partial parsing - when one field has a type
		// mismatch, other correctly-typed fields may still be populated. We verify:
		// 1. No panic occurs
		// 2. For empty objects, both fields are nil (zero value)
		// 3. For objects with only unrelated keys, both fields are nil
		rapid.Check(t, func(t *rapid.T) {
			// Pick a category of invalid/incomplete JSON structure
			category := rapid.IntRange(0, 4).Draw(t, "category")

			var data []byte
			var expectBothNil bool
			switch category {
			case 0:
				// Empty object - both fields should be nil
				data = []byte(`{}`)
				expectBothNil = true
			case 1:
				// menus is not an array - unmarshal error, but permissions may be partially parsed
				data = []byte(`{"menus": "not_an_array", "permissions": []}`)
				expectBothNil = false
			case 2:
				// permissions is not an array - unmarshal error, but menus may be partially parsed
				data = []byte(`{"menus": [], "permissions": "not_an_array"}`)
				expectBothNil = false
			case 3:
				// Random keys that don't match schema - both fields should be nil
				// Avoid generating keys that match actual field names or json tags
				key := "rnd_" + rapid.String().Draw(t, "randomKey")
				val := rapid.IntRange(0, 9999).Draw(t, "randomVal")
				data = []byte(`{"` + escapeJSONString(key) + `": ` + intToStr(val) + `}`)
				expectBothNil = true
			case 4:
				// Both fields have wrong types - unmarshal error
				data = []byte(`{"menus": 123, "permissions": true}`)
				expectBothNil = false
			}

			var result MenuPermissionIO
			func() {
				defer func() {
					if r := recover(); r != nil {
						t.Fatalf("json.Unmarshal panicked on invalid structure: %v\ninput: %s", r, string(data))
					}
				}()
				_ = json.Unmarshal(data, &result)
			}()

			if expectBothNil {
				// For empty objects or unrelated keys, fields should remain zero-valued
				if result.Menus != nil {
					t.Fatalf("expected nil Menus for input %s, got: %+v", string(data), result.Menus)
				}
				if result.Permissions != nil {
					t.Fatalf("expected nil Permissions for input %s, got: %+v", string(data), result.Permissions)
				}
			}
			// For wrong-typed fields (categories 1, 2, 4), we only verify no panic.
			// Go's json.Unmarshal may partially populate fields before encountering
			// a type mismatch, which is correct behavior.
		})
	})
}

// escapeJSONString escapes a string for safe embedding in a JSON string literal.
func escapeJSONString(s string) string {
	b, err := json.Marshal(s)
	if err != nil {
		return "key"
	}
	// json.Marshal wraps in quotes, strip them
	if len(b) >= 2 {
		return string(b[1 : len(b)-1])
	}
	return "key"
}

// intToStr converts an int to its string representation.
func intToStr(n int) string {
	return json.Number(rapidIntToString(n)).String()
}

func rapidIntToString(n int) string {
	if n == 0 {
		return "0"
	}
	neg := false
	if n < 0 {
		neg = true
		n = -n
	}
	digits := make([]byte, 0, 20)
	for n > 0 {
		digits = append(digits, byte('0'+n%10))
		n /= 10
	}
	// reverse
	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}
	if neg {
		return "-" + string(digits)
	}
	return string(digits)
}
