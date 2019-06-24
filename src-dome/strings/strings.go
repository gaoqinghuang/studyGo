package main

func main()  {
	//总结。主要用的字符串剪切

	//群体替换
	//r := strings.NewReplacer("<", "&lt;", ">", "&gt;")
	//fmt.Println(r.Replace("This is <b>HTML</b>!"))

	//等看到io操作就能看到了@todo 未知
	//reader := strings.NewReader("ssss")

	//按指定字符拼凑
	//fmt.Println(strings.Join([]string{"foo", "bar", "baz"}, ", "))

	//按指定符号分隔符号保留
	//fmt.Println(strings.SplitAfter("a,b,c", ","))

	//按指定符号分隔
	//fmt.Println(strings.Split("a,b,c", ","))

	//按空白符切割
	//fmt.Println(strings.Fields("  foo bar  baz   "))

	//去除前缀
	//fmt.Println(strings.TrimPrefix("Goodbye,, world!", "Goodbye,"))

	//去除前后空白符
	//fmt.Println(strings.TrimSpace(" \t\n a lone gopher \n\t\r\n"))

	//返回count个s串联的字符串。
	//fmt.Println("ba" + strings.Repeat("na", 8))//banana

	//每个单词都改成大写,有bug不能很好的处理Unicode标点符号
	//fmt.Println(strings.Title("her's royal highness"))

	//s中第一个满足函数f的位置i
	//f := func(c rune) bool {
	//	return unicode.Is(unicode.Han, c)
	//}
	//fmt.Println(strings.IndexFunc("Hello, 世界", f))//7
	//fmt.Println(strings.IndexFunc("Hello, world", f))//0


	//字符c在s中第一次出现的位置，不存在则返回-1。
	//fmt.Println(strings.IndexByte("0nde1Byte",110))//这里的110是n,true

	//子串sep在字符串s中第一次出现的位置，不存在则返回-1。
	//fmt.Println(strings.Index("chicken", "ken"))//4

	//返回字符串s中有几个不重复的sep子串
	//fmt.Println(strings.Count("cheese", "e"))//3

	//判断字符串s是否包含字符串chars中的任一字符
	//fmt.Println(strings.ContainsAny("team", "ti"))//true

	//判断是否包含,rune
	//fmt.Println(strings.ContainsRune("11131",'1'))//true

	//判断是否包含
	//fmt.Println(strings.Contains("seafood", "foo")) //true

	//判断是否有后缀
	//fmt.Println(strings.HasSuffix("aaa","a"))//true

	//判断是否有前缀
	//fmt.Println(strings.HasPrefix("aaa","a"))//true

	//不管大小写的比较
	//fmt.Println(strings.EqualFold("gO","Go"))//true
}
