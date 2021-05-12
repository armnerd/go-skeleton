package config

import "os"

type EmailConfig struct {
	Host     string
	Port     int
	User     string
	Password string
}

var MailConfig EmailConfig

var AlertEmailList = []string{
	"pythonup@hotmail.com",
}

type Kafka struct {
	Broker []string
}

var KafkaConfig Kafka

var KafkaTopic map[string]string

type RabbitMQ struct {
	Host string
	Port string
	User string
	Pass string
}

var RabbitMQConfig RabbitMQ

var RabbitMQQueue map[string]string

func init() {
	if os.Getenv("GIN_MODE") == "release" {
		KafkaConfig = Kafka{
			[]string{
				"127.0.0.1:9092",
			},
		}
		RabbitMQConfig = RabbitMQ{
			"localhost",
			"5672",
			"guest",
			"guest",
		}
	} else {
		KafkaConfig = Kafka{
			[]string{
				"127.0.0.1:9092",
			},
		}
		RabbitMQConfig = RabbitMQ{
			"localhost",
			"5672",
			"guest",
			"guest",
		}
	}

	KafkaTopic = map[string]string{
		"topic": "topic",
	}

	RabbitMQQueue = map[string]string{
		"queue": "queue",
	}

	MailConfig = EmailConfig{
		"smtp.163.com",
		465,
		"username",
		"password",
	}
}
