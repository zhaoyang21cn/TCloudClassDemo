# 空中课堂集成化demo
1. 在运行该demo前，先向空中课堂接口人申请接入密钥secretid和secretkey，并在src/class-demo文件中替换
2. utils文件内的代码是辅助方法
3. 空中课堂云API接口调用方法在class-demo中，步骤如下：
	
	第一步是计算签名方法`calcAuthorization`，
	
	第二步在发起请求前需要计算密钥签名，并将其写入请求header中`req.Header.Set("Authorization", sign)`
	
	第三步根据API文档请求不同的接口连接
4. 运行测试，在本地运行成功后服务会监听8080端口，代码默认为8080端口，可自己修改。main.go中写了对外路由，本地测试可根据路由请求，例如创建学习接口请求为：
	`http://127.0.0.1:8080/demo/v1/new_enter/create`