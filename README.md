# go-gitee #

go-gitee is a Go client library for accessing the [Gitee API v5](https://gitee.com/api/v5/swagger).

gitee 的 golang 版本的 API 实现 模仿go-github实现的



# API 文档

* API 文档
    * [动态通知(Activity)]()
    * [邮箱(Emails)]()
    * [企业(Enterprises)]()
    * [任务(Issues)]()
    * [标签(Labels)]()
    * [里程碑(Milestones)]()
    * [杂项(Miscellaneous)]()
    * [组织(Organizations)]()
    * [PR操作(Pull Requests)]()
    * [仓库(Repositories)]()
    * [搜索(Search)]()
    * [用户账号(Users)]()
    * [钩子(Webhooks)]()


# TODO


大部分 xxxOptions 传给 addOpttins 用的, 参数都是 指针. 结构体里面 都是 普通类型，拼接url使用的。


大部分 xxxRequest 传给 NewRequest 用的, 参数也是 指针. 结构体里面 都是 指针类型， json 的。   有时候传给 NewRequest 也叫什么什么 Options

后续希望统一一下命名，统一一下指针还是普通类型

