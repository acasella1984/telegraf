package inputs

import (
	"encoding/json"
	"fmt"
	"net"
	"io/ioutil"
	"net/http"
	"time"
	"strconv"
	"os"
	"github.com/golang/glog"
	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/inputs"
)


type Platform struct {
	ObjectID string `json:"ObjectId"`
	Object struct {
		ObjName string `json:"ObjName"`
		ProductName string `json:"ProductName"`
		SerialNum string `json:"SerialNum"`
		Manufacturer string `json:"Manufacturer"`
		Vendor string `json:"Vendor"`
		Release string `json:"Release"`
		PlatformName string `json:"PlatformName"`
		Version string `json:"Version"`
	} `json:"Object"`
}


type PSUState struct {
	MoreExist bool `json:"MoreExist"`
	ObjCount int64`json:"ObjCount"`
	CurrentMarker int64`json:"CurrentMarker"`
	NextMarker int64`json:"NextMarker"`
	Objects []struct {
		ObjectID string `json:"ObjectId"`
		Object struct {
			PsuID int `json:"PsuId"`
			AdminState string `json:"AdminState"`
			ModelNum string `json:"ModelNum"`
			SerialNum string `json:"SerialNum"`
			Vin int64`json:"Vin"`
			Vout int64`json:"Vout"`
			Iin int64`json:"Iin"`
			Iout int64`json:"Iout"`
			Pin int64`json:"Pin"`
			Pout int64`json:"Pout"`
			Fan string `json:"Fan"`
			FanID int64`json:"FanId"`
			LedID int64`json:"LedId"`
		} `json:"Object"`
	} `json:"Objects"`
}

type SFPState struct {
	MoreExist bool `json:"MoreExist"`
	ObjCount int64`json:"ObjCount"`
	CurrentMarker int64`json:"CurrentMarker"`
	NextMarker int64`json:"NextMarker"`
	Objects []struct {
		ObjectID string `json:"ObjectId"`
		Object struct {
			SfpID int `json:"SfpId"`
			SfpSpeed string `json:"SfpSpeed"`
			SfpLOS string `json:"SfpLOS"`
			SfpPresent string `json:"SfpPresent"`
			SfpType string `json:"SfpType"`
			SerialNum string `json:"SerialNum"`
			EEPROM string `json:"EEPROM"`
		} `json:"Object"`
	} `json:"Objects"`
}

type CoppState struct {
	MoreExist bool `json:"MoreExist"`
	ObjCount int64`json:"ObjCount"`
	CurrentMarker int64`json:"CurrentMarker"`
	NextMarker int64`json:"NextMarker"`
	Objects []struct {
		ObjectID string `json:"ObjectId"`
		Object struct {
			Protocol string `json:"Protocol"`
			PeakRate int64`json:"PeakRate"`
			BurstRate int64`json:"BurstRate"`
			GreenPackets int64`json:"GreenPackets"`
			RedPackets int64`json:"RedPackets"`
		} `json:"Object"`
	} `json:"Objects"`
}

type IPv4IntfStates struct {
	MoreExist bool `json:"MoreExist"`
	ObjCount int64`json:"ObjCount"`
	CurrentMarker int64`json:"CurrentMarker"`
	NextMarker int64`json:"NextMarker"`
	Objects []struct {
		ObjectID string `json:"ObjectId"`
		Object struct {
			IntfRef string `json:"IntfRef"`
			IfIndex int64`json:"IfIndex"`
			IPAddr string `json:"IpAddr"`
			OperState string `json:"OperState"`
			NumUpEvents int64`json:"NumUpEvents"`
			LastUpEventTime string `json:"LastUpEventTime"`
			NumDownEvents int64`json:"NumDownEvents"`
			LastDownEventTime string `json:"LastDownEventTime"`
			L2IntfType string `json:"L2IntfType"`
			L2IntfID int64`json:"L2IntfId"`
		} `json:"Object"`
	} `json:"Objects"`
}


