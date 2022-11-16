# go-gitee #

go-gitee is a Go client library for accessing the [Gitee API v5](https://gitee.com/api/v5/swagger).

gitee 的 golang 版本的 API 实现 模仿go-github实现的



# API 文档

* API 文档
    * [动态通知(Activity)]()
    * [邮箱(Emails)](gitee/miscs.go) 接口全部实现
    * [企业(Enterprises)]()
    * [任务(Issues)]()
    * [标签(Labels)]()
    * [里程碑(Milestones)]()
    * [杂项(Miscellaneous)](gitee/miscs.go) 接口全部实现
    * [组织(Organizations)]()
    * [PR操作(Pull Requests)]()
    * [仓库(Repositories)](gitee/repos.go) 接口全部实现
    * [搜索(Search)]()
    * [用户账号(Users)](gitee/users.go) 接口全部实现
    * [钩子(Webhooks)]()


# TODO


大部分 xxxOptions 传给 addOpttins 用的, 参数都是 指针. 结构体里面 都是 普通类型，拼接url使用的。

大部分 xxxRequest 传给 NewRequest 用的, 参数也是 指针. 结构体里面 都是 指针类型， json 的。   有时候传给 NewRequest 也叫什么什么 Options



# DONE

后续希望统一一下命名，统一一下指针还是普通类型.统一命名，传给 addOpttins 用的 都叫 xxxOptions, 传给 NewRequest 用的 都叫 xxxRequest.,这个基本上改造完成


结构体里面字段的命名, 统一一下:url这种缩写 都改成全大写的 URL, htmlurl 这种也改成全大写的 HTMLURL,这个基本上改造完成


# License
```
Copyright magesfc bright.ma

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```
