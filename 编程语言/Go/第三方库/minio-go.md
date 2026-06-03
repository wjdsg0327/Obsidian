```go
package main  
  
import (  
    "context"  
    "fmt"    "github.com/minio/minio-go/v7"    "github.com/minio/minio-go/v7/pkg/credentials"    "log")  
  
func main() {  
    // 1. 自建环境配置  
    endpoint := "192.168.1.80:9000" // 或者 "minio.your-company.com"    accessKey := "admin"            // 你自建时设置的 Access Key    secretKey := "wjdsgMinio"       // 你自建时设置的 Secret Key  
    // 2. SSL 开关  
    // 如果你没有配置 HTTPS 证书（即通过 http:// 访问），这里必须设为 false    useSSL := false  
  
    // 初始化客户端  
    minioClient, err := minio.New(endpoint, &minio.Options{  
       Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),  
       Secure: useSSL,  
    })  
  
    if err != nil {  
       fmt.Println("初始化失败:", err)  
       return  
    }  
  
    fmt.Printf("成功连接到自建 MinIO: %s\n", endpoint)  
  
    // 上传文件示例  
    bucketName := "wjdsgbiji"  
    objectName := "python310.dll"  
    filePath := "./python310.dll"  
  
    info, err := minioClient.FPutObject(context.Background(), bucketName, objectName, filePath, minio.PutObjectOptions{})  
    if err != nil {  
       log.Fatalln(err)  
    }  
    log.Printf("成功上传: %v", info)  
}
```