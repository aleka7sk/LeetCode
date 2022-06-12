package leetcode

import (
	"fmt"
	"regexp"
)

func IsMatch(s string, p string) bool {
	change := fmt.Sprintf("^%v$", p)
	res, _ := regexp.Compile(change)

	if res.MatchString(s) {
		return true
	} else {
		return false
	}
}


#define LED_PIN 9
String message;
 
void setup()
{
  pinMode(LED_PIN, OUTPUT);
  Serial.begin(9600);
}
 
void loop()
{
  while (Serial.available())
  {  
  char incomingChar = Serial.read();
    message+=incomingChar;
  }
  if(message=="on")
  {
   digitalWrite(LED_PIN, HIGH);
  }
  else if (message=="off")
  {
   digitalWrite(LED_PIN, LOW);
  }
  message = "";
  delay(1000);

}