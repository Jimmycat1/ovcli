package ovextra

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"

	"github.com/HewlettPackard/oneview-golang/ov"
)

type SPCol struct {
	Type        string `json:"type"`
	Members     SPList `json:"members"`
	NextPageURI string `json:"nextPageUri"`
	Start       int    `json:"start"`
	PrevPageURI string `json:"prevPageUri"`
	Count       int    `json:"count"`
	Total       int    `json:"total"`
	Category    string `json:"category"`
	Modified    string `json:"modified"`
	ETag        string `json:"eTag"`
	Created     string `json:"created"`
	URI         string `json:"uri"`
}

type SPList []SP

type SP struct {
	Type                     string `json:"type"`
	URI                      string `json:"uri"`
	Name                     string `json:"name"`
	Description              string `json:"description"`
	SerialNumber             string `json:"serialNumber"`
	UUID                     string `json:"uuid"`
	IscsiInitiatorName       string `json:"iscsiInitiatorName"`
	IscsiInitiatorNameType   string `json:"iscsiInitiatorNameType"`
	ServerProfileTemplateURI string `json:"serverProfileTemplateUri"`
	TemplateCompliance       string `json:"templateCompliance"`
	ServerHardwareURI        string `json:"serverHardwareUri"`
	ServerHardwareTypeURI    string `json:"serverHardwareTypeUri"`
	EnclosureGroupURI        string `json:"enclosureGroupUri"`
	EnclosureURI             string `json:"enclosureUri"`
	EnclosureBay             int    `json:"enclosureBay"`
	Affinity                 string `json:"affinity"`
	AssociatedServer         string `json:"associatedServer"`
	HideUnusedFlexNics       bool   `json:"hideUnusedFlexNics"`
	Firmware                 struct {
		FirmwareScheduleDateTime string `json:"firmwareScheduleDateTime"`
		FirmwareActivationType   string `json:"firmwareActivationType"`
		FirmwareInstallType      string `json:"firmwareInstallType"`
		ForceInstallFirmware     bool   `json:"forceInstallFirmware"`
		ManageFirmware           bool   `json:"manageFirmware"`
		FirmwareBaselineURI      string `json:"firmwareBaselineUri"`
	} `json:"firmware"`
	MacType          string       `json:"macType"`
	WwnType          string       `json:"wwnType"`
	SerialNumberType string       `json:"serialNumberType"`
	Category         string       `json:"category"`
	Created          string       `json:"created"`
	Modified         string       `json:"modified"`
	Status           string       `json:"status"`
	State            string       `json:"state"`
	InProgress       bool         `json:"inProgress"`
	TaskURI          string       `json:"taskUri"`
	Connections      []Connection `json:"connections"`
	BootMode         struct {
		ManageMode    bool   `json:"manageMode"`
		PxeBootPolicy string `json:"pxeBootPolicy"`
		Mode          string `json:"mode"`
	} `json:"bootMode"`
	Boot struct {
		ManageBoot bool     `json:"manageBoot"`
		Order      []string `json:"order"`
	} `json:"boot"`
	Bios struct {
		ManageBios         bool `json:"manageBios"`
		OverriddenSettings []struct {
			ID    string `json:"id"`
			Value string `json:"value"`
		} `json:"overriddenSettings"`
	} `json:"bios"`
	LocalStorage struct {
		SasLogicalJBODs []SasLogicalJBOD `json:"sasLogicalJBODs"`
		Controllers     []struct {
			DeviceSlot          string         `json:"deviceSlot"`
			Mode                string         `json:"mode"`
			Initialize          bool           `json:"initialize"`
			ImportConfiguration bool           `json:"importConfiguration"`
			LogicalDrives       []LogicalDrive `json:"logicalDrives"`
		} `json:"controllers"`
	} `json:"localStorage"`
	SanStorage struct {
		VolumeAttachments    []string `json:"volumeAttachments"`
		SanSystemCredentials []string `json:"sanSystemCredentials"`
		ManageSanStorage     bool     `json:"manageSanStorage"`
	} `json:"sanStorage"`
	OsDeploymentSettings string `json:"osDeploymentSettings"`
	RefreshState         string `json:"refreshState"`
	ETag                 string `json:"eTag"`
	SPTemplate           string
	ServerHW             string
	ServerHWType         string
	PowerState           string
}

