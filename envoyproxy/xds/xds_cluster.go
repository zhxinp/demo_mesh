package xds

import (
	"strconv"
	"strings"
)

type Cluster struct {
	FQDN string
	Port int
	Subset string
	Direction string
}

func (cl *Cluster) GetFQDN() string {
	if cl != nil {
		return cl.FQDN
	}
	return ""
}

func (cl *Cluster) GetPort() int {
	if cl != nil {
		return cl.Port
	}
	return 0
}

func (cl *Cluster) GetDirection() string {
	if cl != nil {
		return cl.Direction
	}
	return ""
}

func (cl *Cluster) GetSubset() string {
	if cl != nil {
		return cl.Subset
	}
	return ""
}

func (cl *Cluster) GetCluster() string {
	if cl != nil {
		if cl.Port == 0 {
			return cl.Direction + "||" + cl.Subset + "|" + cl.FQDN
		}
		return cl.Direction + "|" + strconv.Itoa(cl.Port) + "|" + cl.Subset + "|" + cl.FQDN
	}
	return ""
}

func InitCluster(clusterString string) *Cluster {
	newCluster := Cluster{}
	cluster := strings.Split(clusterString,"|")
	if len(cluster) != 4 {
		return nil
	}
	newCluster.Direction = cluster[0]
	if cluster[1] == "" {
		newCluster.Port = 0
	} else {
		port ,err := strconv.Atoi(cluster[1])
		if err == nil {
			newCluster.Port = port
		} else {
			newCluster.Port = 0
		}
	}
	newCluster.Subset = cluster[2]
	newCluster.FQDN = cluster[3]
	return &newCluster
}