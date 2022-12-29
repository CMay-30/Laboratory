## 一、配置环境

### 1.安装VS Code

官网下载 https://code.visualstudio.com/ 下载VS Code，按照步骤安装。

### 2.安装node.js

1. 官网  https://nodejs.org/en/  下载 node.js ，按照步骤安装即可， node.js 安装完成之后会同步安装npm。

2. 配置环境变量：把 node.js 安装路径配置到环境变量Path中，使用 node.js 安装包安装时一般会自动添加环境变量。

3. 查看 node.js 是否安装成功：打开 cmd ，输入 `node -v`  和  `npm -v`  如果显示版本信息，则说明安装成功。

   ![1672212695698](C:\Users\tanwen\AppData\Roaming\Typora\typora-user-images\1672212695698.png)

### 3.安装配置脚手架vue-cli

脚手架可以帮助我们快速配置项目，打开 VSCode 的终端，输入下述命令，回车，等待安装完成。

```npm
npm install -g vue-cli
```

![1672213255441](C:\Users\tanwen\AppData\Roaming\Typora\typora-user-images\1672213255441.png)



> npm默认从[npm官网](https://registry.npmjs.org/)（国外网站）下载模块，下载速度特别慢。
>
> 解决方法：使用国内镜像，打开 cmd  ，输入下述命令，回车。
>
> ```npm
> npm config set registry http://registry.npmmirror.com
> ```
>
> 验证是否配置成功：
>
> ```npm
> npm config get registry
> ```

## 二、创建vue项目

### 1.命令方式创建

同样也是在 VSCode 终端，输入 `vue init webpack laboratory`   ，意思是初始化一个项目，其中 webpack 是构建工具，也就是整个项目是基于 webpack 的， Laboratory_System是整个项目文件夹的名称，可自行更改。

```
vue init webpack laboratory
```

![1672214691827](C:\Users\tanwen\AppData\Roaming\Typora\typora-user-images\1672214691827.png)

![1672214721872](C:\Users\tanwen\AppData\Roaming\Typora\typora-user-images\1672214721872.png)

### 2.重新初始化依赖

1. 按照上图的提示，cd 到刚才项目目录
2. 执行 `npm cache clean --force` 清除缓存
3. 执行 `npm install` 重新初始化依赖。

### 3.启动项目

1. 打开项目里面的 package.json ，在 VSCode 终端执行 start 中的命令 `npm run dev` ，启动成功后会出现访问地址。

   ![1672215439872](C:\Users\tanwen\AppData\Roaming\Typora\typora-user-images\1672215439872.png)

   ![1672215557697](C:\Users\tanwen\AppData\Roaming\Typora\typora-user-images\1672215557697.png)

2. 根据提示，访问[http://localhost:8080](http://localhost:8080/)，会访问到如下界面。至此，VS Code创建Vue.js项目已经完成。

   ![1672215583340](C:\Users\tanwen\AppData\Roaming\Typora\typora-user-images\1672215583340.png)