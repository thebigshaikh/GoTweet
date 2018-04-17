// package main

// import (
//   "fmt"
//   "github.com/sendgrid/sendgrid-go"
// )

// func main() {
//   sg := sendgrid.NewSendGridClient("thebigshaikh", "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")
//   message := sendgrid.NewMail()
//   message.AddTo("bshaikh@tibco.com")
//   message.AddSubject("My first email!")
//   message.AddText("Sending Email from Go using SendGrid")
//   message.AddFrom("bshaikh@tibco.com")
//   if r := sg.Send(message); r == nil {
//     fmt.Println("Email sent!")
//   } else {
//     fmt.Println(r)
//   }
// }


package main

import (
	"fmt"
	"log"
	//"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func main() {

	var(
		senderEmail string 
		receiverEmail string
		subject string
		plainTextContent string
		apiKey string = "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"
		)

	fmt.Println(" *** Enter Senders E-Mail ID ***")
	fmt.Scanf("%s", &senderEmail)

	fmt.Println(" *** Enter Receivers E-Mail ID ***")
	fmt.Scanf("%s", &receiverEmail)

	fmt.Println(" *** Enter Subject ***")
	fmt.Scanf("%s", &subject)

	fmt.Println(" *** Enter Body ***")
	fmt.Scanf("%s", &plainTextContent)


	from := mail.NewEmail("Basil Shaikh", senderEmail)
	//subject := "Sending with SendGrid is Fun"
	to := mail.NewEmail("Basil Shaikh", receiverEmail)
	//plainTextContent := "and easy to do anywhere, even with Go"
	
	htmlContent := " "
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(apiKey)
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}