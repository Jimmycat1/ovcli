package oneview

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"sync"

	"github.com/ghodss/yaml"
)

type LIGCol struct {
	Total       int     `json:"total,omitempty"`       // "total": 1,
	Count       int     `json:"count,omitempty"`       // "count": 1,
	Start       int     `json:"start,omitempty"`       // "start": 0,
	PrevPageURI string  `json:"prevPageUri,omitempty"` // "prevPageUri": null,
	NextPageURI string  `json:"nextPageUri,omitempty"` // "nextPageUri": null,
	URI         string  `json:"uri,omitempty"`         // "uri": "/rest/server-profiles?filter=connectionTemplateUri%20matches%7769cae0-b680-435b-9b87-9b864c81657fsort=name:asc"
	Members     LIGList `json:"members,omitempty"`     // "members":[]
}

type LIGList []LIG

type LIG struct {
	Category                string                  `json:"category,omitempty"`         // "category": "logical-interconnect-groups",
	Created                 string                  `json:"created,omitempty"`          // "created": "20150831T154835.250Z",
	Description             string                  `json:"description,omitempty"`      // "description": "Logical Interconnect Group 1",
	ETAG                    string                  `json:"eTag,omitempty"`             // "eTag": "1441036118675/8",
	EnclosureIndexes        []int                   `json:"enclosureIndexes,omitempty"` // "enclosureIndexes": [1],
	EnclosureType           string                  `json:"enclosureType,omitempty"`    // "enclosureType": "C7000",
	EthernetSettings        *EthernetSettings       `json:"ethernetSettings,omitempty"` // "ethernetSettings": {...},
	FabricUri               string                  `json:"fabricUri,omitempty"`        // "fabricUri": "/rest/fabrics/9b8f7ec0-52b3-475e-84f4-c4eac51c2c20",
	InterconnectBaySet      int                     `json:"interconnectBaySet"`
	InterconnectMapTemplate InterconnectMapTemplate `json:"interconnectMapTemplate"` // "interconnectMapTemplate": {...},
	// InterconnectMapTemplate *InterconnectMapTemplate `json:"interconnectMapTemplate"`          // "interconnectMapTemplate": {...},
	InternalNetworkUris    []string                `json:"internalNetworkUris,omitempty"`    // "internalNetworkUris": []
	Modified               string                  `json:"modified,omitempty"`               // "modified": "20150831T154835.250Z",
	Name                   string                  `json:"name"`                             // "name": "Logical Interconnect Group1",
	QosConfiguration       *QosConfiguration       `json:"qosConfiguration,omitempty"`       // "qosConfiguration": {},
	RedundancyType         string                  `json:"redundancyType,omitempty"`         // "redundancyType": "HighlyAvailable"
	SnmpConfiguration      *SnmpConfiguration      `json:"snmpConfiguration,omitempty"`      // "snmpConfiguration": {...}
	StackingHealth         string                  `json:"stackingHealth,omitempty"`         //"stackingHealth": "Connected",
	StackingMode           string                  `json:"stackingMode,omitempty"`           //"stackingMode": "Enclosure",
	State                  string                  `json:"state,omitempty"`                  // "state": "Normal",
	Status                 string                  `json:"status,omitempty"`                 // "status": "Critical",
	TelemetryConfiguration *TelemetryConfiguration `json:"telemetryConfiguration,omitempty"` // "telemetryConfiguration": {...},
	Type                   string                  `json:"type"`                             // "type": "logical-interconnect-groupsV3",
	UplinkSets             []LIGUplinkSet          `json:"uplinkSets"`                       // "uplinkSets": {...},
	URI                    string                  `json:"uri,omitempty"`                    // "uri": "/rest/logical-interconnect-groups/e2f0031b-52bd-4223-9ac1-d91cb519d548",
	IOBays                 LIGIOBayList            `json:"-"`                                //define []IOBay as named type to use multisort method later
}

type InterconnectMapTemplate struct {
	InterconnectMapEntryTemplates []InterconnectMapEntryTemplate `json:"interconnectMapEntryTemplates"` // "interconnectMapEntryTemplates": {...},
}

type InterconnectMapEntryTemplate struct {
	EnclosureIndex               int             `json:"enclosureIndex,omitempty"`               // "enclosureIndex": 1,
	LogicalDownlinkUri           string          `json:"logicalDownlinkUri,omitempty"`           // "logicalDownlinkUri": "/rest/logical-downlinks/5b33fec1-63e8-40e1-9e3d-3af928917b2f",
	LogicalLocation              LogicalLocation `json:"logicalLocation,omitempty"`              // "logicalLocation": {...},
	PermittedInterconnectTypeUri string          `json:"permittedInterconnectTypeUri,omitempty"` //"permittedSwitchTypeUri": "/rest/switch-types/a2bc8f42-8bb8-4560-b80f-6c3c0e0d66e0",
}

