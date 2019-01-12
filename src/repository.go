package main

import (
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

//TODO: abstrakte Schicht machen, die unabhängig von Dynamo DB ist

const (
	tableName = "bodyweight"
)

var (
	sess *dynamodb.DynamoDB
)

type Entry struct {
	AlexaID       string        `json:"alexa_id"`
	Date          time.Time     `json:"date"`
	UserName      string        `json:"username"`
	TrainingState TrainingState `json:"training_state"`
	Desc          string        `json:"desc"`
}

func getSession() *dynamodb.DynamoDB {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("eu-west-1")},
	)

	if err != nil {
		log.Println(err.Error())
		return nil
	}

	// Create DynamoDB client
	return dynamodb.New(sess)

}

func getLastUsedEntry(alexa_id string) *Entry {
	entries := getEntries(alexa_id)
	if entries == nil {
		return nil
	}

	var latest = (*entries)[0]
	for _, entry := range *entries {
		if entry.Date.UnixNano() > latest.Date.UnixNano() {
			latest = entry
		}
	}
	return &latest
}

func getEntries(alexa_id string) *[]Entry {

	result, err := sess.Scan(&dynamodb.ScanInput{
		TableName:        aws.String(tableName),
		FilterExpression: aws.String("alexa_id = :u"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":u": {
				S: aws.String(alexa_id),
			},
		},
	})
	// something found?
	if len(result.Items) > 0 {
		var entries = []Entry{}
		for _, item := range result.Items {
			e := Entry{}
			dynamodbattribute.Unmarshal(&dynamodb.AttributeValue{M: item}, &e)
			entries = append(entries, e)
		}

		if err != nil {
			fmt.Printf("Error:%+v\n", err)
			return nil
		}

		return &entries
	}
	return nil // no entries
}

func deleteItem(alexaID string, date time.Time) error {
	delitem := &dynamodb.DeleteItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"alexa_id" : {
				S:aws.String(alexaID),
			},
			"date" : {
				S:aws.String(date.Format("2006-01-02T15:04:05.999999-07:00")),
			},
		},
	}

	_, err := sess.DeleteItem(delitem)
	if err != nil {
		log.Println("Error:", err)
		return err
	}
	return nil
}

func deleteAllEntries(alexaID string) error {
	entries:= getEntries(alexaID) 
	if entries!=nil {
		for _,entry:=range *entries{
			err:=deleteItem(entry.AlexaID, entry.Date)
			if err!=nil {
				return err
			}
		}
	}
	return nil
}

func createEntry(alexa_id string, name string, training TrainingState, desc string) error {

	entry := Entry{
		AlexaID:       alexa_id,
		Date:          time.Now(),
		UserName:      name,
		TrainingState: training,
		Desc:          desc,
	}

	av, err := dynamodbattribute.MarshalMap(entry)
	if err != nil {
		log.Println("Error Marshalling: ", err)
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	_, err = sess.PutItem(input)

	if err != nil {
		fmt.Println("Got error calling PutItem:", err)
		return err
	}
	log.Println("Entry created: ", entry.AlexaID)
	return nil
}

func existsDB(name string) bool {

	result, err := sess.ListTables(&dynamodb.ListTablesInput{})

	if err != nil {
		log.Println(err)
		return false
	}
	for _, n := range result.TableNames {
		if *n == name {
			return true
		}
	}
	return false
}

func createDBIfNotExists() {
	if !existsDB(tableName) {
		log.Println("Creating DB-Table: ", tableName)

		input := &dynamodb.CreateTableInput{
			AttributeDefinitions: []*dynamodb.AttributeDefinition{
				{
					AttributeName: aws.String("alexa_id"),
					AttributeType: aws.String("S"),
				},
				{
					AttributeName: aws.String("date"),
					AttributeType: aws.String("S"),
				},
			},
			KeySchema: []*dynamodb.KeySchemaElement{
				{
					AttributeName: aws.String("alexa_id"),
					KeyType:       aws.String("HASH"),
				},
				{
					AttributeName: aws.String("date"),
					KeyType:       aws.String("RANGE"),
				},
			},
			ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
				ReadCapacityUnits:  aws.Int64(5),
				WriteCapacityUnits: aws.Int64(5),
			},
			TableName: aws.String(tableName),
		}

		_, err := sess.CreateTable(input)

		if err != nil {
			log.Println("Got error calling CreateTable:")
			log.Println(err.Error())
			return
		}
	} else {
		log.Println("DB already exists")
	}

}

func init() {
	sess = getSession()
	createDBIfNotExists()
}
