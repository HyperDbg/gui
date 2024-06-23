package constants

type NtstatusKind int

const (
	STATUS_SUCCESS NtstatusKind = 0x00000000
	STATUS_WAIT_1               = 0x00000001
	STATUS_WAIT_2               = 0x00000002
	STATUS_WAIT_3               = 0x00000003
)

func (k NtstatusKind) String() string {
	switch k {
	case 0x00000000:
		return "StatusSuccess"
	case 0x00000001:
		return "StatusWait1"
	case 0x00000002:
		return "StatusWait2"
	case 0x00000003:
		return "StatusWait3"
	default:
		return "unknown"
	}
}

func (k NtstatusKind) Elements() []NtstatusKind {
	return []NtstatusKind{
		STATUS_SUCCESS,
		STATUS_WAIT_1,
		STATUS_WAIT_2,
		STATUS_WAIT_3,
	}
}