type LogicalLocation struct {
	LocationEntries []LocationEntry `json:"locationEntries,omitempty"` // "locationEntries": {...}
}

type LocationEntry struct {
	RelativeValue int    `json:"relativeValue,omitempty"` //"relativeValue": 2,
	Type          string `json:"type,omitempty"`          //"type": "StackingMemberId",
}

type LogicalPortConfigInfo struct {
	DesiredSpeed    string          `json:"desiredSpeed,omitempty"`    // "desiredSpeed": "Auto",
	LogicalLocation LogicalLocation `json:"logicalLocation,omitempty"` // "logicalLocation": {...},
}

type LIGIOBayList []LIGIOBay

type LIGIOBay struct {
	Enclosure   int
	Bay         int
	ModelName   string
	ModelNumber string
}

type LIGUplinkSet struct {
	EthernetNetworkType    string                  `json:"ethernetNetworkType,omitempty"` // "ethernetNetworkType": "Tagged",
	LacpTimer              string                  `json:"lacpTimer,omitempty"`           // "lacpTimer": "Long",
	LogicalPortConfigInfos []LogicalPortConfigInfo `json:"logicalPortConfigInfos"`        // "logicalPortConfigInfos": {...},
	Mode                   string                  `json:"mode,omitempty"`                // "mode": "Auto",
	Name                   string                  `json:"name,omitempty"`                // "name": "Uplink 1",
	NativeNetworkUri       string                  `json:"nativeNetworkUri,omitempty"`    // "nativeNetworkUri": null,
	NetworkType            string                  `json:"networkType,omitempty"`         // "networkType": "Ethernet",
	NetworkUris            []string                `json:"networkUris"`                   // "networkUris": ["/rest/ethernet-networks/f1e38895-721b-4204-8395-ae0caba5e163"]
	PrimaryPort            *LogicalLocation        `json:"primaryPort,omitempty"`         // "primaryPort": {...},
	Reachability           string                  `json:"reachability,omitempty"`        // "reachability": "Reachable",
	UplinkPorts            LIGUplinkPortList       //define named type to use multisort method later
	Networks               []NetworkSummary        //collect network name and vlanid from NetworkURI list
}

type LIGUplinkPortList []LIGUplinkPort

type LIGUplinkPort struct {
	Enclosure int
	Bay       int
	Port      string
}

type NetworkSummary struct {
	Name   string
	Vlanid int
	Type   string
}

func (c *CLIOVClient) GetLIG() []LIG {

	var wg sync.WaitGroup

	rl := []string{"LIG"}

	for _, v := range rl {
		localv := v
		wg.Add(1)

		go func() {
			defer wg.Done()
			c.GetResourceLists(localv, "")
		}()
	}

	wg.Wait()

	l := *(rmap["LIG"].listptr.(*[]LIG))

	sort.Slice(l, func(i, j int) bool { return l[i].Name < l[j].Name })

	return l

}

func (c *CLIOVClient) GetLIGVerbose(name string) []LIG {

	var wg sync.WaitGroup

	rl := []string{"LIG", "ENetwork", "ICType"}

	for _, v := range rl {
		localv := v
		wg.Add(1)

		go func() {
			defer wg.Done()
			c.GetResourceLists(localv, "")
		}()
	}

	wg.Wait()

	l := *(rmap["LIG"].listptr.(*[]LIG))
	netList := *(rmap["ENetwork"].listptr.(*[]ENetwork))
	ictypeList := *(rmap["ICType"].listptr.(*[]ICType))

	log.Printf("[DEBUG] liglist length: %d\n", len(l))
	log.Printf("[DEBUG] netlist length: %d\n", len(netList))
	log.Printf("[DEBUG] ictypelist length: %d\n", len(ictypeList))

	if err := validateName(&l, name); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for i := range l {

		lig := &l[i]
		lig.getIOBay(ictypeList)
		lig.getUplinkPort(ictypeList)
		lig.getNetwork(netList)
	}

	sort.Slice(l, func(i, j int) bool { return l[i].Name < l[j].Name })

	return l

}

func (lig *LIG) getIOBay(ictypeList []ICType) {

	lig.IOBays = make([]LIGIOBay, 0)

	for _, v := range lig.InterconnectMapTemplate.InterconnectMapEntryTemplates {

		var e, s int

		for _, v := range v.LogicalLocation.LocationEntries {
			switch v.Type {
			case "Enclosure":
				e = v.RelativeValue
			case "Bay":
				s = v.RelativeValue
			}
		}

		//convert ICType list to ICType URI mapping to prepare lookup later
		ictypeMap := make(map[string]ICType)
		for _, v := range ictypeList {
			ictypeMap[v.URI] = v
		}

		n := ictypeMap[v.PermittedInterconnectTypeUri].Name
		m := ictypeMap[v.PermittedInterconnectTypeUri].PartNumber

		lig.IOBays = append(lig.IOBays, LIGIOBay{e, s, n, m})

	}

	sort.Slice(lig.IOBays, func(i, j int) bool { return lig.IOBays.multiSort(i, j) })

}

