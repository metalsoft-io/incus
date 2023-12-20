// Code generated by "libovsdb.modelgen"
// DO NOT EDIT.

package ovsmodel

const MACBindingTable = "MAC_Binding"

// MACBinding defines an object in MAC_Binding table
type MACBinding struct {
	UUID        string `ovsdb:"_uuid"`
	Datapath    string `ovsdb:"datapath"`
	IP          string `ovsdb:"ip"`
	LogicalPort string `ovsdb:"logical_port"`
	MAC         string `ovsdb:"mac"`
	Timestamp   int    `ovsdb:"timestamp"`
}