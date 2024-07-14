# SSL CA
这是一个SSL免费证书申请工具，通过调用Let's Encrypt的API来申请SSL证书，支持自动续期。
为了方便使用，我们提供了操作lego来申请和续约，也集成了部分效验和操作证书的方法。

## 调用方法：
    1. 设置证书文件的存放路径
        sslca.SetCertFile("cert.pem")
    2. 获取证书信息
        status, info := sslca.GetCertInfo()
        if status == 0 { // 成功
            fmt.Println(info)
        } else {	// 失败
            fmt.Println(sslca.ErrMsg(status))
        }
## update:
        2021-08-24: v0.01