const shell = require("shelljs");
const dist = "./views";
const backendPath = "../backend/";
const backendViewPath = backendPath + "views";
shell.rm("-rf", dist);
shell.rm("-rf", backendViewPath);
const startportaltime = new Date().getTime();
shell.echo("开始打包");
shell.exec("npm run build");
const endportaltime = new Date().getTime();
shell.echo(
	"构建frontend用时：" + (endportaltime - startportaltime) / 1000 + "s"
);
shell.echo("开始拷贝frontend代码");
shell.cp("-R", dist, backendPath);
shell.echo("完成");
