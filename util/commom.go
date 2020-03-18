package util

func FirstToUpper(insName *string) {

	ascByte := []byte(*insName)
	ascByte[0] -= 32
	*insName = string(ascByte)

}
