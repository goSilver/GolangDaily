package h_adapter

import "fmt"

// ICreateServer 创建主机接口
type ICreateServer interface {
	CreateServer(cpu, mem float64) error
}

// AWSClient aws sdk
type AWSClient struct{}

// RunInstance 	启动实例
func (a *AWSClient) RunInstance(cpu, mem float64) error {
	fmt.Printf("aws client run success, cpu:%f, mem:%f", cpu, mem)
	return nil
}

// AwsClientAdapter aws适配器
type AwsClientAdapter struct {
	Client AWSClient
}

// CreateServer 包一层的启动实例
func (a *AwsClientAdapter) CreateServer(cpu, mem float64) error {
	a.Client.RunInstance(cpu, mem)
	return nil
}

// AliyunClient aliyun sdk
type AliyunClient struct{}

// CreateServer 启动实例
func (a *AliyunClient) CreateServer(cpu, mem int) error {
	fmt.Printf("aliyun client run success, cpu:%d, mem:%d", cpu, mem)
	return nil
}

// AliyunClientAdapter aliyun适配器
type AliyunClientAdapter struct {
	Client AliyunClient
}

// CreateServer 启动实例
func (a *AliyunClientAdapter) CreateServer(cpu, mem float64) error {
	a.Client.CreateServer(int(cpu), int(mem))
	return nil
}