type VlanState struct {
	MoreExist bool `json:"MoreExist"`
	ObjCount int64`json:"ObjCount"`
	CurrentMarker int64`json:"CurrentMarker"`
	NextMarker int64`json:"NextMarker"`
	Objects []struct {
		ObjectID string `json:"ObjectId"`
		Object struct {
			VlanID int `json:"VlanId"`
			Name string `json:"Name"`
			OperState string `json:"OperState"`
			IfIndex int64 `json:"IfIndex"`
			SysInternalDescription string `json:"SysInternalDescription"`
		} `json:"Object"`
	} `json:"Objects"`
}



type RouteStats struct {
	ObjectID string `json:"ObjectId"`
	Object struct {
		Vrf string `json:"Vrf"`
		TotalRouteCount int `json:"TotalRouteCount"`
		ECMPRouteCount int `json:"ECMPRouteCount"`
		V4RouteCount int `json:"V4RouteCount"`
		V6RouteCount int `json:"V6RouteCount"`
		PerProtocolRouteCountList []struct {
			Protocol string `json:"Protocol"`
			RouteCount int `json:"RouteCount"`
			EcmpCount int `json:"EcmpCount"`
		} `json:"PerProtocolRouteCountList"`
	} `json:"Object"`
}


type RouteStatsPerInt struct {
	MoreExist bool `json:"MoreExist"`
	ObjCount int64`json:"ObjCount"`
	CurrentMarker int64`json:"CurrentMarker"`
	NextMarker int64`json:"NextMarker"`
	Objects []struct {
		ObjectID string `json:"ObjectId"`
		Object struct {
			Intfref string `json:"Intfref"`
			V4Routes []string `json:"V4Routes"`
			V6Routes interface{} `json:"V6Routes"`
		} `json:"Object"`
	} `json:"Objects"`
}

type RouteStatsPerProto struct {
	MoreExist bool `json:"MoreExist"`
	ObjCount int64`json:"ObjCount"`
	CurrentMarker int64`json:"CurrentMarker"`
	NextMarker int64`json:"NextMarker"`
	Objects []struct {
		ObjectID string `json:"ObjectId"`
		Object struct {
			Protocol string `json:"Protocol"`
			V4Routes []struct {
				DestinationNw string `json:"DestinationNw"`
				IsInstalledInHw bool `json:"IsInstalledInHw"`
				NextHopList []struct {
					NextHopIP string `json:"NextHopIp"`
					NextHopIntRef string `json:"NextHopIntRef"`
					Weight int64`json:"Weight"`
				} `json:"NextHopList"`
			} `json:"V4Routes"`
			V6Routes interface{} `json:"V6Routes"`
		} `json:"Object"`
	} `json:"Objects"`
}

type AsicSummary struct {
	ObjectId string `json:"ObjectId"`
	Object   struct {
			ModuleId      int64  `json:"ModuleId"`
			NumPortsUp    int64  `json:"NumPortsUp"`
			NumPortsDown  int64  `json:"NumPortsDown"`
			NumVlans      int64  `json:"NumVlans"`
			NumV4Intfs    int64  `json:"NumV4Intfs"`
			NumV6Intfs    int64  `json:"NumV6Intfs"`
			NumV4Adjs     int64  `json:"NumV4Adjs"`
			NumV6Adjs     int64  `json:"NumV6Adjs"`
			NumV4Routes   int64  `json:"NumV4Routes"`
			NumV6Routes   int64  `json:"NumV6Routes"`
			NumECMPRoutes int64  `json:"NumECMPRoutes"`			
	} `json:"Object"`
}



type Status struct {
	MoreExist bool `json:"MoreExist"`
	ObjCount int64`json:"ObjCount"`
	CurrentMarker int64`json:"CurrentMarker"`
	NextMarker int64`json:"NextMarker"`
	ObjectID string `json:"ObjectId"`
	Object struct {
		Name string `json:"Name"`
		Ready bool `json:"Ready"`
		Reason string `json:"Reason"`
		UpTime string `json:"UpTime"`
		NumCreateCalls string `json:"NumCreateCalls"`
		NumDeleteCalls string `json:"NumDeleteCalls"`
		NumUpdateCalls string `json:"NumUpdateCalls"`
		NumGetCalls string `json:"NumGetCalls"`
		NumActionCalls string `json:"NumActionCalls"`
		FlexDaemons []struct {
			Name string `json:"Name"`
			Enable bool `json:"Enable"`
			State string `json:"State"`
			Reason string `json:"Reason"`
			StartTime string `json:"StartTime"`
			KeepAlive string `json:"KeepAlive"`
			RestartCount int64`json:"RestartCount"`
			RestartTime string `json:"RestartTime"`
			RestartReason string `json:"RestartReason"`
		} `json:"FlexDaemons"`
	} `json:"Object"`
}


