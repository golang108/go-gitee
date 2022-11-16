//Copyright magesfc bright.ma
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

package test

import (
	"fmt"
	"github.com/mamh-mixed/go-gitee/gitee"
	"testing"
)

func TestGetUser(t *testing.T) {
	user, response, err := client.Users.Get(ctx, "")
	fmt.Println(user)
	fmt.Println(response)
	fmt.Println(err)
}

func TestListSshKeys(t *testing.T) {
	var opts = &gitee.ListOptions{
		Page:    1,  // 这个key目前只有2个，这里一页就能获取全部的了
		PerPage: 10, // perPage 表示每页的总数
	}
	for {
		keys, response, err := client.Users.ListKeys(ctx, "", opts)
		if err != nil {
			fmt.Println(err)
			return
		}
		for index, key := range keys {
			fmt.Println(index, len(keys), *key.ID, *key.Title)
		}
		if response.NextPage == 0 {
			break
		}
		opts.Page = response.NextPage
	}
}

func TestGetSshKey(t *testing.T) {
	keys, response, err := client.Users.GetKey(ctx, 3544397)
	fmt.Println(keys)
	fmt.Println(response)
	fmt.Println(err)
}

func TestListFollowers(t *testing.T) {
	var opts = &gitee.ListOptions{
		Page:    90,  // page 表示从第几页 开始，一般从第 1 页开始，然后第 2 页，然后第 3 页，到最后一页
		PerPage: 100, // perPage 表示每页的总数
	}
	for { //分页 循环 获取 所有的，这里有 几百个值的，需要循环的
		users, response, err := client.Users.ListFollowers(ctx, "y_project", opts)
		if err != nil {
			fmt.Println(err)
			return
		}
		for index, user := range users {
			fmt.Println(index, "len: ", len(users), response.NextPage, *user.ID, *user.Login, *user.Name)
		}

		if response.NextPage == 0 {
			break
		}
		opts.Page = response.NextPage
	}

}

func TestListFollowers1(t *testing.T) {
	var opts = &gitee.ListOptions{
		Page:    1,
		PerPage: 10,
	}
	users, response, err := client.Users.ListFollowers(ctx, "", opts)
	fmt.Println(users)
	fmt.Println(response)
	fmt.Println(err)
}

func TestListFollowings(t *testing.T) {
	var opts = &gitee.ListOptions{
		Page:    1,
		PerPage: 10,
	}
	// 这里获取mamh这个账号关注了哪几个人
	users, response, err := client.Users.ListFollowings(ctx, "mamh", opts)
	fmt.Println(users)
	fmt.Println(response)
	fmt.Println(err)
}

func TestGetUserFollowings1(t *testing.T) {
	var opts = &gitee.ListOptions{
		Page:    1,
		PerPage: 10,
	}
	// user参数设置空字符串，就是 当前授权账号 关注了哪几个人
	users, response, err := client.Users.ListFollowings(ctx, "", opts)
	fmt.Println(users)
	fmt.Println(response)
	fmt.Println(err)
}

func TestListNamespaces(t *testing.T) {
	var opts = &gitee.NamespacesOptions{
		Mode: "project",
	}
	names, response, err := client.Users.ListNamespaces(ctx, opts)
	fmt.Println(names) //  project 类型的，只会获取2组数据
	fmt.Println(response)
	fmt.Println(err)
}

func TestGetNamespace(t *testing.T) {
	var opts = &gitee.NamespaceOptions{
		Path: "mamh-java",
	}
	names, response, err := client.Users.GetNamespace(ctx, opts)
	fmt.Println(names)
	fmt.Println(response)
	fmt.Println(err)
}

func TestGetNamespace1(t *testing.T) {
	var opts = &gitee.NamespaceOptions{
		Path: "mamh",
	}
	names, response, err := client.Users.GetNamespace(ctx, opts)
	fmt.Println(names)
	fmt.Println(response)
	fmt.Println(err)
}

func TestFollow(t *testing.T) {
	user := "y_project"
	response, err := client.Users.Follow(ctx, user)
	fmt.Println(response)
	fmt.Println(err)
}
