package customerpref

import(
	"fmt"
	"time"
	"github.com/Shopify/sarama"
)

type Event struct {
	name	string
	time	string
	id		string
	pref 	string
}


func SendUpdateCustomerPref(id string, pref string) {
	t := time.Now().UTC()
	e := Event{"Update_Customer_Preference", t.Format("20060102150405") , id, pref}

	client, err := sarama.NewClient("customer_pref", []string{"localhost:9092"}, sarama.NewClientConfig())
	if err != nil {
		panic(err)
	} else {
		fmt.Println("> Connected")
	}

	defer client.Close()

	producer, err := sarama.NewProducer(client, nil)
	if err != nil {
		panic(err)
	}

	defer producer.Close()
	nd := fmt.Sprintf("name:%s;time:%s;id:%s;pref:%s", e.name, e.time, e.id, e.pref)
	select {
		case producer.Input() <- &sarama.MessageToSend{Topic: "eventbus", Key: sarama.StringEncoder("Update_Customer_Preference"), Value: sarama.StringEncoder(nd)}:
			fmt.Println("> message queued")
		case err := <-producer.Errors():
			panic(err.Err)
	}
}



