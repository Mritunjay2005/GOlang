package auth
// keeping the capital letter of the first letter of the package is known as scope 
// this is applied to every scope in the go lang that keeping this first letter capital will allow the variable, package , identifer in difefrent folder acorse the program/project


func extractSession()string{// this  is a antive function -->this is private for this package
	return "logged in"
}

func GetSession()string{// this is a golabal function 
	return extractSession()
}