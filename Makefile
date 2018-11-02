#定义变量
project_name=user_srv


.PHONY : fmt
fmt :
	@echo "格式化代码"
	@gofmt -l -w ./



.PHONY : test
test :
	@echo "检查代码"
	@go vet ./...
	@echo "测试代码"
	@go test -race -coverprofile=coverage.txt -covermode=atomic ./...


#project:user account auth
#type:api srv web

.PHONY : build
build :
	@echo "部分编译开始:"$(project)_$(type)
	@chmod +x ./scripts/build.sh && ./scripts/build.sh build $(project) $(type)
	@echo "部分编译结束"



.PHONY : allbuild
allbuild :

	@echo "全部编译开始"
	@chmod +x ./scripts/build.sh && ./scripts/build.sh allbuild
	@echo "全部编译结束"



#生成pb文件

#生成pb文件

.PHONY : proto
proto :

	@echo "生成proto开始"
	@chmod +x ./scripts/proto.sh && ./scripts/proto.sh
	@echo "生成proto结束"

#生成dockerfile
.PHONY : dockerfile
dockerfile :

	@echo "开始生成dockerfile"
	@echo "FROM alpine:3.2" >$(docker_file_name)
	@echo "RUN set -xe && apk add --no-cache tzdata && cp -r -f /usr/share/zoneinfo/Asia/Shanghai /etc/localtime" >>$(docker_file_name)
	@echo "ADD $(project_name) /$(project_name)" >>$(docker_file_name)
	@echo "ENTRYPOINT [ "/$(project_name)" ] " >>$(docker_file_name)
	@echo "生成Dockerfile"

#编辑k8s配置
.PHONY : k8sconfig
k8sconfig :

	@echo "配置k8s"
	@chmod +x k8sconf.sh && ./k8sconf.sh