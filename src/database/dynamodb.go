package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// DynamoDB base Structure
type DynamoDB struct {
	tableName string
	sess      *dynamodb.DynamoDB
}

// DynamoDBAccessor gets a new instance
func DynamoDBAccessor() *DynamoDB {

	dbName := os.Getenv("TEST_DB")
	if dbName == "" {
		dbName = "bodyweight"
	}
	log.Println("Instantiating Table:", dbName)
	db := &DynamoDB{
		tableName: dbName,
		sess:      getDynamoDBSession(),
	}
	db.CreateDBIfNotExists()
	return db
}

// CreateDBIfNotExists creates a new DB
func (d DynamoDB) CreateDBIfNotExists() error {
	if !d.existsDB() {
		log.Println("Creating DB-Table: ", d.tableName)

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
			TableName: aws.String(d.tableName),
		}

		_, err := d.sess.CreateTable(input)

		if err != nil {
			log.Println("Got error calling CreateTable:")
			log.Println(err.Error())
			return err
		}

		d.waitForTableState("ACTIVE", 10, false)

	} else {
		log.Println("DB already exists")
	}
	return nil

}

// DeleteDB deletes a DB
func (d DynamoDB) DeleteDB() error {
	log.Println("Deleting Table:", d)
	_, err := d.sess.DeleteTable(&dynamodb.DeleteTableInput{
		TableName: aws.String(d.tableName),
	})

	if err != nil {
		log.Println("Error: ", err)
		return err
	}

	d.waitForTableState("DELETING", 100, true)

	return nil

	// TODO: wait for Table-State
}

// CreateEntry creates an Entry
func (d DynamoDB) CreateEntry(entry *Entry) error {

	av, err := dynamodbattribute.MarshalMap(entry)
	if err != nil {
		log.Println("Error Marshalling: ", err)
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(d.tableName),
	}

	_, err = d.sess.PutItem(input)

	if err != nil {
		log.Println("Got error calling PutItem:", err)
		return err
	}
	log.Println("Entry created: ", entry.AlexaID, entry.Date, entry.UserName)
	return nil
}

// GetLastUsedEntry gets the last used Entry
func (d DynamoDB) GetLastUsedEntry(alexaID string) (*Entry, error) {
	entries, err := d.GetEntries(alexaID)
	if err != nil {
		return nil, err
	}
	if entries == nil {
		return nil, nil
	}

	var latest = (*entries)[0]
	for _, entry := range *entries {
		if entry.Date.UnixNano() > latest.Date.UnixNano() {
			latest = entry
		}
	}
	return &latest, nil
}

// GetEntries gets all Entries
func (d DynamoDB) GetEntries(alexaID string) (*[]Entry, error) {
	result, err := d.sess.Scan(&dynamodb.ScanInput{
		TableName:        aws.String(d.tableName),
		FilterExpression: aws.String("alexa_id = :u"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":u": {
				S: aws.String(alexaID),
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
			log.Printf("Error:%+v\n", err)
			return nil, err
		}

		return &entries, nil
	}
	return nil, nil // no entries
}

// DeleteItem deletes a Item
func (d DynamoDB) DeleteItem(entry PK) error {
	av, err := dynamodbattribute.MarshalMap(entry)
	if err != nil {
		log.Println(err)
		return err
	}

	r, err := d.sess.DeleteItem(&dynamodb.DeleteItemInput{
		Key:          av,
		TableName:    aws.String(d.tableName),
		ReturnValues: aws.String("ALL_OLD"),
	})
	if err != nil {
		log.Println("Error:", err)
		return err
	}

	e := Entry{}
	err = dynamodbattribute.Unmarshal(&dynamodb.AttributeValue{M: r.Attributes}, &e)
	if err != nil {
		log.Println("error unmarshalling:", err)
		return err
	}

	log.Println("Entry deleted:", e.AlexaID, e.Date, e.UserName)
	if e.AlexaID != entry.AlexaID || e.Date != entry.Date {
		return fmt.Errorf("unable to delete: %s, %v", entry.AlexaID, entry.Date)
	}
	return nil
}

//DeleteAllEntries deletes all entries
func (d DynamoDB) DeleteAllEntries(alexaID string) error {
	entries, err := d.GetEntries(alexaID)

	if err != nil {
		return err
	}

	if entries != nil {
		for _, entry := range *entries {
			err := d.DeleteItem(entry.PK)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (d DynamoDB) waitForTableState(desiredState string, timeout int, until bool) {
	var state string
	waitfn := func() {
		log.Println("we have to wait for table state:", desiredState)

		time.Sleep(1 * time.Second)
		out, err := d.sess.DescribeTable(&dynamodb.DescribeTableInput{
			TableName: aws.String(d.tableName),
		})
		if err != nil {
			log.Println("we got an error: ", err)
			state = err.Error()
			return
		}

		state = *out.Table.TableStatus

		if timeout--; timeout < 0 {
			log.Panic("Timeout Table State:", desiredState)
		}
	}

	if until {
		state = desiredState
		for state == desiredState {
			waitfn()
		}
	} else {
		for state != desiredState {
			waitfn()
		}
	}

}

func (d DynamoDB) existsDB() bool {

	result, err := d.sess.ListTables(&dynamodb.ListTablesInput{})

	if err != nil {
		log.Println(err)
		return false
	}
	for _, n := range result.TableNames {
		if *n == d.tableName {
			return true
		}
	}
	return false
}

func getDynamoDBSession() *dynamodb.DynamoDB {
	sess, err := session.NewSession(&aws.Config{
		// LogLevel: aws.LogLevel(aws.LogDebugWithHTTPBody),
		Region: aws.String("eu-west-1")},
	)

	if err != nil {
		log.Println(err.Error())
		return nil
	}

	// Create DynamoDB client
	return dynamodb.New(sess)

}
