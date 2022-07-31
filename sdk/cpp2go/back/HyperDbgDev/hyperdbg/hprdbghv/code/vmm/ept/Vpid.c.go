package ept
type (
	Vpid interface {
		VpidInvvpid() (ok bool)                              //col:38
		VpidInvvpidIndividualAddress() (ok bool)             //col:53
		VpidInvvpidSingleContext() (ok bool)                 //col:66
		VpidInvvpidAllContext() (ok bool)                    //col:77
		VpidInvvpidSingleContextRetainingGlobals() (ok bool) //col:90
	}
)
func NewVpid() { return &vpid{} }
func (v *vpid) VpidInvvpid() (ok bool) { //col:38
	return true
}

func (v *vpid) VpidInvvpidIndividualAddress() (ok bool) { //col:53
	return true
}

func (v *vpid) VpidInvvpidSingleContext() (ok bool) { //col:66
	return true
}

func (v *vpid) VpidInvvpidAllContext() (ok bool) { //col:77
	return true
}

func (v *vpid) VpidInvvpidSingleContextRetainingGlobals() (ok bool) { //col:90
	return true
}

