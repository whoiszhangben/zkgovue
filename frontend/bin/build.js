const shell = require("shelljs");
const dist = "./dist";
const backendPath = "../backend/";
const backendViewPath = backendPath + "views";
shell.rm("-rf", dist);
shell.rm("-rf", backendViewPath);
const startportaltime = new Date().getTime();
shell.echo("开始打包前端");
shell.exec("npm run build");
const endportaltime = new Date().getTime();
shell.echo(
	"构建frontend用时：" + (endportaltime - startportaltime) / 1000 + "s"
);
shell.echo("开始拷贝frontend代码");
shell.cp("-R", dist+"/", backendViewPath);
// shell.echo("开始打包后端");
// shell.exec("cd ..");
// shell.exec("cd backend");
// shell.exec("SET CGO_ENABLED=0");
// shell.exec("SET GOOS=linux");
// shell.exec("SET GOARCH=amd64");
// shell.exec("go build ./main.go");
shell.echo("完成");
