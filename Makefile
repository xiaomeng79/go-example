#定义变量

#project:user account auth
#type:api srv web

.PHONY : fmt
fmt :
	@echo "格式化代码"
	@gofmt -l -w ./

.PHONY : vendor
vendor :
	@echo "创建vendor"
	@go mod vendor
	@echo "结束vendor"

.PHONY : test
test :
	@echo "检查代码"
	@go vet  ./...
	@echo "测试代码"
	@go test -mod=vendor -race -coverprofile=coverage.txt -covermode=atomic ./...


.PHONY : build
build : proto dockerfile builddata
	@echo "部分编译开始:"$(project)_$(type)
	@chmod +x ./scripts/build.sh && ./scripts/build.sh build $(type) $(project)
	@echo "部分编译结束"



.PHONY : allbuild
allbuild : proto alldockerfile  builddata

	@echo "全部编译开始"
	@chmod +x ./scripts/build.sh && ./scripts/build.sh allbuild
	@echo "全部编译结束"



#生成pb文件

.PHONY : proto
proto :

	@echo "生成proto开始"
	@chmod +x ./scripts/proto.sh && ./scripts/proto.sh
	@echo "生成proto结束"

#生成dockerfile
.PHONY : dockerfile
dockerfile :

	@echo "部分生成dockerfile开始"
	@chmod +x ./scripts/dockerfile.sh && ./scripts/dockerfile.sh df $(project) $(type)
	@echo "部分生成Dockerfile结束"


.PHONY : alldockerfile
alldockerfile :

	@echo "全部生成dockerfile开始"
	@chmod +x ./scripts/dockerfile.sh && ./scripts/dockerfile.sh alldf
	@echo "全部生成Dockerfile结束"


#compose命令 bin:up stop restart kill rm ps
.PHONY : compose
compose :

	@chmod +x ./scripts/docker-compose.sh && ./scripts/docker-compose.sh $(bin)


.PHONY : builddata
builddata :

	@chmod +x ./scripts/builddata.sh && ./scripts/builddata.sh

#编辑k8s配置
.PHONY : k8sconfig
k8sconfig :

	@echo "配置k8s"
	@chmod +x k8sconf.sh && ./k8sconf.sh