func (lig *LIG) getUplinkPort(ictypeList []ICType) {

	//prepare enc/bay lookup map to find out model number, 1st step loopup to convert port from "83" to "Q4:1"
	slotModel := make(map[struct{ enc, slot int }]string)
	for _, v := range lig.IOBays {
		slotModel[struct{ enc, slot int }{v.Enclosure, v.Bay}] = v.ModelNumber
	}

	//prepare modelnumber/portnumber lookup map to find out portname, 1st step loopup to convert port from "83" to "Q4:1"
	type ModelPort struct {
		model string
		port  int
	}
	modelPort := make(map[ModelPort]string)
	for _, t := range ictypeList {
		for _, p := range t.PortInfos {
			modelPort[ModelPort{t.PartNumber, p.PortNumber}] = p.PortName
		}
	}

	//get all uplinkport list for all uplinksets, like []{UplinkPort{1,2,67},{2,3,72}}
	for i, v := range lig.UplinkSets {

		lig.UplinkSets[i].UplinkPorts = make(LIGUplinkPortList, 0)
		uplinkports := lig.UplinkSets[i].UplinkPorts

		for _, v := range v.LogicalPortConfigInfos {

			var e, b, p int

			for _, v := range v.LogicalLocation.LocationEntries {
				switch v.Type {
				case "Enclosure":
					e = v.RelativeValue
				case "Bay":
					b = v.RelativeValue
				case "Port":
					p = v.RelativeValue
				}
			}

			//use above 2-step map lookups to convert final port number from "67" to "Q3:1"
			model := slotModel[struct{ enc, slot int }{e, b}]
			port := modelPort[ModelPort{model, p}]

			//update lig uplinkset uplink port list
			uplinkports = append(uplinkports, LIGUplinkPort{e, b, port})
			lig.UplinkSets[i].UplinkPorts = uplinkports

		}

		//use x,y to avoice conflict with existing i.
		sort.Slice(uplinkports, func(x, y int) bool { return uplinkports.multiSort(x, y) })

	}
}

func (lig *LIG) getNetwork(networkList []ENetwork) {

	networkMap := make(map[string]ENetwork)
	for _, v := range networkList {
		networkMap[v.URI] = v
	}

	for i, v := range lig.UplinkSets {
		lig.UplinkSets[i].Networks = make([]NetworkSummary, 0)
		networklist := lig.UplinkSets[i].Networks

		for _, v := range v.NetworkUris {
			vlanname := networkMap[v].Name
			vlanid := networkMap[v].VlanId
			vlantype := networkMap[v].EthernetNetworkType
			//lig.UplinkSets[i].Networks = append(lig.UplinkSets[i].Networks, NetworkSummary{vlanname, vlanid})
			networklist = append(networklist, NetworkSummary{vlanname, vlanid, vlantype})
		}

		sort.Slice(networklist, func(i, j int) bool { return networklist[i].Name < networklist[j].Name })
		lig.UplinkSets[i].Networks = networklist
		//sort.Slice(lig.UplinkSets[i].Networks, func(x, y int) bool { return lig.UplinkSets[i].Networks[x].Name < lig.UplinkSets[i].Networks[y].Name })

	}
}

func (x LIGUplinkPortList) multiSort(i, j int) bool {
	switch {
	case x[i].Enclosure < x[j].Enclosure:
		return true
	case x[i].Enclosure > x[j].Enclosure:
		return false
	case x[i].Bay < x[j].Bay:
		return true
	case x[i].Bay > x[j].Bay:
		return false
	case x[i].Port < x[j].Port:
		return true
	}
	return false
}

func (x LIGIOBayList) multiSort(i, j int) bool {
	switch {
	case x[i].Enclosure < x[j].Enclosure:
		return true
	case x[i].Enclosure > x[j].Enclosure:
		return false
	case x[i].Bay < x[j].Bay:
		return true
	case x[i].Bay > x[j].Bay:
		return false
	}
	return false
}

