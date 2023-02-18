package internal

import (
	"testing"
)

func TestTargetHost_ParseUrl(t *testing.T) {
	type fields struct {
		Host string
		Port int64
	}
	tests := []struct {
		name     string
		fields   fields
		arg      string
		expected TargetHost
	}{
		{
			name: "Empty string",
			arg:  "",
			expected: TargetHost{
				Host: "",
				Port: 443,
			},
		},
		{
			name: "www.nos.nl",
			arg:  "www.nos.nl",
			expected: TargetHost{
				Host: "www.nos.nl",
				Port: 443,
			},
		},
		{
			name: "https://www.nos.nl",
			arg:  "https://www.nos.nl",
			expected: TargetHost{
				Host: "www.nos.nl",
				Port: 443,
			},
		},
		{
			name: "www.nos.nl Port 4433",
			arg:  "www.nos.nl",
			fields: fields{
				Port: 4433,
			},
			expected: TargetHost{
				Host: "www.nos.nl",
				Port: 4433,
			},
		},
		{
			name: "www.nos.nl:4433",
			arg:  "www.nos.nl:4433",
			fields: fields{
				Port: 4433,
			},
			expected: TargetHost{
				Host: "www.nos.nl",
				Port: 4433,
			},
		},
		{
			name: "52.222.149.103",
			arg:  "52.222.149.103",
			fields: fields{
				Port: 443,
			},
			expected: TargetHost{
				Host: "52.222.149.103",
				Port: 443,
			},
		},
		{
			name: "52.222.149.103:4433",
			arg:  "52.222.149.103:4433",
			fields: fields{
				Port: 4433,
			},
			expected: TargetHost{
				Host: "52.222.149.103",
				Port: 4433,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &TargetHost{}
			s.SetPort(tt.fields.Port)
			s.ParseUrl(tt.arg)

			if s.Port != tt.expected.Port {
				t.Fatalf("Port not equal %d want %d", s.Port, tt.expected.Port)
			}

			if s.Host != tt.expected.Host {
				t.Fatalf("Host not equal %s want %s", s.Host, tt.expected.Host)
			}
		})
	}
}
