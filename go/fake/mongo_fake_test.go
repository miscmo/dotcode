package fake

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	//"go.mongodb.org/mongo-driver/x/mongo/driver/topology"
)

// 定义一个模拟的MongoDB客户端
type mockClient struct {
	*mongo.Client
}

// 使用go-mockgen自动生成的模拟方法
func (mc *mockClient) Database(name string, opts ...*options.DatabaseOptions) *mongo.Database {
	return mc.Client.Database(name, opts...)
}

func TestMongoDB(t *testing.T) {
	// 创建一个gomock控制器
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// 创建一个模拟的MongoDB客户端
	mockCli := &mockClient{}

	// 创建一个内存拓扑，作为mockClient的底层驱动程序
	//topology.NewMockServer(ctrl)

	// 将mockClient与内存拓扑连接
	mockCli.Client, _ = mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))

	// 使用模拟方法模拟MongoDB的行为
	db := mockCli.Database("mydb")

	// 在模拟的数据库中插入数据
	_, err := db.Collection("mycollection").InsertOne(context.Background(), bson.M{"name": "John Doe"})
	assert.NoError(t, err)

	// 在模拟的数据库中查找数据
	var result bson.M
	err = db.Collection("mycollection").FindOne(context.Background(), bson.M{"name": "John Doe"}).Decode(&result)
	assert.NoError(t, err)
	assert.Equal(t, "John Doe", result["name"])
}


	func NewInMemoryMongoDB(ctx context.Context, dbName string) (*mongo.Database, error) {
		clientOptions := options.Client().ApplyURI("mongodb://localhost:27017").SetInMemory(true)
		client, err := mongo.NewClient(clientOptions)
		if err != nil {
			return nil, err
		}

		err = client.Connect(ctx)
		if err != nil {
			return nil, err
		}

		db := client.Database(dbName)

		return db, nil
	}