func CreateLIGConfigParse(fileName string) {

	if fileName == "" {
		fmt.Println(`Please specify config YAML filename by using "-f". 
	
	A sample config file format is as below:
	→ cat config.yml
	servertemplates:
		- name: hj-sptemplate1
		enclosuregroup: DCA-SolCenter-EG
		serverhardwaretype: "SY 480 Gen9 3"
		connections:
			- id: 1
			name: nic1
			network: TE-Testing-300
	
		- name: hj-sptemplate2
		enclosuregroup: DCA-SolCenter-EG
		serverhardwaretype: "SY 480 Gen9 1"
	
	
	networks:
		- name: hj-test1
		vlanId: 671
		- name: hj-test2
		vlanId: 672
	`)
		os.Exit(1)
	}

	y := YAMLConfig{}

	yamlFile, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := yaml.Unmarshal(yamlFile, &y); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	c := NewCLIOVClient()
	c.GetResourceLists("ICType", "")
	ictypeList := *(rmap["ICType"].listptr.(*[]ICType))

	ictypeMap := make(map[string]ICType)
	for _, v := range ictypeList {
		ictypeMap[v.PartNumber] = v
	}

	icTypeTable := map[string]string{"VC40F8": "794502-B23", "ILM20": "779218-B21"}

	for _, v := range y.LIGs {

		//fmt.Println(v.Name, v.FrameCount, v.InterConnectSet, v.Interconnects[0].Frame, v.Interconnects[0].Bay, v.Interconnects[0].Interconnect)

		var lig LIG
		lig.Name = v.Name
		lig.UplinkSets = []LIGUplinkSet{}
		lig.Type = "logical-interconnect-groupV300"
		lig.EnclosureType = "SY12000"
		lig.RedundancyType = "HighlyAvailable"
		lig.InterconnectBaySet = v.InterConnectSet
		lig.EnclosureIndexes = make([]int, v.FrameCount)

		for i := 0; i < v.FrameCount; i++ {
			lig.EnclosureIndexes[i] = i + 1
		}

		lig.InterconnectMapTemplate.InterconnectMapEntryTemplates = make([]InterconnectMapEntryTemplate, len(v.Interconnects))

		for i, v := range v.Interconnects {

			lf := LocationEntry{
				Type:          "Enclosure",
				RelativeValue: v.Frame,
			}
			lb := LocationEntry{
				Type:          "Bay",
				RelativeValue: v.Bay,
			}

			//lEntries := make([]LocationEntry, 2)
			lEntries := []LocationEntry{lf, lb}
			lLocation := LogicalLocation{LocationEntries: lEntries}

			ictypePN, ok := icTypeTable[v.Interconnect]
			if !ok {
				fmt.Printf("can't find matching IC type part number for config %q\n", v.Interconnect)
				os.Exit(1)
			}

			ictype, ok := ictypeMap[ictypePN]
			if !ok {
				fmt.Printf("can't find matching IC part numbber %q from server response\n", v.Interconnect)
				os.Exit(1)
			}

			lig.InterconnectMapTemplate.InterconnectMapEntryTemplates[i] = InterconnectMapEntryTemplate{
				EnclosureIndex:               v.Frame,
				LogicalLocation:              lLocation,
				PermittedInterconnectTypeUri: ictype.URI,
			}

			//fmt.Printf("%#v\n", lig.InterconnectMapTemplate.InterconnectMapEntryTemplates[i])

		}

		// b, err := json.MarshalIndent(lig, "", "  ")
		// if err != nil {
		// 	fmt.Println(err)
		// }
		//fmt.Printf("%s\n", b)

		fmt.Printf("Creating Logical Interconnect Group: %q\n", v.Name)

		if _, err := c.SendHTTPRequest("POST", LIGURL, "", "", lig); err != nil {
			fmt.Println(err)
			os.Exit(1)

		}

	}
}

func DeleteLIG(name string) error {

	if name == "" {
		fmt.Println("Neet to specify LIG name using \"-n\" flag")
		os.Exit(1)
	}

	c := NewCLIOVClient()

	c.GetResourceLists("LIG", name)

	list := *(rmap["LIG"].listptr.(*[]LIG))

	if len(list) == 0 {
		fmt.Printf("Can't find LIG %v to delete", name)
		os.Exit(1)
	}

	for _, v := range list {
		fmt.Printf("Deleting LIG: %q\n", v.Name)
		_, err := c.SendHTTPRequest("DELETE", v.URI, "", "", nil)
		if err != nil {
			fmt.Printf("Error submitting delete server LIG request: %v", err)
		}
	}
	return nil
}

