package auth

import("fmt")




// if we want to keep a package to a =native file and don't want to export it we need to keep its name starting with 
// small case letters 
//but if we want to have an package which is needed to be used in out side of the folder we start with capital letters 
// this is somthing like export feture used in javascript

func LoginWithCrediential(userName string , passWord string){
	fmt.Println("user logins are ",userName,passWord)
}