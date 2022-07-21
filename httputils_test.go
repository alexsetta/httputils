package httputils

import "testing"

func TestGetBody(t *testing.T) {
	tests := []struct {
		name string
		url  string
	}{
		{"Google", "https://www.google.com.br/"},
		{"Microsoft", "https://www.microsoft.com/"},
		{"AlexSetta", "http://alexsetta.com/"},
		{"Mega-sena", "https://servicebus2.caixa.gov.br/portaldeloterias/api/megasena"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetBody(tt.url, false)
			if err != nil {
				t.Errorf("GetBody() error = %v", err)
				return
			}
			if len(got) == 0 {
				t.Errorf("GetBody() got = %v", got)
			}
		})
	}
}