type PortState struct {
	MoreExist bool `json:"MoreExist"`
	ObjCount int64`json:"ObjCount"`
	CurrentMarker int64`json:"CurrentMarker"`
	NextMarker int64`json:"NextMarker"`
	Objects []struct {
		ObjectID string `json:"ObjectId"`
		Object struct {
			IntfRef string `json:"IntfRef"`
			IfIndex int64`json:"IfIndex"`
			Name string `json:"Name"`
			OperState string `json:"OperState"`
			NumUpEvents int64`json:"NumUpEvents"`
			LastUpEventTime string `json:"LastUpEventTime"`
			NumDownEvents int64`json:"NumDownEvents"`
			LastDownEventTime string `json:"LastDownEventTime"`
			Pvid int64`json:"Pvid"`
			IfInOctets int64`json:"IfInOctets"`
			IfInUcastPkts int64`json:"IfInUcastPkts"`
			IfInDiscards int64`json:"IfInDiscards"`
			IfInErrors int64`json:"IfInErrors"`
			IfInUnknownProtos int64`json:"IfInUnknownProtos"`
			IfOutOctets int64`json:"IfOutOctets"`
			IfOutUcastPkts int64`json:"IfOutUcastPkts"`
			IfOutDiscards int64`json:"IfOutDiscards"`
			IfOutErrors int64`json:"IfOutErrors"`
			IfEtherUnderSizePktCnt int64`json:"IfEtherUnderSizePktCnt"`
			IfEtherOverSizePktCnt int64`json:"IfEtherOverSizePktCnt"`
			IfEtherFragments int64`json:"IfEtherFragments"`
			IfEtherCRCAlignError int64`json:"IfEtherCRCAlignError"`
			IfEtherJabber int64`json:"IfEtherJabber"`
			IfEtherPkts int64`json:"IfEtherPkts"`
			IfEtherMCPkts int64`json:"IfEtherMCPkts"`
			IfEtherBcastPkts int64`json:"IfEtherBcastPkts"`
			IfEtherPkts64OrLessOctets int64`json:"IfEtherPkts64OrLessOctets"`
			IfEtherPkts65To127Octets int64`json:"IfEtherPkts65To127Octets"`
			IfEtherPkts128To255Octets int64`json:"IfEtherPkts128To255Octets"`
			IfEtherPkts256To511Octets int64`json:"IfEtherPkts256To511Octets"`
			IfEtherPkts512To1023Octets int64`json:"IfEtherPkts512To1023Octets"`
			IfEtherPkts1024To1518Octets int64`json:"IfEtherPkts1024To1518Octets"`
			ErrDisableReason string `json:"ErrDisableReason"`
			PresentInHW string `json:"PresentInHW"`
			ConfigMode string `json:"ConfigMode"`
			PRBSRxErrCnt int64`json:"PRBSRxErrCnt"`
			PcpToCosProfileRef string `json:"PcpToCosProfileRef"`
			DscpToCosProfileRef string `json:"DscpToCosProfileRef"`
			SchedProfileRef string `json:"SchedProfileRef"`
			OperSpeed int64`json:"OperSpeed"`
			OperDuplex string `json:"OperDuplex"`
		} `json:"Object"`
	} `json:"Objects"`
}

type BufferPortStats struct {
	MoreExist bool `json:"MoreExist"`
	ObjCount int64`json:"ObjCount"`
	CurrentMarker int64`json:"CurrentMarker"`
	NextMarker int64`json:"NextMarker"`
	Objects []struct {
		ObjectId string `json:"ObjectId"`
		Object   struct {
			IntfRef        string `json:"IntfRef"`
			IfIndex        int64  `json:"IfIndex"`
			EgressPort     int64  `json:"EgressPort"`
			IngressPort    int64  `json:"IngressPort"`
			PortBufferStat int64  `json:"PortBufferStat"`
		} `json:"Object"`
	} `json:"Objects"`
}

