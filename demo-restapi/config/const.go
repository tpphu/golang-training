package config

var IdentityKey string = "identity"
var JWTSecretKey []byte = []byte("ThisIsAVerySecretKey")
var DBConnectString = "default:secret@/notes?charset=utf8&parseTime=True&loc=Local"
