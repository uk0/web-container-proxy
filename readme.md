# 这是一个简单的 静态服务器，自带反向代理.


### quick

  * 1，如何配置
  ```bash
    {
        "firsh.cc": {# 这个url需要在 host文件里面修改一下下面会讲
        "/api": "localhost:8005",  # 需要代理的服务器地址 与mock一致即可 忽略POST，GET DELETE等方法
        "*": "/Users/zhangjianxin/home/GO_LIB/src/github.com/breakEval13/libs/lib/staticweb/root"
        # 上面的文件路径为dist目录
        }
    }
  ```


  * 2 host如何进行映射

  ```bash

➜  Desktop cat /etc/hosts
##
# Host Database
#
127.0.0.1    firsh.cc # 这里

  ```
  * windows 同理进行修改hosts文件



### 如何启动

  * start
  ```bash
    #windows 直接双击
    #mac or linux
    chmod +x  xxxx
    ./xxx #即可
    
  ```
  * 配置文件和可执行文件放在统一目录(目录级别一致)))
    

  * 需要修改代理的服务器IP直接编辑文件即可,程序会自动更新代理。



###  演示
  * GIF

  ![](http://zmatsh.b0.upaiyun.com/demos/1f4abb63-ed08-4004-a612-427a32ee6fb5.gif)



## ！！！小白手段，大佬勿喷。

