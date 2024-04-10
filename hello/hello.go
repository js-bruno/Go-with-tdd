package main

import "fmt"

const prefixHelloInFrench  = "bonjour, " 
const prefixHelloInSpanish  = "hola, " 
const prefixHelloInEnglish = "hello, "

func Hello(name string, language string) string {
  if name  == ""{
    name = "World!!!"
  }

 return salutePrefix(language) + name
}

func salutePrefix(language string) string {
  switch language {
  case"french" :
    return prefixHelloInFrench 

  case "spanish" :
    return prefixHelloInSpanish 
  
  default:
    return prefixHelloInEnglish
}

}

func main() {
	fmt.Println(Hello("World", ""))
}
