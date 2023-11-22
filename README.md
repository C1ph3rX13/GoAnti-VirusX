# GoAnti-VirusX
Go Anti-Virus Module Collection

## Description

本项目会不断添加各种免杀的技术，但是**不适合直接不做任何修改的编译和使用**，即使是有随机特征的编译

### 特别说明

1. 学习Go免杀的代码集合，顺手做了模块化处理，**实际开发未结束，还在持续更新**；

2. 实际使用效果，国内杀软都可以过；

3. 想要实现最好的免杀效果还需要自行修改渲染编译模板，代码提供了三种加载方式的渲染模板；

4. 如果只想获取免杀的最终结果，只需要关注下列文件夹，直接执行即可在`compilation`文件夹获取已编译的免杀可执行文件。

   ```go
   // 生成 payload.bin 文件放置在对应文件夹下运行即可编译 - 建议导入IDE运行
   basetmpl：baseBuilder.go
   
   remotetmpl：remoteBuilder.go
   
   localtmpl：localBuilder.go
   ```

## Module

```go
│  go.mod
│  go.sum
│  main.go
│
├─AntiSandBox // 反沙箱
│      BootTime.go
│      CheckLanguage.go
│      CheckVirtual.go
│      CheckWallPaperMD5.go
│      CheckWeChatExist.go
│      CompareProcessNames.go
│      GetCountOfCPU.go
│      GetDesktopFiles.go
│      GetPhysicalMemory.go
│      GetTimeZone.go
│
├─API // Windows API  
│      WindowsAPI.go
│
├─Build // 模板渲染和编译
│  ├─basetmpl // 本地加载模板渲染编译
│  │      baseBuilder.go
│  │      baseTmpl.tmpl.bak
│  │      payload.bin
│  │
│  ├─compilation // 最终编译输出
│  │  └─winres
│  │          winres.json
│  │
│  ├─compileutils
│  │      encwriter.go
│  │      encwriter.go.bak
│  │      exeutils.go
│  │      getfilesutils.go
│  │      rawreader.go
│  │      strwriter.go
│  │      tmplbuilder.go
│  │      winresutils.go
│  │
│  ├─localtmpl // 分离加载模板渲染编译
│  │      localBuilder.go
│  │      localTmpl.tmpl
│  │      payload.bin
│  │
│  ├─remotetmpl // 远程加载模板渲染编译
│  │      payload.bin
│  │      remoteBuilder.go
│  │      remoteTmpl.tmpl
│  │
│  └─selectfunctmpl // 加密输出模板渲染 - 未完成（已实现可选加密的加载器模板渲染）
│          encwriter.tmpl
│          exportfunc.go
│          selectdec.go
│          selectenc.go
│          selectfunc.tmpl
│
├─Crypto // 加密方法
│      AES.go
│      Base85.go
│      DES.go
│      RC4.go
│      SM4.go
│      XOR.go
│      XorBase64.go
│      XorRC4Base85.go
│      XorRC4HexBase85.go
│
├─Loader // shellcode加载器
│  │  BaseLoader.go
│  │  CreateRemoteThreadNativeLoader.go // Windows 10 测试不成功，会造成桌面重启
│  │  CreateThreadLoader.go
│  │  EtwpCreateEtwThreadLoader.go
│  │  FiberLoader.go
│  │  NtQueueApcThreadExLoader.go
│  │  RemoteLibs.go
│  │  StaneAloneLoader.go
│  │  UuidFromStringLoader.go
│  │
│  └─halosgate
│          asm_x64.s
│          Gate.go
│          GateLoader.go
│
└─Utils // 全局工具类
        deleteutil.go
        logconfig.go
        randomutil.go
        randpnghash.go
```

### Thanks

https://github.com/Ne0nd0g/go-shellcode

https://github.com/safe6Sec/GolangBypassAV

https://github.com/afwu/GoBypass

https://github.com/piiperxyz/AniYa

https://github.com/wumansgy/goEncrypt

https://github.com/TideSec/GoBypassAV

https://github.com/Pizz33/GobypassAV-shellcode

https://github.com/timwhitez/Doge-Gabh

后继添加 ……