type Connection struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	FunctionType    string `json:"functionType"`
	NetworkURI      string `json:"networkUri"`
	PortID          string `json:"portId"`
	RequestedVFs    string `json:"requestedVFs"`
	AllocatedVFs    int    `json:"allocatedVFs"`
	InterconnectURI string `json:"interconnectUri"`
	MacType         string `json:"macType"`
	WwpnType        string `json:"wwpnType"`
	Mac             string `json:"mac"`
	Wwnn            string `json:"wwnn"`
	Wwpn            string `json:"wwpn"`
	RequestedMbps   string `json:"requestedMbps"`
	AllocatedMbps   int    `json:"allocatedMbps"`
	MaximumMbps     int    `json:"maximumMbps"`
	Ipv4            string `json:"ipv4"`
	Boot            struct {
		Priority string `json:"priority"`
	} `json:"boot"`
	State       string `json:"state"`
	Status      string `json:"status"`
	NetworkName string `json:"networkName"`
	NetworkVlan string
	ICName      string
}

type LogicalDrive struct {
	Name              string `json:"name"`
	RaidLevel         string `json:"raidLevel"`
	Bootable          bool   `json:"bootable"`
	NumPhysicalDrives int    `json:"numPhysicalDrives"`
	DriveTechnology   string `json:"driveTechnology"`
	SasLogicalJBODID  int    `json:"sasLogicalJBODId"`
	DriveNumber       int    `json:"driveNumber"`
}

type SasLogicalJBOD struct {
	ID                int    `json:"id"`
	DeviceSlot        string `json:"deviceSlot"`
	Name              string `json:"name"`
	NumPhysicalDrives int    `json:"numPhysicalDrives"`
	DriveMinSizeGB    int    `json:"driveMinSizeGB"`
	DriveMaxSizeGB    int    `json:"driveMaxSizeGB"`
	DriveTechnology   string `json:"driveTechnology"`
	SasLogicalJBODURI string `json:"sasLogicalJBODUri"`
	Status            string `json:"status"`
}

type serverprofileDetailPrint struct {
	ov.ServerProfile
	ServerProfileTemplate string
	ServerHardware        string
	ServerPower           string
	ServerHardwareType    string
	EnclosureGroup        string
	ProfileConnectionList
}

type ProfileConnectionList []ProfileConnection

type ProfileConnection struct {
	CID           int
	CName         string
	CNetwork      string
	CVLAN         string
	CMAC          string
	CPort         string
	CInterconnect string
	CBoot         string
}

type queryAttribute struct {
	attributeName  string
	attributeValue string
}

type serverProfilePrint struct {
	Name               string
	ServerHardware     string
	ServerHardwareType string
	ConsistencyState   string
}

type serverHardwarePrint struct {
	Enclosure int
	ServerBay int
}

type serverHardwareTypePrint struct {
	Name  string
	Model string
}

//var ovextra.OVClient *ov.OVClient

var serverProfilePrintlist []serverProfilePrint
var serverHardwarePrintList []serverHardwarePrint
var serverHardwareTypePrintList []serverHardwareTypePrint

func GetSP() SPList {

	var wg sync.WaitGroup

	rl := []string{"SP", "SPTemplate", "ServerHW", "ServerHWType"}

	for _, v := range rl {
		localv := v
		wg.Add(1)

		go func() {
			defer wg.Done()
			c := NewCLIOVClient()
			c.GetResourceLists(localv, "")
		}()
	}

	wg.Wait()

	spList := *(rmap["SP"].listptr.(*SPList))
	sptList := *(rmap["SPTemplate"].listptr.(*[]SPTemplate))
	hwList := *(rmap["ServerHW"].listptr.(*[]ServerHW))
	hwtList := *(rmap["ServerHWType"].listptr.(*[]ServerHWType))

	sptMap := make(map[string]SPTemplate)

	for _, v := range sptList {
		sptMap[v.URI] = v
	}

	hwMap := make(map[string]ServerHW)

	for _, v := range hwList {
		hwMap[v.URI] = v
	}

	hwtMap := make(map[string]ServerHWType)

	for _, v := range hwtList {
		hwtMap[v.URI] = v
	}

	for i, v := range spList {
		spList[i].SPTemplate = sptMap[v.ServerProfileTemplateURI].Name
		spList[i].ServerHW = hwMap[v.ServerHardwareURI].Name
		spList[i].ServerHWType = hwtMap[v.ServerHardwareTypeURI].Name

	}

	sort.Slice(spList, func(i, j int) bool { return spList[i].Name < spList[j].Name })

	return spList

}

