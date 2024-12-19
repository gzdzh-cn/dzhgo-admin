// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// BaseEpsApp is the golang structure for table base_eps_app.
type BaseEpsApp struct {
	Id      string `json:"id"      orm:"id"      ` //
	Module  string `json:"module"  orm:"module"  ` //
	Method  string `json:"method"  orm:"method"  ` //
	Path    string `json:"path"    orm:"path"    ` //
	Prefix  string `json:"prefix"  orm:"prefix"  ` //
	Summary string `json:"summary" orm:"summary" ` //
	Tag     string `json:"tag"     orm:"tag"     ` //
	Dts     string `json:"dts"     orm:"dts"     ` //
}
