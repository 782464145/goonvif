# Goonvif
易于管理IP设备（包括摄像机）。 Goonvif是用于管理IP设备的ONVIF协议的实现。该库的目的是方便，轻松地控制IP摄像机和其他支持ONVIF标准的设备。

## 安装
要安装该库，您需要使用go get实用程序：
```
go get github.com/782464145/goonvif
```
## 支持的服务
以下服务已完全实现：
- Device
- Media
- PTZ
- Imaging

## 使用

### 一般概念
1）连接到设备
2）身份验证（如有必要）
3）数据类型的定义
4）执行所需的方法

#### 设备连接
如果设备位于网络上的地址*192.168.13.42*，并且其ONVIF服务使用*1234* 端口，则可以通过以下方式连接到设备：
```
dev, err := goonvif.NewDevice("192.168.13.42:1234")
```

*ONVIF端口可能因设备而异，要确定要使用哪个端口，可以转到设备的Web界面。 **通常这是端口80。**

#### 认证方式
如果ONVIF服务之一的任何功能需要认证，则必须使用`Authenticate`方法。
```
device := onvif.NewDevice("192.168.13.42:1234")
device.Authenticate("username", "password")
```

#### 数据类型的定义
该库中的每个ONVIF服务都有其自己的程序包，其中定义了该服务的所有数据类型，并且程序包名称与服务名称相同，并以大写字母开头。
Goonvif为该库支持的每个ONVIF服务的每个功能定义结构。
定义设备服务的GetCapabilities函数的数据类型。这样做如下：
```
capabilities := Device.GetCapabilities{Category:"All"}
```
为什么GetCapabilities结构具有“类别”字段，并且此字段的值为何为All？

下图显示了[GetCapabilities]函数的文档（https://www.onvif.org/ver10/device/wsdl/devicemgmt.wsdl）。可以看出该函数接受一个Category参数，并且其值必须是以下之一：``All``，``Analytics``，``Device``，``Events``，``Imaging``，``Media``或``PTZ``。 

![Device GetCapabilities](img/exmp_GetCapabilities.png)

确定[PTZ]服务的GetServiceCapabilities函数的数据类型的示例(https://www.onvif.org/ver20/ptz/wsdl/ptz.wsdl):
```
ptzCapabilities := PTZ.GetServiceCapabilities{}
```
下图显示GetServiceCapabilities不接受任何参数。 

![PTZ GetServiceCapabilities](img/GetServiceCapabilities.png)

*常见的数据类型在xsd /onvif包中。所有服务可以共享的数据类型（结构）在onvif包中定义。

确定[Device]服务的CreateUsers函数的数据类型的示例(https://www.onvif.org/ver10/device/wsdl/devicemgmt.wsdl):
```
createUsers := Device.CreateUsers{User: onvif.User{Username:"admin", Password:"qwerty", UserLevel:"User"}}
```

下图显示，在此示例中，CreateUsers结构的字段应为User，其数据类型为User结构，其中包含Username，Password，UserLevel字段和可选的Extension。用户结构在onvif包中。

![Device CreateUsers](img/exmp_CreateUsers.png)

#### 执行所需的方法
要执行已定义其结构的ONVIF服务之一的任何功能，必须使用设备对象的“ CallMethod”。
```
createUsers := Device.CreateUsers{User: onvif.User{Username:"admin", Password:"qwerty", UserLevel:"User"}}
device := onvif.NewDevice("192.168.13.42:1234")
device.Authenticate("username", "password")
resp, err := dev.CallMethod(createUsers)
```