type EthernetSettings struct {
	Category                    string `json:"category,omitempty"`                    // "category": null,
	Created                     string `json:"created,omitempty"`                     // "created": "20150831T154835.250Z",
	DependentResourceUri        string `json:"dependentResourceUri,omitempty"`        // dependentResourceUri": "/rest/logical-interconnect-groups/b7b144e9-1f5e-4d52-8534-2e39280f9e86",
	Description                 string `json:"description,omitempty,omitempty"`       // "description": "Ethernet Settings",
	ETAG                        string `json:"eTag,omitempty"`                        // "eTag": "1441036118675/8",
	EnableFastMacCacheFailover  *bool  `json:"enableFastMacCacheFailover,omitempty"`  //"enableFastMacCacheFailover": false,
	EnableIgmpSnooping          *bool  `json:"enableIgmpSnooping,omitempty"`          // "enableIgmpSnooping": false,
	EnableNetworkLoopProtection *bool  `json:"enableNetworkLoopProtection,omitempty"` // "enableNetworkLoopProtection": false,
	EnablePauseFloodProtection  *bool  `json:"enablePauseFloodProtection,omitempty"`  // "enablePauseFloodProtection": false,
	EnableRichTLV               *bool  `json:"enableRichTLV,omitempty"`               // "enableRichTLV": false,
	ID                          string `json:"id,omitempty"`                          //"id": "0c398238-2d35-48eb-9eb5-7560d59f94b3",
	IgmpIdleTimeoutInterval     int    `json:"igmpIdleTimeoutInterval,omitempty"`     // "igmpIdleTimeoutInterval": 260,
	InterconnectType            string `json:"interconnectType,omitempty"`            // "interconnectType": "Ethernet",
	MacRefreshInterval          int    `json:"macRefreshInterval,omitempty"`          // "macRefreshInterval": 5,
	Modified                    string `json:"modified,omitempty"`                    // "modified": "20150831T154835.250Z",
	Name                        string `json:"name,omitempty"`                        // "name": "ethernetSettings 1",
	State                       string `json:"state,omitempty"`                       // "state": "Normal",
	Status                      string `json:"status,omitempty"`                      // "status": "Critical",
	Type                        string `json:"type,omitempty"`                        // "EthernetInterconnectSettingsV3",
	URI                         string `json:"uri,omitempty"`                         // "uri": "/rest/logical-interconnect-groups/b7b144e9-1f5e-4d52-8534-2e39280f9e86/ethernetSettings"
}

type QosConfiguration struct {
	ActiveQosConfig          ActiveQosConfig           `json:"activeQosConfig,omitempty"`          //"activeQosConfig": {...},
	Category                 string                    `json:"category,omitempty"`                 // "category": "qos-aggregated-configuration",
	Created                  string                    `json:"created,omitempty"`                  // "created": "20150831T154835.250Z",
	Description              string                    `json:"description,omitempty,omitempty"`    // "description": null,
	ETAG                     string                    `json:"eTag,omitempty"`                     // "eTag": "1441036118675/8",
	InactiveFCoEQosConfig    *InactiveFCoEQosConfig    `json:"inactiveFCoEQosConfig,omitempty"`    // "inactiveFCoEQosConfig": {...},
	InactiveNonFCoEQosConfig *InactiveNonFCoEQosConfig `json:"inactiveNonFCoEQosConfig,omitempty"` // "inactiveNonFCoEQosConfig": {...},
	Modified                 string                    `json:"modified,omitempty"`                 // "modified": "20150831T154835.250Z",
	Name                     string                    `json:"name,omitempty"`                     // "name": "Qos Config 1",
	State                    string                    `json:"state,omitempty"`                    // "state": "Normal",
	Status                   string                    `json:"status,omitempty"`                   // "status": "Critical",
	Type                     string                    `json:"type,omitempty"`                     // "qos-aggregated-configuration",
	URI                      string                    `json:"uri,omitempty"`                      // "uri": null
}

type ActiveQosConfig struct {
	Category                   string                 `json:"category,omitempty"`                   // "category": "null",
	ConfigType                 string                 `json:"configType,omitempty"`                 // "configType": "CustomWithFCoE",
	Created                    string                 `json:"created,omitempty"`                    // "created": "20150831T154835.250Z",
	Description                string                 `json:"description,omitempty,omitempty"`      // "description": "Ethernet Settings",
	DownlinkClassificationType string                 `json:"downlinkClassificationType,omitempty"` //"downlinkClassifcationType": "DOT1P_AND_DSCP",
	ETAG                       string                 `json:"eTag,omitempty"`                       // "eTag": "1441036118675/8",
	Modified                   string                 `json:"modified,omitempty"`                   // "modified": "20150831T154835.250Z",
	Name                       string                 `json:"name,omitempty"`                       // "name": "active QOS Config 1",
	QosTrafficClassifiers      []QosTrafficClassifier `json:"qosTrafficClassifiers"`                // "qosTrafficClassifiers": {...},
	State                      string                 `json:"state,omitempty"`                      // "state": "Normal",
	Status                     string                 `json:"status,omitempty"`                     // "status": "Critical",
	Type                       string                 `json:"type,omitempty"`                       // "type": "QosConfiguration",
	UplinkClassificationType   string                 `json:"uplinkClassificationType,omitempty"`   // "uplinkClassificationType": "DOT1P"
	URI                        string                 `json:"uri,omitempty"`                        // "uri": null
}

