package auth

import (
	"crypto/rsa"
	"shManager/util"
	"testing"
)

func TestCreateJwt(t *testing.T) {
	type args struct {
		privateKey *rsa.PrivateKey
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "创建jwt",
			args: args{
				privateKey: util.LoadRSAPrivateKeyFromDisk("../private.pem"),
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateJwt(tt.args.privateKey); got != tt.want {
				t.Errorf("CreateJwt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseJwt(t *testing.T) {
	type args struct {
		tokenString string
		publicKey   *rsa.PublicKey
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "解析jwt",
			args: args{
				tokenString: "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTE0MDA2MzMsImlzcyI6InRlc3QiLCJmb28iOiJteSBGb28ifQ.aYrMLUALh6efhFkmo5aprh1Fe-c3HUyo_TuNqDeELOSNZ_RNk56QATJOulB1Ytaoig9OpDrOkYRhjbCNcYZ0eqbkQI2Hr7G2WAFkE_jBcYR-HvloDEE4yCXXs1mgoU-jpmDSIfTJn_zmXATpn3-PVc_Ow6lD1PDImK--b4b98LA",
				publicKey: util.LoadRSAPublicKeyFromDisk("../public.pem"),
			},
		},	
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ParseJwt(tt.args.tokenString, tt.args.publicKey)
		})
	}
}