type ConfigLogs struct {
	MoreExist bool `json:"MoreExist"`
	ObjCount int64 `json:"ObjCount"`
	CurrentMarker int64 `json:"CurrentMarker"`
	NextMarker int64 `json:"NextMarker"`
	Objects []struct {
		ObjectID string `json:"ObjectId"`
		Object struct {
			SeqNum int64 `json:"SeqNum"`
			Time string `json:"Time"`
			API string `json:"API"`
			Operation string `json:"Operation"`
			Data string `json:"Data"`
			Result string `json:"Result"`
			UserAddr string `json:"UserAddr"`
			UserName string `json:"UserName"`
		} `json:"Object"`
	} `json:"Objects"`
}


type SnapRoute struct {
	Url               string
	IsBarefoot        bool
	lastTime          time.Time
	lastIfInUcastPkts    [256]int64
	lastIfOutUcastPkts   [256]int64
	lastInDiscards    [256]int64
	lastOutDiscards   [256]int64
	lastEtherPkts     [256]int64
	lastEtherMCPkts   [256]int64
	lastEtherBCPkts   [256]int64
}

var sampleConfig = `
  ## URL-prefix for SnapRoute
  url = "http://localhost:8080/public/v1/"
  ## Set if box is Barefoot
  isBarefoot = false
`


func (_ *SnapRoute) Description() string {
	return `Gather SnapRoute Metrics`
}

func (_ *SnapRoute) SampleConfig() string {
	return sampleConfig
}

