package ovextra

import (
	"encoding/json"
	"log"
	"time"
)

type LICol struct {
	Type        string      `json:"type"`
	Members     []LI        `json:"members"`
	NextPageURI string      `json:"nextPageUri"`
	Start       int         `json:"start"`
	PrevPageURI interface{} `json:"prevPageUri"`
	Count       int         `json:"count"`
	Total       int         `json:"total"`
	Created     interface{} `json:"created"`
	ETag        interface{} `json:"eTag"`
	Modified    interface{} `json:"modified"`
	Category    string      `json:"category"`
	URI         string      `json:"uri"`
}

type LI struct {
	Type                        string `json:"type"`
	LogicalInterconnectGroupURI string `json:"logicalInterconnectGroupUri"`
	// SnmpConfiguration           struct {
	// 	Type             string        `json:"type"`
	// 	ReadCommunity    string        `json:"readCommunity"`
	// 	TrapDestinations []interface{} `json:"trapDestinations"`
	// 	SystemContact    string        `json:"systemContact"`
	// 	SnmpAccess       []interface{} `json:"snmpAccess"`
	// 	Enabled          bool          `json:"enabled"`
	// 	Description      interface{}   `json:"description"`
	// 	Status           interface{}   `json:"status"`
	// 	Name             interface{}   `json:"name"`
	// 	State            interface{}   `json:"state"`
	// 	Created          time.Time     `json:"created"`
	// 	ETag             interface{}   `json:"eTag"`
	// 	Modified         time.Time     `json:"modified"`
	// 	Category         string        `json:"category"`
	// 	URI              interface{}   `json:"uri"`
	// } `json:"snmpConfiguration"`
	// TelemetryConfiguration struct {
	// 	Type            string      `json:"type"`
	// 	SampleCount     int         `json:"sampleCount"`
	// 	SampleInterval  int         `json:"sampleInterval"`
	// 	EnableTelemetry bool        `json:"enableTelemetry"`
	// 	Description     interface{} `json:"description"`
	// 	Status          interface{} `json:"status"`
	// 	Name            string      `json:"name"`
	// 	State           interface{} `json:"state"`
	// 	Created         interface{} `json:"created"`
	// 	ETag            interface{} `json:"eTag"`
	// 	Modified        interface{} `json:"modified"`
	// 	Category        string      `json:"category"`
	// 	URI             string      `json:"uri"`
	// } `json:"telemetryConfiguration"`
	EnclosureUris  []string `json:"enclosureUris"`
	EnclosureType  string   `json:"enclosureType"`
	StackingHealth string   `json:"stackingHealth"`
	Interconnects  []string `json:"interconnects"`
	// QosConfiguration struct {
	// 	Type            string `json:"type"`
	// 	ActiveQosConfig struct {
	// 		Type                       string        `json:"type"`
	// 		ConfigType                 string        `json:"configType"`
	// 		DownlinkClassificationType interface{}   `json:"downlinkClassificationType"`
	// 		UplinkClassificationType   interface{}   `json:"uplinkClassificationType"`
	// 		QosTrafficClassifiers      []interface{} `json:"qosTrafficClassifiers"`
	// 		Description                interface{}   `json:"description"`
	// 		Status                     interface{}   `json:"status"`
	// 		Name                       interface{}   `json:"name"`
	// 		State                      interface{}   `json:"state"`
	// 		Created                    interface{}   `json:"created"`
	// 		ETag                       interface{}   `json:"eTag"`
	// 		Modified                   interface{}   `json:"modified"`
	// 		Category                   string        `json:"category"`
	// 		URI                        interface{}   `json:"uri"`
	// 	} `json:"activeQosConfig"`
	// 	InactiveFCoEQosConfig    interface{} `json:"inactiveFCoEQosConfig"`
	// 	InactiveNonFCoEQosConfig interface{} `json:"inactiveNonFCoEQosConfig"`
	// 	Description              interface{} `json:"description"`
	// 	Status                   interface{} `json:"status"`
	// 	Name                     interface{} `json:"name"`
	// 	State                    interface{} `json:"state"`
	// 	Created                  time.Time   `json:"created"`
	// 	ETag                     interface{} `json:"eTag"`
	// 	Modified                 time.Time   `json:"modified"`
	// 	Category                 string      `json:"category"`
	// 	URI                      interface{} `json:"uri"`
	// } `json:"qosConfiguration"`
	InternalNetworkUris []string `json:"internalNetworkUris"`
	ICMap               struct {
		InterconnectMapEntries []struct {
			InterconnectURI              string `json:"interconnectUri"`
			EnclosureIndex               int    `json:"enclosureIndex"`
			PermittedInterconnectTypeURI string `json:"permittedInterconnectTypeUri"`
			LogicalDownlinkURI           string `json:"logicalDownlinkUri"`
			Location                     struct {
				LocationEntries []struct {
					Value string `json:"value"`
					Type  string `json:"type"`
				} `json:"locationEntries"`
			} `json:"location"`
		} `json:"interconnectMapEntries"`
	} `json:"ICMap"`
	IcmLicenses struct {
		License []struct {
			RequiredCount int         `json:"requiredCount"`
			LicenseType   string      `json:"licenseType"`
			ConsumedCount int         `json:"consumedCount"`
			State         interface{} `json:"state"`
		} `json:"license"`
	} `json:"icmLicenses"`
	ConsistencyStatus string `json:"consistencyStatus"`
	EthernetSettings  struct {
		Type                        string      `json:"type"`
		InterconnectType            string      `json:"interconnectType"`
		LldpIpv4Address             string      `json:"lldpIpv4Address"`
		LldpIpv6Address             string      `json:"lldpIpv6Address"`
		EnableIgmpSnooping          bool        `json:"enableIgmpSnooping"`
		IgmpIdleTimeoutInterval     int         `json:"igmpIdleTimeoutInterval"`
		EnableFastMacCacheFailover  bool        `json:"enableFastMacCacheFailover"`
		MacRefreshInterval          int         `json:"macRefreshInterval"`
		EnableNetworkLoopProtection bool        `json:"enableNetworkLoopProtection"`
		EnablePauseFloodProtection  bool        `json:"enablePauseFloodProtection"`
		EnableRichTLV               bool        `json:"enableRichTLV"`
		EnableTaggedLldp            bool        `json:"enableTaggedLldp"`
		DependentResourceURI        string      `json:"dependentResourceUri"`
		Name                        string      `json:"name"`
		ID                          string      `json:"id"`
		Description                 interface{} `json:"description"`
		Status                      interface{} `json:"status"`
		State                       interface{} `json:"state"`
		Created                     time.Time   `json:"created"`
		ETag                        interface{} `json:"eTag"`
		Modified                    time.Time   `json:"modified"`
		Category                    interface{} `json:"category"`
		URI                         string      `json:"uri"`
	} `json:"ethernetSettings"`
	FabricURI   string `json:"fabricUri"`
	PortMonitor struct {
		Type              string        `json:"type"`
		AnalyzerPort      interface{}   `json:"analyzerPort"`
		MonitoredPorts    []interface{} `json:"monitoredPorts"`
		EnablePortMonitor bool          `json:"enablePortMonitor"`
		Description       interface{}   `json:"description"`
		Status            interface{}   `json:"status"`
		Name              string        `json:"name"`
		State             interface{}   `json:"state"`
		Created           interface{}   `json:"created"`
		ETag              string        `json:"eTag"`
		Modified          interface{}   `json:"modified"`
		Category          string        `json:"category"`
		URI               string        `json:"uri"`
	} `json:"portMonitor"`
	DomainURI   string        `json:"domainUri"`
	ScopeUris   []interface{} `json:"scopeUris"`
	Description interface{}   `json:"description"`
	Status      string        `json:"status"`
	Name        string        `json:"name"`
	State       string        `json:"state"`
	Created     time.Time     `json:"created"`
	ETag        string        `json:"eTag"`
	Modified    time.Time     `json:"modified"`
	Category    string        `json:"category"`
	URI         string        `json:"uri"`
	LIGName     string
}

