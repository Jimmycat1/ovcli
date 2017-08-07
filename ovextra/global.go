package ovextra

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
var OVClient = NewCLIOVClient()
var taskuri string

const (
	LIGURL          = "/rest/logical-interconnect-groups"
	LIURL           = "/rest/logical-interconnects"
	UplinkSetURL    = "/rest/uplink-sets"
	ICURL           = "/rest/interconnects"
	SFPURL          = "/rest/interconnects/pluggableModuleInformation/"
	ICTypeURL       = "/rest/interconnect-types"
	ENetworkURL     = "/rest/ethernet-networks"
	NetSetURL       = "/rest/network-sets"
	EnclosureURL    = "/rest/enclosures"
	EGURL           = "/rest/enclosure-groups"
	SPURL           = "/rest/server-profiles"
	SPTemplateURL   = "/rest/server-profile-templates"
	ServerHWURL     = "/rest/server-hardware"
	ServerHWTypeURL = "/rest/server-hardware-types"

	DefaultConfigFile = "appliance-credential.yaml"
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

	"EG": resource{
		listptr: &[]EG{},
		colptr:  &EGCol{},
		uri:     EGURL,
		logmsg:  "get EG",
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

	"LI": resource{
		listptr: &[]LI{},
		colptr:  &LICol{},
		uri:     LIURL,
		logmsg:  "get LI",
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

type YamlConfig struct {
	Networks        []ENetwork   //`yaml:"networks"`
	ServerTemplates []SPTemplate //`yaml:"serverTemplates"`
}