func (s *SnapRoute) Gather(acc telegraf.Accumulator) error {
	defer func() {
		hostname,err := os.Hostname()
		if err != nil {
    		fmt.Println("Hostname get error","Error:",err)
  			}
		if r := recover(); r != nil {
			glog.Error("E! Problem reading from SnapRoute")
			acc.AddFields("status", map[string]interface{}{"ready": false}, map[string]string{"hostname":hostname}, time.Now())
		}
	}()

	now := time.Now()
	hostname,err := os.Hostname()
	mgmtip := ""
	mgmtipv6 := ""
	list, err := net.Interfaces()
	if err != nil {
		panic(err)
	}   

	for _, iface := range list {
		addrs, err := iface.Addrs()
		if err != nil {
			panic(err)
		}
		if iface.Name == "ma1" {
			for _, addr := range addrs {
				ip, _, _ := net.ParseCIDR(addr.String())
				if ip.To4() != nil {
				    mgmtip = addr.String()				 
				}else {
					if net.IP.IsGlobalUnicast(ip){
						mgmtipv6 = addr.String()
					}
				}
			}
		}
		if iface.Name == "eth0" {
			for _, addr := range addrs {
				ip := net.ParseIP(addr.String())	
				if ip.To4() != nil {
				      mgmtip = addr.String()
				}else {
					if net.IP.IsGlobalUnicast(ip){
						mgmtipv6 = addr.String()
					}
				}
			}
		}		
 	}
	fmt.Printf("ipv4 =%s ipv6 = %s\n" ,mgmtip, mgmtipv6)	
	tags := map[string]string{"hostname":hostname, "mgmt-ip":mgmtip, "mgmt-ipv6":mgmtipv6}
	if err != nil {
    	fmt.Println("Hostname get error","Error:",err)
  	}


//Platform

		var platform Platform
		var requestURL = fmt.Sprint(s.Url, "state/platform")
		content, err := getContent(requestURL)
		if err != nil {
			glog.Error("Error talking to SnapRoute:", err)
			acc.AddFields("status", map[string]interface{}{"ready": false}, tags, time.Now())
			return err
		}

		err = json.Unmarshal(content, &platform)
		if err != nil {
			glog.Error("Error Umarshalling:", err)
			glog.Error("content:", content)
			return err
		}
		acc.AddFields("platform", map[string]interface{}{"ProductName": platform.Object.ProductName}, tags, now)
		acc.AddFields("platform", map[string]interface{}{"SerialNum": platform.Object.SerialNum}, tags, now)
		acc.AddFields("platform", map[string]interface{}{"Manufacturer": platform.Object.Manufacturer}, tags, now)
		acc.AddFields("platform", map[string]interface{}{"Vendor": platform.Object.Vendor}, tags, now)
		acc.AddFields("platform", map[string]interface{}{"Release": platform.Object.Release}, tags, now)
		acc.AddFields("platform", map[string]interface{}{"PlatformName": platform.Object.PlatformName}, tags, now)
		acc.AddFields("platform", map[string]interface{}{"Version": platform.Object.Version}, tags, now)
		

//Environmentals (Fans, PSUState, Thermals)


	//PSU
		var psustate PSUState 
		requestURL = fmt.Sprint(s.Url, "state/psus")
		content, err = getContent(requestURL)
		if err != nil {
			glog.Error("Error talking to SnapRoute:", err)
			acc.AddFields("status", map[string]interface{}{"ready": false}, map[string]string{"hostname":hostname}, time.Now())
			return err
		}

		err = json.Unmarshal(content, &psustate)
		if err != nil {
			glog.Error("Error Umarshalling:", err)
			glog.Error("content:", content)
			return err
		}
		for _, psu := range psustate.Objects {
				psuid := strconv.Itoa(psu.Object.PsuID)
				tags := map[string]string{
					"PsuId": psuid,
					"Hostname": hostname,
					"mgmtip":mgmtip,
					"mgmtipv6":mgmtipv6,
				}
			acc.AddFields("psu", map[string]interface{}{"AdminState": psu.Object.AdminState}, tags, now)
			acc.AddFields("psu", map[string]interface{}{"ModelNum": psu.Object.ModelNum}, tags, now)
			acc.AddGauge("psu", map[string]interface{}{"SerialNum": psu.Object.SerialNum}, tags, now)
			acc.AddGauge("psu", map[string]interface{}{"Volts in": psu.Object.Vin}, tags, now)
			acc.AddGauge("psu", map[string]interface{}{"Volts out": psu.Object.Vout}, tags, now)
			acc.AddGauge("psu", map[string]interface{}{"Amps In": psu.Object.Iin}, tags, now)
			acc.AddGauge("psu", map[string]interface{}{"Amps Out": psu.Object.Iout}, tags, now)
			acc.AddGauge("psu", map[string]interface{}{"Power In": psu.Object.Pin}, tags, now)
			acc.AddGauge("psu", map[string]interface{}{"Power Out": psu.Object.Pout}, tags, now)
			acc.AddFields("psu", map[string]interface{}{"Fan Status": psu.Object.Fan}, tags, now)
			acc.AddGauge("psu", map[string]interface{}{"Fan ID": psu.Object.FanID}, tags, now)
			acc.AddGauge("psu", map[string]interface{}{"Led ID": psu.Object.LedID}, tags, now)
		}
	//SFPs

		var sfps SFPState 
		requestURL = fmt.Sprint(s.Url, "state/sfps")
		content, err = getContent(requestURL)
		if err != nil {
			glog.Error("Error talking to SnapRoute:", err)
			acc.AddFields("status", map[string]interface{}{"ready": false}, map[string]string{"hostname":hostname}, time.Now())
			return err
		}

		err = json.Unmarshal(content, &sfps)
		if err != nil {
			glog.Error("Error Umarshalling:", err)
			glog.Error("content:", content)
			return err
		}
		for _, sfp := range sfps.Objects {
				sfpid := strconv.Itoa(sfp.Object.SfpID)
				tags := map[string]string{
					"SfpId": sfpid,
					"Hostname": hostname,
					"mgmtip":mgmtip,
					"mgmtipv6":mgmtipv6,
				}
		acc.AddFields("sfp", map[string]interface{}{"SfpSpeed": sfp.Object.SfpSpeed}, tags, now)
		acc.AddFields("sfp", map[string]interface{}{"SfpLOS": sfp.Object.SfpLOS}, tags, now)
		acc.AddFields("sfp", map[string]interface{}{"SfpPresent": sfp.Object.SfpPresent}, tags, now)
		acc.AddFields("sfp", map[string]interface{}{"SfpType": sfp.Object.SfpType}, tags, now)
		acc.AddFields("sfp", map[string]interface{}{"SerialNum": sfp.Object.SerialNum}, tags, now)
		acc.AddFields("sfp", map[string]interface{}{"EEPROM": sfp.Object.EEPROM}, tags, now)
		}
	//Thermals
	
//CoPP

		var copp CoppState 
		requestURL = fmt.Sprint(s.Url, "state/coppstate")
		content, err = getContent(requestURL)
		if err != nil {
			glog.Error("Error talking to SnapRoute:", err)
			acc.AddFields("status", map[string]interface{}{"ready": false}, map[string]string{"hostname":hostname}, time.Now())
			return err
		}

		err = json.Unmarshal(content, &copp)
		if err != nil {
			glog.Error("Error Umarshalling:", err)
			glog.Error("content:", content)
			return err
		}
		for _, coppstate := range copp.Objects {
						tags := map[string]string{
							"Protocol": coppstate.Object.Protocol,
							"Hostname": hostname,
							"mgmtip":mgmtip,
							"mgmtipv6":mgmtipv6,
						}
						
		acc.AddGauge("copp", map[string]interface{}{"PeakRate": coppstate.Object.PeakRate}, tags, now)
		acc.AddGauge("copp", map[string]interface{}{"BurstRate": coppstate.Object.BurstRate}, tags, now)
		acc.AddCounter("copp", map[string]interface{}{"GreenPackets": coppstate.Object.GreenPackets}, tags, now)
		acc.AddCounter("copp", map[string]interface{}{"RedPackets": coppstate.Object.RedPackets}, tags, now)
	}
	
//VlanState

		var vlans VlanState
		requestURL = fmt.Sprint(s.Url, "state/vlans")
		content, err = getContent(requestURL)
		if err != nil {
			glog.Error("Error talking to SnapRoute:", err)
			acc.AddFields("status", map[string]interface{}{"ready": false}, map[string]string{"hostname":hostname}, time.Now())
			return err
		}

		err = json.Unmarshal(content, &vlans)
		if err != nil {
			glog.Error("Error Umarshalling:", err)
			glog.Error("content:", content)
			return err
		}

		for _, vlanstate := range vlans.Objects {
				vlanid := strconv.Itoa(vlanstate.Object.VlanID)

				tags := map[string]string{
					"VlanId": vlanid,
					"Hostname": hostname,
					"mgmtip":mgmtip,
					"mgmtipv6":mgmtipv6,
				}
					acc.AddFields("ports", map[string]interface{}{"Name": vlanstate.Object.Name}, tags, now)
					acc.AddFields("ports", map[string]interface{}{"OperState": vlanstate.Object.OperState}, tags, now)
					acc.AddGauge("ports", map[string]interface{}{"IfIndex": vlanstate.Object.IfIndex}, tags, now)
					acc.AddFields("ports", map[string]interface{}{"SysInternal Description": vlanstate.Object.SysInternalDescription}, tags, now)
				
		}



//PortState

		var ports PortState
		requestURL = fmt.Sprint(s.Url, "state/Ports")
		content, err = getContent(requestURL)
		if err != nil {
			glog.Error("Error talking to SnapRoute:", err)
			acc.AddFields("status", map[string]interface{}{"ready": false}, map[string]string{"hostname":hostname}, time.Now())
			return err
		}

		err = json.Unmarshal(content, &ports)
		if err != nil {
			glog.Error("Error Umarshalling:", err)
			glog.Error("content:", content)
			return err
		}

		for _, port := range ports.Objects {
				tags := map[string]string{
					"port": port.Object.IntfRef,
					"Hostname": hostname,
					"mgmtip":mgmtip,
					"mgmtipv6":mgmtipv6,
				}
					acc.AddFields("ports", map[string]interface{}{"OperState": port.Object.OperState}, tags, now)
					acc.AddFields("ports", map[string]interface{}{"OperSpeed": port.Object.OperState}, tags, now)
					acc.AddCounter("ports", map[string]interface{}{"IfInUcastPkts": port.Object.IfInUcastPkts}, tags, now)
					acc.AddCounter("ports", map[string]interface{}{"IfOutUcastPkts": port.Object.IfOutUcastPkts}, tags, now)
					acc.AddCounter("ports", map[string]interface{}{"IfinDiscards": port.Object.IfInDiscards}, tags, now)
					acc.AddCounter("ports", map[string]interface{}{"IfoutDiscards": port.Object.IfOutDiscards}, tags, now)
  					acc.AddCounter("ports", map[string]interface{}{"IfInOctets": port.Object.IfInOctets}, tags, now)
  					acc.AddCounter("ports", map[string]interface{}{"IfOutOctets": port.Object.IfOutOctets}, tags, now)
					acc.AddCounter("ports", map[string]interface{}{"IfEtherPkts": port.Object.IfEtherPkts}, tags, now)
					acc.AddCounter("ports", map[string]interface{}{"IfEtherMCPkts": port.Object.IfEtherMCPkts}, tags, now)
					acc.AddCounter("ports", map[string]interface{}{"IfEtherBcastPkts":  port.Object.IfEtherBcastPkts}, tags, now)
		}


//AsicSummary

	s.lastTime = now

	var asicsum AsicSummary
	
	requestURL = fmt.Sprint(s.Url, "state/asicsummary")
	content, err = getContent(requestURL)
	if err != nil {
		glog.Error("Error talking to SnapRoute:", err)
		return err
	}

	err = json.Unmarshal(content, &asicsum)
	if err != nil {
		glog.Error("Error Unmarshalling:", err)
		glog.Error("content:", content)
		return err
	}
	acc.AddCounter("asicsum", map[string]interface{}{"ModuleId": asicsum.Object.ModuleId},map[string]string{"hostname":hostname}, time.Now())
	acc.AddCounter("asicsum", map[string]interface{}{"NumPortsUp": asicsum.Object.NumPortsUp},map[string]string{"hostname":hostname}, time.Now())
	acc.AddCounter("asicsum", map[string]interface{}{"NumPortsDown": asicsum.Object.NumPortsDown},map[string]string{"hostname":hostname},time.Now())
	acc.AddCounter("asicsum", map[string]interface{}{"NumVlans": asicsum.Object.NumVlans}, map[string]string{"hostname":hostname}, time.Now())
	acc.AddCounter("asicsum", map[string]interface{}{"NumV4Intfs": asicsum.Object.NumV4Intfs}, map[string]string{"hostname":hostname}, time.Now())
	acc.AddCounter("asicsum", map[string]interface{}{"NumV6Intfs": asicsum.Object.NumV6Intfs}, map[string]string{"hostname":hostname}, time.Now())
	acc.AddCounter("asicsum", map[string]interface{}{"NumV4Adjs": asicsum.Object.NumV4Adjs}, map[string]string{"hostname":hostname}, time.Now())
	acc.AddCounter("asicsum", map[string]interface{}{"NumV6Adjs": asicsum.Object.NumV6Adjs}, map[string]string{"hostname":hostname}, time.Now())
	acc.AddCounter("asicsum", map[string]interface{}{"NumV4Routes": asicsum.Object.NumV4Routes}, map[string]string{"hostname":hostname}, time.Now())
	acc.AddCounter("asicsum", map[string]interface{}{"NumV6Routes": asicsum.Object.NumV6Routes}, map[string]string{"hostname":hostname}, time.Now())
	acc.AddCounter("asicsum", map[string]interface{}{"NumECMPRoutes": asicsum.Object.NumECMPRoutes}, map[string]string{"hostname":hostname}, time.Now())
	
//SystemStatus
	s.lastTime = now

	var status Status
	requestURL = fmt.Sprint(s.Url, "state/SystemStatus")
	content, err = getContent(requestURL)
	if err != nil {
		glog.Error("Error talking to SnapRoute:", err)
		acc.AddFields("status", map[string]interface{}{"ready": false}, map[string]string{"hostname":hostname}, time.Now())
		return err
	}

	err = json.Unmarshal(content, &status)
	if err != nil {
		glog.Error("Error Unmarshalling:", err)
		glog.Error("content:", content)
		return err
	}

	acc.AddFields("status", map[string]interface{}{"Ready": status.Object.Ready}, map[string]string{"hostname":hostname}, time.Now())
	acc.AddFields("status", map[string]interface{}{"ReadyReason": status.Object.Reason},map[string]string{"hostname":hostname}, time.Now())
	acc.AddFields("status", map[string]interface{}{"Uptime": status.Object.UpTime}, map[string]string{"hostname":hostname}, time.Now())
	for _, daemon := range status.Object.FlexDaemons {
				tags := map[string]string{
					"Daemon": daemon.Name,
					"Hostname": hostname,
					"mgmtip":mgmtip,
					"mgmtipv6":mgmtipv6,
				}

			acc.AddFields("DaemonStats", map[string]interface{}{"Enable": daemon.Enable}, tags, now)
			acc.AddFields("DaemonStats", map[string]interface{}{"State": daemon.State}, tags, now)
			acc.AddFields("DaemonStats", map[string]interface{}{"Reason": daemon.Reason}, tags, now)
			acc.AddFields("DaemonStats", map[string]interface{}{"StartTime": daemon.StartTime}, tags, now)
			acc.AddFields("DaemonStats", map[string]interface{}{"KeepAlive": daemon.KeepAlive}, tags, now)
			acc.AddCounter("DaemonStats", map[string]interface{}{"RestartCount": daemon.RestartCount}, tags, now)
			acc.AddFields("DaemonStats", map[string]interface{}{"RestartTime": daemon.RestartTime}, tags, now)
			acc.AddFields("DaemonStats", map[string]interface{}{"RestartReason": daemon.RestartReason}, tags, now)			
			
		}


//RouteStats


s.lastTime = now

	var routestat RouteStats
	
	requestURL = fmt.Sprint(s.Url, "state/routestat")
	content, err = getContent(requestURL)
	if err != nil {
		glog.Error("Error talking to SnapRoute:", err)
		return err
	}

	err = json.Unmarshal(content, &routestat)
	if err != nil {
		glog.Error("Error Unmarshalling:", err)
		glog.Error("content:", content)
		return err
	}

	acc.AddCounter("routestats", map[string]interface{}{"TotalRouteCount": routestat.Object.TotalRouteCount}, map[string]string{"hostname":hostname}, now)
	acc.AddCounter("routestats", map[string]interface{}{"ECMPRouteCount": routestat.Object.ECMPRouteCount}, map[string]string{"hostname":hostname}, now)
	acc.AddCounter("routestats", map[string]interface{}{"IPv4RouteCount": routestat.Object.V4RouteCount}, map[string]string{"hostname":hostname}, now)
	acc.AddCounter("routestats", map[string]interface{}{"IPv6RouteCount": routestat.Object.V6RouteCount}, map[string]string{"hostname":hostname}, now)
	
//BufferPortStats (BST)

s.lastTime = now

	var buffstat BufferPortStats
	
	requestURL = fmt.Sprint(s.Url, "state/bufferportstats")
	content, err = getContent(requestURL)
	if err != nil {
		glog.Error("Error talking to SnapRoute:", err)
		return err
	}

	err = json.Unmarshal(content, &buffstat)
	if err != nil {
		glog.Error("Error Unmarshalling:", err)
		glog.Error("content:", content)
		return err
	}


		for _, buffport := range buffstat.Objects {
				tags := map[string]string{
					"IntfRef": buffport.Object.IntfRef,
					"Hostname": hostname,
					"mgmtip":mgmtip,
					"mgmtipv6":mgmtipv6,
				}

			acc.AddCounter("bufferportstats", map[string]interface{}{"EgressPort": buffport.Object.EgressPort}, tags, now)
			acc.AddCounter("bufferportstats", map[string]interface{}{"IngressPort": buffport.Object.IngressPort}, tags, now)
			acc.AddCounter("bufferportstats", map[string]interface{}{"PortBufferStat": buffport.Object.PortBufferStat}, tags, now)

		}

	return nil
}


func getContent(url string) ([]byte, error) {

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}

func init() {
	inputs.Add("snaproute", func() telegraf.Input {
		return &SnapRoute{}
	})
}
