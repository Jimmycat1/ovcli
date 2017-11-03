package oneview

//SFPMap is from conversion of raw SFPList(a slice) to mapping struct with port names as keys and the pointers of SFP structs as values. Each module has its own "sfpMap" to pass to channel
type SFPMap map[string]*SFP

//ICSFPStruct can give us information of Modulename, which we can't get from simple map key for IC module
type ICSFPStruct struct {
	ModuleName string
	SFPMapping *SFPMap
}

//ICSFPMap is mapping between each module and its own port mapping table, such as map["module 1, top frame"]*map[d1]struct{for d1}
//type ICSFPMap map[string]*SFPMap
//create extract struct inside map to give us information on Module Name
type ICSFPMap map[string]*ICSFPStruct

//UplinkSetMap is mapping between each uplinkset name/URI and its struct
type UplinkSetMap map[string]*UplinkSet

//LIUplinkSetMap is mapping between each LI and its own Uplinkset maps
type LIUplinkSetMap map[string]UplinkSetMap

// var ovAddress = os.Getenv("OneView_address")
// var ovUsername = os.Getenv("OneView_username")
// var ovPassword = os.Getenv("OneView_password")

//OVClient is the sole OV client for all CLI commands
//var OVClient = NewCLIOVClient()
var taskuri string

const (
	LIGURL          = "/rest/logical-interconnect-groups"
	LIURL           = "/rest/logical-interconnects"
	UplinkSetURL    = "/rest/uplink-sets"
	ICURL           = "/rest/interconnects"
	SFPURL          = "/pluggableModuleInformation"
	ICTypeURL       = "/rest/interconnect-types"
	ENetworkURL     = "/rest/ethernet-networks"
	NetSetURL       = "/rest/network-sets"
	EnclosureURL    = "/rest/enclosures"
	EGURL           = "/rest/enclosure-groups"
	LEURL           = "/rest/logical-enclosures"
	SPURL           = "/rest/server-profiles"
	SPTemplateURL   = "/rest/server-profile-templates"
	ServerHWURL     = "/rest/server-hardware"
	ServerHWTypeURL = "/rest/server-hardware-types"
	VersionURL      = "/rest/version"

	DefaultConfigFile = "appliance-credential.yml"
)

type resource struct {
	listptr interface{}
	colptr  interface{}
	uri     string
	logmsg  string
}

type resourceMap map[string]resource

var rmap = resourceMap{
	"SPTemplate": resource{
		listptr: &[]SPTemplate{},
		colptr:  &SPTemplateCol{},
		uri:     SPTemplateURL,
		logmsg:  "get SPTemplate",
	},

	"SP": resource{
		listptr: &SPList{},
		colptr:  &SPCol{},
		uri:     SPURL,
		logmsg:  "get Server Profile",
	},

	"LE": resource{
		listptr: &[]LE{},
		colptr:  &LECol{},
		uri:     LEURL,
		logmsg:  "get LE",
	},

	"EG": resource{
		listptr: &[]EG{},
		colptr:  &EGCol{},
		uri:     EGURL,
		logmsg:  "get EG",
	},

	"ServerHW": resource{
		listptr: &[]ServerHW{},
		colptr:  &ServerHWCol{},
		uri:     ServerHWURL,
		logmsg:  "get Server Hardware ",
	},

	"ServerHWType": resource{
		listptr: &[]ServerHWType{},
		colptr:  &ServerHWTypeCol{},
		uri:     ServerHWTypeURL,
		logmsg:  "get ServerHW Type",
	},

	"IC": resource{
		listptr: &[]IC{},
		colptr:  &ICCol{},
		uri:     ICURL,
		logmsg:  "get IC",
	},

	"ICType": resource{
		listptr: &[]ICType{},
		colptr:  &ICTypeCol{},
		uri:     ICTypeURL,
		logmsg:  "get ICType",
	},

	"LI": resource{
		listptr: &[]LI{},
		colptr:  &LICol{},
		uri:     LIURL,
		logmsg:  "get LI",
	},

	"LIG": resource{
		listptr: &[]LIG{},
		colptr:  &LIGCol{},
		uri:     LIGURL,
		logmsg:  "get LIG",
	},

	"ENetwork": resource{
		listptr: &[]ENetwork{},
		colptr:  &ENetworkCol{},
		uri:     ENetworkURL,
		logmsg:  "get Ethernet Network",
	},

	"NetSet": resource{
		listptr: &[]NetSet{},
		colptr:  &NetSetCol{},
		uri:     NetSetURL,
		logmsg:  "get Network Set",
	},

	"UplinkSet": resource{
		listptr: &UplinkSetList{},
		colptr:  &UplinkSetCol{},
		uri:     UplinkSetURL,
		logmsg:  "get UplinkSet",
	},

	"Enclosure": resource{
		listptr: &[]Enclosure{},
		colptr:  &EncCol{},
		uri:     EnclosureURL,
		logmsg:  "get Enclosure",
	},
}

type cred struct {
	Ip   string `json:"ipaddress"`
	User string `json:"username"`
	Pass string `json:"password"`
}

type OVCol interface {
	GetMap(c *CLIOVClient)
}

func ColToMap(x OVCol, c *CLIOVClient) {
	x.GetMap(c)
}

type YAMLConfig struct {
	ENetworks   []ENetwork       `json:"networks"`
	SPTemplates []YAMLSPTemplate `json:"servertemplates"`
	LIGs        []YAMLLIG        `json:"ligs"`
	EGs         []YAMLEG         `json:"egs"`
	LEs         []YAMLLE         `json:"les"`
}

type YAMLLE struct {
	Name       string
	Enclosures []string
	EG         string
}

type YAMLEG struct {
	Name       string
	FrameCount int `json:"framecount"`
	Frames     []YAMLFrame
}

type YAMLFrame struct {
	ID   int
	LIGs []string `json:"ligs"`
}

type YAMLUplinkSet struct {
	Name        string
	Type        string
	Networks    []string
	UplinkPorts []string `json:"uplinkports"`
}

type YAMLSPTemplate struct {
	Name            string           `json:"name,omitempty"`
	EG              string           `json:"enclosuregroup,omitempty"`
	ServerHWType    string           `json:"serverhardwaretype,omitempty"`
	YAMLConnections []YAMLConnection `json:"connections,omiempty"`
}

type YAMLConnection struct {
	ID      int    `json:"id,omiempty"`
	Name    string `json:"name,omiempty"`
	Network string `json:"network,omiempty"`
}

type YAMLLIG struct {
	Name            string `json:"name"`
	FrameCount      int    `json:"framecount`
	InterConnectSet int    `json:"interconnectset`
	Interconnects   []YAMLInterconnect
	UplinkSets      []YAMLUplinkSet `json:"uplinksets"`
}

type YAMLInterconnect struct {
	Frame        int    `json:"frame"`
	Bay          int    `json:"bay"`
	Interconnect string `json:"interconnect"`
}

// ligs:
// - name: test2
//   framecount: 2
//   interconnectset: 3
//   interconnects:
// 	- { frame: 1, bay: 3, interconnect: VC40F8 }
// 	# - { frame: 1, bay: 6, interconnect: ILM
// 	# - frame: 1, bay: 3, interconnect: ILM
// 	# - frame: 1, bay: 6, interconnect: VC40F8
