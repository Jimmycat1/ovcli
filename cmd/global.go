package cmd

import (
	"github.com/HewlettPackard/oneview-golang/ov"
	"github.com/docker/machine/libmachine/log"
)

var serverProfileList ov.ServerProfileList
var serverHardwareList ov.ServerHardwareList
var serverHardwareTypeList ov.ServerHardwareTypeList
var profileTemplateList ov.ServerProfileList
var enclosureGroupList ov.EnclosureGroupList

var logicalInterconnectGroupList ov.LogicalInterconnectGroupList

var interconnectTypeList ov.InterconnectTypeList
var ethernetNetworkList ov.EthernetNetworkList
var networkSetList ov.NetworkSetList

var empty_query_string = make(map[string]interface{})

//var OVClient *ov.OVClien

var ligModuleList []LIGModule

type uplinkPortListType []UplinkPort

var (
	profileNamePtr          *string
	ligNamePtr              *string
	createNetworkNamePtr    *string
	createNetworkTypePtr    *string
	createNetworkPurposePtr *string
	createNetworkVlanIDPtr  *int
	porttype                string
)

func init() {
	log.SetDebug(true)

}