func GetSPVerbose(name string) SPList {

	var wg sync.WaitGroup

	rl := []string{"SP", "SPTemplate", "ServerHW", "ServerHWType", "IC", "ENetwork", "NetSet"}

	for _, v := range rl {
		localv := v
		wg.Add(1)

		go func() {
			defer wg.Done()
			c := NewCLIOVClient()
			c.GetResourceLists(localv, "")
		}()
	}

	wg.Wait()

	spList := *(rmap["SP"].listptr.(*SPList))
	sptList := *(rmap["SPTemplate"].listptr.(*[]SPTemplate))
	hwList := *(rmap["ServerHW"].listptr.(*[]ServerHW))
	hwtList := *(rmap["ServerHWType"].listptr.(*[]ServerHWType))
	icList := *(rmap["IC"].listptr.(*[]IC))
	netList := *(rmap["ENetwork"].listptr.(*[]ENetwork))
	netsetList := *(rmap["NetSet"].listptr.(*[]NetSet))

	if err := (&spList).validateName(name); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	sptMap := make(map[string]SPTemplate)

	for _, v := range sptList {
		sptMap[v.URI] = v
	}

	hwMap := make(map[string]ServerHW)

	for _, v := range hwList {
		hwMap[v.URI] = v
	}

	hwtMap := make(map[string]ServerHWType)

	for _, v := range hwtList {
		hwtMap[v.URI] = v
	}

	for i, v := range spList {
		spList[i].SPTemplate = sptMap[v.ServerProfileTemplateURI].Name
		spList[i].ServerHW = hwMap[v.ServerHardwareURI].Name
		spList[i].PowerState = hwMap[v.ServerHardwareURI].PowerState
		spList[i].ServerHWType = hwtMap[v.ServerHardwareTypeURI].Name

		spList[i].conns(icList, netList, netsetList)

	}

	return spList

}

func (sp *SP) conns(icList []IC, netList []ENetwork, netsetList []NetSet) {

	icMap := make(map[string]IC)

	for _, v := range icList {
		icMap[v.URI] = v
	}

	netMap := make(map[string]ENetwork)
	for _, v := range netList {
		netMap[v.URI] = v
	}

	netsetMap := make(map[string]NetSet)
	for _, v := range netsetList {
		netsetMap[v.URI] = v
	}

	for i, v := range sp.Connections {

		if strings.Contains(v.NetworkURI, "ethernet-networks") {

			sp.Connections[i].NetworkName = netMap[v.NetworkURI].Name
			sp.Connections[i].NetworkVlan = strconv.Itoa(netMap[v.NetworkURI].VlanId)
		} else {
			sp.Connections[i].NetworkName = netsetMap[v.NetworkURI].Name
			sp.Connections[i].NetworkVlan = "NetworkSet"
		}

		sp.Connections[i].ICName = icMap[v.InterconnectURI].Name

	}

}

func (list *SPList) validateName(name string) error {

	if name == "all" {
		return nil //if name is all, don't touch *list, directly return
	}

	localslice := *list //define a localslice to avoid too many *list in the following

	for i, v := range localslice {
		if name == v.Name {
			localslice = localslice[i : i+1] //if name is valid, only display one LIG instead of whole list
			*list = localslice               //update list pointer to point to new shortened slice
			return nil
		}
	}

	return fmt.Errorf("no profile matching name: \"%v\" was found, please check spelling and syntax, valid syntax example: \"show serverprofile --name profile1\" ", name)

}
