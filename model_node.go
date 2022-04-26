/*
 * Tencent Docs SDK
 *
 * This is tencent docs editorapi sdk
 *
 * API version: 0.1.0
 * Contact: tencentdocs@tencent.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type Node struct {
	Begin    int32                  `json:"begin,omitempty"`
	End      int32                  `json:"end,omitempty"`
	Type_    string                 `json:"type,omitempty"`
	Children []Node                 `json:"children,omitempty"`
	Text     string                 `json:"text,omitempty"`
	Property map[string]interface{} `json:"property,omitempty"`
}
