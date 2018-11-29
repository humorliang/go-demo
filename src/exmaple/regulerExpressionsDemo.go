package main

import (
	"regexp"
	"fmt"
)

/*
正则表达式：
单字符：
		.              任意字符（标志s==true时还包括换行符）
        [xyz]          字符族
        [^xyz]         反向字符族
        \d             Perl预定义字符族
        \D             反向Perl预定义字符族
        [:alpha:]      ASCII字符族
        [:^alpha:]     反向ASCII字符族
        \pN            Unicode字符族（单字符名），参见unicode包
        \PN            反向Unicode字符族（单字符名）
        \p{Greek}      Unicode字符族（完整字符名）
        \P{Greek}      反向Unicode字符族（完整字符名）
组合：
		xy             匹配x后接着匹配y
        x|y            匹配x或y（优先匹配x）
重复：
        x*             重复>=0次匹配x，越多越好（优先重复匹配x）
        x+             重复>=1次匹配x，越多越好（优先重复匹配x）
        x?             0或1次匹配x，优先1次
        x{n,m}         n到m次匹配x，越多越好（优先重复匹配x）
        x{n,}          重复>=n次匹配x，越多越好（优先重复匹配x）
        x{n}           重复n次匹配x
        x*?            重复>=0次匹配x，越少越好（优先跳出重复）
        x+?            重复>=1次匹配x，越少越好（优先跳出重复）
        x??            0或1次匹配x，优先0次
        x{n,m}?        n到m次匹配x，越少越好（优先跳出重复）
        x{n,}?         重复>=n次匹配x，越少越好（优先跳出重复）
        x{n}?          重复n次匹配x
分组：
        (re)           编号的捕获分组
        (?P<name>re)   命名并编号的捕获分组
        (?:re)         不捕获的分组
        (?flags)       设置当前所在分组的标志，不捕获也不匹配
        (?flags:re)    设置re段的标志，不捕获的分组
        I              大小写敏感（默认关闭）
        m              ^和$在匹配文本开始和结尾之外，还可以匹配行首和行尾（默认开启）
        s              让.可以匹配\n（默认关闭）
        U              非贪婪的：交换x*和x*?、x+和x+?……的含义（默认关闭）
边界：
        ^              匹配文本开始，标志m为真时，还匹配行首
        $              匹配文本结尾，标志m为真时，还匹配行尾
        \A             匹配文本开始
        \b             单词边界（一边字符属于\w，另一边为文首、文尾、行首、行尾或属于\W）
        \B             非单词边界
        \z             匹配文本结尾
*/
func main() {
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)
	//创建正则模板
	r, _ := regexp.Compile("p([a-z]+)ch")
	fmt.Println(r.FindString("peach punch"))
	fmt.Println(r.FindStringIndex("peach punch"))


	fmt.Println(r.ReplaceAllString("a peach","<fruit>"))
}