//GetLI is the function called from ovcli cmd package to get information on "show li", it in turn calls RestGet
func GetLI() LIMap {

	liMapC := make(chan LIMap)
	go LIGetURI(liMapC, "Name")
	//liMap := <-liMapC

	ligMapC := make(chan LIGMap)
	go LIGGetURI(ligMapC, "URI")
	//ligMap := <-ligMapC

	var liMap LIMap
	var ligMap LIGMap

	for i := 0; i < 2; i++ {
		select {
		case ligMap = <-ligMapC:
		case liMap = <-liMapC:
		}
	}

	//Save LIG value to LI entrie field, LIGName is manually added in JSON struct. use LI's LIG URI as index to find among LIG Map
	for k := range liMap {
		liMap[k].LIGName = ligMap[liMap[k].LogicalInterconnectGroupURI].Name
	}

	return liMap

}

//LIGetURI is the function to get raw structs from all json next pages
func LIGetURI(x chan LIMap, key string) {

	log.Println("Rest Get LI")

	defer timeTrack(time.Now(), "Rest Get LI")

	c := NewCLIOVClient()

	liMap := LIMap{}
	pages := make([]LICol, 5)

	for i, uri := 0, LIURL; uri != ""; i++ {

		data, err := c.GetURI("", "", uri)
		if err != nil {
			log.Fatal(err)
		}

		err = json.Unmarshal(data, &pages[i])

		if err != nil {
			log.Fatal(err)
		}

		for k := range pages[i].Members {
			switch key {
			case "Name":
				liMap[pages[i].Members[k].Name] = &pages[i].Members[k]
			case "URI":
				liMap[pages[i].Members[k].URI] = &pages[i].Members[k]
			}
		}

		uri = pages[i].NextPageURI
	}

	//
	x <- liMap

}