type InactiveFCoEQosConfig struct {
	Category                   string                 `json:"category,omitempty"`                   // "category": "null",
	ConfigType                 string                 `json:"configType,omitempty"`                 // "configType": "CustomWithFCoE",
	Created                    string                 `json:"created,omitempty"`                    // "created": "20150831T154835.250Z",
	Description                string                 `json:"description,omitempty,omitempty"`      // "description": "Ethernet Settings",
	DownlinkClassificationType string                 `json:"downlinkClassificationType,omitempty"` //"downlinkClassifcationType": "DOT1P_AND_DSCP",
	ETAG                       string                 `json:"eTag,omitempty"`                       // "eTag": "1441036118675/8",
	Modified                   string                 `json:"modified,omitempty"`                   // "modified": "20150831T154835.250Z",
	Name                       string                 `json:"name,omitempty"`                       // "name": "active QOS Config 1",
	QosTrafficClassifiers      []QosTrafficClassifier `json:"qosTrafficClassifiers,omitempty"`      // "qosTrafficClassifiers": {...},
	State                      string                 `json:"state,omitempty"`                      // "state": "Normal",
	Status                     string                 `json:"status,omitempty"`                     // "status": "Critical",
	Type                       string                 `json:"type,omitempty"`                       // "type": "QosConfiguration",
	UplinkClassificationType   string                 `json:"uplinkClassificationType,omitempty"`   // "uplinkClassificationType": "DOT1P"
	URI                        string                 `json:"uri,omitempty"`                        // "uri": null
}

type InactiveNonFCoEQosConfig struct {
	Category                   string                 `json:"category,omitempty"`                   // "category": "null",
	ConfigType                 string                 `json:"configType,omitempty"`                 // "configType": "CustomWithFCoE",
	Created                    string                 `json:"created,omitempty"`                    // "created": "20150831T154835.250Z",
	Description                string                 `json:"description,omitempty,omitempty"`      // "description": "Ethernet Settings",
	DownlinkClassificationType string                 `json:"downlinkClassificationType,omitempty"` //"downlinkClassifcationType": "DOT1P_AND_DSCP",
	ETAG                       string                 `json:"eTag,omitempty"`                       // "eTag": "1441036118675/8",
	Modified                   string                 `json:"modified,omitempty"`                   // "modified": "20150831T154835.250Z",
	Name                       string                 `json:"name,omitempty"`                       // "name": "active QOS Config 1",
	QosTrafficClassifiers      []QosTrafficClassifier `json:"qosTrafficClassifiers,omitempty"`      // "qosTrafficClassifiers": {...},
	State                      string                 `json:"state,omitempty"`                      // "state": "Normal",
	Status                     string                 `json:"status,omitempty"`                     // "status": "Critical",
	Type                       string                 `json:"type,omitempty"`                       // "type": "QosConfiguration",
	UplinkClassificationType   string                 `json:"uplinkClassificationType,omitempty"`   // "uplinkClassificationType": "DOT1P"
	URI                        string                 `json:"uri,omitempty"`                        // "uri": null
}

type QosTrafficClassifier struct {
	QosClassificationMapping *QosClassificationMap `json:"qosClassificationMapping"`  // "qosClassificationMapping": {...},
	QosTrafficClass          QosTrafficClass       `json:"qosTrafficClass,omitempty"` // "qosTrafficClass": {...},
}

type QosClassificationMap struct {
	Dot1pClassMapping []int    `json:"dot1pClassMapping"` // "dot1pClassMapping": [3],
	DscpClassMapping  []string `json:"dscpClassMapping"`  // "dscpClassMapping": [],
}

type QosTrafficClass struct {
	BandwidthShare   string `json:"bandwidthShare,omitempty"` // "bandwidthShare": "fcoe",
	ClassName        string `json:"className"`                // "className": "FCoE lossless",
	EgressDot1pValue int    `json:"egressDot1pValue"`         // "egressDot1pValue": 3,
	Enabled          *bool  `json:"enabled,omitempty"`        // "enabled": true,
	MaxBandwidth     int    `json:"maxBandwidth"`             // "maxBandwidth": 100,
	RealTime         *bool  `json:"realTime,omitempty"`       // "realTime": true,
}

