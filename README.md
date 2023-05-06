# ChatGPT-Plus

基于 OpenAI API 实现的 ChatGPT Web 应用，一共分为两个版本：

* 通用版：交互体验跟 ChatGPT 官方一致，聊天记录保存在客户端（浏览器）
* 角色版：内置了各种预训练好的角色，比如小红书写手，英语翻译大师，苏格拉底，孔子，乔布斯，周报助手等。轻松满足你的各种聊天和应用需求。
聊天记录保存在云端（可以配置是否保存聊天记录）
每个版本都有 PC 版和移动版，PC 版本的体验要略优于移动版。

**本项目基于 MIT 协议，免费开放全部源代码，可以作为个人学习使用或者商用。如需商用建议联系作者登记，仅做统计使用，优秀项目我们将在项目首页为您展示。**

## 项目介绍
这一套完整的系统，包括两套前端聊天应用和一个后台管理系统。系统有用户鉴权，你可以自己使用，也可以部署直接给 C 端用户提供 ChatGPT 的服务。

项目的技术架构是

> Go + Vue3 + element-plus

后端采用的是 Go 语言开发的 Gin Web 框架。前端用的是 Vue3 + element-plus UI 框架

目前已经实现了以下功能：
1. 通用版的 ChatGPT 聊天界面和功能，聊天记录保存在客户端。
2. 口令机制：输入口令才可以访问，支持设置口令的对话次数，有效期。
3. 角色版的聊天界面和功能，角色设定，预设一些角色，比如程序员，客服，作家，老师，艺术家... 
4. 保存聊天记录，支持聊天上下文。 
5. OpenAI API 负载均衡，限制每个 API Key 每分钟之内调用次数不超过 15次，防止被封。 
6. 支持配置多个代理，保证高可用。 
7. 实现 markdown 语法解析和代码高亮，支持复制回复内容功能。 
8. 后台管理功能，实现系统的动态配置，用户和角色的动态管理。

## 功能截图

### 1. 角色版PC端
![ChatGPT Plus](docs/imgs/chat-plus.png)

![ChatGPT admin](docs/imgs/gpt.gif)

### 2. 角色版移动端
![ChatGPT Mobile](docs/imgs/chat-mobile.png)

### 3. 通用版
![ChatGPT Free](docs/imgs/chat-free.png)

### 4. 管理后台
![ChatGPT admin](docs/imgs/admin-01.png)

![ChatGPT admin](docs/imgs/admin-02.png)

![ChatGPT admin](docs/imgs/admin-03.png)

### 5. 体验地址
> 体验地址：[https://www.chat-plus.net/chat/#/free](https://www.chat-plus.net/chat/#/free) </br>
> 口令：GeekMaster

## TODOLIST
* [ ] 让用户配置自己的 API KEY，调用自己的 API Key，将不记 Token 的使用次数
* [ ] 嵌入 AI 绘画功能，支持根据描述词生成图片
* [ ] 接入自己训练的开源大语言模型
* [ ] 接入 Google 语音 API，支持语音聊天

## 本地部署
## 线上发布

## 注意事项

## 参与贡献


