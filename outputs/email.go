package outputs

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"html/template"
	"log"

	"gopkg.in/gomail.v2"
)

type Config struct {
	Host        string
	Username    string
	Password    string
	Port        int
	Sender      string
	SenderEmail string
	Vault       *VaultConfig
	AlertUrl    string
	HeaderLogo  string
}
type RecipentConfig struct {
	To  []string
	Cc  []string
	Bcc []string
}
type VaultConfig struct {
	SecretPath  string
	UsernameKey string
	PasswordKey string
}

var (
	config *Config
)

// SendMessageToEmail - this function is used to send alerts to email .
func SendMessageToEmail(triggerName, mapvalue, tenantID string, alertmessage map[string]interface{}) error {

	// Setting up the request Body for email .
	emailbody, err := setEmailbody(AlertTemplate, alertmessage, triggerName, tenantID)
	if err != nil {
		log.Fatalf("error while setting up the email body for Alerts Template . error : %v ", err)
		return err
	}

	var recipient RecipentConfig
	err = json.Unmarshal([]byte(mapvalue), &recipient)
	if err != nil {
		log.Fatalf("error while unmarshalling the email recipients . error : %s ", err)
		return err
	}
	subject := "Alert : " + triggerName

	// Send email
	err = sendEmail(recipient.To, recipient.Cc, recipient.Bcc, subject, emailbody)
	if err != nil {
		return err
	}
	return nil
}

// setemailbody - this takes emailTemplate,alert and dataObjects and executes those dataobjects into the template .
func setEmailbody(emailTemplate string, logs map[string]interface{}, triggerName string, tenantID string) (string, error) {

	t, err := template.New("emailTemplate").Parse(emailTemplate)

	if err != nil {
		log.Fatalf("error while parsing the email template . error : %v", err)
		return "", err
	}
	//update
	tenantName := "getTenantName(tenantID)"
	if err != nil {
		return "", err
	}
	var Severity, PolicyName, Message, Cluster, Action, Result interface{}

	if logs != nil {
		Severity = logs["Severity"]
		PolicyName = logs["PolicyName"]
		Message = logs["Message"]
		Cluster = logs["ClusterName"]
		Action = logs["Action"]
		Result = logs["Result"]
	}

	var body bytes.Buffer

	err = t.Execute(&body, struct {
		TriggerName interface{}
		Severity    interface{}
		PolicyName  interface{}
		Message     interface{}
		Cluster     interface{}
		Action      interface{}
		Result      interface{}
		TenantName  interface{}
		Link        interface{}
		HeaderLogo  interface{}
	}{
		TriggerName: triggerName,
		Severity:    Severity,
		PolicyName:  PolicyName,
		Message:     Message,
		Cluster:     Cluster,
		Action:      Action,
		Result:      Result,
		TenantName:  tenantName,
		Link:        config.AlertUrl,
		HeaderLogo:  config.HeaderLogo,
	})
	if err != nil {
		log.Fatalf("error while executing the data object to email Template . error : %v", err)
		return "", err
	}
	return body.String(), nil
}

// sendEmail - this will take recipients,subject and body as input and forward the email respectively .
func sendEmail(To, Cc, Bcc []string, Subject string, Body string) error {

	dialer := gomail.NewDialer(config.Host, config.Port, config.Username, config.Password)
	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	m := gomail.NewMessage()
	m.AddAlternative("text/html", Body)

	m.SetHeaders(map[string][]string{
		"From":    {m.FormatAddress(config.SenderEmail, config.Sender)},
		"To":      To,
		"Subject": {Subject},
		"Cc":      Cc,
		"Bcc":     Bcc,
	})
	// Now send E-Mail
	if err := dialer.DialAndSend(m); err != nil {
		log.Fatalf("error while sending email . error : %v | recipients  => To : %s | Cc : %s | Bcc : %s ", err, To, Cc, Bcc)
		return err
	}

	fmt.Println("email sent successfully . recipients  => To : %s | Cc : %s | Bcc : %s ", To, Cc, Bcc)
	return nil
}
