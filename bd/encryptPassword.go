package bd

import "golang.org/x/crypto/bcrypt"

/* EncryptPassword() sirve para encriptar la contraseña */
func EncryptPassword(pass string) (string, error) {
	costo := 8 // es la cantidad de pasadas que va a hacer sobre el password, 6 para un usuario normal, 8 para un superusuario
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)
	return string(bytes), err

}