//TODO SNMPConfiguration
type SnmpConfiguration struct {
	Category         string            `json:"category,omitempty"`         // "category": "snmp-configuration",
	Created          string            `json:"created,omitempty"`          // "created": "20150831T154835.250Z",
	Description      string            `json:"description,omitempty"`      // "description": null,
	ETAG             string            `json:"eTag,omitempty"`             // "eTag": "1441036118675/8",
	Enabled          *bool             `json:"enabled,omitempty"`          // "enabled": true,
	Modified         string            `json:"modified,omitempty"`         // "modified": "20150831T154835.250Z",
	Name             string            `json:"name,omitempty"`             // "name": "Snmp Config",
	ReadCommunity    string            `json:"readCommunity,omitempty"`    // "readCommunity": "public",
	SnmpAccess       []string          `json:"snmpAccess,omitempty"`       // "snmpAccess": [],
	State            string            `json:"state,omitempty"`            // "state": "Normal",
	Status           string            `json:"status,omitempty"`           // "status": "Critical",
	SystemContact    string            `json:"systemContact,omitempty"`    // "systemContact": "",
	TrapDestinations []TrapDestination `json:"trapDestinations,omitempty"` // "trapDestinations": {...}
	Type             string            `json:"type,omitempty"`             // "type": "snmp-configuration",
	URI              string            `json:"uri,omitempty"`              // "uri": null
}

type TrapDestination struct {
	CommunityString    string   `json:"communityString,omitempty"`    //"communityString": "public",
	EnetTrapCategories []string `json:"enetTrapCategories,omitempty"` //"enetTrapCategories": ["PortStatus", "Other"],
	FcTrapCategories   []string `json:"fcTrapCategories,omitempty"`   //"fcTrapCategories": ["PortStatus", "Other"]
	TrapDestination    string   `json:"trapDestination,omitempty"`    //"trapDestination": "127.0.0.1",
	TrapFormat         string   `json:"trapFormat,omitempty"`         //"trapFormat", "SNMPv1",
	TrapSeverities     []string `json:"trapSeverities,omitempty"`     //"trapSeverities": "Info",
	VcmTrapCategories  []string `json:"vcmTrapCategories,omitempty"`  // "vcmTrapCategories": ["Legacy"],
}

type TelemetryConfiguration struct {
	Category        string `json:"category,omitempty"`        // "category": "telemetry-configuration",
	Created         string `json:"created,omitempty"`         // "created": "20150831T154835.250Z",
	Description     string `json:"description,omitempty"`     // "description": null,
	ETAG            string `json:"eTag,omitempty"`            // "eTag": "1441036118675/8",
	EnableTelemetry *bool  `json:"enableTelemetry,omitempty"` // "enableTelemetry": false,
	Modified        string `json:"modified,omitempty"`        // "modified": "20150831T154835.250Z",
	Name            string `json:"name,omitempty"`            // "name": "telemetry configuration",
	SampleCount     int    `json:"sampleCount,omitempty"`     // "sampleCount": 12
	SampleInterval  int    `json:"sampleInterval,omitempty"`  // "sampleInterval": 300,
	State           string `json:"state,omitempty"`           // "state": "Normal",
	Status          string `json:"status,omitempty"`          // "status": "Critical",
	Type            string `json:"type,omitempty"`            // "type": "telemetry-configuration",
	URI             string `json:"uri,omitempty"`             // "uri": null
}

