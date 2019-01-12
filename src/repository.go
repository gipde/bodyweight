package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
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

type EntryType int

func (t EntryType) Str() string {
	return strconv.Itoa(int(t))
}

const (
	LOG_ENTRY EntryType = iota
	STATE_ENTRY
	START_ENTRY
)

type Entry struct {
	AlexaID       string        `json:"alexa_id"`
	Date          time.Time     `json:"date"`
	UserName      string        `json:"username"`
	Type          EntryType     `json:"typ"`
	TrainingState TrainingState `json:"training_state"`
	Desc          string        `json:"desc"`
}

const (
	letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

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

func getLastUsedEntry(uid string) *Entry {
	entries := getEntries(uid, STATE_ENTRY)
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

func getEntries(uid string, typ EntryType) *[]Entry {
	result, err := sess.Scan(&dynamodb.ScanInput{
		TableName:        aws.String("bodyweight"),
		FilterExpression: aws.String("typ = :t AND uid = :u"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":u": {
				S: aws.String(uid),
			},
			":t": {N: aws.String(typ.Str())},
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

func createEntry(uid string, name string, t EntryType, level TrainingLevel, week, day, unit int, desc string) error {

	entry := Entry{
		AlexaID:  uid,
		Date:     time.Now(),
		UserName: name,
		Type:     t,
		TrainingState: TrainingState{
			Level: level,
			Week:  week,
			Day:   day,
			Unit:  unit,
		},
		Desc: desc,
	}
	av, err := dynamodbattribute.MarshalMap(entry)
	log.Printf("Entry: %+v",av)
	if err != nil {
		log.Println("Error Marshalling: ", err)
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("bodyweight"),
	}

	_, err = sess.PutItem(input)

	if err != nil {
		fmt.Println("Got error calling PutItem:", err)
		return err
	}
	log.Println("Entry created: ", entry.AlexaID)
	return nil
}

func genRowId() string {
	result := make([]byte, 3)
	for i := 0; i < 3; i++ {
		result[i] = letterBytes[rand.Intn(52)]
	}
	num := time.Now().Unix()<<12 + rand.Int63n(512)
	return string(result) + strconv.FormatInt(num, 10)
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

var sess *dynamodb.DynamoDB

func initDB() {
	// get DB-Session
	sess = getSession()
	createDBIfNotExists()

}

func init() {
	rand.Seed(time.Now().Unix())
	initDB()
}
