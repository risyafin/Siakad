package module

import "github.com/golang-jwt/jwt/v4"

type Admin struct {
	Username string
	Password string
}

type Mahasiswa struct{
	NIM string
	Password string
}

type Dosen struct{
	NIDN string 
	Password string
}

type MyClaims struct{
	jwt.StandardClaims
	NIM string
	NIDN string
	Username string
}