// type LIG struct {
// 	Type          string      `json:"type"`
// 	EnclosureType string      `json:"enclosureType"`
// 	UplinkSets    []LIGUplinkSet `json:"uplinkSets"`
// QosConfiguration struct {
// 	Type            string `json:"type"`
// 	ActiveQosConfig struct {
// 		Type                       string        `json:"type"`
// 		ConfigType                 string        `json:"configType"`
// 		UplinkClassificationType   interface{}   `json:"uplinkClassificationType"`
// 		DownlinkClassificationType interface{}   `json:"downlinkClassificationType"`
// 		QosTrafficClassifiers      []interface{} `json:"qosTrafficClassifiers"`
// 		Description                interface{}   `json:"description"`
// 		Name                       interface{}   `json:"name"`
// 		State                      interface{}   `json:"state"`
// 		Status                     interface{}   `json:"status"`
// 		Category                   string        `json:"category"`
// 		ETag                       interface{}   `json:"eTag"`
// 		Created                    interface{}   `json:"created"`
// 		Modified                   interface{}   `json:"modified"`
// 		URI                        interface{}   `json:"uri"`
// 	} `json:"activeQosConfig"`
// 	InactiveFCoEQosConfig    interface{} `json:"inactiveFCoEQosConfig"`
// 	InactiveNonFCoEQosConfig interface{} `json:"inactiveNonFCoEQosConfig"`
// 	Description              interface{} `json:"description"`
// 	Name                     interface{} `json:"name"`
// 	State                    interface{} `json:"state"`
// 	Status                   interface{} `json:"status"`
// 	Category                 string      `json:"category"`
// 	ETag                     interface{} `json:"eTag"`
// 	Created                  time.Time   `json:"created"`
// 	Modified                 time.Time   `json:"modified"`
// 	URI                      interface{} `json:"uri"`
// } `json:"qosConfiguration"`
// SnmpConfiguration struct {
// 	Type             string        `json:"type"`
// 	ReadCommunity    string        `json:"readCommunity"`
// 	SystemContact    string        `json:"systemContact"`
// 	TrapDestinations []interface{} `json:"trapDestinations"`
// 	SnmpAccess       []interface{} `json:"snmpAccess"`
// 	Enabled          bool          `json:"enabled"`
// 	Description      interface{}   `json:"description"`
// 	Name             interface{}   `json:"name"`
// 	State            interface{}   `json:"state"`
// 	Status           interface{}   `json:"status"`
// 	Category         string        `json:"category"`
// 	ETag             interface{}   `json:"eTag"`
// 	Created          time.Time     `json:"created"`
// 	Modified         time.Time     `json:"modified"`
// 	URI              interface{}   `json:"uri"`
// } `json:"snmpConfiguration"`
// TelemetryConfiguration struct {
// 	Type            string      `json:"type"`
// 	EnableTelemetry bool        `json:"enableTelemetry"`
// 	SampleCount     int         `json:"sampleCount"`
// 	SampleInterval  int         `json:"sampleInterval"`
// 	Description     interface{} `json:"description"`
// 	Name            interface{} `json:"name"`
// 	State           interface{} `json:"state"`
// 	Status          interface{} `json:"status"`
// 	Category        string      `json:"category"`
// 	ETag            interface{} `json:"eTag"`
// 	Created         time.Time   `json:"created"`
// 	Modified        time.Time   `json:"modified"`
// 	URI             interface{} `json:"uri"`
// } `json:"telemetryConfiguration"`
// StackingHealth          string `json:"stackingHealth"`
// InterconnectMapTemplate InterconnectMapTemplate `json:"interconnectMapTemplate"`
// FabricURI          string `json:"fabricUri"`
// InterconnectBaySet int    `json:"interconnectBaySet"`
// RedundancyType     string `json:"redundancyType"`
// EthernetSettings   struct {
// 	Type                        string      `json:"type"`
// 	LldpIpv4Address             string      `json:"lldpIpv4Address"`
// 	LldpIpv6Address             string      `json:"lldpIpv6Address"`
// 	EnableIgmpSnooping          bool        `json:"enableIgmpSnooping"`
// 	IgmpIdleTimeoutInterval     int         `json:"igmpIdleTimeoutInterval"`
// 	EnableFastMacCacheFailover  bool        `json:"enableFastMacCacheFailover"`
// 	MacRefreshInterval          int         `json:"macRefreshInterval"`
// 	EnableNetworkLoopProtection bool        `json:"enableNetworkLoopProtection"`
// 	EnablePauseFloodProtection  bool        `json:"enablePauseFloodProtection"`
// 	EnableRichTLV               bool        `json:"enableRichTLV"`
// 	EnableTaggedLldp            bool        `json:"enableTaggedLldp"`
// 	InterconnectType            string      `json:"interconnectType"`
// 	DependentResourceURI        string      `json:"dependentResourceUri"`
// 	Name                        string      `json:"name"`
// 	ID                          string      `json:"id"`
// 	Description                 interface{} `json:"description"`
// 	State                       interface{} `json:"state"`
// 	Status                      interface{} `json:"status"`
// 	Category                    interface{} `json:"category"`
// 	ETag                        interface{} `json:"eTag"`
// 	Created                     time.Time   `json:"created"`
// 	Modified                    time.Time   `json:"modified"`
// 	URI                         string      `json:"uri"`
// } `json:"ethernetSettings"`
// 	EnclosureIndexes    []int        `json:"enclosureIndexes"`
// 	InternalNetworkUris []string     `json:"internalNetworkUris"`
// 	StackingMode        string       `json:"stackingMode"`
// 	ScopeUris           []string     `json:"scopeUris"`
// 	Description         string       `json:"description"`
// 	Name                string       `json:"name"`
// 	State               string       `json:"state"`
// 	Status              string       `json:"status"`
// 	Category            string       `json:"category"`
// 	ETag                string       `json:"eTag"`
// 	Created             string       `json:"created"`
// 	Modified            string       `json:"modified"`
// 	URI                 string       `json:"uri"`
// 	IOBays              LIGIOBayList //define []IOBay as named type to use multisort method later
// }

// type InterconnectMapTemplate struct{
// 	InterconnectMapEntryTemplates []InterconnectMapEntryTemplate `json:"interconnectMapEntryTemplates"`
// }

// type InterconnectMapEntryTemplate struct {
// 	EnclosureIndex               int    `json:"enclosureIndex"`
// 	PermittedInterconnectTypeURI string `json:"permittedInterconnectTypeUri"`
// 	LogicalLocation              struct {
// 		LocationEntries []struct {
// 			RelativeValue int    `json:"relativeValue"`
// 			Type          string `json:"type"`
// 		} `json:"locationEntries"`
// 	} `json:"logicalLocation"`
// 	LogicalDownlinkURI string `json:"logicalDownlinkUri"`
// }
