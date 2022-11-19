package Headers

import "testing"

func TestCTL_CODE(t *testing.T) {
	type args struct {
		deviceType uint32
		function   uint32
		method     uint32
		access     uint32
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CTL_CODE(tt.args.deviceType, tt.args.function, tt.args.method, tt.args.access); got != tt.want {
				t.Errorf("CTL_CODE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestErrorCodes_String(t *testing.T) {
	tests := []struct {
		name string
		e    ErrorCodes
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHIBYTE(t *testing.T) {
	type args struct {
		l uint32
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HIBYTE(tt.args.l); got != tt.want {
				t.Errorf("HIBYTE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHIWORD(t *testing.T) {
	type args struct {
		l uint32
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HIWORD(tt.args.l); got != tt.want {
				t.Errorf("HIWORD() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLOBYTE(t *testing.T) {
	type args struct {
		l uint32
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LOBYTE(tt.args.l); got != tt.want {
				t.Errorf("LOBYTE() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLOWORD(t *testing.T) {
	type args struct {
		l uint32
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LOWORD(tt.args.l); got != tt.want {
				t.Errorf("LOWORD() = %v, want %v", got, tt.want)
			}
		})
	}
}
