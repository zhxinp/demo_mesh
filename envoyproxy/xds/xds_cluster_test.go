// Author: Zhang Xin Peng
// This demonstration project is programmed for newly learners of service mesh.
// The purpose of this demo project is to demonstrate the basic logic of service mesh,
// Enable the beginners to obtain a basic understanding of service mesh in a short of time.
// Any Question please kindly contact me through the following email :
// 898178433@qq.com
package xds

import (
	"fmt"
	"testing"
)

func TestCluster(t *testing.T) {
	testCluster := InitCluster("outbound|9080|v2|reviews.default.svc.cluster.local")
	fmt.Printf("DIRECTION: %s \n" ,testCluster.GetDirection())
	fmt.Printf("PORT: %d \n" , testCluster.GetPort())
	fmt.Printf("SUBSET: %s \n",testCluster.GetSubset())
	fmt.Printf("FQDN: %s \n" , testCluster.GetFQDN())
	fmt.Printf("CLUSTER: %s \n" , testCluster.GetCluster())

	testCluster = InitCluster("outbound|||reviews.default.svc.cluster.local")
	fmt.Printf("DIRECTION: %s \n" ,testCluster.GetDirection())
	fmt.Printf("PORT: %d \n" , testCluster.GetPort())
	fmt.Printf("SUBSET: %s \n",testCluster.GetSubset())
	fmt.Printf("FQDN: %s \n" , testCluster.GetFQDN())
	fmt.Printf("CLUSTER: %s \n" , testCluster.GetCluster())

	testCluster01 := InitCluster("outbound||v2|reviews.default.svc.cluster.local")
	fmt.Printf("DIRECTION: %s \n" ,testCluster01.GetDirection())
	fmt.Printf("PORT: %d \n" , testCluster01.GetPort())
	fmt.Printf("SUBSET: %s \n",testCluster01.GetSubset())
	fmt.Printf("FQDN: %s \n" , testCluster01.GetFQDN())
	fmt.Printf("CLUSTER: %s \n" , testCluster01.GetCluster())


}
